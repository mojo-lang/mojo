package compiler

import (
	"sort"
	"strings"
	"text/template"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/http/go/pkg/mojo/http"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/compiler/transformer"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/openapi/generator/compiler"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/decompiler"
)

type Services struct {
	Data []*data.Service

	Interfaces []*data.Interface
}

func CompilePackage(ctx context.Context, pkg *lang.Package) ([]*data.Service, error) {
	services := &Services{}
	if err := plugin.CompilePackage(services, ctx, pkg); err != nil {
		return nil, err
	}

	for _, s := range services.Data {
		s.AllInterfaces = services.Interfaces
	}

	return services.Data, nil
}

func (s *Services) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	thisCtx := context.WithType(ctx, decl)
	pkg := context.Package(ctx)

	pkgName := lang.GetGoPackageName(decl.PackageName)
	service := &data.Service{
		PackageName:     pkg.Name,
		PackageFullName: pkg.FullName,
		Interface: &data.Interface{
			Decl:       decl,
			Name:       decl.Name,
			BaredName:  data.GetInterfaceBaredName(decl.Name),
			ServerName: data.GetInterfaceServerName(decl.Name),
			Extensions: make(map[string]interface{}),
		},
		FuncMap: template.FuncMap{
			"ToLower":       strings.ToLower,
			"ToUpper":       strings.ToUpper,
			"GoName":        GoName,
			"ToSnake":       strcase.ToSnake,
			"ToKebab":       strcase.ToKebab,
			"ToCamel":       strcase.ToCamel,
			"ToLowerCamel":  strcase.ToLowerCamel,
			"Plural":        transformer.Plural,
			"GoPackageName": GoPackageName,
		},
		Go: &data.GoService{
			PackageName: pkgName,
		},
		Extensions: make(map[string]interface{}),
	}

	if pkg.Version != nil {
		service.Version = pkg.Version.Format()
	}

	if path, ok := GoPackageImport(ctx, decl.PackageName).(string); ok {
		service.Go.ApiImportPath = path
	}

	for _, d := range decl.Type.Methods {
		if err := s.CompileMethod(thisCtx, d, service); err != nil {
			return err
		}
	}

	for _, m := range service.Interface.Methods {
		if len(m.Subscriptions) > 0 {
			service.Interface.HasSubscription = true
			break
		}
	}

	service.Go.ImportedTypePaths = unifyStringArray(service.Go.ImportedTypePaths)

	s.Data = append(s.Data, service)
	s.Interfaces = append(s.Interfaces, service.Interface)

	return nil
}

func unifyStringArray(array []string) []string {
	index := make(map[string]bool)
	for _, item := range array {
		if _, ok := index[item]; !ok {
			index[item] = true
		}
	}
	var result []string
	for key := range index {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}

func (s *Services) CompileMethod(ctx context.Context, decl *lang.FunctionDecl, service *data.Service) error {
	thisCtx := context.WithType(ctx, decl)
	m := &data.Method{
		Decl:       decl,
		Name:       decl.Name,
		Extensions: make(map[string]interface{}),
	}
	registerType := func(t *lang.NominalType) {
		if !t.IsScalar() && !t.IsMapType() && !t.IsArrayType() && !t.IsUnionType() && (len(t.PackageName) > 0 && t.PackageName != service.PackageFullName) {
			goPackageName := lang.GetGoPackageName(t.PackageName)
			if RegisterMessageGoPackageName(service.Interface.Name, t.Name, goPackageName) {
				return
			}

			if len(t.GenericArguments) == 0 {
				if t.TypeDeclaration != nil && t.TypeDeclaration.GetEnumDecl() != nil {
					service.ImportedEnums = append(service.ImportedEnums, &data.Enum{
						Decl:        t.GetTypeDeclaration().GetEnumDecl(),
						Name:        t.Name,
						PackageName: t.PackageName,
						Go: &data.GoEnum{
							PackageName: goPackageName,
						},
						Extensions: make(map[string]interface{}),
					})
				} else {
					service.ImportedMessages = append(service.ImportedMessages, &data.Message{
						Decl:        t.GetTypeDeclaration().GetStructDecl(),
						Name:        t.Name,
						PackageName: t.PackageName,
						Go: &data.GoMessage{
							PackageName: goPackageName,
						},
						Extensions: make(map[string]interface{}),
					})
				}

				if path, ok := GoPackageImport(ctx, t.PackageName).(string); ok {
					service.Go.ImportedTypePaths = append(service.Go.ImportedTypePaths, path)
				}
			}
		}
	}

	req, resp, err := decompiler.CompileMethod(ctx, decl)
	if err != nil {
		return err
	}
	if m.Request, err = s.CompileMessage(thisCtx, req); err != nil {
		return err
	}
	for _, field := range req.Type.Fields {
		t := field.Type
		if t.IsMapType() {
			for _, gt := range t.GenericArguments {
				registerType(gt)
			}
		} else if !t.IsScalar() {
			registerType(t)
		}
	}

	result := decl.GetSignature().GetResultType()
	if result.GetTypeDeclaration().GetStructDecl() != resp {
		result = &lang.NominalType{
			PackageName:     resp.PackageName,
			Name:            resp.Name,
			TypeDeclaration: lang.NewStructTypeDeclaration(resp),
		}
	}
	if m.Response, err = s.CompileMessage(thisCtx, resp); err != nil {
		return err
	}

	var httpResponse *data.HTTPResponse
	if result != nil {
		if v, e := result.GetStringAttribute(http.BodyAttributeFullName); e == nil && len(v) > 0 {
			httpResponse = &data.HTTPResponse{
				Headers: make(map[string]string),
			}

			for _, attribute := range result.Attributes {
				if attribute.IsSameName(http.HeaderAttributeFullName) {
					if len(attribute.Arguments) == 2 {
						key, _ := attribute.Arguments[0].GetString()
						value, _ := attribute.Arguments[1].GetString()
						if len(key) > 0 && len(value) > 0 {
							httpResponse.Headers[key] = value
						}
					}
				}
			}

			for _, field := range m.Response.Fields {
				if strings.ToLower(v) == strings.ToLower(field.Name) {
					httpResponse.Body = field
				}
			}
		}

		registerType(result)
	}

	// if len(decl.PackageName) > 0 && decl.PackageName != service.PackageFullName {
	//    goPackageName := lang.GetGoPackageName(decl.PackageName)
	//    if !RegisterMessageGoPackageName(decl.Name, goPackageName) && result != nil && len(result.GenericArguments) == 0 {
	//        if result.TypeDeclaration != nil && result.TypeDeclaration.GetEnumDecl() != nil {
	//            service.ImportedEnums = append(service.ImportedEnums, &data.Enum{
	//                Decl:        t.GetTypeDeclaration().GetEnumDecl(),
	//                Name:        t.Name,
	//                PackageName: t.PackageName,
	//                Go: data.GoEnum{
	//                    PackageName: goPackageName,
	//                },
	//            })
	//
	//            service.ImportedEnums = append(service.ImportedEnums, &data.Enum{
	//                Name: decl.Name,
	//            })
	//        } else {
	//            service.ImportedMessages = append(service.ImportedMessages, &data.Message{Name: decl.Name})
	//        }
	//
	//        if path, ok := GoPackageImport(ctx, decl.PackageName).(string); ok {
	//            service.Go.ImportedTypePaths = append(service.Go.ImportedTypePaths, path)
	//        }
	//    }
	// }

	index := 0
	for _, attribute := range decl.Attributes {
		if attribute.PackageName != "http" {
			continue
		}

		switch attribute.Name {
		case http.GetAttributeName, http.PostAttributeName, http.PutAttributeName, http.PatchAttributeName,
			http.DeleteAttributeName, http.OptionsAttributeName, http.HeadAttributeName, http.TraceAttributeName:
			if len(attribute.Arguments) > 0 {
				if value, err := attribute.Arguments[0].GetString(); err == nil && len(value) > 0 {
					thisCtx = WithDataMethod(context.WithValues(ctx, "index", index), m)
					binding, err := s.compileBindings(thisCtx, attribute.Name, value, decl)
					if err != nil {
						continue
					}

					sort.Sort(binding.Parameters)

					m.Bindings = append(m.Bindings, binding)
					index++
				}
			}
		}
	}

	for _, attribute := range decl.Attributes {
		if attribute.IsSameName("messaging.subscription") {
			subscription := &data.MessagingSubscription{}
			if len(attribute.Arguments) == 1 {
				if value, err := attribute.Arguments[0].GetString(); err == nil && len(value) > 0 {
					subscription.Topic = value
				}
			}
			if len(subscription.Topic) == 0 {
				subscription.Topic = m.Name
			}

			m.Subscriptions = append(m.Subscriptions, subscription)
		}
	}

	service.Interface.Methods = append(service.Interface.Methods, m)
	return nil
}

func (s *Services) CompileMessage(ctx context.Context, decl *lang.StructDecl) (*data.Message, error) {
	_ = ctx
	msg := &data.Message{
		Decl:        decl,
		PackageName: decl.PackageName,
		Name:        decl.Name,
		Go:          &data.GoMessage{},
		Extensions:  make(map[string]interface{}),
	}

	err := decl.EachField(func(f *lang.ValueDecl) error {
		t := f.GetType()
		fieldType := &data.FieldType{
			Go:         &data.GoFieldType{},
			Extensions: make(map[string]interface{}),
		}

		switch t.GetFullName() {
		case core.ArrayTypeFullName:
			if len(t.GenericArguments) == 1 {
				fieldType.Name = t.GenericArguments[0].Name
				fieldType.IsArray = true
			}
		case core.MapTypeFullName:
			if len(t.GenericArguments) == 2 {
				fieldType.KeyType = &data.FieldType{
					Name: t.GenericArguments[0].Name,
				}
				fieldType.Name = t.GenericArguments[1].Name
				fieldType.IsMap = true
			}
		default:
			fieldType.Name = t.Name
		}

		if t.TypeDeclaration.GetStructDecl() != nil {
			fieldType.Go.IsPointer = true
		}

		msg.Fields = append(msg.Fields, &data.Field{
			Name: f.Name,
			Type: fieldType,
		})

		return nil
	})
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *Services) compileBindings(ctx context.Context, methodName string, path string, method *lang.FunctionDecl) (*data.HTTPBinding, error) {
	dm := DataMethod(ctx)
	p, pathParams := compiler.CompilePath(path)

	index := 0
	if value, ok := ctx.Value("index").(int); ok {
		index = value
	}

	binding := &data.HTTPBinding{
		Verb:       methodName,
		Path:       p,
		Label:      strcase.ToCamel(method.Name) + data.EnglishNumber(index),
		Parent:     dm,
		Extensions: make(map[string]interface{}),
	}

	if len(method.Signature.GetParameters()) == 1 {
		decl := method.Signature.ParameterDecl(0)
		if b, _ := decl.GetBoolAttribute(protobuf.MethodRequestTypeAttributeName); b {
			if structDecl := decl.Type.GetTypeDeclaration().GetStructDecl(); structDecl != nil {
				err := structDecl.EachField(func(decl *lang.ValueDecl) error {
					return s.compileBindingParameter(ctx, decl, pathParams, nil, binding)
				})
				if err != nil {
					return nil, err
				}
				return binding, nil
			} else {
				return nil, logs.NewErrorw("invalid request type", "method", method.FullName, "http-bind", methodName+":"+path)
			}
		}
	}

	for _, param := range method.Signature.GetParameters() {
		if err := s.compileBindingParameter(ctx, param, pathParams, nil, binding); err != nil {
			return nil, err
		}
	}

	if binding.Body == nil && http.CanCarryBody(binding.Verb) {
		for _, parameter := range binding.Parameters {
			goTypeName := parameter.GetGoTypeName()
			if parameter.Location == "query" && len(GoPackageName(goTypeName)) > 0 &&
				!(strings.HasPrefix(goTypeName, "map[") || strings.HasPrefix(goTypeName, "[]")) {
				binding.Body = parameter
				break
			}
		}
	}

	return binding, nil
}

func (s *Services) compileBindingParameter(ctx context.Context, decl *lang.ValueDecl, pathParams map[string]bool, enclosing *data.HTTPParameter, binding *data.HTTPBinding) error {
	c := &compiler.NominalTypeCompiler{}

	schema, err := c.Compile(decl.Type)
	if err != nil {
		return err
	}

	param := &data.HTTPParameter{
		Name: decl.Name,

		Field: s.compileBindingField(ctx, schema, context.OpenAPIComponents(c.Context).GetSchemas()),

		Go:         &data.GoHTTPParameter{},
		Extensions: make(map[string]interface{}),
	}

	if param.Field == nil {
		return logs.NewErrorw("failed to compile the binding field")
	}

	param.Style, _ = decl.GetStringAttribute(http.StyleAttributeFullName)
	param.Field.Name = decl.Name

	if enclosing != nil {
		param.Field.Enclosing = enclosing.Field
		param.Field.FullName = enclosing.Field.FullName + "." + decl.Name
	} else {
		param.Field.FullName = decl.Name
	}
	param.FullName = strcase.ToSnake(strings.ReplaceAll(param.Field.FullName, ".", "_"))

	if v, e := lang.GetStringAttribute(decl.Type.Attributes, core.ExplodedAttributeFullName); e == nil {
		param.Field.Exploded = true
		param.Field.ExplodePrefix = v
		if len(v) == 0 {
			param.Field.ExplodePrefix = decl.Name
		}

		structDecl := decl.Type.GetTypeDeclaration().GetStructDecl()
		if len(decl.Type.GenericArguments) == 0 && structDecl != nil {
			err = structDecl.EachField(func(decl *lang.ValueDecl) error {
				if err = s.compileBindingParameter(ctx, decl, pathParams, param, binding); err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	// TODO update sync with openapi, add request body
	if pathParams[param.Field.Name] || pathParams[param.Field.FullName] {
		param.Location = "path"
	} else {
		param.Location = "query"
	}

	if body := decl.Type.HasAttribute(http.BodyAttributeFullName); body {
		binding.Body = param
	}
	binding.Parameters = append(binding.Parameters, param)

	if fieldType := param.Field.GetType(); fieldType != nil && fieldType.Message != nil {
		for pathParam := range pathParams {
			if strings.Contains(pathParam, ".") {
				for _, field := range fieldType.Message.Fields {
					fullName := param.Field.Name + "." + field.Name
					if pathParam != fullName {
						fullName = param.Field.Name + "." + strcase.ToSnake(field.Name)
					}

					if pathParam == fullName {
						parameter := &data.HTTPParameter{
							Name:     field.Name,
							FullName: HttpPathParameterName(fullName),
							Field: &data.Field{
								Name:      field.Name,
								FullName:  fullName,
								Type:      field.Type,
								Enclosing: param.Field,
								Decl:      field.Decl,
							},
							Location:   "path",
							Go:         &data.GoHTTPParameter{},
							Extensions: make(map[string]interface{}),
						}
						binding.Parameters = append(binding.Parameters, parameter)
					}
				}
			}
		}
	}
	return nil
}

func (s *Services) compileBindingField(ctx context.Context, schema *openapi.Schema, index map[string]*openapi.Schema) *data.Field {
	if schema == nil {
		// error
		return nil
	}

	field := &data.Field{}
	switch schema.Type {
	case openapi.Schema_TYPE_BOOLEAN:
		if schema.Title == core.BoolValueTypeFullName {
			field.Type = &data.FieldType{
				Name:        "BoolValue",
				PackageName: "mojo.core",
				Go: &data.GoFieldType{
					Name:      "core.BoolValue",
					IsPointer: true,
				},
			}
		} else {
			field.Type = &data.FieldType{
				Name:        "Bool",
				PackageName: "mojo.core",
				IsScalar:    true,
				Go: &data.GoFieldType{
					Name:      "bool",
					IsPointer: false,
				},
			}
		}
	case openapi.Schema_TYPE_INTEGER:
		if len(schema.Format) > 0 {
			switch schema.Format {
			case "Int8", "Int16":
				field.Type = &data.FieldType{
					Name: "Int32",
				}
			default:
				field.Type = &data.FieldType{
					Name: schema.Format,
				}
			}
			field.Type.IsScalar = true
			field.Type.Go = &data.GoFieldType{
				Name:      strings.ToLower(field.Type.Name),
				IsPointer: false,
			}
		} else {
			field.Type = &data.FieldType{
				Name:     "Int64",
				IsScalar: true,
				Go: &data.GoFieldType{
					Name: "int64",
				},
			}
		}
	case openapi.Schema_TYPE_NUMBER:
		field.Type = &data.FieldType{
			Name:     "Float64",
			IsScalar: true,
			Go: &data.GoFieldType{
				Name: "float64",
			},
		}
	case openapi.Schema_TYPE_STRING:
		if len(schema.Enum) > 0 {
			field.Type = &data.FieldType{
				PackageName: lang.GetTypePackageName(schema.Title),
				Name:        lang.GetTypeTypeName(schema.Title),
				IsEnum:      true,
				Go: &data.GoFieldType{
					Name: getGoTypeName(schema.Title),
				},
			}
		} else {
			field.Type = &data.FieldType{
				Name:     "String",
				IsScalar: true,
				Go: &data.GoFieldType{
					Name: "string",
				},
			}
		}
	case openapi.Schema_TYPE_ARRAY:
		if schema.Title == core.ValuesTypeFullName {
			field.Type = &data.FieldType{
				Name:        core.ValuesTypeName,
				PackageName: "mojo.core",
				Go: &data.GoFieldType{
					Name:      "core.Values",
					IsPointer: true,
				},
				Extensions: make(map[string]interface{}),
			}
		} else {
			sch := schema.Items.GetSchemaOf(index)
			f := s.compileBindingField(ctx, sch, index)
			if f != nil {
				field.Type = &data.FieldType{
					Name:    "Array<" + f.Type.Name + ">",
					IsArray: true,
					ElementType: &data.FieldType{
						Name:        f.Type.Name,
						PackageName: f.Type.PackageName,
						Enum:        f.Type.Enum,
						Message:     f.Type.Message,
						KeyType:     f.Type.KeyType,
						ElementType: f.Type.ElementType,
						IsEnum:      f.Type.IsEnum,
						IsMap:       f.Type.IsMap,
						IsArray:     f.Type.IsArray,
						IsScalar:    f.Type.IsScalar,
						Go:          f.Type.Go,
						Extensions:  f.Type.Extensions,
					},
					Go:         &data.GoFieldType{},
					Extensions: make(map[string]interface{}),
				}

				if f.Type.Go.IsPointer {
					field.Type.Go.Name = "[]*" + f.Type.Go.Name
				} else {
					field.Type.Go.Name = "[]" + f.Type.Go.Name
				}
			} else {
				logs.Errorw("failed to compile the schema", "schema", schema.Title)
			}
		}
	case openapi.Schema_TYPE_OBJECT:
		if schema.AdditionalProperties != nil { // map
			typ := s.compileBindingField(ctx, schema.GetAdditionalProperties().GetSchemaOf(index), index).GetType()
			field.Type = &data.FieldType{
				Name:  "Map<String," + typ.Name + ">",
				IsMap: true,
				KeyType: &data.FieldType{
					Name:     "String",
					IsScalar: true,
					Go: &data.GoFieldType{
						Name: "string",
					},
				},
				ElementType: &data.FieldType{
					Name:        typ.Name,
					PackageName: typ.PackageName,
					Enum:        typ.Enum,
					Message:     typ.Message,
					KeyType:     typ.KeyType,
					ElementType: typ.ElementType,
					IsEnum:      typ.IsEnum,
					IsMap:       typ.IsMap,
					IsArray:     typ.IsArray,
					IsScalar:    typ.IsScalar,
					Go:          typ.Go,
					Extensions:  typ.Extensions,
				},
				Go:         &data.GoFieldType{},
				Extensions: make(map[string]interface{}),
			}
			if field.Type.Go.IsPointer {
				field.Type.Go.Name = "map[string]*" + typ.Go.Name
			} else {
				field.Type.Go.Name = "map[string]" + typ.Go.Name
			}
		} else {
			field.Type = &data.FieldType{
				PackageName: lang.GetTypePackageName(schema.Title),
				Name:        lang.GetTypeTypeName(schema.Title),
				IsScalar:    false,
				Message: &data.Message{
					PackageName: lang.GetTypePackageName(schema.Title),
					Name:        lang.GetTypeTypeName(schema.Title),
				},
				Go: &data.GoFieldType{
					Name:      getGoTypeName(schema.Title),
					IsPointer: true,
				},
				Extensions: make(map[string]interface{}),
			}

			for name, item := range schema.Properties {
				f := s.compileBindingField(ctx, item.GetSchemaOf(index), index)
				if f == nil {
					println("")
				}
				f.Name = strcase.ToSnake(name)
				field.Type.Message.Fields = append(field.Type.Message.Fields, f)
			}
		}
	}

	// FIXME add the well known types to here
	switch schema.Title {
	case core.ValueTypeFullName:
		field.Type = &data.FieldType{
			Name:        core.ValueTypeName,
			PackageName: "mojo.core",
			Go: &data.GoFieldType{
				Name:      "core.Value",
				IsPointer: true,
			},
			Extensions: make(map[string]interface{}),
		}
	case core.DomainTypeFullName:
		field.Type = &data.FieldType{
			Name:        core.DomainTypeName,
			PackageName: "mojo.core",
			Go: &data.GoFieldType{
				Name:      "core.Domain",
				IsPointer: true,
			},
			Extensions: make(map[string]interface{}),
		}
	case core.EmailAddressTypeFullName:
		field.Type = &data.FieldType{
			Name:        core.EmailAddressTypeName,
			PackageName: "mojo.core",
			Go: &data.GoFieldType{
				Name:      "core.EmailAddress",
				IsPointer: true,
			},
			Extensions: make(map[string]interface{}),
		}
	case core.FieldMaskTypeFullName:
		field.Type = &data.FieldType{
			Name:        core.FieldMaskTypeName,
			PackageName: "mojo.core",
			Go: &data.GoFieldType{
				Name:      "core.FieldMask",
				IsPointer: true,
			},
			Extensions: make(map[string]interface{}),
		}
	case core.MediaTypeTypeFullName:
		field.Type = &data.FieldType{
			Name:        core.MediaTypeTypeName,
			PackageName: "mojo.core",
			Go: &data.GoFieldType{
				Name:      "core.MediaType",
				IsPointer: true,
			},
			Extensions: make(map[string]interface{}),
		}
	case core.OrderingTypeFullName:
		field.Type = &data.FieldType{
			Name:        core.OrderingTypeName,
			PackageName: "mojo.core",
			Go: &data.GoFieldType{
				Name:      "core.Ordering",
				IsPointer: true,
			},
			Extensions: make(map[string]interface{}),
		}
	case core.PlatformTypeFullName:
		field.Type = &data.FieldType{
			Name:        core.PlatformTypeName,
			PackageName: "mojo.core",
			Go: &data.GoFieldType{
				Name:      "core.Platform",
				IsPointer: true,
			},
			Extensions: make(map[string]interface{}),
		}
	case core.TimestampTypeFullName:
		field.Type = &data.FieldType{
			Name:        core.TimestampTypeName,
			PackageName: "mojo.core",
			Go: &data.GoFieldType{
				Name:      "core.Timestamp",
				IsPointer: true,
			},
			Extensions: make(map[string]interface{}),
		}
	case core.UrlTypeFullName:
		field.Type = &data.FieldType{
			Name:        core.UrlTypeName,
			PackageName: "mojo.core",
			Go: &data.GoFieldType{
				Name:      "core.Url",
				IsPointer: true,
			},
			Extensions: make(map[string]interface{}),
		}
	case core.VersionTypeFullName:
		field.Type = &data.FieldType{
			Name:        core.VersionTypeName,
			PackageName: "mojo.core",
			Go: &data.GoFieldType{
				Name:      "core.Version",
				IsPointer: true,
			},
			Extensions: make(map[string]interface{}),
		}
	}

	if field.Type != nil {
		return field
	}
	return nil
}

func getGoTypeName(title string) string {
	pkg := lang.GetTypeGoPackageName(title)
	name := lang.GetTypeGoTypeName(title)
	if len(pkg) > 0 {
		return strings.Join([]string{pkg, name}, ".")
	}
	return name
}
