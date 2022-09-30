package format

import (
	"bytes"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getExpression(file *lang.SourceFile) *lang.Expression {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		return statement.GetExpression()
	}
	return nil
}

func parseExpression(t *testing.T, str string) *lang.Expression {
	parser := &syntax.Parser{}
	file, err := parser.ParseString(str)
	assert.NoError(t, err)

	expr := getExpression(file)
	assert.NotNil(t, expr)
	return expr
}

func TestPrinter_PrintExpression(t *testing.T) {
	const expr = `
"test_data"
`
	const expect = `"test_data"`

	decl := parseExpression(t, expr)
	buffer := bytes.NewBuffer(nil)
	NewPrinter(Config{}, buffer).PrintExpression(context.Empty(), decl)
	assert.Equal(t, expect, buffer.String())
}
