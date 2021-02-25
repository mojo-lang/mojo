package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type PackageCompiler interface {
	Compile(ctx *context.Context, pkg *lang.Package) error
}
