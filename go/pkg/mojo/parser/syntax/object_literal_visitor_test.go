package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

func TestObjectLiteralVisitor_VisitObjectLiteral(t *testing.T) {
	const typeAttribute = `{"key": "value"}`
	expr := parseExpression(t, typeAttribute)

	object := make(map[string]string)
	err := expr.EvalStringMapLiteral(func(key string, value *lang.Expression) error {
		v, err := value.EvalStringLiteral()
		if err != nil {
			return err
		}
		object[key] = v
		return nil
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, object)
	assert.Equal(t, "value", object["key"])
}
