package client

import (
	"fmt"
	"go/token"
	"go/types"
	"path"
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
	names := map[string]string{service.Go.ApiImportPath: service.Go.PackageName}
	for _, method := range service.Interface.Methods {
		for _, msg := range []*data.Message{method.Request, method.Response} {
			if msg == nil || msg.Go == nil || msg.Go.ImportPath == "" || msg.Decl == nil {
				return fmt.Errorf("client %s.%s: missing request or response Go type", service.Interface.Name, method.Name)
			}
			paths[msg.Go.ImportPath] = true
			names[msg.Go.ImportPath] = msg.Go.PackageName
		}
	}
	var sorted []string
	for p := range names {
		sorted = append(sorted, p)
	}
	sort.Strings(sorted)
	aliases := packageAliases(sorted, names, service.Go.ApiImportPath)
	service.Extensions["ClientAPIAlias"] = aliases[service.Go.ApiImportPath]
	var imports, external []Import
	for _, p := range sorted {
		alias := aliases[p]
		if p != service.Go.ApiImportPath {
			external = append(external, Import{alias, p})
		}
		if paths[p] {
			imports = append(imports, Import{alias, p})
		}
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

// Reserve pb for the service's generated Protobuf/gRPC API, matching server
// code. Keep other native package names unless they conflict with generated code.
func packageAliases(paths []string, names map[string]string, apiPath string) map[string]string {
	used := map[string]bool{"pb": true}
	for _, name := range strings.Fields(`bytes context jsoniter fmt io http url reflect strings
		endpoint httptransport grpc metadata Client ClientOption Endpoints GrpcClient
		HTTPError FullServiceName clientConfig ctx request req response remote config
		endpoints options conn err c e r target body instance base clientOptions result ok client`) {
		used[name] = true
	}
	preferred := make(map[string]bool)
	for _, p := range paths {
		if p == apiPath {
			continue
		}
		name := names[p]
		if name == "" {
			name = strings.ReplaceAll(path.Base(p), "-", "_")
		}
		if !token.IsIdentifier(name) || name == "_" {
			name = "pkg"
		}
		names[p] = name
		preferred[name] = true
	}
	aliases := map[string]string{apiPath: "pb"}
	for _, p := range paths {
		if p == apiPath {
			continue
		}
		name := names[p]
		alias := name
		for suffix := 2; used[alias] || types.Universe.Lookup(alias) != nil; suffix++ {
			alias = fmt.Sprintf("%s%d", name, suffix)
			// Do not consume another dependency's native name.
			for preferred[alias] {
				suffix++
				alias = fmt.Sprintf("%s%d", name, suffix)
			}
		}
		aliases[p] = alias
		used[alias] = true
	}
	return aliases
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
