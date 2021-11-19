package builder

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type Builder struct {
	PWD     string
	Path    string
	Package *lang.Package

	APIEnabled bool
}
