package semantic

import (
	"errors"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic/circle"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic/identifier"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
	"sort"
)

var _ = identifier.Namer{}
var _ = identifier.Resolver{}
var _ = circle.Resolver{}

func ParsePackages(packages lang.Packages, global *lang.Package, options core.Options) (*lang.Package, error) {
	if len(packages) == 0 {
		return nil, errors.New("empty packages to parse")
	}
	global = treePackages(packages, global)

	// sort plugin by priority from higher to lower
	plugin.Sort(plugin.ParserPlugins)
	for _, p := range plugin.ParserPlugins {
		err := plugin.ParsePackage(p, context.WithOptions(context.Empty(), options), global)
		if err != nil {
			return nil, err
		}
	}

	return global, nil
}

func treePackages(packages lang.Packages, global *lang.Package) *lang.Package {
	var names []string
	for _, pkg := range packages {
		if pkg.IsGlobal() {
			global = pkg
		}
		names = append(names, pkg.FullName)
	}
	sort.Strings(names)

	var existPackages map[string]*lang.Package

	if global == nil {
		// add an empty package as the global package for maintain the global scope
		global = lang.NewGlobalPackage()
		existPackages = make(map[string]*lang.Package)
		packages = append(packages, global)
	} else {
		existPackages = global.GetAllPackage()
		for _, pkg := range existPackages {
			packages = append(packages, pkg)
		}
	}

	newPkgs := make(map[string]bool)
	for i := len(names) - 1; i >= 0; i-- {
		parent := lang.GetPackageParentName(names[i])
		if len(parent) == 0 {
			continue
		}
		_, exist := existPackages[parent]
		for (i == 0 || (parent != names[i-1])) && len(parent) > 0 && !newPkgs[parent] && !exist {
			p := &lang.Package{
				Name:     lang.GetPackageName(parent),
				FullName: parent,
				Scope:    nil,
			}
			packages = append(packages, p)
			newPkgs[p.FullName] = true

			parent = lang.GetPackageParentName(parent)
		}
	}

	sort.Sort(packages)

	for i, pkg := range packages {
		pkg.Children = nil
		for j := i + 1; j < len(packages); j++ {
			parent := lang.GetPackageParentName(packages[j].FullName)
			if pkg.FullName == parent {
				pkg.Children = append(pkg.Children, packages[j])
			}
		}
	}

	return global
}
