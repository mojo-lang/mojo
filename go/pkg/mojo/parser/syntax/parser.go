package syntax

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin/parser"
	"io/fs"
	"path"
	"strings"
)

func init() {
	plugin.RegisterPlugin(New(nil))
}

const pluginName = "syntax.parser"

type Parser struct {
	plugin.BasicPlugin
}

func New(options core.Options) *Parser {
	return &Parser{
		BasicPlugin: plugin.BasicPlugin{
			Name:          pluginName,
			Group:         "syntax",
			GroupPriority: 2,
			Priority:      1,
			Creator: func(options core.Options) plugin.Plugin {
				return New(options)
			},
		},
	}
}

func (p Parser) ParseExpression(expr string) (*lang.Expression, error) {
	if file, err := p.ParseString(expr); err != nil {
		return nil, err
	} else {
		if len(file.Statements) > 0 {
			if expression := file.Statements[0].GetExpression(); expression != nil {
				return expression, nil
			}
		}
	}
	return nil, errors.New("not a valid expression")
}

func (p Parser) ParseString(mojo string) (*lang.SourceFile, error) {
	input := antlr.NewInputStream(mojo)
	return p.ParseStream("", input)
}

func (p Parser) ParseStream(fileName string, input *antlr.InputStream) (*lang.SourceFile, error) {
	lexer := NewMojoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	errorListener := NewMojoErrorListener()
	errorListener.FileName = fileName

	parser := NewMojoParser(stream)
	parser.AddErrorListener(errorListener)
	parser.BuildParseTrees = true

	tree := parser.MojoFile()
	visitor := NewMojoFileVisitor()
	result := visitor.Visit(tree).(bool)

	if result && errorListener.Error == nil {
		CommentParser{}.Parse(stream)

		return visitor.SourceFile, nil
	}

	if errorListener.Error != nil {
		return nil, errorListener.Error
	} else {
		const msg = "failed to parse mojo file"
		logs.Errorw(msg, "file", fileName)
		return nil, errors.New(msg)
	}
}

func (p Parser) ParseFile(ctx context.Context, fileName string, fileSys fs.FS) (*lang.SourceFile, error) {
	if bytes, err := fs.ReadFile(fileSys, fileName); err != nil {
		return nil, err
	} else {
		sourceFile, err := p.ParseStream(fileName, antlr.NewInputStream(string(bytes)))
		if err != nil {
			return nil, err
		}

		sourceFile.Name = path.Base(fileName)
		sourceFile.FullName = fileName

		return sourceFile, nil
	}
}

func (p Parser) ParsePackagePath(ctx context.Context, pkgPath string, fileSys fs.FS) (*lang.Package, error) {
	currentPkg := parser.ContextDeclaredPackage(ctx)
	currentPkgName := ""
	if currentPkg != nil {
		if currentPkg.GetExtraBool(pluginName) {
			return currentPkg, nil
		}
		currentPkgName = currentPkg.FullName
	} else {
		currentPkgName = parser.ContextPackageName(ctx)
		currentPkg = &lang.Package{
			Name:     lang.GetPackageName(currentPkgName),
			FullName: currentPkgName,
		}
	}

	currentPath := path.Join(pkgPath, lang.PackageNameToPath(currentPkgName))
	files, err := fs.ReadDir(fileSys, currentPath)
	if err != nil {
		logs.Errorw("failed to read source directory", "path", currentPath, "error", err.Error())
		return nil, err
	}

	parentPkg := context.Package(ctx)
	if parentPkg != nil {
		parentPkg.Children = append(parentPkg.Children, currentPkg)

		currentPkg.Authors = parentPkg.Authors
		currentPkg.Repository = parentPkg.Repository
		currentPkg.Version = parentPkg.Version
	}

	thisCtx := context.WithType(ctx, currentPkg)
	for _, f := range files {
		if f.IsDir() {
			pkgName := ""
			if len(currentPkgName) == 0 {
				pkgName = f.Name()
			} else {
				pkgName = currentPkgName + "." + f.Name()
			}

			newCtx := parser.WithDeclaredPackage(context.WithType(thisCtx, currentPkg), &lang.Package{
				Name:     lang.GetPackageName(pkgName),
				FullName: pkgName,
			})
			if _, err = p.ParsePackagePath(newCtx, pkgPath, fileSys); err != nil {
				return nil, err
			}
		} else {
			if !strings.HasSuffix(f.Name(), ".mojo") {
				continue
			}

			sourceFile, err := p.ParseFile(thisCtx, path.Join(currentPath, f.Name()), fileSys)
			if err != nil {
				return nil, err
			}

			sourceFile.PackageName = currentPkgName
			sourceFile.FullName = path.Join(lang.PackageNameToPath(currentPkgName), sourceFile.Name)
			currentPkg.SourceFiles = append(currentPkg.SourceFiles, sourceFile)
		}
	}

	currentPkg.SetExtraBool(pluginName, true)
	return currentPkg, nil
}
