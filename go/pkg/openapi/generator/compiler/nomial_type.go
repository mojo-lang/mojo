package compiler

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/packages/openapi/go/pkg/mojo/openapi"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

var schemaCompilers []SchemaCompiler

var (
	float32Schema = &openapi.Schema{
		Title:  core.Float32TypeFullName,
		Type:   openapi.Schema_TYPE_NUMBER,
		Format: core.Float32TypeName,
	}

	float64Schema = &openapi.Schema{
		Title:  core.Float64TypeFullName,
		Type:   openapi.Schema_TYPE_NUMBER,
		Format: core.Float64TypeName,
	}

	int64Schema = &openapi.Schema{
		Title:  core.Int64TypeFullName,
		Type:   openapi.Schema_TYPE_INTEGER,
		Format: core.Int64TypeName,
	}
)

func init() {
	schemaCompilers = append(schemaCompilers,
		&WellKnowTypeCompiler{},
		&ReferenceCompiler{},
	)
}

type NominalTypeCompiler struct {
	Context context.Context
}

func (n *NominalTypeCompiler) Compile(nominalType *lang.NominalType) (*openapi.Schema, error) {
	if n.Context == nil {
		n.Context = context.WithOpenAPIComponents(context.Empty(), openapi.NewComponents())
	}
	schema, err := compileNominalType(n.Context, nominalType)
	if err != nil {
		return nil, err
	}

	components := context.OpenAPIComponents(n.Context)
	if components == nil || len(components.Schemas) == 0 {
		return schema.GetSchema(), err
	}
	return schema.GetSchemaOf(components.Schemas), nil
}

func (n *NominalTypeCompiler) GetSchema(s *openapi.ReferenceableSchema) *openapi.Schema {
	if s != nil {
		components := context.OpenAPIComponents(n.Context)
		return s.GetSchemaOf(components.Schemas)
	}
	return nil
}

func compileNominalType(ctx context.Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error) {
	schema := &openapi.Schema{}

	attribute := lang.GetAttribute(nominalType.Attributes, core.TypeFormatAttributeName)
	if attribute != nil && len(attribute.GenericArguments) > 0 {
		return compileNominalType(ctx, attribute.GenericArguments[0])
	}

	fullName := nominalType.GetFullName()
	schema.Title = fullName

	switch fullName {
	case core.UInt8TypeFullName, core.UInt16TypeFullName, core.UInt32TypeFullName, core.UInt64TypeFullName,
		core.Int8TypeFullName, core.Int16TypeFullName, core.Int32TypeFullName, core.Int64TypeFullName,
		core.IntTypeFullName, core.UIntTypeFullName, core.ByteTypeFullName, core.SizeTypeFullName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		if nominalType.Name != "Int" && nominalType.Name != "UInt" {
			schema.Format = nominalType.Name
		}
	case core.PositiveTypeFullName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		schema.Format = core.PositiveTypeName
	case core.NegativeTypeFullName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		schema.Format = core.NegativeTypeName
	case core.Float32TypeFullName, core.FloatTypeFullName:
		schema = float32Schema.Clone()
	case core.Float64TypeFullName, core.DoubleTypeFullName:
		schema = float64Schema.Clone()
	case core.NullTypeFullName:
		schema.Type = openapi.Schema_TYPE_NULL
	case core.BoolTypeFullName:
		schema.Type = openapi.Schema_TYPE_BOOLEAN
	case core.StringTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		if constant, _ := lang.GetStringAttribute(nominalType.Attributes, core.ConstAttributeName); len(constant) > 0 {
			schema.Enum = []*core.Value{
				core.NewStringValue(constant),
			}
		}
	case core.BytesTypeFullName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = core.BytesTypeName
		// schema.Format = core.BytesTypeName
		// schema.Description = &openapi.CachedDocument{Cache: "the format is: `b64.{base64 encoded bytes}`"}
	case core.ArrayTypeFullName:
		s, err := compileNominalType(ctx, nominalType.GenericArguments[0])
		if err != nil {
			return nil, err
		}

		schema.Type = openapi.Schema_TYPE_ARRAY
		schema.Items = s
		if val, err := lang.GetIntegerAttribute(nominalType.Attributes, core.MinLengthAttributeName); err == nil {
			schema.MinItems = uint64(val)
		}
		if val, err := lang.GetIntegerAttribute(nominalType.Attributes, core.MaxLengthAttributeName); err == nil {
			schema.MaxItems = uint64(val)
		}
	case core.MapTypeFullName:
		schema.Type = openapi.Schema_TYPE_OBJECT
		s, err := compileNominalType(ctx, nominalType.GenericArguments[1])
		if err != nil {
			return nil, err
		}
		schema.AdditionalProperties = s
	case core.UnionTypeFullName:
		var schemas []*openapi.ReferenceableSchema
		for _, argument := range nominalType.GenericArguments {
			s, err := compileNominalType(ctx, argument)
			if err != nil {
				return nil, err
			}
			schemas = append(schemas, s)
		}
		if alias, _ := lang.GetStringAttribute(nominalType.Attributes, lang.OriginalTypeAliasName); len(alias) > 0 {
			typeAlias := context.TypeAliasDecl(ctx)
			if typeAlias != nil {
				schema.Title = typeAlias.GetFullName()
			}
		}
		schema.OneOf = schemas
	default:
		if enumDecl := nominalType.TypeDeclaration.GetEnumDecl(); enumDecl != nil {
			return compileEnumDecl(ctx, enumDecl)
		}

		for _, compiler := range schemaCompilers {
			s, err := compiler.Compile(ctx, nominalType)
			if err != nil {
				return nil, err
			}
			if s != nil {
				return s, nil
			}
		}
	}

	if schema.IsEmpty() {
		return nil, logs.NewErrorw("failed to compile the nominal type", "type", nominalType.GetFullName())
	}

	return openapi.NewReferenceableSchema(schema), nil
}
