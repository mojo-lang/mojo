package syntax

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"io/fs"
	path2 "path"
	"strings"
)

type Parser struct {
	Path     string
	Package  string
	FileName string
}

func New() *Parser {
	return &Parser{}
}

func (p Parser) ParseString(mojo string) (*lang.SourceFile, error) {
	input := antlr.NewInputStream(mojo)
	return p.ParseStream(input)
}

func (p Parser) ParseFile(filename string, fileSys fs.FS) (*lang.SourceFile, error) {
	if bytes, err := fs.ReadFile(fileSys, filename); err != nil {
		return nil, err
	} else {
		return p.parseFile(filename, antlr.NewInputStream(string(bytes)))
	}
}

func (p Parser) parseFile(name string, input *antlr.InputStream) (*lang.SourceFile, error) {
	p.FileName = name
	sourceFile, err := p.ParseStream(input)
	if err != nil {
		return nil, err
	}

	sourceFile.Name = path2.Base(name)
	sourceFile.FullName = name
	return sourceFile, nil
}

func (p Parser) ParseStream(input *antlr.InputStream) (*lang.SourceFile, error) {
	lexer := NewMojoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	errorListener := NewMojoErrorListener()
	errorListener.FileName = p.FileName

	parser := NewMojoParser(stream)
	parser.AddErrorListener(errorListener)
	parser.BuildParseTrees = true

	tree := parser.MojoFile()
	visitor := NewMojoFileVisitor()
	result := visitor.Visit(tree).(bool)
	if result && errorListener.Error == nil {
		return visitor.SourceFile, nil
	}

	if errorListener.Error != nil {
		return nil, errorListener.Error
	} else {
		const msg = "failed to parse mojo file"
		logs.Errorw(msg, "file", p.FileName, "path", p.Path, "package", p.Package)
		return nil, errors.New(msg)
	}
}

/// path: root path for the package
func (p Parser) ParsePackage(path string, pkg string, fileSys fs.FS) (map[string]*lang.Package, error) {
	p.Path = path
	p.Package = pkg

	packages, err := p.parsePackage(path, pkg, fileSys)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func (p Parser) parsePackage(path string, pkg string, fileSys fs.FS) (map[string]*lang.Package, error) {
	packages := make(map[string]*lang.Package)

	currentPath := path2.Join(path, lang.PackageNameToPath(pkg))
	files, err := fs.ReadDir(fileSys, currentPath)
	if err != nil {
		logs.Errorw("failed to read source directory", "path", currentPath, "error", err.Error())
		return nil, err
	}

	currentPkgName := pkg
	for _, f := range files {
		if f.IsDir() {
			pkgName := ""
			if len(currentPkgName) == 0 {
				pkgName = f.Name()
			} else {
				pkgName = currentPkgName + "." + f.Name()
			}
			pkgs, err := p.parsePackage(path, pkgName, fileSys)
			if err != nil {
				return nil, err
			}

			for k, v := range pkgs {
				packages[k] = v
			}
		} else {
			if !strings.HasSuffix(f.Name(), ".mojo") {
				continue
			}

			sourceFile, err := p.ParseFile(path2.Join(currentPath, f.Name()), fileSys)
			if err != nil {
				return nil, err
			}

			if packages[currentPkgName] == nil {
				packages[currentPkgName] = &lang.Package{}
			}

			sourceFile.PackageName = currentPkgName
			sourceFile.FullName = path2.Join(lang.PackageNameToPath(currentPkgName), sourceFile.Name)
			packages[currentPkgName].SourceFiles = append(packages[currentPkgName].SourceFiles, sourceFile)
			packages[currentPkgName].Name = path2.Base(currentPath)
			packages[currentPkgName].FullName = currentPkgName
		}
	}
	return packages, nil
}
