package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjectLiteralVisitor_VisitObjectLiteral(t *testing.T) {
	const typeAttribute = `{"key": "value"}`

	parser := &Parser{}
	file, err := parser.ParseString(typeAttribute)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)

	object := make(map[string]string)
	expr.EvalStringMapLiteral(func(key string, value *lang.Expression) error {
		v, err := value.EvalStringLiteral()
		if err != nil {
			return err
		}
		object[key] = v
		return nil
	})

	assert.NotEmpty(t, object)
	assert.Equal(t, "value", object["key"])
}
