package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type PathParser interface {
	ParsePath(ctx context.Context, pkgPath string) (*lang.Package, error)
}
