package plugin

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type FunctionParser interface {
	ParseFunction(ctx context.Context, decl *lang.FunctionDecl) error
}
