package hook

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

type InterfacePreHook interface {
	PreInterface(ctx context.Context, pkg *lang.InterfaceDecl) error
}

type InterfacePostHook interface {
	PostInterface(ctx context.Context, pkg *lang.InterfaceDecl)
}
