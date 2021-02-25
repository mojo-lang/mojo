package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
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

func compileNominalType(ctx *Context, nominalType *lang.NominalType) (*openapi.ReferenceableSchema, error) {
	schema := &openapi.Schema{
	}

	switch nominalType.GetFullName() {
	case core.UInt8TypeName, core.UInt16TypeName, core.UInt32TypeName, core.UInt64TypeName,
		core.Int8TypeName, core.Int16TypeName, core.Int32TypeName, core.Int64TypeName,
		core.IntTypeName, core.UIntTypeName:
		schema.Type = openapi.Schema_TYPE_INTEGER
		if nominalType.Name != "Int" || nominalType.Name != "UInt" {
			schema.Format = strings.ToLower(nominalType.Name)
		}
	case core.Float32TypeName, core.Float64TypeName:
		schema.Type = openapi.Schema_TYPE_NUMBER
	case core.BoolTypeName:
		schema.Type = openapi.Schema_TYPE_BOOLEAN
	case core.StringTypeName:
		schema.Type = openapi.Schema_TYPE_STRING
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
	case core.DictionaryTypeName:
		schema.Type = openapi.Schema_TYPE_OBJECT
		s, err := compileNominalType(ctx, nominalType.GenericArguments[0])
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
