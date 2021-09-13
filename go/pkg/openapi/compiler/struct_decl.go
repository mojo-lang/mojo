package compiler

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

var PrimeTypes = map[string]bool{
	core.UInt8TypeName:      true,
	core.UInt16TypeName:     true,
	core.UInt32TypeName:     true,
	core.UInt64TypeName:     true,
	core.Int8TypeName:       true,
	core.Int16TypeName:      true,
	core.Int32TypeName:      true,
	core.Int64TypeName:      true,
	core.IntTypeName:        true,
	core.UIntTypeName:       true,
	core.Float32TypeName:    true,
	core.Float64TypeName:    true,
	core.BoolTypeName:       true,
	core.StringTypeName:     true,
	core.BytesTypeName:      true,
	core.ArrayTypeName:      true,
	core.DictionaryTypeName: true,
	core.UnionTypeName:      true,
}

func CompileStructDecl(ctx *Context, decl *lang.StructDecl) error {
	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	if PrimeTypes[decl.GetFullName()] {
		return nil
	}

	for _, s := range decl.StructDecls {
		err := CompileStructDecl(ctx, s)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to parse the inner struct decl %s in %s: %s", s.Name, decl.Name, err.Error()))
		}
	}

	// add an dummy schema first for recursion reference
	ctx.Components.Schemas[decl.GetFullName()] = &openapi.Schema{
		Title: "DUMMY",
		Type:  0,
	}

	schema, err := compileStructDecl(ctx, decl)
	if err != nil {
		return err
	}

	ctx.Components.Schemas[decl.GetFullName()] = schema.GetSchema()
	return nil
}

func compileStructDecl(ctx *Context, decl *lang.StructDecl) (*openapi.ReferenceableSchema, error) {
	if decl.Type == nil {
		logs.Warnw("compile an empty object because of the struct has no type", ctx.GetNamesForLogs()...)
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
				s := ctx.Components.Schemas[reference.GetSchemaName()]
				return openapi.NewReferenceableSchema(s), nil
			} else {
				return schemas[0], nil
			}
		} else {
			s := &openapi.Schema{}
			schemas = append(schemas, openapi.NewReferenceableSchema(schema))
			s.AllOf = schemas
			return openapi.NewReferenceableSchema(s), nil
		}
	}

	return openapi.NewReferenceableSchema(schema), nil
}
