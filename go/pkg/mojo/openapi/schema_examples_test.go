package openapi

import (
	"testing"
	"time"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaExampleScalarFormats(t *testing.T) {
	for _, generate := range []struct {
		name string
		call func(*Schema) *core.Value
	}{
		{"generate", func(s *Schema) *core.Value { return s.GenerateExample(nil) }},
		{"supplement", func(s *Schema) *core.Value { return s.SupplementExample(nil) }},
	} {
		t.Run(generate.name, func(t *testing.T) {
			assert.Nil(t, generate.call(nil))
			assert.Equal(t, core.ValueKind_VALUE_KIND_BOOLEAN, generate.call(&Schema{Type: Schema_TYPE_BOOLEAN}).GetKind())
			for _, format := range []string{"int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64", ""} {
				value := generate.call(&Schema{Type: Schema_TYPE_INTEGER, Format: format, Enum: []*core.Value{core.NewIntValue(12)}})
				require.NotNil(t, value)
				assert.Equal(t, int64(12), value.GetInt64(), format)
			}
			for _, format := range []string{"float32", "float64", ""} {
				value := generate.call(&Schema{Type: Schema_TYPE_NUMBER, Format: format, Enum: []*core.Value{core.NewFloat64Value(1.25)}})
				assert.Equal(t, 1.25, value.GetFloat64(), format)
			}
			before := time.Now().Unix()
			timestamp := generate.call(&Schema{Type: Schema_TYPE_INTEGER, Format: "timestamp"}).GetInt64()
			assert.GreaterOrEqual(t, timestamp, before)
			assert.LessOrEqual(t, timestamp, time.Now().Unix())
		})
	}
}

func TestSchemaExampleContainerSizes(t *testing.T) {
	for _, tt := range []struct {
		name     string
		readOnly bool
		min, max uint64
		count    int
	}{
		{"default", false, 0, 0, 3}, {"read only", true, 0, 0, 1}, {"maximum", false, 0, 2, 2}, {"minimum", false, 4, 0, 4}, {"equal bounds", false, 2, 2, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			array := &Schema{Type: Schema_TYPE_ARRAY, Items: exampleSchema(), ReadOnly: tt.readOnly, MinItems: tt.min, MaxItems: tt.max}
			dictionary := &Schema{Type: Schema_TYPE_OBJECT, AdditionalProperties: exampleSchema(), ReadOnly: tt.readOnly, MinProperties: tt.min, MaxProperties: tt.max}
			for _, value := range []*core.Value{array.GenerateExample(nil), array.SupplementExample(nil)} {
				require.Len(t, value.GetValues(), tt.count)
				for _, element := range value.GetValues() {
					assert.Equal(t, "generated", element.GetString())
				}
			}
			for _, value := range []*core.Value{dictionary.GenerateExample(nil), dictionary.SupplementExample(nil)} {
				require.Len(t, value.GetObject().Vals, tt.count)
				for _, element := range value.GetObject().Vals {
					assert.Equal(t, "generated", element.GetString())
				}
			}
		})
	}
}

func TestSchemaExampleRecursionAndExistingValues(t *testing.T) {
	leaf := exampleSchema().GetSchema()
	existing := core.NewStringValue("existing")
	leaf.Example = existing
	assert.Same(t, existing, leaf.SupplementExample(nil))
	assert.Equal(t, "generated", leaf.GenerateExample(nil).GetString())
	node := &Schema{Type: Schema_TYPE_OBJECT, Title: "Node", Properties: map[string]*ReferenceableSchema{"name": NewReferenceableSchema(leaf)}}
	node.Properties["self"] = NewReferencedSchema(testReference(t, SchemaReferenceRoot+"Node"))
	node.Properties["missing"] = NewReferencedSchema(testReference(t, SchemaReferenceRoot+"Missing"))
	index := map[string]*Schema{"Node": node}
	generated := node.GenerateExample(index).GetObject()
	require.NotNil(t, generated)
	assert.Len(t, generated.Vals, 1)
	assert.Equal(t, "generated", generated.GetString("name"))
	supplemented := node.SupplementExample(index).GetObject()
	assert.Len(t, supplemented.Vals, 1)
	assert.Equal(t, "existing", supplemented.GetString("name"))
}
