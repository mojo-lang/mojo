package compiler

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/document"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/openapi"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type OperationCompiler struct {
	Package    *lang.Package
	Components *openapi.Components
}

// ### 请求参数
//
// #### Path 参数
//
// #### Query 参数
//
// | 参数 | 类型 | 是否必须 | 默认值 | 说明 |
// |:---  |:----|:-----   | :---- | :----|
// {{range .TableList}}| {{.Index}}|[ {{.TableName}}](#{{.Index}}{{.TableName}})       |{{.Comment}}| |
// {{end}}
//
// #### Body 请求对象
//
// ##### Body 请求示例
//
// {{ range .Schemas}}
// {{end}}
//
// #### 完整请求示例
//
// ### 返回值
//
// #### 返回对象
//
// {{ range .Schemas}}
// #### {{}}
//
// | 参数 | 类型 | 是否必须 | 默认值 | 说明 |
// |:---  |:----|:-----   | :---- | :----|
// {{range .TableList}}| {{.Index}}|[ {{.TableName}}](#{{.Index}}{{.TableName}})       |{{.Comment}}| |
// {{end}}
//
// #####
//
// {{end}}
//
// ### API示例

func (o *OperationCompiler) Compile(operation *openapi.Operation) (*document.Document, error) {
	doc := &document.Document{}
	ctx := context.Empty()

	doc.AppendHeaderFromText(3, "请求参数")

	parameters := operation.GetLocationParameters()
	if ps := parameters[openapi.Parameter_LOCATION_PATH]; len(ps) > 0 {
		doc.AppendHeaderFromText(4, "Path 参数")
		err := o.compilePathParameters(ctx, ps, doc)
		if err != nil {
			return nil, err
		}
	}

	if ps := parameters[openapi.Parameter_LOCATION_QUERY]; len(ps) > 0 {
		doc.AppendHeaderFromText(4, "Query 参数")
		err := o.compileQueryParameters(ctx, ps, doc)
		if err != nil {
			return nil, err
		}
	}

	if operation.GetRequestBody() != nil {
		doc.AppendHeaderFromText(4, "Body 请求对象")
		err := o.compileRequestBody(ctx, operation.GetRequestBody().GetRequestBody(), doc)
		if err != nil {
			return nil, err
		}
	}

	doc.AppendHeaderFromText(3, "返回值")
	if resp, ok := operation.Responses.Vals["200"]; ok {
		response := resp.GetResponse()
		doc.AppendHeaderFromText(4, "返回对象")
		if len(response.Content) == 0 {
			doc.AppendBlock(document.NewTextPlainBlock("对象为空"))
		} else {
			err := o.compileResponse(ctx, response, doc)
			if err != nil {
				return nil, err
			}
		}
	}

	return doc, nil
}

func (o *OperationCompiler) compilePathParameters(ctx context.Context, parameters []*openapi.Parameter, doc *document.Document) error {
	_ = ctx
	table := &document.Table{
		Caption:   nil,
		Alignment: 0,
		Header:    document.NewTextTableHeader("参数名", "参数类型", "格式类型", "说明"),
	}

	for _, parameter := range parameters {
		row := &document.Table_Row{}

		row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(parameter.Name)))

		typeName := parameter.Schema.GetTypeName(o.Components.Schemas)
		row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(typeName)))

		typeFormat := parameter.Schema.GetFormat(o.Components.Schemas)
		row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(typeFormat))) // 格式类型

		row.Vals = append(row.Vals, document.NewTextTableCell(parameter.Description))

		table.Rows = append(table.Rows, row)
	}

	doc.AppendTable(table)
	return nil
}

func (o *OperationCompiler) compileQueryParameters(ctx context.Context, parameters []*openapi.Parameter, doc *document.Document) error {
	_ = ctx
	table := &document.Table{
		Caption:   nil,
		Alignment: 0,
		Header:    document.NewTextTableHeader("参数名", "参数类型", "格式类型", "是否必须", "默认值", "说明"),
	}

	for _, parameter := range parameters {
		row := &document.Table_Row{}

		row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(parameter.Name)))

		typeName := parameter.Schema.GetTypeName(o.Components.Schemas)
		row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(typeName)))

		typeFormat := parameter.Schema.GetFormat(o.Components.Schemas)
		row.Vals = append(row.Vals, document.NewTableCell(wrapCodeToBlock(typeFormat))) // 格式类型

		required := "否"
		if parameter.Required {
			required = "是"
		}
		row.Vals = append(row.Vals, document.NewTextTableCell(required)) // 是否必须

		row.Vals = append(row.Vals, document.NewTextTableCell("")) // 默认值

		row.Vals = append(row.Vals, document.NewTextTableCell(parameter.Description))

		table.Rows = append(table.Rows, row)
	}

	doc.AppendTable(table)
	return nil
}

func (o *OperationCompiler) compileRequestBody(ctx context.Context, body *openapi.RequestBody, doc *document.Document) error {
	_ = ctx
	if body == nil || doc == nil {
		return nil
	}

	if mediaType, ok := body.Content[core.ApplicationJson]; ok {
		return o.compileMediaType(mediaType, doc)
	}
	return nil
}

func (o *OperationCompiler) compileMediaType(mediaType *openapi.MediaType, doc *document.Document) error {
	compiler := &SchemaCompiler{Components: o.Components}
	schema := mediaType.Schema.GetSchemaOf(o.Components.Schemas)

	decl := o.Package.GetIdentifier(schema.Title).GetDeclaration()
	if d, err := compiler.Compile(decl, schema); err != nil {
		return err
	} else if d != nil {
		doc.Blocks = append(doc.Blocks, d.GetBlocks()...)
	}

	deps := schema.Dependencies(o.Components.Schemas)
	for _, dep := range deps {
		decl = o.Package.GetIdentifier(dep.Title).GetDeclaration()
		if d, err := compiler.Compile(decl, dep); err != nil {
			return err
		} else if d != nil && len(d.Blocks) > 0 {
			doc.AppendHeaderFrom(4, wrapCode(dep.Title))
			doc.Blocks = append(doc.Blocks, d.GetBlocks()...)
		}
	}
	return nil
}

func (o *OperationCompiler) compileResponse(ctx context.Context, response *openapi.Response, doc *document.Document) error {
	_ = ctx
	if response == nil || doc == nil {
		return nil
	}

	if mediaType, ok := response.Content[core.ApplicationJson]; ok {
		return o.compileMediaType(mediaType, doc)
	}

	return nil
}

func (o *OperationCompiler) compileErrorResponse(ctx context.Context) {
	_ = ctx
}
