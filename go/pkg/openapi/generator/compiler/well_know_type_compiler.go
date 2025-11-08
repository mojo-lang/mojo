package compiler

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/db/go/pkg/mojo/db"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/packages/openapi/go/pkg/mojo/openapi"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

type WellKnowTypeCompiler struct {
}

// Compile some well know types will have different schema when as the nominal type, aka, when in struct field
func (w *WellKnowTypeCompiler) Compile(ctx context.Context, decl *lang.NominalType) (*openapi.ReferenceableSchema, error) {
	_ = ctx

	fullName := decl.GetFullName()
	schema := &openapi.Schema{
		Title: fullName,
	}
	switch fullName {
	case core.ErrorCodeTypeFullName:
		if !decl.HasAttribute(core.EncodingAsStringAttributeName) {
			return nil, nil
		}
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.ErrorCodeTypeName
	default:
		return nil, nil
	}
	return openapi.NewReferenceableSchema(schema), nil
}

func (w *WellKnowTypeCompiler) CompileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) (*openapi.ReferenceableSchema, error) {
	_ = ctx
	fullName := decl.GetFullName()
	switch fullName {
	case core.ValueTypeFullName:
		return openapi.NewReferenceableSchema(&openapi.Schema{
			Title: fullName,
			OneOf: []*openapi.ReferenceableSchema{
				openapi.NewReferenceableSchema(&openapi.Schema{
					Title: core.NullTypeFullName,
					Type:  openapi.Schema_TYPE_NULL,
				}),
				openapi.NewReferenceableSchema(&openapi.Schema{
					Title: core.BoolTypeFullName,
					Type:  openapi.Schema_TYPE_BOOLEAN,
				}),
				openapi.NewReferenceableSchema(&openapi.Schema{
					Title: core.StringTypeFullName,
					Type:  openapi.Schema_TYPE_STRING,
				}),
				openapi.NewReferenceableSchema(&openapi.Schema{
					Title:       core.BytesTypeFullName,
					Type:        openapi.Schema_TYPE_STRING,
					Format:      core.BytesTypeName,
					Description: &openapi.CachedDocument{Cache: "the format is: `b64.{base64 encoded bytes}`"},
				}),
				openapi.NewReferenceableSchema(int64Schema.Clone()),
				openapi.NewReferenceableSchema(float64Schema.Clone()),
				openapi.NewReferenceableSchema(&openapi.Schema{
					Title: core.ObjectTypeFullName,
					Type:  openapi.Schema_TYPE_OBJECT,
				}),
				openapi.NewReferenceableSchema(&openapi.Schema{
					Title: core.ValuesTypeFullName,
					Type:  openapi.Schema_TYPE_ARRAY,
				}),
			},
		}), nil
	}
	return nil, nil
}

func (w *WellKnowTypeCompiler) CompileStruct(ctx context.Context, decl *lang.StructDecl) (*openapi.ReferenceableSchema, error) {
	_ = ctx

	fullName := decl.GetFullName()
	schema := &openapi.Schema{
		ExternalDocs: nil,
		Example:      nil,
		Title:        fullName,
		Description:  &openapi.CachedDocument{Cache: decl.GetDocument().GetContent()},
	}

	switch fullName {
	case core.AnyTypeFullName:
		schema.Type = openapi.Schema_TYPE_OBJECT
		schema.Properties = map[string]*openapi.ReferenceableSchema{
			"@type": openapi.NewReferenceableSchema(&openapi.Schema{
				Title: core.StringTypeFullName,
				Type:  openapi.Schema_TYPE_STRING,
			}),
		}
		schema.Required = append(schema.Required, "@type")
	case core.BoolValueTypeFullName:
		schema.Type = openapi.Schema_TYPE_BOOLEAN
	case core.Int32ValueTypeFullName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		schema.Format = core.Int32TypeName
	case core.UInt32ValueTypeFullName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		schema.Format = core.UInt32TypeName
	case core.Int64ValueTypeFullName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		schema.Format = core.Int64TypeName
	case core.UInt64ValueTypeFullName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		schema.Format = core.UInt64TypeName
	case core.PercentageTypeFullName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		schema.Format = core.PercentageTypeName
	case core.Float32ValueTypeFullName:
		schema.Type = openapi.Schema_TYPE_NUMBER
		schema.Format = core.Float32TypeName
	case core.Float64ValueTypeFullName:
		schema.Type = openapi.Schema_TYPE_NUMBER
		schema.Format = core.Float64TypeName
	case core.RatioTypeFullName:
		schema.Type = openapi.Schema_TYPE_NUMBER
		schema.Format = core.RatioTypeName
	case core.StringValueTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
	case core.BytesValueTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.BytesTypeName
	case core.UrlTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.UrlTypeName
	case core.DurationTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.DurationTypeName
	case core.TimestampTypeFullName, db.DeleteTimeTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.TimestampTypeName
	case core.DateTimeTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.DateTimeTypeName
	case core.DateTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.DateTypeName
	case core.ColorTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.ColorTypeName
	case core.ChecksumTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.ChecksumTypeName
	case core.EmailAddressTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.EmailAddressTypeName
	case core.FieldMaskTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.FieldMaskTypeName
	case core.PlatformTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.PlatformTypeName
	case core.VersionTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.VersionTypeName
	case core.UuidTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.UuidTypeName
	case core.ObjectTypeFullName:
		schema.Type = openapi.Schema_TYPE_OBJECT
	case core.ValuesTypeFullName:
		return openapi.NewReferenceableSchema(&openapi.Schema{
			Title: fullName,
			Type:  openapi.Schema_TYPE_ARRAY,
		}), nil
	// case core.ErrorTypeFullName:
	//    return openapi.NewReferenceableSchema(&openapi.Schema{
	//        Title: fullName,
	//        Type:  openapi.Schema_TYPE_OBJECT,
	//        Properties: map[string]*openapi.ReferenceableSchema{
	//            "code": openapi.NewReferenceableSchema(&openapi.Schema{
	//                Title:  core.ErrorCodeTypeFullName,
	//                Type:   openapi.Schema_TYPE_STRING,
	//                Format: core.ErrorCodeTypeName,
	//            }),
	//            "message": openapi.NewReferenceableSchema(&openapi.Schema{
	//                Type:        openapi.Schema_TYPE_STRING,
	//                Description: &openapi.CachedDocument{Cache: "A developer-facing error message, which should be in English."},
	//            }),
	//            "details": openapi.NewReferenceableSchema(&openapi.Schema{
	//                Title:       "Array<" + core.AnyTypeName + ">",
	//                Type:        openapi.Schema_TYPE_ARRAY,
	//                Items:       openapi.NewReferenceableSchema(anySchema),
	//                Description: &openapi.CachedDocument{Cache: "A list of messages that carry the error details.  There is a common set of message types for APIs to use."},
	//            }),
	//        },
	//    }), nil
	default:
		return nil, nil
	}

	return openapi.NewReferenceableSchema(schema), nil
}
