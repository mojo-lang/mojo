package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

type WellKnowTypeCompiler struct {
}

func (w *WellKnowTypeCompiler) Compile(ctx context.Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error) {
    switch nominalType.GetFullName() {
    case core.UrlTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Type:   openapi.Schema_TYPE_STRING,
            Format: "url",
        }), nil
    case core.DurationTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Type:   openapi.Schema_TYPE_STRING,
            Format: "Duration",
        }), nil
    case core.ObjectTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Type: openapi.Schema_TYPE_OBJECT,
        }), nil
    case core.ValuesTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Type: openapi.Schema_TYPE_ARRAY,
        }), nil
    case core.ValueTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            OneOf: []*openapi.ReferenceableSchema{
                openapi.NewReferenceableSchema(&openapi.Schema{Type: openapi.Schema_TYPE_NULL}),
                openapi.NewReferenceableSchema(&openapi.Schema{Type: openapi.Schema_TYPE_BOOLEAN}),
                openapi.NewReferenceableSchema(&openapi.Schema{Type: openapi.Schema_TYPE_NUMBER}),
                openapi.NewReferenceableSchema(&openapi.Schema{Type: openapi.Schema_TYPE_STRING}),
                openapi.NewReferenceableSchema(&openapi.Schema{Type: openapi.Schema_TYPE_ARRAY}),
                openapi.NewReferenceableSchema(&openapi.Schema{Type: openapi.Schema_TYPE_OBJECT}),
            },
        }), nil
    default:
        return nil, nil
    }
}
