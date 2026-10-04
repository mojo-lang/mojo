package openapi

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComponents_GetSchema(t *testing.T) {
	c := NewComponents()
	c.Schemas["foo.bar"] = &Schema{Title: "Foo"}
	url, _ := core.NewUrl("#/components/schemas/foo.bar")
	schema := c.GetSchema(url)

	if assert.NotNil(t, schema) {
		assert.Equal(t, "Foo", schema.Title)
	}
}

func TestComponentsMissingSchema(t *testing.T) {
	components := NewComponents()
	components.Schemas["Pet"] = &Schema{Title: "Pet"}
	for _, input := range []string{"#/components/schemas/Missing", "#/components/parameters/Pet", "#/schemas/Pet", ""} {
		url, err := core.NewUrl(input)
		assert.NoError(t, err)
		assert.Nil(t, components.GetSchema(url), input)
	}
	assert.Nil(t, components.GetSchema(nil))
	assert.Nil(t, (*Components)(nil).GetSchema(nil))
	assert.Empty(t, NewComponents().Schemas)
}
