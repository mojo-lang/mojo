package compiler

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
)

const goPackageNameName = "compiler.go-package-name"

func init() {
	plugin.RegisterPlugin(NewGoPackageNameCompiler(nil))
}

type GoPackageNameCompiler struct {
	plugin.BasicPlugin
}

func NewGoPackageNameCompiler(options core.Options) *GoPackageNameCompiler {
	return &GoPackageNameCompiler{
		BasicPlugin: plugin.BasicPlugin{
			Name:          goPackageNameName,
			Group:         "compiler",
			GroupPriority: 10,
			Priority:      11,
			Creator: func(options core.Options) plugin.Plugin {
				return NewGoPackageNameCompiler(options)
			},
		},
	}
}

func (c *GoPackageNameCompiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompilePackage", "pkg", pkg.FullName)

	goPkgName := pkg.GoPackageName()
	if goPkgName == pkg.Name {
		return nil
	}

	//thisCtx := context.WithType(ctx, pkg)
	for _, sourceFile := range pkg.SourceFiles {
		//fileCtx := context.WithType(ctx, thisCtx)
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
	}
	return nil
}
