package compiler

import (
	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type ExpressionCompiler interface {
	CompileExpression(ctx context.Context, expr *sql.Expression) (*lang.Expression, error)
}
