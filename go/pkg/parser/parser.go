package parser

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mojo-lang/lang/go/pkg/lang"
	"io/ioutil"
	"log"
	path2 "path"
	"strings"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p Parser) Parse(mojo string) (*lang.SourceFile, error) {
	input := antlr.NewInputStream(mojo)
	return p.ParseStream(input)
}

func (p Parser) ParseStream(input *antlr.InputStream) (*lang.SourceFile, error) {
	lexer := NewMojoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	parser := NewMojoParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true

	tree := parser.MojoFile()
	visitor := NewMojoFileVisitor()
	result := visitor.Visit(tree).(bool)
	if result {
		print(visitor.SourceFile.String())
		return visitor.SourceFile, nil
	} else {
		print("parser failed")
	}

	return nil, errors.New("parse failed")
}

func (p Parser) ParseFile(filename string) (*lang.SourceFile, error) {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}
	return p.ParseStream(input.InputStream)
}

func (p Parser) ParsePackage(path string) (map[string]*lang.Package, error) {
	packages := make(map[string]*lang.Package)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	currentPkgName := strings.TrimPrefix(path, "./")
	currentPkgName = strings.TrimPrefix(currentPkgName, "../")
	currentPkgName = strings.ReplaceAll(currentPkgName, "/", ".")
	currentPkgName = strings.ReplaceAll(currentPkgName, "\\", ".")

	for _, f := range files {
		if f.IsDir() {
			pkgs, err := p.ParsePackage(path2.Join(path, f.Name()))
			if err != nil {
				return nil, err
			}

			for k, v := range pkgs {
				packages[k] = v
			}
		} else {
			sourceFile, err := p.ParseFile(path2.Join(path, f.Name()))
			if err != nil {
				return nil, err
			}

			if packages[currentPkgName] == nil {
				packages[currentPkgName] = &lang.Package{}
			}

			packages[currentPkgName].SourceFiles = append(packages[currentPkgName].SourceFiles, sourceFile)
		}
	}
	return packages, nil
}
