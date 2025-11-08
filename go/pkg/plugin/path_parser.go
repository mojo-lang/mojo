package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type PathParser interface {
	ParsePath(ctx context.Context, pkgPath string) (*lang.Package, error)
}
