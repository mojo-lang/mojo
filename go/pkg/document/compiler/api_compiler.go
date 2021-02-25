package compiler

import (
	"errors"
	"github.com/mojo-lang/document/go/pkg/mojo/document"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
	"net/http"
)

type ApiCompiler struct {
}

func (a *ApiCompiler) Compile(api *openapi.OpenAPI) (*document.Document, error) {
	ctx := &Context{Context: &context.Context{}}
	doc := &document.Document{}

	a.compileHeader(ctx, api, doc)

	for path, item := range api.Paths.Values {
		d, err := a.compilePathItem(ctx, path, item)
		if err != nil {
			return nil, err
		}
		for _, b := range d.Blocks {
			doc.Blocks = append(doc.Blocks, b)
		}
	}

	return doc, nil
}

func (a *ApiCompiler) compileHeader(ctx *Context, api *openapi.OpenAPI, doc *document.Document) error {

	return nil
}

func (a *ApiCompiler) compilePathItem(ctx *Context, path string, item *openapi.PathItem) (*document.Document, error) {
	if item.GetGet() != nil {
		return a.compileMethod(ctx, path, http.MethodGet, item.GetGet())
	} else if item.GetPut() != nil {
		return a.compileMethod(ctx, path, http.MethodPut, item.GetPut())
	} else if item.GetPost() != nil {
		return a.compileMethod(ctx, path, http.MethodPost, item.GetPost())
	} else if item.GetDelete() != nil {
		return a.compileMethod(ctx, path, http.MethodDelete, item.GetDelete())
	} else if item.GetPatch() != nil {
		return a.compileMethod(ctx, path, http.MethodPatch, item.GetPatch())
	} else if item.GetHead() != nil {
		return a.compileMethod(ctx, path, http.MethodHead, item.GetHead())
	}
	return nil, errors.New("no supported method set")
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
	doc := &document.Document{}
	doc.AppendHeaderFrom(2, operation.Summary)

	description := operation.Description.GetDocument()
	if description != nil {
		doc.AppendBlocks(description.Blocks...)
	}

	doc.AppendHeaderFrom(3, "请求路径")

	requestPath := document.NewCodeBlock("http", method+" "+path)
	doc.AppendCodeBlock(requestPath)

	compiler := &OperationCompiler{}
	op, err := compiler.Compile(operation)
	if err != nil {
		return nil, err
	}
	doc.AppendBlocks(op.Blocks...)

	return doc, nil
}
