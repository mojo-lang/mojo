package openapi

import (
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

func NewComponents() *Components {
	return &Components{
		Schemas:         make(map[string]*Schema),
		Responses:       nil,
		Parameters:      nil,
		Examples:        nil,
		RequestBodies:   nil,
		Headers:         nil,
		SecuritySchemes: nil,
		Links:           nil,
		Callbacks:       nil,
		PathItems:       nil,
	}
}

func (x *Components) GetSchema(url *core.Url) *Schema {
	if x == nil {
		return nil
	}

	segments := strings.Split(url.GetFragment(), "/")
	if len(segments) > 0 && len(segments[0]) == 0 {
		segments = segments[1:]
	}

	if len(segments) > 2 && segments[0] == "components" && segments[1] == "schemas" {
		object := segments[len(segments)-1]
		return x.Schemas[object]
	}
	return nil
}
