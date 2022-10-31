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
			MarkedPackages: make(map[string]bool),
		},
	}
}

func (c *GoPackageNameCompiler) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
	pkg := context.Package(ctx)
	pkgFullName := pkg.GetFullName()

	if !c.MarkedPackages[pkgFullName] {
		c.MarkedPackages[pkgFullName] = true
		logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompilePackage", "pkg", pkg.FullName)
	}

	goPkgName := pkg.GoPackageName()
	if goPkgName == pkg.Name {
		return nil
	}

	decl.SetStringAttribute("go_package_name", goPkgName)
	return nil
}
