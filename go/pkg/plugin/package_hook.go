package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type PackagePreHook interface {
	PrePackage(ctx context.Context, pkg *lang.Package) error
}

type PackagePostHook interface {
	PostPackage(ctx context.Context, pkg *lang.Package)
}
