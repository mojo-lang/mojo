package compiler

import (
    "fmt"
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/document/go/pkg/mojo/document"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
    "sort"
    "strings"
)

type SchemaCompiler struct {
    Components *openapi.Components
}

func wrapCode(code string) *document.Inline {
    if len(code) > 0 {
        return document.NewCodeInlineFrom(code)
    }
    return document.NewTextInline(code)
}

func wrapCodeToBlock(code string) *document.Block {
    return document.NewPainBlock(wrapCode(code))
}

func (s *SchemaCompiler) Compile(decl *lang.Declaration, schema *openapi.Schema) (*document.Document, error) {
    doc := &document.Document{}

    if schema.Type == openapi.Schema_TYPE_OBJECT || len(schema.AllOf) > 0 {
        if schema.AdditionalProperties != nil {
            //FIXME for the map type
            return doc, nil
        }

        table := &document.Table{
            Caption:   nil,
            Alignment: 0,
            Header:    document.NewTextTableHeader("field", "type", "format", "required", "default", "description"),
        }

        fieldNames := decl.GetStructDecl().GetAllFieldNames(lang.FieldNamOptionUseAlias)
        if len(fieldNames) == 0 {
            fieldNames = schema.FieldNames(s.Components.Schemas)
            sort.Strings(fieldNames)
        }

        if len(fieldNames) > 0 {
            s.compileFields(context.Empty(), fieldNames, schema, table)
            doc.AppendTable(table)
        }
    } else if schema.Type == openapi.Schema_TYPE_ARRAY {
        if schema.Items == nil {
            logs.Infow("do not generate the document for empty array", "type", schema.Title)
            return nil, nil
        }

        table := &document.Table{
            Caption:   nil,
            Alignment: 0,
            Header:    document.NewTextTableHeader("type", "description"),
        }
        row := &document.Table_Row{}

        //row.Vals = append(row.Vals, document.NewTextTableCell(""))

        typeName := schema.GetTypeName(s.Components.Schemas)
        row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(typeName)))

        // description
        d := schema.Items.GetDescription(s.Components.Schemas)
        if d != nil {
            row.Vals = append(row.Vals, &document.Table_Cell{Vals: doc.Blocks})
        }

        table.Rows = append(table.Rows, row)
        doc.AppendTable(table)
    } else if len(schema.OneOf) > 0 {
        table := &document.Table{
            Caption:   nil,
            Alignment: 0,
            Header:    document.NewTextTableHeader("type", "format", "description"),
        }

        for _, item := range schema.OneOf {
            row := &document.Table_Row{}

            typeName := item.GetTypeName(s.Components.Schemas)
            row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(typeName)))

            typeFormat := item.GetFormat(s.Components.Schemas)
            row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(typeFormat))) // format

            // description
            summary := item.GetSummary(s.Components.Schemas)
            row.Vals = append(row.Vals, document.NewTextTableCell(summary))

            table.Rows = append(table.Rows, row)
        }

        doc.AppendTable(table)
    }

    return doc, nil
}

func (s *SchemaCompiler) compileFields(ctx context.Context, fieldNames []string, schema *openapi.Schema, table *document.Table) {
    if len(schema.AllOf) > 0 {
        for _, sch := range schema.AllOf {
            url := sch.GetReferenceUrl()
            if url != nil {
                sc := s.Components.GetSchema(url)
                if sc != nil {
                    s.compileFields(ctx, fieldNames, sc, table)
                } else {
                    logs.Warnw("failed to find referenced schema", "url", url.Format(), "schema", schema.Title)
                }
            } else {
                sc := sch.GetSchema()
                if sc != nil {
                    s.compileFields(ctx, fieldNames, sc, table)
                } else {
                    logs.Warnw("failed to find schema from the allOf schemas", "schema", schema.Title)
                }
            }
        }
    }

    for _, fieldName := range fieldNames {
        if strings.HasPrefix(fieldName, "@") {
            fieldName = "@" + strcase.ToLowerCamel(fieldName)
        } else {
            fieldName = strcase.ToLowerCamel(fieldName)
        }

        property := schema.Properties[fieldName]
        if property == nil {
            continue
        }

        row := &document.Table_Row{}
        row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(fieldName)))

        typeName := property.GetTypeName(s.Components.Schemas)
        row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(typeName)))

        typeFormat := property.GetFormat(s.Components.Schemas)
        row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(typeFormat))) // format

        required := "N"
        if schema.IsPropertyRequired(fieldName) {
            required = "Y"
        }
        row.Vals = append(row.Vals, document.NewTextTableCell(required)) // required

        row.Vals = append(row.Vals, document.NewTextTableCell("")) // default

        // const
        constVal := ""
        if constSchema := property.GetSchemaOf(s.Components.Schemas); constSchema != nil && len(constSchema.Enum) == 1 {
            constVal = constSchema.Enum[0].GetString()
            if len(constVal) > 0 {
                constVal = fmt.Sprintf("the value must be const to \"%s\"", constVal)
            }
        }

        // description
        doc := property.GetDescription(s.Components.Schemas)
        if doc != nil {
            if len(constVal) > 0 {
                doc.Blocks = append(doc.Blocks, document.NewTextPlainBlock(constVal))
            }
            row.Vals = append(row.Vals, &document.Table_Cell{Vals: doc.Blocks})
        } else {
            if len(constVal) > 0 {
                row.Vals = append(row.Vals, document.NewTextTableCell(constVal))
            }
        }

        table.Rows = append(table.Rows, row)
    }
}
