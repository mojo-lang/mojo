package pkg

import (
	"errors"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/parser"
	path2 "path"
)

func ParsePackageFile(path string) (*lang.Package, error) {
	parser := parser.New()
	packageFile := path2.Join(path, "package.mojo")
	file, err := parser.ParseFile(packageFile)
	if err != nil {
		logs.Errorw("failed to parse package file", "file", packageFile, "error", err.Error())
		return nil, err
	}

	if len(file.Statements) == 0 {
		logs.Errorw("not a valid package.mojo file, has no statement include", "file", packageFile)
		return nil, errors.New("there is no package declaration in mojo file")
	}

	decl := file.Statements[0].GetDeclaration()
	if decl == nil {
		logs.Errorw("not a valid package.mojo file, has no declaration statement include", "file", packageFile)
		return nil, errors.New("there is no package declaration in mojo file")
	}

	pkgDecl := decl.GetPackageDecl()
	if pkgDecl == nil {
		logs.Errorw("not a valid package.mojo file, has no declaration statement include", "file", packageFile)
		return nil, errors.New("there is no package declaration in mojo file")
	}

	return pkgDecl.Package, nil
}