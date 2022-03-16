package hook

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

type StructPreHook interface {
    PreStruct(ctx context.Context, pkg *lang.StructDecl) error
}

type StructPostHook interface {
    PostStruct(ctx context.Context, pkg *lang.StructDecl)
}
