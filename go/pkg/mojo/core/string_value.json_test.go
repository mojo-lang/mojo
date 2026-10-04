package core

import (
	"testing"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringValueJSON(t *testing.T) {
	for _, text := range []string{"", "hello", "你好\n\"world\"\\"} {
		encoded, err := jsoniter.MarshalToString(&StringValue{Val: text})
		require.NoError(t, err)
		var decoded StringValue
		require.NoError(t, jsoniter.UnmarshalFromString(encoded, &decoded))
		assert.Equal(t, text, decoded.Val)
		var plain string
		require.NoError(t, jsoniter.UnmarshalFromString(encoded, &plain))
		assert.Equal(t, text, plain)
	}
}

func TestStringValueJSONEmpty(t *testing.T) {
	codec := &StringValueCodec{}
	require.NotPanics(t, func() { assert.True(t, codec.IsEmpty(nil)) })
	assert.True(t, codec.IsEmpty(unsafe.Pointer(&StringValue{})))
	assert.False(t, codec.IsEmpty(unsafe.Pointer(&StringValue{Val: "x"})))
	type payload struct {
		Value StringValue `json:"value,omitempty"`
	}
	encoded, err := jsoniter.MarshalToString(&payload{})
	require.NoError(t, err)
	assert.Equal(t, `{}`, encoded)
	encoded, err = jsoniter.MarshalToString(&payload{Value: StringValue{Val: "x"}})
	require.NoError(t, err)
	assert.JSONEq(t, `{"value":"x"}`, encoded)
}
