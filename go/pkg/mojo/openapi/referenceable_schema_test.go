package openapi

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReferenceableSchemaResolution(t *testing.T) {
	value := &Schema{Title: "inline"}
	inline := NewReferenceableSchema(value)
	assert.Same(t, value, inline.GetSchemaOf(nil))
	url, err := core.NewUrl("#" + SchemaReferenceRoot + "Thing")
	require.NoError(t, err)
	reference := &Reference{Ref: url}
	wrapper := NewReferencedSchema(reference)
	assert.Same(t, reference, wrapper.GetReference())
	assert.Nil(t, wrapper.GetSchemaOf(nil))
	assert.Nil(t, wrapper.GetSchemaOf(map[string]*Schema{}))
	assert.Same(t, value, wrapper.GetSchemaOf(map[string]*Schema{"Thing": value}))
	assert.Nil(t, (*ReferenceableSchema)(nil).GetSchemaOf(nil))
	assert.Nil(t, (&ReferenceableSchema{}).GetSchemaOf(nil))
	assert.Nil(t, NewReferencedSchema(&Reference{}).GetSchemaOf(nil))
	wrapper.SetSchema(value)
	assert.Nil(t, wrapper.GetReference())
	assert.Same(t, value, wrapper.GetSchema())
	wrapper.SetReference(reference)
	assert.Nil(t, wrapper.GetSchema())
	assert.Same(t, reference, wrapper.GetReference())
	assert.NotPanics(t, func() {
		(*ReferenceableSchema)(nil).SetSchema(value)
		(*ReferenceableSchema)(nil).SetReference(reference)
	})
	assert.Nil(t, NewReferenceableSchema(nil).GetSchema())
	assert.Nil(t, NewReferencedSchema(nil).GetReference())
}

func TestReferenceableSchemaMetadata(t *testing.T) {
	schema := &Schema{Title: "Pet", Type: Schema_TYPE_STRING, Format: "Name", Description: &CachedDocument{Cache: "First paragraph.\n\nSecond paragraph."}}
	inline := NewReferenceableSchema(schema)
	assert.Equal(t, "Pet", inline.GetSchemaName())
	assert.Equal(t, "Pet", inline.GetTitle(nil))
	assert.Equal(t, "string", inline.GetTypeName(nil))
	assert.Equal(t, "Name", inline.GetFormat(nil))
	assert.Contains(t, inline.GetSummary(nil), "First paragraph.")
	assert.NotContains(t, inline.GetSummary(nil), "Second paragraph.")
	assert.Same(t, schema.Description.GetDocument(), inline.GetDescription(nil))
	inline.SetSchema(nil)
	inline.SetReference(nil)
	assert.Same(t, schema, inline.GetSchema())
	ref := &ReferenceableSchema{}
	require.NoError(t, ref.SetReferenceUrl("#/components/schemas/Pet"))
	assert.Equal(t, "Pet", ref.GetSchemaName())
	assert.Equal(t, "#/components/schemas/Pet", ref.GetReferenceUrl().Format())
	assert.Error(t, ref.SetReferenceUrl("http://%"))
	assert.Equal(t, "Pet", ref.GetSchemaName())
	index := map[string]*Schema{"Pet": schema}
	assert.Equal(t, "Pet", ref.GetTitle(index))
	assert.Same(t, schema.Description.GetDocument(), ref.GetDescription(index))
	ref.GetReference().Summary = "override"
	assert.Equal(t, "override", ref.GetSummary(index))
	description := &CachedDocument{Cache: "new description"}
	ref.SetDescription(description)
	assert.Same(t, description.GetDocument(), ref.GetDescription(index))
	inline.SetDescription(description)
	assert.Same(t, description, schema.Description)
	var empty *ReferenceableSchema
	assert.Empty(t, empty.GetSchemaName())
	assert.Empty(t, empty.GetTitle(nil))
	assert.Empty(t, empty.GetTypeName(nil))
	assert.Empty(t, empty.GetFormat(nil))
	assert.Empty(t, empty.GetSummary(nil))
	assert.Nil(t, empty.GetDescription(nil))
	assert.Nil(t, empty.GetReferenceUrl())
	assert.NoError(t, empty.SetReferenceUrl("ignored"))
	assert.NotPanics(t, func() { empty.SetDescription(description) })
}
