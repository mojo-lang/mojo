package parser

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type PackageParser interface {
	Parse(ctx context.Context, pkg *lang.Package) error
}
