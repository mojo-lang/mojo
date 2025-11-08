package openapi

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
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
