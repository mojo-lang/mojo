package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
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
	Context *Context
}

func (n *NominalTypeCompiler) Compile(nominalType *lang.NominalType) (*openapi.Schema, error) {
	if n.Context == nil {
		n.Context = &Context{Context: &context.Context{}}
	}
	schema, err := compileNominalType(n.Context, nominalType)
	if err != nil {
		return nil, err
	}

	if n.Context.Components == nil {
		return schema.GetSchema(), err
	}
	return schema.GetSchemaOf(n.Context.Components.Schemas), nil
}

func (n *NominalTypeCompiler) GetSchema(s *openapi.ReferenceableSchema) *openapi.Schema {
	if s != nil {
		return s.GetSchemaOf(n.Context.Components.Schemas)
	}
	return nil
}

func compileNominalType(ctx *Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error) {
	schema := &openapi.Schema{}

	attribute := lang.GetAttribute(nominalType.Attributes, "type_format")
	if attribute != nil && len(attribute.GenericArguments) > 0 {
		return compileNominalType(ctx, attribute.GenericArguments[0])
	}

	switch nominalType.GetFullName() {
	case core.UInt8TypeName, core.UInt16TypeName, core.UInt32TypeName, core.UInt64TypeName,
		core.Int8TypeName, core.Int16TypeName, core.Int32TypeName, core.Int64TypeName,
		core.IntTypeName, core.UIntTypeName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		if nominalType.Name != "Int" || nominalType.Name != "UInt" {
			schema.Format = strings.ToLower(nominalType.Name)
		}
	case core.Float32TypeName, core.Float64TypeName, core.FloatTypeName, core.DoubleTypeName:
		schema.Type = openapi.Schema_TYPE_NUMBER
	case core.BoolTypeName:
		schema.Type = openapi.Schema_TYPE_BOOLEAN
	case core.StringTypeName:
		schema.Type = openapi.Schema_TYPE_STRING
		if constant, _ := lang.GetStringAttribute(nominalType.Attributes, "const"); len(constant) > 0 {
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
		schema.OneOf = schemas
	case core.TimestampTypeName, core.DateTimeTypeName:
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
