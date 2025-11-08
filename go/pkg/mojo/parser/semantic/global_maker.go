package semantic

import (
	"errors"
	"sort"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

func init() {
	plugin.RegisterPlugin(NewGlobalMaker(nil))
}

type GlobalMaker struct {
	plugin.BasicPlugin
	Options core.Options
	Global  *lang.Package
}

func NewGlobalMaker(options core.Options) *GlobalMaker {
	return &GlobalMaker{
		BasicPlugin: plugin.BasicPlugin{
			Name:          "semantic.global-maker",
			Group:         "semantic",
			GroupPriority: 3,
			Priority:      0,
			Creator: func(options core.Options) plugin.Plugin {
				return NewGlobalMaker(options)
			},
		},
		Options: options,
	}
}

func (r *GlobalMaker) ParsePackage(ctx context.Context, pkg *lang.Package) error {
	logs.Infow("enter to the plugin", "plugin", r.Name, "pkg", pkg.FullName)

	if pkg == nil {
		return errors.New("empty packages to parse")
	}

	pkgs := pkg.GetAllPackageArray()
	r.Global = treePackages(pkgs, r.Global)

	if plugins := plugin.ContextPlugins(ctx); plugins != nil && plugins.Next() != nil {
		if err := plugins.ParsePackage(ctx, r.Global); err != nil {
			return err
		}
	}

	return nil
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
		existPackages = global.GetAllPackages()
		inputLen := len(packages)
		for _, pkg := range existPackages {
			packages = append(packages, pkg)
		}

		for i := 0; i < inputLen; i++ {
			existPackages[packages[i].GetFullName()] = packages[i]
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
				Implicit: true,
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
