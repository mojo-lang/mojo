package openapi

import (
	"io/ioutil"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
)

func TestSchema_Parse(t *testing.T) {
	s := `{
			"maxLength": 32,
			"minLength": 1,
			"pattern": "^[\\u4E00-\\u9FA5A-Za-z0-9_.]{1,32}$",
			"type": "string",
			"description": "名称",
			"format": "Name"
		}`

	schema := &Schema{}
	err := jsoniter.Unmarshal([]byte(s), schema)

	assert.NoError(t, err)
	assert.Equal(t, Schema_TYPE_STRING, schema.Type)
}

func TestSchema_ValidateValue(t *testing.T) {
	schema := &Schema{Type: Schema_TYPE_INTEGER}
	intv := core.NewInt32Value(100)
	strv := core.NewStringValue("foo")

	assert.NoError(t, schema.ValidateValue(intv, nil))
	assert.Error(t, schema.ValidateValue(strv, nil))
}

func TestSchema_ValidateExample(t *testing.T) {
	iSchema := &Schema{Type: Schema_TYPE_INTEGER, Example: core.NewInt32Value(100)}
	sSchema := &Schema{Type: Schema_TYPE_INTEGER, Example: core.NewStringValue("foo")}

	assert.NoError(t, iSchema.ValidateExample(nil))
	assert.Error(t, sSchema.ValidateExample(nil))
}

func TestSchema_GenerateExample(t *testing.T) {
	schema := &Schema{Type: Schema_TYPE_INTEGER}
	v := schema.GenerateExample(nil)
	assert.NotNil(t, v)
	assert.Equal(t, core.ValueKind_VALUE_KIND_INTEGER, v.GetKind())

	schema = &Schema{Type: Schema_TYPE_NUMBER}
	v = schema.GenerateExample(nil)
	assert.NotNil(t, v)
	assert.Equal(t, core.ValueKind_VALUE_KIND_NUMBER, v.GetKind())

	schema = &Schema{Type: Schema_TYPE_STRING}
	v = schema.GenerateExample(nil)
	assert.NotNil(t, v)
	assert.Equal(t, core.ValueKind_VALUE_KIND_STRING, v.GetKind())

	schema = &Schema{Type: Schema_TYPE_OBJECT, Properties: map[string]*ReferenceableSchema{
		"foo": NewReferenceableSchema(&Schema{Type: Schema_TYPE_INTEGER}),
		"bar": NewReferenceableSchema(&Schema{Type: Schema_TYPE_STRING}),
	}}
	v = schema.GenerateExample(nil)
	assert.NotNil(t, v)
	assert.Equal(t, core.ValueKind_VALUE_KIND_OBJECT, v.GetKind())
	assert.Equal(t, core.ValueKind_VALUE_KIND_INTEGER, v.GetObject().GetValue("foo").GetKind())
	assert.Equal(t, core.ValueKind_VALUE_KIND_STRING, v.GetObject().GetValue("bar").GetKind())
}

func TestSchema_SupplementExample(t *testing.T) {
	schema := &Schema{Type: Schema_TYPE_OBJECT, Properties: map[string]*ReferenceableSchema{
		"foo": NewReferenceableSchema(&Schema{Type: Schema_TYPE_INTEGER}),
		"bar": NewReferenceableSchema(&Schema{Type: Schema_TYPE_STRING, Example: core.NewStringValue("baz")}),
	}}
	v := schema.SupplementExample(nil)
	assert.NotNil(t, v)
	assert.Equal(t, core.ValueKind_VALUE_KIND_OBJECT, v.GetKind())
	assert.Equal(t, core.ValueKind_VALUE_KIND_INTEGER, v.GetObject().GetValue("foo").GetKind())
	assert.Equal(t, core.ValueKind_VALUE_KIND_STRING, v.GetObject().GetValue("bar").GetKind())
	assert.Equal(t, "baz", v.GetObject().GetValue("bar").GetString())
}

func TestSchema_SupplementExample2(t *testing.T) {
	schema := &Schema{Type: Schema_TYPE_OBJECT, Properties: map[string]*ReferenceableSchema{
		"foo": NewReferenceableSchema(&Schema{Type: Schema_TYPE_INTEGER, Format: "int16", Maximum: core.NewInt16Value(100), Minimum: core.NewInt16Value(0), MultipleOf: 5}),
		"bar": NewReferenceableSchema(&Schema{Type: Schema_TYPE_INTEGER, Format: "uint32", Maximum: core.NewUInt32Value(1000)}),
		"baz": NewReferenceableSchema(&Schema{Type: Schema_TYPE_NUMBER, Maximum: core.NewFloat64Value(1000.0)}),
	}}
	v := schema.SupplementExample(nil)
	assert.NotNil(t, v)
	assert.Equal(t, core.ValueKind_VALUE_KIND_OBJECT, v.GetKind())
	assert.Equal(t, core.ValueKind_VALUE_KIND_INTEGER, v.GetObject().GetValue("foo").GetKind())

	value1 := v.GetObject().GetValue("foo").GetInt()
	assert.True(t, value1 >= 0 && value1 <= 100)

	value2 := v.GetObject().GetValue("bar").GetUint()
	assert.True(t, value2 <= 1000)

	value3 := v.GetObject().GetValue("baz").GetFloat64()
	assert.True(t, value3 <= 1000.0)
}

func TestSchema_GetMaximum(t *testing.T) {
	api := &OpenAPI{}
	content, err := ioutil.ReadFile("./testdata/supplement_example_test.json")
	if assert.NoError(t, err) {
		err = jsoniter.Unmarshal(content, api)
		if assert.NoError(t, err) {
			schema := api.Components.Schemas["EveArchiveFileInfoFree"]
			s := schema.Properties["mainContainFileNums"].GetSchema()
			v := s.Maximum.GetUint64()
			assert.Equal(t, uint64(18446744073709551615), v)
		}
	}
}

func TestSchema_SupplementExample_Array(t *testing.T) {
	api := &OpenAPI{}
	content, err := ioutil.ReadFile("./testdata/supplement_example_test.json")
	if assert.NoError(t, err) {
		err = jsoniter.Unmarshal(content, api)
		if assert.NoError(t, err) {
			api.SupplementExample()
			schema := api.Components.Schemas["JqGridRequestBean"]
			s := schema.Properties["classyModeIds"].GetSchema()
			vs := s.SupplementExample(api.Components.Schemas).GetValues()
			assert.Equal(t, 1, len(vs))
		}
	}
}

func TestFakeString(t *testing.T) {
	str := FakeString(&Schema{
		Type:   Schema_TYPE_STRING,
		Format: "Date",
	})

	assert.NotEmpty(t, str)
}
