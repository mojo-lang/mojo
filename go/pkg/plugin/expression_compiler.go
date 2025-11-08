package plugin

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type ExpressionCompiler interface {
	CompileExpression(ctx context.Context, decl *lang.Expression) error
}

func IsExpressionCompiler(c interface{}) bool {
	if _, ok := c.(ExpressionCompiler); ok {
		return true
	}
	return false
}
