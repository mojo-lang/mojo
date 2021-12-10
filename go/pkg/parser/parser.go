package parser

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic"
	"github.com/mojo-lang/mojo/go/pkg/parser/syntax"
	"io/fs"
)

type Parser struct {
	Global *lang.Package
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) ParseFile(filename string, fileSys fs.FS) (*lang.SourceFile, error) {
	return syntax.Parser{}.ParseFile(filename, fileSys)
}

func (p *Parser) ParsePackage(path string, pkgName string, fileSys fs.FS) (map[string]*lang.Package, error) {
	syntaxParser := syntax.Parser{}
	pkgs, err := syntaxParser.ParsePackage(path, pkgName, fileSys)
	if err != nil {
		logs.Errorw("failed to parse package", "package", pkgName, "error", err.Error())
		return nil, err
	}

	options := make(map[string]interface{})
	var packages lang.Packages
	for _, pkg := range pkgs {
		packages = append(packages, pkg)
	}

	if !hasRootPackage(pkgs, pkgName) {
		pkgs[pkgName] = &lang.Package{
			Name:     lang.GetPackageName(pkgName),
			FullName: pkgName,
		}
		packages = append(packages, pkgs[pkgName])
	}

	p.Global, err = semantic.ParsePackages(packages, p.Global, options)
	if err != nil {
		return nil, err
	}

	return pkgs, err
}

func hasRootPackage(packages map[string]*lang.Package, rootName string) bool {
	for k, _ := range packages {
		if k == rootName {
			return true
		}
	}
	return false
}
