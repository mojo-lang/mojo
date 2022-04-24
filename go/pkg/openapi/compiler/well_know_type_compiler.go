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
    anySchema := openapi.NewReferenceableSchema(&openapi.Schema{
        Title: core.AnyTypeName,
        Type:  openapi.Schema_TYPE_OBJECT,
    })
    objectSchema := openapi.NewReferenceableSchema(&openapi.Schema{
        Title: core.ObjectTypeName,
        Type:  openapi.Schema_TYPE_OBJECT,
    })
    valuesSchema := openapi.NewReferenceableSchema(&openapi.Schema{
        Title: core.ValuesTypeName,
        Type:  openapi.Schema_TYPE_ARRAY,
    })

    switch nominalType.GetFullName() {
    case core.AnyTypeFullName:
        return anySchema, nil
    case core.BoolValueTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title: core.BoolTypeName,
            Type:  openapi.Schema_TYPE_BOOLEAN,
        }), nil
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
    case core.ChecksumTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:       core.ChecksumTypeName,
            Type:        openapi.Schema_TYPE_STRING,
            Format:      core.ChecksumTypeName,
            Description: &openapi.CachedDocument{Cache: "the format is: `{algorithm} {value}` and the algorithm should be one of the `MD5`, `SHA1`, `SHA256` and `SHA512`"},
        }), nil
    case core.EmailAddressTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.EmailAddressTypeName,
            Type:   openapi.Schema_TYPE_STRING,
            Format: core.EmailAddressTypeName,
        }), nil
    case core.FieldMaskTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:       core.FieldMaskTypeName,
            Type:        openapi.Schema_TYPE_STRING,
            Format:      core.FieldMaskTypeName,
            Description: &openapi.CachedDocument{Cache: "field mask support google protobuf format like : 'foo.bar,boo.baz', also support like this: 'foo{bar,baz}'."},
        }), nil
    case core.ObjectTypeFullName:
        return objectSchema, nil
    case core.PlatformTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title:  core.PlatformTypeName,
            Type:   openapi.Schema_TYPE_STRING,
            Format: core.PlatformTypeName,
        }), nil
    case core.ValuesTypeFullName:
        return valuesSchema, nil
    case core.ValueTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title: core.ValueTypeName,
            OneOf: []*openapi.ReferenceableSchema{
                openapi.NewReferenceableSchema(&openapi.Schema{Type: openapi.Schema_TYPE_NULL}),
                openapi.NewReferenceableSchema(&openapi.Schema{Type: openapi.Schema_TYPE_BOOLEAN}),
                openapi.NewReferenceableSchema(&openapi.Schema{Type: openapi.Schema_TYPE_STRING}),
                openapi.NewReferenceableSchema(&openapi.Schema{
                    Title:       core.BytesTypeFullName,
                    Type:        openapi.Schema_TYPE_STRING,
                    Format:      core.BytesTypeName,
                    Description: &openapi.CachedDocument{Cache: "the format is: `b64.{base64 encoded bytes}`"},
                }),
                openapi.NewReferenceableSchema(int64Schema.Clone()),
                openapi.NewReferenceableSchema(float64Schema.Clone()),
                valuesSchema,
                objectSchema,
            },
        }), nil
    case core.ErrorTypeFullName:
        return openapi.NewReferenceableSchema(&openapi.Schema{
            Title: core.ErrorTypeName,
            Type:  openapi.Schema_TYPE_OBJECT,
            Properties: map[string]*openapi.ReferenceableSchema{
                "code": openapi.NewReferenceableSchema(&openapi.Schema{
                    Title:  core.ErrorCodeTypeName,
                    Type:   openapi.Schema_TYPE_STRING,
                    Format: core.ErrorCodeTypeName,
                }),
                "message": openapi.NewReferenceableSchema(&openapi.Schema{
                    Type:        openapi.Schema_TYPE_STRING,
                    Description: &openapi.CachedDocument{Cache: "A developer-facing error message, which should be in English."},
                }),
                "details": openapi.NewReferenceableSchema(&openapi.Schema{
                    Type:        openapi.Schema_TYPE_ARRAY,
                    Items:       anySchema,
                    Description: &openapi.CachedDocument{Cache: "A list of messages that carry the error details.  There is a common set of message types for APIs to use."},
                }),
            },
        }), nil
    default:
        return nil, nil
    }
}
