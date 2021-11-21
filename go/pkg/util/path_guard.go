package util

import (
	"strings"
)

type PathGuard struct {
	OnlyClearGenerated bool
	Suffixes           []string

	paths []string
}

func (g *PathGuard) Check(path string) error {
	if IsExist(path) {
		err := g.Clear(path)
		if err != nil {
			return err
		}
		return nil
	} else {
		err := CreateDir(path)
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
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}
