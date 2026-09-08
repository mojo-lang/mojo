package builder

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

type Builder struct {
	PWD     string
	Path    string
	Package *lang.Package

	APIEnabled bool
}

func (b Builder) GetAbsolutePath() string {
	if b.Package != nil && b.Package.GetExtraString("path") != "" {
		return b.Package.GetExtraString("path")
	}
	return util.GetAbsolutePath(b.PWD, b.Path)
}
