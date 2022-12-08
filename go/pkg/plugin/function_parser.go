package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type FunctionParser interface {
	ParseFunction(ctx context.Context, decl *lang.FunctionDecl) error
}
