package semantic

import (
	"errors"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic/circle"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic/identifier"
	"github.com/mojo-lang/mojo/go/pkg/parser/semantic/plugin"
	"sort"
)

var _ = identifier.Namer{}
var _ = identifier.Resolver{}
var _ = circle.Resolver{}

func ParsePackages(packages lang.Packages, options map[string]interface{}) error {
	if len(packages) == 0 {
		return errors.New("empty packages to parse")
	}

	global := treePackages(packages)

	// sort plugin by priority from higher to lower
	plugin.SortPlugins()
	for _, p := range plugin.Plugins {
		ctx := plugin.NewContext()
		p.Parse(ctx, global, options)
	}

	return nil
}

func treePackages(packages lang.Packages) *lang.Package {
	var names []string
	var global *lang.Package
	for _, pkg := range packages {
		if pkg.IsGlobal() {
			global = pkg
		}
		names = append(names, pkg.FullName)
	}
	sort.Strings(names)

	if global == nil {
		// add an empty package as the global package for maintain the global scope
		global = lang.NewGlobalPackage()
		packages = append(packages, global)
	}

	newPkgs := make(map[string]bool)
	for i := len(names) - 1; i >= 0; i-- {
		parent := lang.GetPackageParentName(names[i])
		if len(parent) == 0 {
			continue
		}

		for (i == 0 || (parent != names[i-1])) && len(parent) > 0 && !newPkgs[parent] {
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
		for j := i + 1; j < len(packages); j++ {
			parent := lang.GetPackageParentName(packages[j].FullName)
			if pkg.FullName == parent {
				pkg.Children = append(pkg.Children, packages[j])
			}
		}
	}

	return global
}
