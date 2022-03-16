package syntax

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPackageDeclarationVisitor_VisitPackageDeclaration(t *testing.T) {
    const packageDecl = `
	package mojo.openapi {
    	license: 'Apache'
    	dependencies: {
        	'mojo.core': {path: '../core', version: '^0.1'}
        	'mojo.document': {path: '../document', version: '^0.0.2'}
    	}
	}`
    parser := &Parser{}
    file, err := parser.ParseString(packageDecl)
    assert.NoError(t, err)

    pkg := getPackage(file)
    assert.NotNil(t, pkg)

    assert.Equal(t, "Apache", pkg.License)
    assert.Equal(t, "../core", pkg.Dependencies["mojo.core"].Path)
    assert.Equal(t, "../document", pkg.Dependencies["mojo.document"].Path)
    assert.Equal(t, uint64(1), pkg.Dependencies["mojo.core"].Version.Range.Min.Minor)
    assert.Equal(t, uint64(2), pkg.Dependencies["mojo.core"].Version.Range.Max.Minor)
}

func TestPackageDeclarationVisitor_VisitPackageDeclaration2(t *testing.T) {
    const packageDecl = `
	package mojo.openapi {
    	license: 'Apache'
    	dependencies: {
        	'mojo.core': '^0.1'
        	'mojo.document': '^0.0.2'
    	}
	}`
    parser := &Parser{}
    file, err := parser.ParseString(packageDecl)
    assert.NoError(t, err)

    pkg := getPackage(file)
    assert.NotNil(t, pkg)

    assert.Equal(t, "Apache", pkg.License)
    assert.Equal(t, uint64(1), pkg.Dependencies["mojo.core"].Version.Range.Min.Minor)
    assert.Equal(t, uint64(2), pkg.Dependencies["mojo.core"].Version.Range.Max.Minor)
}

func getPackage(file *lang.SourceFile) *lang.Package {
    if len(file.Statements) > 0 {
        statement := file.Statements[0]
        if decl := statement.GetDeclaration(); decl != nil {
            if pkgDecl := decl.GetPackageDecl(); pkgDecl != nil {
                return pkgDecl.Package
            }
        }
    }
    return nil
}
