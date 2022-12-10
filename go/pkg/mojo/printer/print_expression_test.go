package printer

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
)

func getExpression(file *lang.SourceFile) *lang.Expression {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		return statement.GetExpression()
	}
	return nil
}

func parseExpression(t *testing.T, str string) *lang.Expression {
	file, err := syntax.New(nil).ParseString(context.Empty(), str)
	assert.NoError(t, err)

	expr := getExpression(file)
	assert.NotNil(t, expr)
	return expr
}

func TestPrinter_PrintExpression(t *testing.T) {
	const expr = `
"testdata"
`
	const expect = `"testdata"`

	decl := parseExpression(t, expr)
	p := New(&Config{}).PrintExpression(context.Empty(), decl)
	assert.Equal(t, expect, p.Buffer.String())
}
