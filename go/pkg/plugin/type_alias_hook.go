package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type TypeAliasPreHook interface {
	PreTypeAlias(ctx context.Context, pkg *lang.TypeAliasDecl) error
}

type TypeAliasPostHook interface {
	PostTypeAlias(ctx context.Context, pkg *lang.TypeAliasDecl)
}
