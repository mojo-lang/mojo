package decompiler

import (
	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

// Decompiler decompile the Mojo AST to Mojo AST for SQL
type Decompiler struct {
}

func (d *Decompiler) Compile(expr *lang.Expression) (*sql.Expression, error) {
	switch expression := expr.Expression.(type) {
	case *lang.Expression_StringPrefixLiteralExpr:
	case *lang.Expression_BinaryExpr:
	case *lang.Expression_FunctionCallExpr:
	case *lang.Expression_TupleExpr:
		_ = expression
	}
	return nil, nil
}
