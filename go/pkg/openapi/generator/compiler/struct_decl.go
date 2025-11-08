package compiler

import (
	"errors"
	"fmt"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/packages/openapi/go/pkg/mojo/openapi"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

var PrimeTypes = map[string]bool{
	core.UInt8TypeFullName:    true,
	core.UInt16TypeFullName:   true,
	core.UInt32TypeFullName:   true,
	core.UInt64TypeFullName:   true,
	core.Int8TypeFullName:     true,
	core.Int16TypeFullName:    true,
	core.Int32TypeFullName:    true,
	core.Int64TypeFullName:    true,
	core.IntTypeFullName:      true,
	core.UIntTypeFullName:     true,
	core.PositiveTypeFullName: true,
	core.NegativeTypeFullName: true,
	core.ByteTypeFullName:     true,
	core.SizeTypeFullName:     true,
	core.Float32TypeFullName:  true,
	core.FloatTypeFullName:    true,
	core.Float64TypeFullName:  true,
	core.DoubleTypeFullName:   true,
	core.NullTypeFullName:     true,
	core.BoolTypeFullName:     true,
	core.StringTypeFullName:   true,
	core.BytesTypeFullName:    true,
	core.ArrayTypeFullName:    true,
	core.MapTypeFullName:      true,
	core.UnionTypeFullName:    true,
}

func IsPrimeType(typeFullName string) bool {
	if _, ok := PrimeTypes[typeFullName]; ok {
		return true
	}
	return false
}

func CompileStructDecl(ctx context.Context, decl *lang.StructDecl) error {
	thisCtx := context.WithType(ctx, decl)

	if IsPrimeType(decl.GetFullName()) {
		return nil
	}

	for _, s := range decl.StructDecls {
		err := CompileStructDecl(thisCtx, s)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to parse the inner struct decl %s in %s: %s", s.Name, decl.Name, err.Error()))
		}
	}

	// for _, e := range decl.EnumDecls {
	//    err := compileEnumDecl(thisCtx, e)
	// }

	components := context.OpenAPIComponents(thisCtx)

	// add a dummy schema first for recursion reference
	components.Schemas[decl.GetFullName()] = &openapi.Schema{
		Title: "DUMMY",
		Type:  0,
	}

	schema, err := compileStructDecl(thisCtx, decl)
	if err != nil {
		return err
	}

	components.Schemas[decl.GetFullName()] = schema.GetSchema()
	return nil
}

func compileStructDecl(ctx context.Context, decl *lang.StructDecl) (*openapi.ReferenceableSchema, error) {
	wellKnow := &WellKnowTypeCompiler{}
	if s, err := wellKnow.CompileStruct(ctx, decl); err != nil {
		return nil, err
	} else if s != nil {
		return s, nil
	}

	if decl.Type == nil {
		logs.Warnw("compile an empty object because of the struct has no type" /*, ctx.GetNamesForLogs()...*/)
		return openapi.NewReferenceableSchema(&openapi.Schema{
			Title: decl.GetFullName(),
			Type:  openapi.Schema_TYPE_OBJECT,
		}), nil
	}

	schema := &openapi.Schema{
		Discriminator: nil,
		ExternalDocs:  nil,
		Example:       nil,
		Title:         decl.GetFullName(),
		Description:   &openapi.CachedDocument{Cache: decl.GetDocument().GetContent()},
	}

	if len(decl.Type.Fields) > 0 {
		schema.Type = openapi.Schema_TYPE_OBJECT
		schema.Properties = make(map[string]*openapi.ReferenceableSchema)
		for _, field := range decl.Type.Fields {
			if private, _ := field.GetBoolAttribute("private"); private {
				continue
			}

			s, err := compileNominalType(ctx, field.Type)
			if err != nil {
				return nil, err
			}

			name := strcase.ToLowerCamel(field.Name)
			if alias, _ := field.GetStringAttribute("alias"); len(alias) > 0 {
				name = strcase.ToLowerCamel(alias)
			}
			schema.Properties[name] = s

			if field.Document != nil {
				s.SetDescription(&openapi.CachedDocument{Cache: field.GetDocument().GetContent()})
			}

			if required, _ := lang.GetBoolAttribute(field.Type.Attributes, "required"); required {
				schema.Required = append(schema.Required, name)
			}
		}
	}

	if inheritCount := len(decl.Type.Inherits); inheritCount > 0 {
		var schemas []*openapi.ReferenceableSchema
		for _, inherit := range decl.Type.Inherits {
			s, err := compileNominalType(ctx, inherit)
			if err != nil {
				return nil, err
			}
			schemas = append(schemas, s)
		}

		if inheritCount == 1 && len(decl.Type.Fields) == 0 {
			reference := schemas[0].GetReference()
			if reference != nil {
				schema.AllOf = []*openapi.ReferenceableSchema{schemas[0]}
			} else {
				s := schemas[0].GetSchema()
				s.Title = schema.Title
				s.Description = schema.Description
				return openapi.NewReferenceableSchema(s), nil
			}
		} else {
			s := &openapi.Schema{}
			s.Title = schema.Title
			s.Type = openapi.Schema_TYPE_OBJECT
			s.Description = schema.Description

			if len(schema.Properties) > 0 {
				schema.Description = nil
				schemas = append(schemas, openapi.NewReferenceableSchema(schema))
			}

			s.AllOf = schemas
			return openapi.NewReferenceableSchema(s), nil
		}
	}

	return openapi.NewReferenceableSchema(schema), nil
}
