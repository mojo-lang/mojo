package compiler

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/db/sql"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type ExpressionCompiler interface {
	CompileExpression(ctx context.Context, expr *sql.Expression) (*lang.Expression, error)
}
