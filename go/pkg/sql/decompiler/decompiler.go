package decompiler

import (
	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"

	"github.com/mojo-lang/mojo/go/pkg/sql/decompiler/ansi"
)

// Decompiler decompile the Mojo AST to Mojo AST for SQL
type Decompiler struct {
	ExpressionDecompiler
}

func New(dialect sql.Dialect) *Decompiler {
	c := &Decompiler{}
	switch dialect {
	case sql.Dialect_DIALECT_SQLITE:
	default:
		c.ExpressionDecompiler = &ansi.ExpressionDecompiler{}
	}

	return c
}
