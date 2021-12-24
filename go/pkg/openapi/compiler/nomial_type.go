package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
	"strings"
)

var schemaCompilers []SchemaCompiler

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
		n.Context = context.WithComponents(context.Empty(), openapi.NewComponents())
	}
	schema, err := compileNominalType(n.Context, nominalType)
	if err != nil {
		return nil, err
	}

	components := context.Components(n.Context)
	if components == nil || len(components.Schemas) == 0 {
		return schema.GetSchema(), err
	}
	return schema.GetSchemaOf(components.Schemas), nil
}

func (n *NominalTypeCompiler) GetSchema(s *openapi.ReferenceableSchema) *openapi.Schema {
	if s != nil {
		components := context.Components(n.Context)
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

	switch nominalType.GetFullName() {
	case core.UInt8TypeName, core.UInt16TypeName, core.UInt32TypeName, core.UInt64TypeName,
		core.Int8TypeName, core.Int16TypeName, core.Int32TypeName, core.Int64TypeName,
		core.IntTypeName, core.UIntTypeName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		if nominalType.Name != "Int" && nominalType.Name != "UInt" {
			schema.Format = strings.ToLower(nominalType.Name)
		}
	case core.Float32TypeName, core.FloatTypeName:
		schema.Type = openapi.Schema_TYPE_NUMBER
		schema.Format = "float32"
	case core.Float64TypeName, core.DoubleTypeName:
		schema.Type = openapi.Schema_TYPE_NUMBER
		schema.Format = "float64"
	case core.BoolTypeName:
		schema.Type = openapi.Schema_TYPE_BOOLEAN
	case core.StringTypeName:
		schema.Type = openapi.Schema_TYPE_STRING
		if constant, _ := lang.GetStringAttribute(nominalType.Attributes, core.ConstAttributeName); len(constant) > 0 {
			schema.Enum = []*core.Value{
				core.NewStringValue(constant),
			}
		}
	case core.BytesTypeName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = "base64"
	case core.ArrayTypeName:
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
	case core.MapTypeName:
		schema.Type = openapi.Schema_TYPE_OBJECT
		s, err := compileNominalType(ctx, nominalType.GenericArguments[1])
		if err != nil {
			return nil, err
		}
		schema.AdditionalProperties = s
	case core.UnionTypeName:
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
	case core.TimestampTypeName, core.DateTimeTypeName, db.DeleteTimeTypeName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = "DateTime"
	case core.DateTypeName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = "Date"
	case core.EmailAddressTypeName:
		schema.Type = openapi.Schema_TYPE_STRING
		schema.Format = "EmailAddress"
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

	return openapi.NewReferenceableSchema(schema), nil
}
