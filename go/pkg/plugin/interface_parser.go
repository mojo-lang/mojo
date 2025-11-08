package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type InterfaceParser interface {
	ParseInterface(ctx context.Context, decl *lang.InterfaceDecl) error
}
