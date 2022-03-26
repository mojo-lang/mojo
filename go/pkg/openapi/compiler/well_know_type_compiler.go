package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/db/go/pkg/mojo/db"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
    "strings"
)

type WellKnowTypeCompiler struct {
}

func (w *WellKnowTypeCompiler) Compile(ctx context.Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error) {
    switch nominalType.GetFullName() {
    case core.Int32ValueTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.Int32TypeName,
            Type:   openapi.Schema_TYPE_NUMBER,
            Format: strings.ToLower(core.Int32TypeName),
        }), nil
    case core.UInt32ValueTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.UInt32TypeName,
            Type:   openapi.Schema_TYPE_NUMBER,
            Format: strings.ToLower(core.UInt32TypeName),
        }), nil
    case core.Int64ValueTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.Int64TypeName,
            Type:   openapi.Schema_TYPE_NUMBER,
            Format: strings.ToLower(core.Int64TypeName),
        }), nil
    case core.UInt64ValueTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.UInt64TypeName,
            Type:   openapi.Schema_TYPE_NUMBER,
            Format: strings.ToLower(core.UInt64TypeName),
        }), nil
    case core.Float32ValueTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.Float32TypeName,
            Type:   openapi.Schema_TYPE_NUMBER,
            Format: strings.ToLower(core.Float32TypeName),
        }), nil
    case core.Float64ValueTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.Float64TypeName,
            Type:   openapi.Schema_TYPE_NUMBER,
            Format: strings.ToLower(core.Float64TypeName),
        }), nil
    case core.UrlTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.UrlTypeName,
            Type:   openapi.Schema_TYPE_STRING,
            Format: core.UrlTypeName,
        }), nil
    case core.DurationTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.DurationTypeName,
            Type:   openapi.Schema_TYPE_STRING,
            Format: core.DurationTypeName,
        }), nil
    case core.TimestampTypeFullName, core.DateTimeTypeFullName, db.DeleteTimeTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.DateTimeTypeName,
            Type:   openapi.Schema_TYPE_STRING,
            Format: core.DateTimeTypeName,
        }), nil
    case core.DateTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.DateTypeName,
            Type:   openapi.Schema_TYPE_STRING,
            Format: core.DateTypeName,
        }), nil
    case core.ColorTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.ColorTypeName,
            Type:   openapi.Schema_TYPE_STRING,
            Format: core.ColorTypeName,
        }), nil
    case core.EmailAddressTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.EmailAddressTypeName,
            Type:   openapi.Schema_TYPE_STRING,
            Format: core.EmailAddressTypeName,
        }), nil
    case core.ObjectTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title: "object",
            Type:  openapi.Schema_TYPE_OBJECT,
        }), nil
    case core.PlatformTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.PlatformTypeName,
            Type:   openapi.Schema_TYPE_STRING,
            Format: core.PlatformTypeName,
        }), nil
    case core.ValuesTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title: core.ValuesTypeName,
            Type:  openapi.Schema_TYPE_ARRAY,
        }), nil
    case core.ValueTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title: core.ValueTypeName,
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
