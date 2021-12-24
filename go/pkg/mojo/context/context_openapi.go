package context

import (
	"context"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

const OpenAPIKey = "@openapi"
const ComponentsOpenAPIKey = "@openapi/Components"

func WithComponents(ctx context.Context, components *openapi.Components) context.Context {
	return WithValues(ctx, ComponentsOpenAPIKey, components)
}

func Components(ctx context.Context) *openapi.Components {
	if components, ok := ctx.Value(ComponentsOpenAPIKey).(*openapi.Components); ok {
		return components
	}
	return nil
}
