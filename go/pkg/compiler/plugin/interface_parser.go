package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type InterfaceParser interface {
	ParseInterface(ctx context.Context, decl *lang.InterfaceDecl) error
}
