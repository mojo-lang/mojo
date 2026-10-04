package openapi

import (
	"testing"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestCachedDocumentJSON(t *testing.T) {
	cached := &CachedDocument{Cache: "old"}
	require.NotNil(t, cached.GetDocument())
	require.NoError(t, jsoniter.UnmarshalFromString(`"new"`, cached))
	assert.Nil(t, cached.Val, "new Markdown must invalidate the parsed document")
	assert.Equal(t, "new", cached.GetStringDocument())
	encoded, err := jsoniter.MarshalToString(cached)
	require.NoError(t, err)
	assert.Equal(t, `"new"`, encoded)
	for _, input := range []string{`true`, `[]`, `42`, `{"blocks":42}`, `"unterminated`} {
		before := proto.Clone(cached)
		assert.Error(t, jsoniter.UnmarshalFromString(input, cached), input)
		assert.True(t, proto.Equal(before, cached))
	}
	require.NoError(t, jsoniter.UnmarshalFromString(`{}`, cached))
	require.NotNil(t, cached.Val)
	assert.Empty(t, cached.Cache)
	require.NoError(t, jsoniter.UnmarshalFromString(`null`, cached))
	assert.Nil(t, cached.Val)
	assert.Empty(t, cached.Cache)
	codec := &CachedDocumentCodec{}
	assert.True(t, codec.IsEmpty(nil))
	assert.True(t, codec.IsEmpty(unsafe.Pointer(cached)))
}
