package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type TypeAliasParser interface {
	ParseTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error
}
