package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjectLiteralVisitor_VisitObjectLiteral (t *testing.T) {
	const typeAttribute = `
@default({"key": "value"})
type Mailbox{}`

	parser := &Parser{}
	file, err := parser.ParseString(typeAttribute)

	assert.NoError(t, err)
	expr := getExpression(file)
	assert.NotNil(t, expr)

	dictionary := make(map[string]string)
	expr.EvalStringDictionaryLiteral(func(key string, value *lang.Expression) error {
		v, err := value.EvalStringLiteral()
		if err != nil {
			return err
		}
		dictionary[key] = v
		return nil
	})

	assert.NotEmpty(t, dictionary)
	assert.Equal(t, "value", dictionary["key"])
}
