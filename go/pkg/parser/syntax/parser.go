package syntax

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"io/ioutil"
	path2 "path"
	"strings"
)

type Parser struct {
	Path string
	Package string
	FileName string
}

func New() *Parser {
	return &Parser{}
}

func (p Parser) ParseString(mojo string) (*lang.SourceFile, error) {
	input := antlr.NewInputStream(mojo)
	return p.ParseStream(input)
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
	if result {
		return visitor.SourceFile, nil
	}

	logs.Errorw("parse file failed", "file", p.FileName, "path", p.Path, "package", p.Package)
	return nil, errors.New("parse file failed")
}

func (p Parser) ParseFile(filename string) (*lang.SourceFile, error) {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}

	p.FileName = filename
	sourceFile, err := p.ParseStream(input.InputStream)
	if err != nil {
		return nil, err
	}

	sourceFile.Name = path2.Base(filename)
	sourceFile.FullName = filename
	return sourceFile, nil
}

/// path: root path for the package
func (p Parser) ParsePackage(path string, pkg string) (map[string]*lang.Package, error) {
	p.Path = path
	p.Package = pkg

	packages, err := p.parsePackage(path, pkg)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func (p Parser) parsePackage(path string, pkg string) (map[string]*lang.Package, error) {
	packages := make(map[string]*lang.Package)

	currentPath := path2.Join(path, strings.ReplaceAll(pkg, ".", "/"))
	files, err := ioutil.ReadDir(currentPath)
	if err != nil {
		logs.Error("failed to read source directory", "path", currentPath, "error", err.Error())
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
			pkgs, err := p.parsePackage(path, pkgName)
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

			sourceFile, err := p.ParseFile(path2.Join(currentPath, f.Name()))
			if err != nil {
				return nil, err
			}

			if packages[currentPkgName] == nil {
				packages[currentPkgName] = &lang.Package{}
			}

			sourceFile.PackageName = currentPkgName
			sourceFile.FullName = strings.ReplaceAll(currentPkgName, ".", "/") + "/" + sourceFile.Name
			packages[currentPkgName].SourceFiles = append(packages[currentPkgName].SourceFiles, sourceFile)
			packages[currentPkgName].Name = path2.Base(currentPath)
			packages[currentPkgName].FullName = currentPkgName
		}
	}
	return packages, nil
}
