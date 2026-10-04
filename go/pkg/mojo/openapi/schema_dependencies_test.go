package openapi

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaDependencies(t *testing.T) {
	address := &Schema{Type: Schema_TYPE_OBJECT, Title: "Address"}
	person := &Schema{Type: Schema_TYPE_OBJECT, Title: "Person", Properties: map[string]*ReferenceableSchema{"address": NewReferenceableSchema(address)}}
	index := map[string]*Schema{"Person": person}
	personRef := NewReferencedSchema(testReference(t, SchemaReferenceRoot+"Person"))
	for _, schema := range []*Schema{
		{Type: Schema_TYPE_ARRAY, Items: personRef},
		{Type: Schema_TYPE_OBJECT, AdditionalProperties: personRef},
		{Type: Schema_TYPE_OBJECT, Properties: map[string]*ReferenceableSchema{"first": personRef, "second": personRef, "scalar": exampleSchema()}},
		{OneOf: []*ReferenceableSchema{personRef, NewReferenceableSchema(address)}},
	} {
		dependencies := schema.Dependencies(index)
		require.Len(t, dependencies, 2)
		assert.Same(t, address, dependencies[0])
		assert.Same(t, person, dependencies[1])
	}
	combined := &Schema{AllOf: []*ReferenceableSchema{personRef}}
	assert.Equal(t, []*Schema{address}, combined.Dependencies(index))
	for _, schema := range []*Schema{nil, {Type: Schema_TYPE_ARRAY}, {Type: Schema_TYPE_OBJECT, Properties: map[string]*ReferenceableSchema{"missing": NewReferencedSchema(testReference(t, SchemaReferenceRoot+"Missing"))}}} {
		assert.NotPanics(t, func() { assert.Empty(t, schema.Dependencies(nil)) })
	}
	person.Properties["self"] = personRef
	dependencies := person.Dependencies(index)
	assert.Equal(t, []*Schema{address, person}, dependencies)
}

func TestSchemasSort(t *testing.T) {
	schemas := Schemas{{Title: "z"}, {Title: "a"}, {Title: "m"}}
	sort.Sort(schemas)
	assert.Equal(t, []string{"a", "m", "z"}, []string{schemas[0].Title, schemas[1].Title, schemas[2].Title})
	assert.Zero(t, Schemas(nil).Len())
}
