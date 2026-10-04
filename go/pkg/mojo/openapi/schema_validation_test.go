package openapi

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaValidateScalarValues(t *testing.T) {
	for _, tt := range []struct {
		name           string
		typ            Schema_Type
		valid, invalid []*core.Value
	}{
		{"null", Schema_TYPE_NULL, []*core.Value{core.NewNullValue()}, []*core.Value{core.NewStringValue("null")}},
		{"boolean", Schema_TYPE_BOOLEAN, []*core.Value{core.NewBoolValue(false)}, []*core.Value{core.NewStringValue("false")}},
		{"integer", Schema_TYPE_INTEGER, []*core.Value{core.NewIntValue(0), core.NewStringValue("-42")}, []*core.Value{core.NewStringValue("4.2"), core.NewBoolValue(true)}},
		{"number", Schema_TYPE_NUMBER, []*core.Value{core.NewFloat64Value(1.5), core.NewIntValue(1), core.NewStringValue("1.5")}, []*core.Value{core.NewStringValue("invalid"), core.NewBoolValue(true)}},
		{"string", Schema_TYPE_STRING, []*core.Value{core.NewStringValue(""), core.NewBytesValue([]byte("bytes")), core.NewIntValue(1), core.NewFloat64Value(1.5)}, []*core.Value{core.NewBoolValue(true)}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			schema := &Schema{Type: tt.typ}
			for _, value := range tt.valid {
				assert.NoError(t, schema.ValidateValue(value, nil))
			}
			for _, value := range tt.invalid {
				assert.Error(t, schema.ValidateValue(value, nil))
			}
			assert.NoError(t, schema.ValidateValue(nil, nil))
		})
	}
	assert.NoError(t, (*Schema)(nil).ValidateValue(core.NewBoolValue(true), nil))
	assert.NoError(t, (*Schema)(nil).ValidateExample(nil))
	assert.NoError(t, (&Schema{}).ValidateExample(nil))
}

func TestSchemaValidateContainers(t *testing.T) {
	integer := &Schema{Type: Schema_TYPE_INTEGER}
	index := map[string]*Schema{"Integer": integer}
	ref := NewReferencedSchema(testReference(t, SchemaReferenceRoot+"Integer"))
	array := &Schema{Type: Schema_TYPE_ARRAY, Items: ref}
	assert.NoError(t, array.ValidateValue(core.NewArrayValue(core.NewIntValue(1)), index))
	assert.Error(t, array.ValidateValue(core.NewArrayValue(core.NewBoolValue(true)), index))
	assert.Error(t, array.ValidateValue(core.NewStringValue("array"), index))
	object := &Schema{Type: Schema_TYPE_OBJECT, Properties: map[string]*ReferenceableSchema{"id": ref}, Required: []string{"id"}}
	assert.NoError(t, object.ValidateValue(core.NewMapValue(map[string]*core.Value{"id": core.NewIntValue(1), "unknown": core.NewBoolValue(true)}), index))
	assert.Error(t, object.ValidateValue(core.NewMapValue(map[string]*core.Value{}), index))
	assert.NotPanics(t, func() { assert.Error(t, object.ValidateValue(core.NewMapValue(nil), index)) })
	assert.Error(t, object.ValidateValue(core.NewMapValue(map[string]*core.Value{"id": core.NewBoolValue(true)}), index))
	assert.Error(t, object.ValidateValue(core.NewIntValue(1), index))
	dictionary := &Schema{Type: Schema_TYPE_OBJECT, AdditionalProperties: ref}
	assert.NoError(t, dictionary.ValidateValue(core.NewMapValue(map[string]*core.Value{"first": core.NewIntValue(1)}), index))
	assert.Error(t, dictionary.ValidateValue(core.NewMapValue(map[string]*core.Value{"first": core.NewBoolValue(true)}), index))
}

func TestSchemaMetadataAndClone(t *testing.T) {
	stringType := NewReferenceableSchema(&Schema{Type: Schema_TYPE_STRING})
	for _, tt := range []struct {
		schema *Schema
		name   string
	}{
		{nil, ""}, {&Schema{}, ""}, {&Schema{Type: Schema_TYPE_BOOLEAN}, "boolean"},
		{&Schema{Type: Schema_TYPE_ARRAY}, "array"}, {&Schema{Type: Schema_TYPE_ARRAY, Items: stringType}, "Array<string>"},
		{&Schema{Type: Schema_TYPE_OBJECT}, "object"}, {&Schema{Type: Schema_TYPE_OBJECT, Title: "Pet"}, "Pet"},
		{&Schema{Type: Schema_TYPE_OBJECT, AdditionalProperties: stringType}, "Map<string, string>"},
		{&Schema{Title: "Named"}, "Named"}, {&Schema{OneOf: []*ReferenceableSchema{stringType, NewReferenceableSchema(&Schema{Type: Schema_TYPE_INTEGER})}}, "Union<string,integer>"},
	} {
		assert.Equal(t, tt.name, tt.schema.GetTypeName(nil))
	}
	for _, typ := range []Schema_Type{Schema_TYPE_NULL, Schema_TYPE_BOOLEAN, Schema_TYPE_INTEGER, Schema_TYPE_NUMBER, Schema_TYPE_STRING} {
		assert.True(t, (&Schema{Type: typ}).IsScalar())
	}
	assert.False(t, (&Schema{Type: Schema_TYPE_OBJECT}).IsScalar())
	assert.False(t, (*Schema)(nil).IsScalar())
	assert.True(t, (&Schema{}).IsEmpty())
	assert.False(t, (*Schema)(nil).IsEmpty())
	for _, schema := range []*Schema{{Type: Schema_TYPE_STRING}, {Enum: []*core.Value{core.NewStringValue("x")}}, {AllOf: []*ReferenceableSchema{stringType}}, {OneOf: []*ReferenceableSchema{stringType}}, {AnyOf: []*ReferenceableSchema{stringType}}, {Not: stringType}} {
		assert.False(t, schema.IsEmpty())
	}
	original := &Schema{Type: Schema_TYPE_OBJECT, Required: []string{"name"}, Properties: map[string]*ReferenceableSchema{"name": stringType}}
	assert.True(t, original.IsPropertyRequired("name"))
	assert.False(t, original.IsPropertyRequired("missing"))
	assert.False(t, (*Schema)(nil).IsPropertyRequired("name"))
	clone := original.Clone()
	require.NotNil(t, clone)
	clone.Required[0] = "changed"
	clone.Properties["name"].GetSchema().Title = "changed"
	assert.Equal(t, "name", original.Required[0])
	assert.Empty(t, original.Properties["name"].GetSchema().Title)
	assert.Nil(t, (*Schema)(nil).Clone())
	assert.ElementsMatch(t, []string{"name"}, original.FieldNames(nil))
	assert.Empty(t, (*Schema)(nil).FieldNames(nil))
	combined := &Schema{AllOf: []*ReferenceableSchema{NewReferenceableSchema(original), NewReferenceableSchema(&Schema{Properties: map[string]*ReferenceableSchema{"age": stringType}})}}
	assert.ElementsMatch(t, []string{"name", "age"}, combined.FieldNames(nil))
}
