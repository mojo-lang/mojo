package context

import (
	"context"
	"testing"

	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
	"github.com/stretchr/testify/assert"
)

func TestWithOpenAPI(t *testing.T) {
	ctx := Empty()
	ctx = WithOpenAPI(ctx, &openapi.OpenAPI{})
	ctx = context.WithValue(ctx, "foo", "bar")

	api := OpenAPI(ctx)
	assert.NotNil(t, api)
}

func TestWithComponents(t *testing.T) {
	ctx := Empty()
	ctx = WithOpenAPIComponents(ctx, &openapi.Components{})
	ctx = context.WithValue(ctx, "foo", "bar")

	components := OpenAPIComponents(ctx)
	assert.NotNil(t, components)
}
