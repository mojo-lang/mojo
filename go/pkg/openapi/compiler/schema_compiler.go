package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

type SchemaCompiler interface {
	Compile(ctx context.Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error)
}

type ReferenceCompiler struct {
}

func (r *ReferenceCompiler) Compile(ctx context.Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error) {
	components := context.Components(ctx)
	thisCtx := ctx
	if components == nil {
		components = openapi.NewComponents()
		thisCtx = context.WithComponents(ctx, components)
	}

	fullName := nominalType.GetFullName()
	schema := components.Schemas[fullName]
	if schema != nil {
		return openapi.NewReferencedSchema(&openapi.Reference{
			Ref: &core.Url{
				Fragment: openapi.ReferenceRoot + fullName,
			},
		}), nil
	} else if structDecl := nominalType.TypeDeclaration.GetStructDecl(); structDecl != nil {
		if err := CompileStructDecl(thisCtx, structDecl); err != nil {
			return nil, err
		}
		return r.Compile(ctx, nominalType)
	} else if aliasDecl := nominalType.TypeDeclaration.GetTypeAliasDecl(); aliasDecl != nil {
		return CompileTypeAliasDecl(thisCtx, aliasDecl)
	}

	return nil, nil
}
