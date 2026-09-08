package compiler

import (
	path2 "path"
	"sort"

	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	data2 "github.com/mojo-lang/mojo/go/pkg/compiler/go/generator/data"
)

type GoMod struct {
	*data2.Data
}

func (g *GoMod) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	_ = ctx
	gm := &data2.GoMod{}
	gm.Name = pkg.GoModName()
	gm.Version = "1.16"
	seen := make(map[string]bool)

	var names []string
	for name := range pkg.Dependencies {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, k := range names {
		d := pkg.Dependencies[k]
		dep := &data2.Dependency{}
		resolved := pkg.ResolvedDependencies[k]
		dep.Name = resolved.GoModName()
		if dep.Name == gm.Name || seen[dep.Name] {
			continue
		}
		seen[dep.Name] = true

		if len(d.Path) > 0 {
			dep.Version = "v0.0.0-00010101000000-000000000000"
			dep.Path = path2.Join("..", d.Path, "go")
		} else {
			minVersion := d.GetVersion().GetRange().GetStart().Format()
			if len(minVersion) == 0 {
				minVersion = "0.0.0"
			}
			dep.Version = "v" + minVersion
		}

		gm.Dependencies = append(gm.Dependencies, dep)
	}

	g.Data.GoMod = gm
	return nil
}
