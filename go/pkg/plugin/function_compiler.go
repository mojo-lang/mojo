package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type FunctionCompiler interface {
	CompileFunction(ctx context.Context, decl *lang.FunctionDecl) error
}

func IsFunctionCompiler(c interface{}) bool {
	if _, ok := c.(FunctionCompiler); ok {
		return true
	}
	return false
}
