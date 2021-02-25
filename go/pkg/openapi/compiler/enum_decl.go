package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

func compileEnumDecl(ctx *Context, decl *lang.EnumDecl) (*openapi.ReferenceableSchema, error) {
	_ = ctx

	schema := &openapi.Schema{
		Type: openapi.Schema_TYPE_STRING,
	}

	for _, e := range decl.Type.Enumerators {
		if e.Name == "unspecified" {
			continue
		}

		schema.Enum = append(schema.Enum, core.NewStringValue(e.Name))
	}

	return openapi.NewReferenceableSchema(schema), nil
}
