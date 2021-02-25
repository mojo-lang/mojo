package compiler

import (
	"bytes"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/document/go/pkg/markdown"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
	"strings"
)

var methods = []string{"http.get", "http.post", "http.put", "http.delete", "http.patch", "http.options", "http.head", "http.trace"}

func CompileInterface(ctx *Context, decl *lang.InterfaceDecl) (*openapi.OpenAPI, error) {
	api := openapi.NewOpenApi()
	api.Openapi = &core.Version{
		Major: 3,
		Minor: 0,
		Patch: 3,
	}

	var version *core.Version
	pkg := ctx.GetPackage()
	if pkg != nil {
		version = pkg.Version
	}

	title := ""
	document := decl.Document.Parse().GetStructured()
	if document != nil && len(document.Blocks) > 0 {
		if header := document.Blocks[0].GetHeader(); header != nil {
			h, err := markdown.New().RenderInlinesToString(header.Text...)
			if err != nil {
				logs.Warnw("failed to render header text", "error", err.Error())
			} else {
				title = h
			}
		}
	}

	api.Info = &openapi.Info{
		Title:          title,
		Description:    nil,
		TermsOfService: nil,
		Contact:        nil,
		License:        nil,
		Version:        version,
	}

	for _, method := range decl.Type.Methods {
		err := compileMethod(ctx, method, api)
		if err != nil {
			return nil, err
		}
	}
	return api, nil
}

func compileMethod(ctx *Context, method *lang.FunctionDecl, api *openapi.OpenAPI) error {
	for _, name := range methods {
		path, _ := lang.GetStringAttribute(method.Attributes, name)
		if len(path) == 0 {
			continue
		}

		var tags []string
		resource, _ := lang.GetStringAttribute(method.Attributes, "resource")
		if len(resource) > 0 {
			tags = append(tags, resource)
		}

		paths, pathParams := compilePath(path)
		for i, p := range paths {
			ctx.SetOption("pathParams", pathParams[i])
			item := &openapi.PathItem{}
			summary, description := compileDescription(ctx, method)

			parameters, err := compileParameters(ctx, method, item)
			if err != nil {
				return err
			}

			switch name {
			case "http.get":
				item.Get = &openapi.Operation{
					Summary:     summary,
					Description: description,
					OperationId: method.Name,
					Tags:        tags,
					Parameters:  parameters,
					Responses:   compileResponses(ctx, method, item),
				}
			case "http.post", "http.put", "http.patch":
				item.Post = &openapi.Operation{
					Summary:     summary,
					Description: description,
					OperationId: method.Name,
					Tags:        tags,
					Parameters:  parameters,
					RequestBody: nil,
					Responses:   compileResponses(ctx, method, item),
				}
			case "http.delete":
				item.Post = &openapi.Operation{
					Summary:     summary,
					Description: description,
					OperationId: method.Name,
					Tags:        tags,
					Parameters:  parameters,
				}
			case "http.options", "http.head", "http.trace":
				item.Post = &openapi.Operation{
					Summary:     summary,
					Description: description,
					OperationId: method.Name,
					Tags:        tags,
					Parameters:  parameters,
				}
			}
			api.Paths.Values[p] = item
		}
	}

	return nil
}

func compileDescription(ctx *Context, method *lang.FunctionDecl) (string, *openapi.CachedDocument) {
	sd := method.Document.Parse().GetStructured()
	md := markdown.New()
	summary := ""
	description := ""
	if len(sd.Blocks) > 0 {
		summary, _ = md.RenderBlocksToString(sd.Blocks[0])
	}
	if len(sd.Blocks) > 1 {
		description, _ = md.RenderBlocksToString(sd.Blocks[1:]...)
	}

	return summary, &openapi.CachedDocument{Cache: description}
}

func compileParameter(ctx *Context, decl *lang.ValueDecl) (*openapi.Parameter, error) {
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

	if v, e := lang.GetBoolAttribute(decl.Type.Attributes, "required"); e == nil && v {
		parameter.Required = true
	}
	if v, e := lang.GetBoolAttribute(decl.Type.Attributes, "deprecated"); e == nil && v {
		parameter.Deprecated = true
	}

	if pathParams, ok := ctx.GetOption("pathParams").(map[string]bool); ok && pathParams[decl.Name] {
		parameter.In = openapi.Parameter_LOCATION_PATH
	} else if unwrappedField, ok := ctx.GetOption("unwrappedField").(string); ok && len(unwrappedField) > 0 {
		fullName := unwrappedField + "." + decl.Name
		parameter.Name = unwrappedField + "_" + decl.Name
		if pathParams[fullName] {
			parameter.In = openapi.Parameter_LOCATION_PATH
			//TODO the required param should get from ValueDecl from context
			parameter.Required = true
		} else {
			parameter.In = openapi.Parameter_LOCATION_QUERY
		}
	} else {
		parameter.In = openapi.Parameter_LOCATION_QUERY
	}

	return parameter, nil
}

func compileParameters(ctx *Context, method *lang.FunctionDecl, item *openapi.PathItem) ([]*openapi.ReferenceableParameter, error) {
	var parameters []*openapi.ReferenceableParameter

	for _, param := range method.Signature.Parameters {
		if v, err := lang.GetBoolAttribute(param.Type.GetAttributes(), "unwrap"); err == nil && v {

			if decl := param.Type.GetTypeDeclaration().GetStructDecl(); decl != nil {
				ctx.SetOption("unwrappedField", param.Name)
				for _, field := range decl.GetType().GetFields() {
					parameter, err := compileParameter(ctx, field)
					if err != nil {
						return nil, err
					}
					parameters = append(parameters, openapi.NewReferenceableParameter(parameter))
				}
				ctx.DeleteOption("unwrappedField")
			}
		} else {
			parameter, err := compileParameter(ctx, param)
			if err != nil {
				return nil, err
			}
			parameters = append(parameters, openapi.NewReferenceableParameter(parameter))
		}
	}
	return parameters, nil
}

func compileResponses(ctx *Context, method *lang.FunctionDecl, item *openapi.PathItem) *openapi.Responses {
	responses := &openapi.Responses{Values: map[string]*openapi.ReferenceableResponse{
		"200": openapi.NewReferenceableResponse(&openapi.Response{
			Description: "OK",
			Headers:     nil,
			Content:     nil,
		}),
	}}

	if method.Signature.Result != nil {
		c := &ReferenceCompiler{}
		schema, err := c.Compile(ctx, method.Signature.Result)
		if err != nil {
			return nil
		}
		responses.Values["200"].GetResponse().Content = map[string]*openapi.MediaType{
			"application/json": {Schema: schema},
		}
	}

	return responses
}

func compilePath(path string) ([]string, []map[string]bool) {
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
	return []string{p.String()}, []map[string]bool{parameters}
}
