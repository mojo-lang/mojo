package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

type SchemaCompiler interface {
    Compile(ctx context.Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error)
}

type ReferenceCompiler struct {
}

// filter ignore types
// filter ignore fields
func FilterSchema(nominalType *lang.NominalType, schema *openapi.Schema) *openapi.ReferenceableSchema {
    if types, _ := lang.GetStringValuesAttribute(nominalType.Attributes, core.IgnoreTypesAttributeName); len(schema.OneOf) > 0 && len(types) > 0 {
        s := &openapi.Schema{
            Discriminator: schema.Discriminator,
            ExternalDocs:  schema.ExternalDocs,
            Example:       schema.Example,
            Title:         schema.Title,
            Description:   schema.Description,
            OneOf:         nil,
        }
        for _, one := range schema.OneOf {
            ignored := false
            for _, ignore := range types {
                typeName := lang.GetTypeTypeName(one.GetSchemaName())
                if typeName == lang.GetTypeTypeName(ignore) {
                    ignored = true
                    break
                }
            }
            if !ignored {
                s.OneOf = append(s.OneOf, one)
            }
        }
        return openapi.NewReferenceableSchema(s)
    }

    return nil
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
        if filtered := FilterSchema(nominalType, schema); filtered != nil {
            return filtered, nil
        }

        return openapi.NewReferencedSchema(&openapi.Reference{
            Ref: &core.Url{
                Fragment: openapi.ReferenceRoot + fullName,
            },
        }), nil
    } else if structDecl := nominalType.TypeDeclaration.GetStructDecl(); structDecl != nil && !IsPrimeType(structDecl.GetFullName()) {
        if err := CompileStructDecl(thisCtx, structDecl); err != nil {
            return nil, err
        }
        return r.Compile(ctx, nominalType)
    } else if aliasDecl := nominalType.TypeDeclaration.GetTypeAliasDecl(); aliasDecl != nil {
        if referenceSchema, err := CompileTypeAliasDecl(thisCtx, aliasDecl); err != nil {
            return nil, err
        } else {
            schema = components.Schemas[fullName]
            if schema != nil {
                if filtered := FilterSchema(nominalType, schema); filtered != nil {
                    return filtered, nil
                }
            }
            return referenceSchema, nil
        }
    }

    return nil, nil
}
