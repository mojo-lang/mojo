package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type FunctionPreHook interface {
	PreFunction(ctx context.Context, pkg *lang.FunctionDecl) error
}

type FunctionPostHook interface {
	PostFunction(ctx context.Context, pkg *lang.FunctionDecl)
}
