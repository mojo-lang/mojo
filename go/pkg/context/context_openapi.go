package context

import (
	"context"

	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

const OpenAPIKey = "@openapi"
const ComponentsOpenAPIKey = "@openapi/Components"

func WithOpenAPI(ctx context.Context, api *openapi.OpenAPI) context.Context {
	return WithValues(ctx, OpenAPIKey, api)
}

func OpenAPI(ctx context.Context) *openapi.OpenAPI {
	if api, ok := ctx.Value(OpenAPIKey).(*openapi.OpenAPI); ok {
		return api
	}
	return nil
}

func WithOpenAPIComponents(ctx context.Context, components *openapi.Components) context.Context {
	return WithValues(ctx, ComponentsOpenAPIKey, components)
}

func OpenAPIComponents(ctx context.Context) *openapi.Components {
	if components, ok := ctx.Value(ComponentsOpenAPIKey).(*openapi.Components); ok {
		return components
	}
	return nil
}
