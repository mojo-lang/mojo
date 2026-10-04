package openapi

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestReferenceableJSONContracts(t *testing.T) {
	for _, tt := range []struct {
		name            string
		newValue        func() proto.Message
		inline, invalid string
	}{
		{"schema", func() proto.Message { return &ReferenceableSchema{} }, `{"type":"string","title":"Name"}`, `{"properties":42}`},
		{"parameter", func() proto.Message { return &ReferenceableParameter{} }, `{"name":"id","in":"query"}`, `{"name":[]}`},
		{"request body", func() proto.Message { return &ReferenceableRequestBody{} }, `{"required":true}`, `{"required":[]}`},
		{"response", func() proto.Message { return &ReferenceableResponse{} }, `{"description":"OK"}`, `{"content":[]}`},
		{"example", func() proto.Message { return &ReferenceableExample{} }, `{"summary":"sample","value":42}`, `{"summary":[]}`},
		{"link", func() proto.Message { return &ReferenceableLink{} }, `{"operationId":"read"}`, `{"operationId":[]}`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			value := tt.newValue()
			for _, input := range []string{tt.inline, `{"$ref":"#/components/schemas/Thing","summary":"reference","description":"details"}`, tt.inline} {
				require.NoError(t, jsoniter.UnmarshalFromString(input, value))
				encoded, err := jsoniter.MarshalToString(value)
				require.NoError(t, err)
				assert.JSONEq(t, input, encoded)
			}
			for _, input := range []string{`true`, `[]`, `42`, `"text"`, `{"$ref":42}`, `{"$ref":"http://%"}`, tt.invalid, `{"$ref":`} {
				before := proto.Clone(value)
				assert.Error(t, jsoniter.UnmarshalFromString(input, value), input)
				assert.True(t, proto.Equal(before, value), "invalid input changed the receiver: %s", input)
			}
			require.NoError(t, jsoniter.UnmarshalFromString(`null`, value))
			assert.True(t, proto.Equal(tt.newValue(), value))
			encoded, err := jsoniter.MarshalToString(value)
			require.NoError(t, err)
			assert.Equal(t, `null`, encoded)
		})
	}
}

func TestMapJSONContracts(t *testing.T) {
	for _, tt := range []struct {
		name           string
		newValue       func() proto.Message
		valid, invalid string
	}{
		{"paths", func() proto.Message { return &Paths{} }, `{"/pets":{"get":{"operationId":"read"}}}`, `{"/pets":{"get":42}}`},
		{"responses", func() proto.Message { return &Responses{} }, `{"200":{"description":"OK"},"default":{"$ref":"#/components/responses/Error"}}`, `{"200":{"content":42}}`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			value := tt.newValue()
			encoded, err := jsoniter.MarshalToString(value)
			require.NoError(t, err)
			assert.JSONEq(t, `{}`, encoded)
			require.NoError(t, jsoniter.UnmarshalFromString(tt.valid, value))
			encoded, err = jsoniter.MarshalToString(value)
			require.NoError(t, err)
			assert.JSONEq(t, tt.valid, encoded)
			for _, input := range []string{`true`, `[]`, `"text"`, tt.invalid} {
				before := proto.Clone(value)
				assert.Error(t, jsoniter.UnmarshalFromString(input, value), input)
				assert.True(t, proto.Equal(before, value))
			}
			require.NoError(t, jsoniter.UnmarshalFromString(`null`, value))
			assert.True(t, proto.Equal(tt.newValue(), value))
		})
	}
}
