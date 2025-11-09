package compiler

import (
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/openapi"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

func compileEnumDecl(ctx context.Context, decl *lang.EnumDecl) (*openapi.ReferenceableSchema, error) {
	_ = ctx

	schema := &openapi.Schema{
		Title: decl.GetFullName(),
		Type:  openapi.Schema_TYPE_STRING,
	}

	var style *lang.CaseStyle
	if attr := lang.GetAttribute(decl.Attributes, "case_style"); attr != nil {
		style = lang.NewCaseStyle(attr)
	}

	for _, e := range decl.Type.Enumerators {
		if strings.ToLower(e.Name) == "unspecified" {
			continue
		}

		enumName := e.Name

		if style != nil {
			enumName = style.ToCase(e.Name)
		}

		if alias, _ := e.GetStringAttribute("alias"); len(alias) > 0 {
			enumName = alias
		}

		schema.Enum = append(schema.Enum, core.NewStringValue(enumName))
	}

	return openapi.NewReferenceableSchema(schema), nil
}
