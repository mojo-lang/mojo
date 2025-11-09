package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type EnumParser interface {
	ParseEnum(ctx context.Context, decl *lang.EnumDecl) error
}
