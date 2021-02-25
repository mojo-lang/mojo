package compiler

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/document/go/pkg/mojo/document"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

type SchemaCompiler struct {
	Components *openapi.Components
}

func (s *SchemaCompiler) Compile(schema *openapi.Schema) (*document.Document, error) {
	doc := &document.Document{}

	if schema.Type == openapi.Schema_TYPE_OBJECT {
		table := &document.Table{
			Caption:   nil,
			Alignment: 0,
			Header: &document.Table_Header{
				Values: []*document.Table_Cell{
					document.NewTextCell("字段"),
					document.NewTextCell("类型"),
					document.NewTextCell("格式类型"),
					document.NewTextCell("是否必须"),
					document.NewTextCell("默认值"),
					document.NewTextCell("说明"),
				},
			},
		}

		ctx := &Context{}
		s.compileFields(ctx, schema, table)

		doc.AppendTable(table)
	}

	return doc, nil
}

func (s *SchemaCompiler) compileFields(ctx *Context, schema *openapi.Schema, table *document.Table) {
	if len(schema.AllOf) > 0 {
		for _, sch := range schema.AllOf {
			url := sch.GetReferenceUrl()
			if url != nil {
				sc := s.Components.GetSchema(url)
				if sc != nil {
					s.compileFields(ctx, sc, table)
				} else {
					logs.Warnw("failed to find referenced schema", "url", url.Encode(), "schema", schema.Title)
				}
			} else {
				sc := sch.GetSchema()
				if sc != nil {
					s.compileFields(ctx, sc, table)
				} else {
					logs.Warnw("failed to find schema from the allOf schemas", "schema", schema.Title)
				}
			}
		}
	}

	for key, property := range schema.Properties {
		row := &document.Table_Row{}

		row.Values = append(row.Values, document.NewTextCell(key))
		row.Values = append(row.Values, document.NewTextCell(property.GetSchemaName()))

		row.Values = append(row.Values, document.NewTextCell("")) // 格式类型
		row.Values = append(row.Values, document.NewTextCell("")) // 是否必须
		row.Values = append(row.Values, document.NewTextCell("")) // 默认值

		// 说明
		doc := property.GetDescription(s.Components)
		if doc != nil {
			row.Values = append(row.Values, &document.Table_Cell{Values: doc.Blocks})
		}

		table.Rows = append(table.Rows, row)
	}
}