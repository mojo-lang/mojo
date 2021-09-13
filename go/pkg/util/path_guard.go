package util

import (
	"strings"
)

type PathGuard struct {
	Paths []string
}

func (g *PathGuard) Check(path string) error {
	if IsExist(path) {
		err := g.Clear(path)
		if err != nil {
			return err
		}
		return nil
	} else {
		return CreateDir(path)
	}
}

func (g *PathGuard) Clear(path string) error {
	if !g.cleared(path) {
		err := ClearGeneratedFiles(path)
		if err != nil {
			return err
		}

		g.Paths = append(g.Paths, path)
	}
	return nil
}

func (g *PathGuard) cleared(path string) bool {
	for _, p := range g.Paths {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}
