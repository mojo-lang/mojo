package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

type SchemaCompiler interface {
	Compile(ctx *Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error)
}

type ReferenceCompiler struct {
}

func (r *ReferenceCompiler) Compile(ctx *Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error) {
	if ctx.Components != nil && ctx.Components.Schemas != nil {
		fullName := nominalType.GetFullName()
		schema := ctx.Components.Schemas[fullName]
		if schema != nil {
			return openapi.NewReferencedSchema(&openapi.Reference{
				Ref: &core.Url{
					Fragment: openapi.ReferenceRoot + fullName,
				},
			}), nil
		} else if structDecl := nominalType.TypeDeclaration.GetStructDecl(); structDecl != nil {
			if err := CompileStructDecl(ctx, structDecl); err != nil {
				return nil, err
			}

			return r.Compile(ctx, nominalType)
		}
	}
	return nil, nil
}
