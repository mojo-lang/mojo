package util

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
)

type PathGuard struct {
	DisableClear       bool
	OnlyClearGenerated bool
	Suffixes           []string

	paths []string
}

func (g *PathGuard) Check(path string) error {
	if core.IsExist(path) {
		if g.DisableClear {
			return nil
		}

		err := g.Clear(path)
		if err != nil {
			return err
		}
		return nil
	} else {
		err := core.CreateDir(path)
		if err != nil {
			return err
		}
		g.paths = append(g.paths, path)
		return nil
	}
}

func (g *PathGuard) Clear(path string) error {
	if !g.cleared(path) {
		var err error
		if g.OnlyClearGenerated {
			err = ClearGeneratedFiles(path, g.Suffixes...)
		} else {
			err = ClearFiles(path, g.Suffixes...)
		}
		if err != nil {
			return err
		}

		g.paths = append(g.paths, path)
	}
	return nil
}

func (g *PathGuard) cleared(path string) bool {
	for _, p := range g.paths {
		if path == p {
			return true
		}
	}
	return false
}
