package builder

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	path2 "path"
)

type Builder struct {
	PWD     string
	Path    string
	Package *lang.Package

	APIEnabled bool
}

func (b Builder) GetAbsolutePath() string {
	return GetAbsolutePath(b.PWD, b.Path)
}

func GetAbsolutePath(pwd string, path string) string {
	if path2.IsAbs(path) {
		return path
	}
	return path2.Join(pwd, path)
}
