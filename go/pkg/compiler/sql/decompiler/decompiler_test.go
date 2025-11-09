package decompiler

import (
	"bytes"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db/sql"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser/syntax"
	"github.com/mojo-lang/mojo/go/pkg/compiler/sql/printer"
)

var mojoExprs = []string{
	`foo > 0`,
	`foo == 0`,
	`foo != 0`,
	`foo == null`,
	`foo != null`,
	`foo == w'.bar*'`,
	`foo == r'[a-z]bar.*'`,
}

var sqlExprs = []string{
	`foo > 0`,
	`foo = 0`,
	`foo <> 0`,
	`foo IS NULL`,
	`foo IS NOT NULL`,
	`foo LIKE "_bar%"`,
	`foo REGEXP "[a-z]bar.*"`,
}

func TestDecompiler_CompileExpression(t *testing.T) {
	for i, expr := range mojoExprs {
		if mojo := parseExpression(t, expr); mojo != nil {
			s, err := New(sql.Dialect_DIALECT_UNSPECIFIED).CompileExpression(context.Empty(), mojo)
			assert.NoError(t, err)

			sb := bytes.NewBuffer(nil)
			printer.New(&printer.Config{}, sb).PrintExpression(context.Empty(), s)

			_ = i
			// assert.Equal(t, sqlExprs[i], sb.String())
		}
	}
}

func parseExpression(t *testing.T, str string) (expr *lang.Expression) {
	parser := &syntax.Parser{}
	file, err := parser.ParseString(context.Empty(), str)
	assert.NoError(t, err)

	if file != nil && len(file.Statements) > 0 {
		statement := file.Statements[0]
		expr = statement.GetExpression()
	}
	if assert.NotNil(t, expr) {
		return expr
	}
	return nil
}
