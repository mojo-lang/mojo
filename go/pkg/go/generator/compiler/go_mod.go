package compiler

import (
	path2 "path"

	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	data2 "github.com/mojo-lang/mojo/go/pkg/go/generator/data"
)

type GoMod struct {
	*data2.Data
}

func (g *GoMod) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	_ = ctx
	gm := &data2.GoMod{}
	gm.Name = pkg.GoModName()
	gm.Version = "1.16"

	for k, d := range pkg.Dependencies {
		dep := &data2.Dependency{}
		resolved := pkg.ResolvedDependencies[k]
		dep.Name = resolved.GoModName()

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
