package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type TypeAliasParser interface {
	ParseTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error
}
