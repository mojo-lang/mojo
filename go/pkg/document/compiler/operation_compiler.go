package compiler

import (
	"github.com/mojo-lang/document/go/pkg/mojo/document"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

type OperationCompiler struct {
}

//### 请求参数
//
//#### Path 参数
//
//#### Query 参数
//
//| 参数 | 类型 | 是否必须 | 默认值 | 说明 |
//|:---  |:----|:-----   | :---- | :----|
//{{range .TableList}}| {{.Index}}|[ {{.TableName}}](#{{.Index}}{{.TableName}})       |{{.Comment}}| |
//{{end}}
//
//#### Body 请求对象
//
//##### Body 请求示例
//
//{{ range .Schemas}}
//{{end}}
//
//#### 完整请求示例
//
//### 返回参数
//
//#### 返回对象
//
//{{ range .Schemas}}
//#### {{}}
//
//| 参数 | 类型 | 是否必须 | 默认值 | 说明 |
//|:---  |:----|:-----   | :---- | :----|
//{{range .TableList}}| {{.Index}}|[ {{.TableName}}](#{{.Index}}{{.TableName}})       |{{.Comment}}| |
//{{end}}
//
//#####
//
//{{end}}
//
//### API示例

func (o *OperationCompiler) Compile(operation *openapi.Operation) (*document.Document, error) {
	doc := &document.Document{}
	ctx := &Context{}

	doc.AppendHeaderFrom(3, "请求参数")

	parameters := operation.GetLocationParameters()
	if ps := parameters[openapi.Parameter_LOCATION_PATH]; len(ps) > 0 {
		doc.AppendHeaderFrom(4, "Path 参数")
		err := o.compilePathParameters(ctx, ps, doc)
		if err != nil {
			return nil, err
		}
	}

	if ps := parameters[openapi.Parameter_LOCATION_QUERY]; len(ps) > 0 {
		doc.AppendHeaderFrom(4, "Query 参数")
		err := o.compileQueryParameters(ctx, ps, doc)
		if err != nil {
			return nil, err
		}
	}

	//if operation.GetRequestBody()

	return doc, nil
}

func (o *OperationCompiler) compilePathParameters(ctx *Context, parameters []*openapi.Parameter, doc *document.Document) error {
	return nil
}

func (o *OperationCompiler) compileQueryParameters(ctx *Context, parameters []*openapi.Parameter, doc *document.Document) error {
	return nil
}

func (o *OperationCompiler) compileRequestBody(ctx *Context, body *openapi.RequestBody, doc *document.Document) error {
	return nil
}

func (o *OperationCompiler) compileResponse(ctx *Context, response *openapi.Response, doc *document.Document) error {
	return nil
}

func (o *OperationCompiler) compileErrorResponse(ctx *Context) {
}
