package compiler

import (
	"bytes"
	"strings"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/document/go/pkg/markdown"
	"github.com/mojo-lang/http/go/pkg/mojo/http"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"

	"github.com/mojo-lang/mojo/go/pkg/context"
	langcompiler "github.com/mojo-lang/mojo/go/pkg/mojo/compiler"
)

var methods = []string{"http.get", "http.post", "http.put", "http.delete", "http.patch", "http.options", "http.head", "http.trace"}

func CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) (*openapi.OpenAPI, error) {
	thisCtx := context.WithType(ctx, decl)

	api := openapi.New()
	api.Openapi = &core.Version{
		Major: 3,
		Minor: 0,
		Patch: 3,
	}

	var version *core.Version
	pkg := context.Package(thisCtx)
	if pkg != nil {
		version = pkg.Version
	}

	title := ""
	description := ""
	document := decl.Document.Parse().GetStructured()
	if document != nil && len(document.Blocks) > 0 {
		if header := document.Blocks[0].GetHeader(); header != nil {
			h, err := markdown.New().RenderInlinesToString(header.Text...)
			if err != nil {
				logs.Warnw("failed to render header text", "error", err.Error())
			} else {
				title = h
			}
		} else if paragraph := document.Blocks[0].GetParagraph(); paragraph != nil {
			p, err := markdown.New().RenderBlocksToString(document.Blocks[0])
			if err != nil {
				logs.Warnw("failed to render paragraph text", "error", err.Error())
			} else {
				title = p
			}
		}

		if len(document.Blocks) > 1 {
			var err error
			description, err = markdown.New().RenderBlocksToString(document.Blocks[1:]...)
			if err != nil {
				logs.Warnw("failed to render description text", "error", err.Error())
			}
		}
	}

	api.Info = &openapi.Info{
		Title:          title,
		Description:    &openapi.CachedDocument{Cache: description},
		TermsOfService: "",
		Contact:        nil,
		License:        nil,
		Version:        version,
	}

	methodCtx := context.WithValues(thisCtx, "interface.attributes", decl.Attributes)
	for _, method := range decl.Type.Methods {
		err := compileMethod(methodCtx, method, api)
		if err != nil {
			return nil, err
		}
	}
	return api, nil
}

func compileMethod(ctx context.Context, method *lang.FunctionDecl, api *openapi.OpenAPI) error {
	thisCtx := context.WithType(ctx, method)

	item := &openapi.PathItem{}
	summary, description := compileDescription(ctx, method)
	op := &openapi.Operation{
		Summary:     summary,
		Description: description,
		OperationId: method.Name,
	}
	setSourceTag := false

	for _, attribute := range method.Attributes {
		setOperation := func(op *openapi.Operation, httpMethod string) (string, error) {
			if len(attribute.Arguments) > 0 {
				if value := attribute.Arguments[0].GetStringLiteralExpr(); value != nil {
					if len(value.Value) == 0 {
						return "", nil
					}

					path, pathParam := CompilePath(value.Value)
					paramCtx := context.WithValues(thisCtx, "pathParams", pathParam)
					parameters, err := CompileParameters(paramCtx, method, httpMethod)
					if err != nil {
						return "", err
					}

					op.Parameters = parameters
					op.Responses = compileResponses(paramCtx, method)
					if httpCanCarryBody(httpMethod) {
						op.RequestBody = compileRequestBody(paramCtx, method)
					}

					return path, nil
				}
			}
			return "", nil
		}

		if attribute.PackageName == "http" {
			switch attribute.Name {
			case "get", "post", "put", "patch", "delete", "options", "head", "trace":
				path, err := setOperation(op, attribute.Name)
				if err != nil {
					return err
				}
				item.SetOperation(attribute.Name, op)
				if i, ok := api.Paths.Vals[path]; ok {
					api.Paths.Vals[path] = i.Merge(item)
				} else {
					api.Paths.Vals[path] = item
				}
			case "resource":
				resource, _ := lang.GetStringAttribute(method.Attributes, "http.resource")
				if len(resource) > 0 {
					op.Tags = append(op.Tags, resource)
					setSourceTag = true
				}
			}
		}
	}

	if !setSourceTag {
		if attributes, ok := ctx.Value("interface.attributes").([]*lang.Attribute); ok {
			resource, _ := lang.GetStringAttribute(attributes, "http.resource")
			if len(resource) > 0 {
				op.Tags = append(op.Tags, resource)
			}
		}
	}

	return nil
}

func compileDescription(ctx context.Context, method *lang.FunctionDecl) (string, *openapi.CachedDocument) {
	_ = ctx
	sd := method.Document.Parse().GetStructured()
	md := markdown.New()
	summary := ""
	description := ""

	if sd != nil {
		if len(sd.Blocks) > 0 {
			summary, _ = md.RenderBlocksToString(sd.Blocks[0])
		}
		if len(sd.Blocks) > 1 {
			description, _ = md.RenderBlocksToString(sd.Blocks[1:]...)
		}
	}

	return summary, &openapi.CachedDocument{Cache: description}
}

func httpCanCarryBody(method string) bool {
	switch method {
	case "post", "put", "patch":
		return true
	}
	return false
}

func compileParameter(ctx context.Context, decl *lang.ValueDecl, method *lang.FunctionDecl, httpMethod string) (*openapi.Parameter, error) {
	parameter := &openapi.Parameter{
		Name: decl.Name,
	}

	parameter.Description = decl.Document.GetContent()
	parameter.Description = strings.TrimSpace(parameter.Description)

	// enum type

	schema, err := compileNominalType(ctx, decl.Type)
	if err != nil {
		return nil, err
	}
	parameter.Schema = schema

	if v, e := decl.GetBoolAttribute(core.RequiredAttributeName); e == nil && v {
		parameter.Required = true
	}
	if v, e := decl.GetBoolAttribute("deprecated"); e == nil && v {
		parameter.Deprecated = true
	}

	if pathParams, ok := ctx.Value("pathParams").(map[string]bool); ok && pathParams[decl.Name] {
		parameter.In = openapi.Parameter_LOCATION_PATH
	} else if unwrappedField, ok := ctx.Value("unwrappedField").(string); ok && len(unwrappedField) > 0 {
		fullName := unwrappedField + "." + decl.Name
		parameter.Name = unwrappedField + "_" + decl.Name
		if pathParams[fullName] {
			parameter.In = openapi.Parameter_LOCATION_PATH
			// TODO the required param should get from ValueDecl from Context
			parameter.Required = true
		} else {
			parameter.In = openapi.Parameter_LOCATION_QUERY
		}
	} else {
		parameter.In = openapi.Parameter_LOCATION_QUERY
	}

	if len(method.Signature.GetParameters()) == 1 &&
		parameter.In != openapi.Parameter_LOCATION_PATH &&
		httpCanCarryBody(httpMethod) &&
		!decl.GetType().IsScalar() {

		decl.SetBoolAttribute("http.body", true)
		return nil, nil
	}

	if decl.HasAttribute("http.body") {
		return nil, nil
	}

	return parameter, nil
}

func CompileParameters(ctx context.Context, method *lang.FunctionDecl, httpMethod string) ([]*openapi.ReferenceableParameter, error) {
	thisCtx := context.WithType(ctx, method)
	var parameters []*openapi.ReferenceableParameter

	params := method.Signature.GetParameters()
	if pagination, _ := lang.GetBoolAttribute(method.Attributes, core.PaginationAttributeName); pagination {
		params = append(params, langcompiler.PaginationRequestFields()...)
	}

	for _, param := range params {
		if param.HasAttribute(http.BodyAttributeFullName) {
			continue
		} else if v, err := param.GetBoolAttribute("unwrap"); err == nil && v {

			if decl := param.Type.GetTypeDeclaration().GetStructDecl(); decl != nil {
				declCtx := context.WithValues(thisCtx, "unwrappedField", param.Name)
				for _, field := range decl.GetType().GetFields() {
					parameter, err := compileParameter(declCtx, field, method, httpMethod)
					if err != nil {
						return nil, err
					}
					if parameter != nil {
						parameters = append(parameters, openapi.NewReferenceableParameter(parameter))
					}
				}
			}
		} else {
			parameter, err := compileParameter(ctx, param, method, httpMethod)
			if err != nil {
				return nil, err
			}
			if parameter != nil {
				parameters = append(parameters, openapi.NewReferenceableParameter(parameter))
			}
		}
	}
	return parameters, nil
}

// must call after compile parameter
func compileRequestBody(ctx context.Context, method *lang.FunctionDecl) *openapi.ReferenceableRequestBody {
	thisCtx := context.WithType(ctx, method)
	var body *lang.ValueDecl
	for _, parameter := range method.Signature.GetParameters() {
		if parameter.HasAttribute(http.BodyAttributeFullName) {
			body = parameter
		}
	}

	if body != nil {
		schema, err := compileNominalType(thisCtx, body.GetType())
		if err != nil {
			return nil
		}

		if style, _ := body.GetStringAttribute(http.StyleAttributeFullName); len(style) > 0 {
			if style == "raw" {
				schema.GetSchema().Format = "Bytes"
			}
		}

		requestBody := &openapi.RequestBody{
			Description: nil,
			Content:     nil,
			Required:    true,
		}

		requestBody.Content = map[string]*openapi.MediaType{
			core.ApplicationJson: {Schema: schema},
		}
		return openapi.NewReferenceableRequestBody(requestBody)
	}

	return nil
}

func compileResponses(ctx context.Context, method *lang.FunctionDecl) *openapi.Responses {
	responses := &openapi.Responses{Vals: map[string]*openapi.ReferenceableResponse{
		"200": openapi.NewReferenceableResponse(&openapi.Response{
			Description: "OK",
			Headers:     nil,
			Content:     nil,
		}),
	}}

	if method.Signature.Result != nil {
		schema, err := compileNominalType(ctx, method.Signature.GetResultType())
		if err != nil {
			return nil
		}
		responses.Vals["200"].GetResponse().Content = map[string]*openapi.MediaType{
			core.ApplicationJson: {Schema: schema},
		}
	}

	return responses
}

func CompilePath(path string) (string, map[string]bool) {
	ts := core.NewTemplateString(path)
	p := bytes.Buffer{}
	parameters := make(map[string]bool)
	for _, segment := range ts.Segments {
		if segment.Templated {
			p.WriteByte('{')
			p.WriteString(strings.ReplaceAll(segment.Content, ".", "_"))
			p.WriteByte('}')

			parameters[segment.Content] = true
		} else {
			p.WriteString(segment.Content)
		}
	}
	return p.String(), parameters
}
