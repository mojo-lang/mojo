package parser

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic"
	"github.com/mojo-lang/mojo/go/pkg/parser/syntax"
)

type Parser struct {
	Dependencies map[string]*lang.Package
}

func New() *Parser {
	return &Parser{}
}

func (p Parser) ParseFile(filename string) (*lang.SourceFile, error) {
	parser := syntax.Parser{}
	return parser.ParseFile(filename)
}

func hasRootPackage(packages map[string]*lang.Package, rootName string) bool {
	for k, _ := range packages {
		if k == rootName {
			return true
		}
	}
	return false
}

func (p Parser) ParsePackage(path string, pkgName string) (map[string]*lang.Package, error) {
	syntaxParser := syntax.Parser{}
	pkgs, err := syntaxParser.ParsePackage(path, pkgName)
	if err != nil {
		logs.Errorw("failed to parse package", "package", pkgName, "error", err.Error())
		return nil, err
	}

	options := make(map[string]interface{})
	var packages lang.Packages
	for _, pkg := range pkgs {
		packages = append(packages, pkg)
	}
	for key, pkg := range p.Dependencies {
		packages = append(packages, pkg)
		options[key+".disable"] = true
	}

	if !hasRootPackage(pkgs, pkgName) {
		pkgs[pkgName] = &lang.Package{
			Name:     lang.GetPackageName(pkgName),
			FullName: pkgName,
		}
		packages = append(packages, pkgs[pkgName])
	}

	err = semantic.ParsePackages(packages, options)
	if err != nil {
		return nil, err
	}

	return pkgs, err
}
