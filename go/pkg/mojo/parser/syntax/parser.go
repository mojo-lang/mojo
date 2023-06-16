package syntax

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

func init() {
	plugin.RegisterPlugin(New(nil))
}

const pluginName = "syntax.parser"

type Parser struct {
	plugin.BasicPlugin

	Options core.Options
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
		Options: options,
	}
}

func (p *Parser) ParseExpression(expr string) (*lang.Expression, error) {
	if file, err := p.ParseString(context.Empty(), expr); err != nil {
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

func (p *Parser) ParseString(ctx context.Context, mojo string) (*lang.SourceFile, error) {
	input := antlr.NewInputStream(mojo)
	return p.ParseStream(plugin.ContextFilename(ctx), input)
}

func (p *Parser) ParseStream(fileName string, input *antlr.InputStream) (*lang.SourceFile, error) {
	lexer := NewMojoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	errorListener := util.NewErrorListener(fileName, false)
	errorListener.FileName = fileName

	parser := NewMojoParser(stream)
	parser.AddErrorListener(errorListener)
	parser.BuildParseTrees = true

	tree := parser.MojoFile()
	if sourceFile, ok := NewMojoFileVisitor().Visit(tree).(*lang.SourceFile); ok {
		if errorListener.Errors == nil {
			comments := CommentParser{}.Parse(stream)
			CommentMerger(comments).Merge(sourceFile)
			return sourceFile, nil
		} else {
			return nil, errorListener.Errors
		}
	}

	if errorListener.Errors != nil {
		return nil, logs.NewErrorw("failed to parse mojo file", "file", fileName, "error", errorListener.Errors.Error())
	}
	return nil, logs.NewErrorw("failed to parse mojo file", "file", fileName)
}

func (p *Parser) ParseFile(ctx context.Context, fileName string) (sourceFile *lang.SourceFile, err error) {
	return plugin.ParseFile(p, ctx, fileName)
}

func (p *Parser) ParsePackage(ctx context.Context, pkg *lang.Package) (err error) {
	if util.IsPackageProcessed(pkg, pluginName) {
		return nil
	}

	pkgPath := pkg.GetExtraString("path")
	if !strings.HasPrefix(pkg.FullName, "mojo.") {
		pkgPath = path.Join(pkgPath, "mojo")
	}

	workingDir := pkg.GetExtraString("workingDir")
	_, err = p.ParsePath(plugin.WithDeclaredPackage(ctx, pkg), path.Join(workingDir, pkgPath))
	util.SetPackageProcessed(pkg, pluginName)
	return
}

func (p *Parser) ParsePath(ctx context.Context, pkgPath string) (*lang.Package, error) {
	currentPkg := plugin.ContextDeclaredPackage(ctx)
	currentPkgName := ""
	if currentPkg != nil {
		if util.IsPackageProcessed(currentPkg, pluginName) {
			return currentPkg, nil
		}
		currentPkgName = currentPkg.FullName
	} else {
		currentPkgName = plugin.ContextPackageName(ctx)
		currentPkg = &lang.Package{
			Name:     lang.GetPackageName(currentPkgName),
			FullName: currentPkgName,
		}
	}

	fileSys := plugin.ContextFs(ctx)
	if fileSys == nil {
		fileSys = os.DirFS(pkgPath)
		ctx = plugin.WithFs(ctx, fileSys)
		pkgPath = ""
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

			newCtx := plugin.WithDeclaredPackage(context.WithType(thisCtx, currentPkg), &lang.Package{
				Name:     lang.GetPackageName(pkgName),
				FullName: pkgName,
			})
			if _, err = p.ParsePath(newCtx, pkgPath); err != nil {
				return nil, err
			}
		} else {
			if !strings.HasSuffix(f.Name(), ".mojo") {
				continue
			}

			if sourceFile, err := p.ParseFile(thisCtx, path.Join(currentPath, f.Name())); err != nil {
				return nil, err
			} else {
				sourceFile.PackageName = currentPkgName
				sourceFile.FullName = path.Join(lang.PackageNameToPath(currentPkgName), sourceFile.Name)
				currentPkg.SourceFiles = append(currentPkg.SourceFiles, sourceFile)
			}
		}
	}

	util.SetPackageProcessed(currentPkg, pluginName)
	return currentPkg, nil
}
