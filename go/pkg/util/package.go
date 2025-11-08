package util

import "github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

func SetPackageProcessed(pkg *lang.Package, plugin string) *lang.Package {
	if pkg != nil {
		pkg.SetExtraBool(plugin, true)
	}
	return pkg
}

func IsPackageProcessed(pkg *lang.Package, plugin string) bool {
	if pkg != nil {
		return pkg.GetExtraBool(plugin)
	}
	return false
}
