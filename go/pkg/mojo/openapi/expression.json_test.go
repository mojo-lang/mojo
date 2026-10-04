package openapi

import (
	"testing"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpressionJSON(t *testing.T) {
	var expression Expression
	require.NoError(t, jsoniter.UnmarshalFromString(`"$request.path.id"`, &expression))
	encoded, err := jsoniter.MarshalToString(&expression)
	require.NoError(t, err)
	assert.Equal(t, `"$request.path.id"`, encoded)
	for _, input := range []string{`"invalid"`, `42`, `true`, `{}`, `[]`, `"unterminated`} {
		assert.Error(t, jsoniter.UnmarshalFromString(input, &expression), input)
		assert.Equal(t, "$request.path.id", expression.Format())
	}
	require.NoError(t, jsoniter.UnmarshalFromString(`null`, &expression))
	assert.Equal(t, Expression_TYPE_UNSPECIFIED, expression.Type)
	assert.Nil(t, expression.Source)
	codec := &ExpressionStringCodec{}
	assert.True(t, codec.IsEmpty(nil))
	assert.True(t, codec.IsEmpty(unsafe.Pointer(&expression)))
	expression.Type = Expression_TYPE_URL
	assert.False(t, codec.IsEmpty(unsafe.Pointer(&expression)))
}
