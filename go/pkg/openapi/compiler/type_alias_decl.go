package compiler

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

func CompileTypeAliasDecl(ctx *Context, decl *lang.TypeAliasDecl) (*openapi.ReferenceableSchema, error) {
	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	if PrimeTypes[decl.GetFullName()] {
		return nil, nil
	}

	// add an dummy schema first for recursion reference
	ctx.Components.Schemas[decl.GetFullName()] = &openapi.Schema{
		Title: "DUMMY",
		Type:  0,
	}

	schema, err := compileTypeAliasDecl(ctx, decl)
	if err != nil {
		return nil, err
	}

	ctx.Components.Schemas[decl.GetFullName()] = schema.GetSchema()
	return schema, nil
}

func compileTypeAliasDecl(ctx *Context, decl *lang.TypeAliasDecl) (*openapi.ReferenceableSchema, error) {
	if decl.Type == nil {
		logs.Warnw("compile an empty object because of the struct has no type", ctx.GetNamesForLogs()...)
		return openapi.NewReferenceableSchema(&openapi.Schema{
			Title: decl.GetFullName(),
			Type:  openapi.Schema_TYPE_OBJECT,
		}), nil
	}

	t, err := compileNominalType(ctx, decl.Type)
	if err != nil {
		return nil, err
	}

	reference := t.GetReference()
	if reference != nil {
		return openapi.NewReferencedSchema(&openapi.Reference{
			Ref:         reference.Ref,
			Description: &openapi.CachedDocument{Cache: decl.GetDocument().GetContent()},
		}), nil
	}

	schema := t.GetSchema()
	if schema != nil {
		return openapi.NewReferenceableSchema(schema), nil
	}

	return nil, err
}
