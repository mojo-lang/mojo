package openapi

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestSchemaPropertiesEncoder_Encode(t *testing.T) {
	schema := &Schema{
		Title: "Test",
		Properties: map[string]*ReferenceableSchema{
			"bcd": NewReferenceableSchema(&Schema{Type: Schema_TYPE_STRING}),
			"efg": NewReferenceableSchema(&Schema{Type: Schema_TYPE_STRING}),
			"abc": NewReferenceableSchema(&Schema{Type: Schema_TYPE_STRING}),
		},
	}
	const expect = `{"title":"Test","properties":{"abc":{"type":"string"},"bcd":{"type":"string"},"efg":{"type":"string"}}}`

	str, err := jsoniter.ConfigFastest.MarshalToString(schema)
	assert.NoError(t, err)
	assert.Equal(t, expect, str)
}
