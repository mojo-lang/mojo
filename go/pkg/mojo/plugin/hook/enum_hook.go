package hook

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type EnumPreHook interface {
	PreEnum(ctx context.Context, pkg *lang.EnumDecl) error
}

type EnumPostHook interface {
	PostEnum(ctx context.Context, pkg *lang.EnumDecl)
}

