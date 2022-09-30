package compiler

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/data"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    path2 "path"
)

type GoMod struct {
    *data.Data
}

func (g *GoMod) CompilePackage(ctx context.Context, pkg *lang.Package) error {

    gm := &data.GoMod{}
    gm.Name = pkg.GoModName()
    gm.Version = "1.16"

    for k, d := range pkg.Dependencies {
        dep := &data.Dependency{}
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
