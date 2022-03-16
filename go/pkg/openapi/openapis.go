package openapi

import "github.com/mojo-lang/openapi/go/pkg/mojo/openapi"

type OpenAPIs struct {
    APIs       map[string]*openapi.OpenAPI
    Components *openapi.Components
}
