package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type EnumParser interface {
	ParseEnum(ctx context.Context, decl *lang.EnumDecl) error
}
