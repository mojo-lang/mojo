package compiler

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    path2 "path"
)

type Dependency struct {
    Name    string
    Version string
    Path    string
}

type GoMod struct {
    Name         string
    Version      string // go
    Dependencies []*Dependency
}

func (g *GoMod) Compile(pkg *lang.Package) error {
    g.Name = pkg.GoModName()
    g.Version = "1.16"

    for k, d := range pkg.Dependencies {
        dep := &Dependency{}
        resolved := pkg.ResolvedDependencies[k]
        dep.Name = resolved.GoModName()

        if len(d.Path) > 0 {
            dep.Version = "v0.0.0-00010101000000-000000000000"
            dep.Path = path2.Join(d.Path, "go")
        } else {
            minVersion := d.GetVersion().GetRange().GetMin().Format()
            if len(minVersion) == 0 {
                minVersion = "0.0.0"
            }
            dep.Version = "v" + minVersion
        }

        g.Dependencies = append(g.Dependencies, dep)
    }

    return nil
}
