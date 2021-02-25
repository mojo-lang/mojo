package compiler

import (
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

type Context struct {
	*context.Context
	Components *openapi.Components
}
