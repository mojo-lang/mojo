package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

func TestObjectLiteralVisitor_VisitMapLiteral(t *testing.T) {
	const typeAttribute = `{"key": "value"}`
	expr := parseExpression(t, typeAttribute)

	dictionary := make(map[string]string)
	err := expr.EvalStringMapLiteral(func(key string, value *lang.Expression) error {
		v, err := value.EvalStringLiteral()
		if err != nil {
			return err
		}
		dictionary[key] = v
		return nil
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, dictionary)
	assert.Equal(t, "value", dictionary["key"])
}
