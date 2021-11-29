package compiler

import (
	"github.com/mojo-lang/document/go/pkg/markdown"
	"github.com/mojo-lang/document/go/pkg/mojo/document"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
	"net/http"
)

type ApiCompiler struct {
	Package *lang.Package
	Api     *openapi.OpenAPI
}

func (a *ApiCompiler) Compile(pkg *lang.Package, api *openapi.OpenAPI) (*document.Document, error) {
	ctx := &Context{Context: &context.Context{}}
	doc := &document.Document{}

	a.Api = api
	a.Package = pkg
	a.compileHeader(ctx, api, doc)

	err := openapi.OrderedPathItemIterator(api.Paths.Vals, func(path string, item *openapi.PathItem) error {
		d, err := a.compilePathItem(ctx, path, item)
		if err != nil {
			return err
		}
		for _, b := range d.Blocks {
			doc.Blocks = append(doc.Blocks, b)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return doc, nil
}

// {{Info.Title}}
//
// {{Info.Description}}
//
//
// 整体说明
//

const Introduce = `
## 整体说明

1. 字符串都为utf8格式;
2. HTTP Headers:
	1. Content-Type设置为：application/json
3. DataTime格式参考RFC3339标准
`

const Error = `
## 错误处理
错误的具体信息将在error字段中返回。

### 错误码示例
`

const ErrorExample = `{
    "code": "400",
    "message": "Param Error"
}`

const ErrorCodes = `
### 状态码列表

| 状态码 | 说明 |
|---|---|
|200|返回正常|
|400|参数错误|
|401|无access key或key无效|
|500|服务器内部错误|
`

func (a *ApiCompiler) compileHeader(ctx *Context, api *openapi.OpenAPI, doc *document.Document) error {
	doc.AppendHeaderFromText(1, api.GetInfo().GetTitle())

	description := api.GetInfo().GetDescription().GetDocument()
	if description != nil {
		doc.Blocks = append(doc.Blocks, description.Blocks...)
	}

	doc.Blocks = append(doc.Blocks, document.NewDivisionBlock(document.NewTocDivision()))

	{
		md := markdown.New()
		introduce, _ := md.Parse(Introduce)
		doc.Blocks = append(doc.Blocks, introduce.Blocks...)
	}

	{
		md := markdown.New()
		err, _ := md.Parse(Error)
		doc.Blocks = append(doc.Blocks, err.Blocks...)

		example := document.NewCodeBlock("json", ErrorExample)
		doc.AppendCodeBlock(example)

		errCodes, _ := md.Parse(ErrorCodes)
		doc.Blocks = append(doc.Blocks, errCodes.Blocks...)
	}

	return nil
}

func (a *ApiCompiler) compilePathItem(ctx *Context, path string, item *openapi.PathItem) (*document.Document, error) {
	doc := &document.Document{}

	compile := func(d *document.Document, e error) error {
		if e == nil && d != nil {
			doc.Blocks = append(doc.Blocks, d.Blocks...)
		}
		return e
	}

	if e := compile(a.compileMethod(ctx, path, http.MethodGet, item.GetGet())); e != nil {
		return nil, e
	}
	if e := compile(a.compileMethod(ctx, path, http.MethodPut, item.GetPut())); e != nil {
		return nil, e
	}
	if e := compile(a.compileMethod(ctx, path, http.MethodPost, item.GetPost())); e != nil {
		return nil, e
	}
	if e := compile(a.compileMethod(ctx, path, http.MethodDelete, item.GetDelete())); e != nil {
		return nil, e
	}
	if e := compile(a.compileMethod(ctx, path, http.MethodPatch, item.GetPatch())); e != nil {
		return nil, e
	}
	if e := compile(a.compileMethod(ctx, path, http.MethodHead, item.GetHead())); e != nil {
		return nil, e
	}
	if e := compile(a.compileMethod(ctx, path, http.MethodOptions, item.GetOptions())); e != nil {
		return nil, e
	}
	if e := compile(a.compileMethod(ctx, path, http.MethodTrace, item.GetTrace())); e != nil {
		return nil, e
	}

	return doc, nil
}

//## {{.Name}}
//
//{{.Description}}
//
//
//### 请求路径
//
//```http
//{{.HttpMethod}} {{.HttpPath}}
//```
func (a *ApiCompiler) compileMethod(ctx *Context, path string, method string, operation *openapi.Operation) (*document.Document, error) {
	if operation == nil {
		return nil, nil
	}

	doc := &document.Document{}
	doc.AppendHeaderFromText(2, operation.Summary)

	description := operation.Description.GetDocument()
	if description != nil {
		doc.AppendBlocks(description.Blocks...)
	}

	doc.AppendHeaderFromText(3, "请求路径")

	requestPath := document.NewCodeBlock("http", method+" "+path)
	doc.AppendCodeBlock(requestPath)

	compiler := &OperationCompiler{
		Package:    a.Package,
		Components: a.Api.Components,
	}
	op, err := compiler.Compile(operation)
	if err != nil {
		return nil, err
	}
	doc.AppendBlocks(op.Blocks...)

	return doc, nil
}
