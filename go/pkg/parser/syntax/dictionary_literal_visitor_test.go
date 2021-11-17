package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjectLiteralVisitor_VisitMapLiteral(t *testing.T) {
	const typeAttribute = `
@default({"key": "value"})
type Mailbox{}`

	parser := &Parser{}
	file, err := parser.ParseString(typeAttribute)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)

	map := make(map[string]string)
	expr.EvalStringMapLiteral(func(key string, value *lang.Expression) error {
		v, err := value.EvalStringLiteral()
		if err != nil {
			return err
		}
		map[key] = v
		return nil
	})

	assert.NotEmpty(t, map)
	assert.Equal(t, "value", map["key"])
}
