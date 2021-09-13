package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type GoPackageNameCompiler struct {
}

func (g *GoPackageNameCompiler) Compile(ctx *context.Context, pkg *lang.Package) error {
	goPkgName := pkg.GoPackageName()
	if goPkgName == pkg.Name {
		return nil
	}

	for _, sourceFile := range pkg.SourceFiles {
		ctx.Open(sourceFile)
		for _, statement := range sourceFile.Statements {
			if decl := statement.GetDeclaration(); decl != nil {
				switch decl.Declaration.(type) {
				case *lang.Declaration_StructDecl:
					if structDecl := decl.GetStructDecl(); structDecl != nil {
						structDecl.Attributes = lang.SetStringAttribute(structDecl.Attributes, "go_package_name", goPkgName)
					}
				case *lang.Declaration_TypeAliasDecl:
				case *lang.Declaration_EnumDecl:
				case *lang.Declaration_InterfaceDecl:
				}
			}
		}
		ctx.Close()
	}
	return nil
}
