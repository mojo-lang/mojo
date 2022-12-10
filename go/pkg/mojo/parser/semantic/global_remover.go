package semantic

import (
	"errors"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
)

func init() {
	plugin.RegisterPlugin(NewGlobalRemover(nil))
}

type GlobalRemover struct {
	plugin.BasicPlugin
	Options core.Options

	processedPackages map[string]bool
}

func NewGlobalRemover(options core.Options) *GlobalRemover {
	return &GlobalRemover{
		BasicPlugin: plugin.BasicPlugin{
			Name:          "semantic.global-remover",
			Group:         "semantic",
			GroupPriority: 3,
			Priority:      100,
			Creator: func(options core.Options) plugin.Plugin {
				return NewGlobalRemover(options)
			},
		},
		Options:           options,
		processedPackages: make(map[string]bool),
	}
}

func (r *GlobalRemover) ParsePackage(ctx context.Context, pkg *lang.Package) error {
	if pkg == nil {
		return errors.New("empty packages to parse")
	}

	if !pkg.IsGlobal() {
		return nil
	}

	logs.Infow("enter to the plugin", "plugin", r.Name)

	pkgName := plugin.ContextPackageName(ctx)
	processPkg := r.removeGlobal(pkg, pkgName)
	if processPkg == nil {
		return errors.New("the GlobalRemover does not match the GlobalMaker")
	}

	if plugins := plugin.ContextPlugins(ctx); plugins != nil && plugins.Next() != nil {
		if err := plugins.ParsePackage(ctx, processPkg); err != nil {
			return err
		}
	}
	return nil
}

func (r *GlobalRemover) removeGlobal(pkg *lang.Package, pkgName string) *lang.Package {
	if !pkg.IsPadding() && !r.processedPackages[pkg.FullName] &&
		(len(pkgName) == 0 || pkg.FullName == pkgName || pkg.Name == pkgName) {
		r.processedPackages[pkg.FullName] = true
		return pkg
	}

	for _, child := range pkg.Children {
		if p := r.removeGlobal(child, pkgName); p != nil {
			return p
		}
	}
	return nil
}
