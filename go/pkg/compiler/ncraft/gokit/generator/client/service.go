package client

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
)

type Import struct{ Alias, Path string }

// PrepareService uses per-service type aliases; client output must not depend
// on the global short-name registry populated while compiling other services.
func PrepareService(service *data.Service) error {
	if service == nil || service.Interface == nil || service.Go == nil || service.Go.ApiImportPath == "" {
		return fmt.Errorf("client generation requires a service and its Go API import")
	}
	if service.Extensions == nil {
		service.Extensions = make(map[string]interface{})
	}
	service.Extensions["ClientPackage"] = strcase.ToSnake(service.Interface.BaredName) + "_client"
	paths := make(map[string]bool)
	for _, method := range service.Interface.Methods {
		for _, msg := range []*data.Message{method.Request, method.Response} {
			if msg == nil || msg.Go == nil || msg.Go.ImportPath == "" || msg.Decl == nil {
				return fmt.Errorf("client %s.%s: missing request or response Go type", service.Interface.Name, method.Name)
			}
			paths[msg.Go.ImportPath] = true
		}
	}
	var sorted []string
	for p := range paths {
		sorted = append(sorted, p)
	}
	sort.Strings(sorted)
	aliases := map[string]string{service.Go.ApiImportPath: "pb"}
	var imports, external []Import
	for _, p := range sorted {
		alias := aliases[p]
		if alias == "" {
			alias = fmt.Sprintf("api%d", len(external))
			aliases[p] = alias
			external = append(external, Import{alias, p})
		}
		imports = append(imports, Import{alias, p})
	}
	service.Extensions["ClientImports"] = imports
	service.Extensions["ClientExternalImports"] = external
	for _, method := range service.Interface.Methods {
		if method.Extensions == nil {
			method.Extensions = make(map[string]interface{})
		}
		method.Extensions["ClientRequestType"] = aliases[method.Request.Go.ImportPath] + "." + lang.GetTypeGoTypeName(method.Request.Decl.GetFullName())
		method.Extensions["ClientResponseType"] = aliases[method.Response.Go.ImportPath] + "." + lang.GetTypeGoTypeName(method.Response.Decl.GetFullName())
		for _, binding := range method.Bindings {
			if binding.Extensions == nil {
				binding.Extensions = make(map[string]interface{})
			}
			if binding.Body != nil {
				binding.Extensions["ClientBodyAccessor"] = accessor(binding.Body.Field)
			}
			for _, param := range binding.Parameters {
				if param.Extensions == nil {
					param.Extensions = make(map[string]interface{})
				}
				param.Extensions["ClientAccessor"] = accessor(param.Field)
				skip := param.Field.Exploded
				if binding.Body != nil {
					body := binding.Body.Field.FullName
					if body == "" {
						body = binding.Body.Field.Name
					}
					field := param.Field.FullName
					if field == "" {
						field = param.Field.Name
					}
					skip = skip || field == body || strings.HasPrefix(field, body+".")
				}
				param.Extensions["ClientSkipQuery"] = skip
			}
		}
	}
	return nil
}

func accessor(field *data.Field) string {
	name := field.FullName
	if name == "" {
		name = field.Name
	}
	parts := strings.Split(name, ".")
	for i, p := range parts {
		parts[i] = "Get" + strcase.ToCamel(p) + "()"
	}
	return "req." + strings.Join(parts, ".")
}
