package core

import (
	"testing"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBytesValueJSON(t *testing.T) {
	for _, tt := range []struct {
		value []byte
		json  string
	}{
		{nil, `""`}, {[]byte{}, `""`}, {[]byte{0, 255, 16}, `"AP8Q"`},
	} {
		encoded, err := jsoniter.MarshalToString(&BytesValue{Val: tt.value})
		require.NoError(t, err)
		assert.Equal(t, tt.json, encoded)
		var decoded BytesValue
		require.NoError(t, jsoniter.UnmarshalFromString(encoded, &decoded))
		assert.Equal(t, string(tt.value), string(decoded.Val))
	}
	var value BytesValue
	assert.Error(t, jsoniter.UnmarshalFromString(`"%%%"`, &value))
}

func TestBytesValueJSONEmpty(t *testing.T) {
	codec := &BytesValueCodec{}
	require.NotPanics(t, func() { assert.True(t, codec.IsEmpty(nil)) })
	assert.True(t, codec.IsEmpty(unsafe.Pointer(&BytesValue{})))
	assert.False(t, codec.IsEmpty(unsafe.Pointer(&BytesValue{Val: []byte{1}})))
	type payload struct {
		Value BytesValue `json:"value,omitempty"`
	}
	encoded, err := jsoniter.MarshalToString(&payload{})
	require.NoError(t, err)
	assert.Equal(t, `{}`, encoded)
	encoded, err = jsoniter.MarshalToString(&payload{Value: BytesValue{Val: []byte{1}}})
	require.NoError(t, err)
	assert.JSONEq(t, `{"value":"AQ=="}`, encoded)
}
