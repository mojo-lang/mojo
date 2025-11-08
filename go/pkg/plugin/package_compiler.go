package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type PackageCompiler interface {
	CompilePackage(ctx context.Context, pkg *lang.Package) error
}

func IsPackageCompiler(c interface{}) bool {
	if _, ok := c.(PackageCompiler); ok {
		return true
	}
	return false
}
