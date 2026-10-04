package openapi

import (
	"testing"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestReferenceJSONRoundTrip(t *testing.T) {
	const input = `{"$ref":"#/components/schemas/Thing","summary":"summary","description":"description"}`
	var reference Reference
	require.NoError(t, jsoniter.UnmarshalFromString(input, &reference))
	assert.Equal(t, "#/components/schemas/Thing", reference.GetRef().Format())
	assert.Equal(t, "summary", reference.Summary)
	assert.Equal(t, "description", reference.Description.GetStringDocument())
	encoded, err := jsoniter.MarshalToString(&reference)
	require.NoError(t, err)
	assert.JSONEq(t, input, encoded)
	for _, input := range []string{`{}`, `[]`, `42`, `{"$ref":false}`, `{"$ref":"http://%"}`, `{"$ref":"#/ok","summary":42}`} {
		before := proto.Clone(&reference)
		assert.Error(t, jsoniter.UnmarshalFromString(input, &reference), input)
		assert.True(t, proto.Equal(before, &reference))
	}
	require.NoError(t, jsoniter.UnmarshalFromString(`null`, &reference))
	assert.Nil(t, reference.Ref)
	encoded, err = jsoniter.MarshalToString(&reference)
	require.NoError(t, err)
	assert.Equal(t, `null`, encoded)
	codec := &ReferenceCodec{}
	assert.True(t, codec.IsEmpty(nil))
	assert.True(t, codec.IsEmpty(unsafe.Pointer(&reference)))
}
