package compiler

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func CompileTypeAliasDecl(ctx context.Context, decl *lang.TypeAliasDecl) (*openapi.ReferenceableSchema, error) {
	if PrimeTypes[decl.GetFullName()] {
		return nil, nil
	}

	thisCtx := context.WithType(ctx, decl)
	components := context.Components(thisCtx)

	disableGenerated := false
	if disableGenerate, _ := lang.GetDisableGenerateAttribute(decl.Attributes); len(disableGenerate) > 0 {
		disableGenerated = disableGenerate.Including("openapi", "")
	}

	// add an dummy schema first for recursion reference
	components.Schemas[decl.GetFullName()] = &openapi.Schema{
		Title: "DUMMY",
		Type:  0,
	}

	schema, err := compileTypeAliasDecl(thisCtx, decl)
	if err != nil {
		return nil, err
	}

	components.Schemas[decl.GetFullName()] = schema.GetSchema()

	if disableGenerated {
		delete(components.Schemas, decl.GetFullName())
		return schema, nil
	} else {
		return openapi.NewReferencedSchema(&openapi.Reference{
			Ref: &core.Url{
				Fragment: openapi.ReferenceRoot + decl.GetFullName(),
			},
		}), nil
	}
}

func compileTypeAliasDecl(ctx context.Context, decl *lang.TypeAliasDecl) (*openapi.ReferenceableSchema, error) {
	wellKnow := &WellKnowTypeCompiler{}
	if s, err := wellKnow.CompileTypeAlias(ctx, decl); err != nil {
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
