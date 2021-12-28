package compiler

import (
	"errors"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/types"
	apicompiler "github.com/mojo-lang/mojo/go/pkg/openapi/compiler"
	"github.com/mojo-lang/mojo/go/pkg/protobuf"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/compiler"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
	"strings"
)

func CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) (*types.Service, error) {
	thisCtx := context.WithType(ctx, decl)

	pkgName := ""
	pkgNames := strings.Split(decl.PackageName, ".")
	for i := len(pkgNames) - 1; i >= 0; i-- {
		if core.IsVersionTag(pkgNames[i]) {
			continue
		}
		pkgName = pkgNames[i]
		break
	}

	service := &types.Service{
		PkgName:     pkgName,
		FullPkgName: decl.PackageName,
		ImportPaths: nil,
		Messages:    nil,
		Enums:       nil,
		Interface: &types.Interface{
			Name:    decl.Name,
			Methods: nil,
		},
	}

	if path, ok := GoPackageImport(ctx, decl.PackageName).(string); ok {
		service.ApiImportPath = path
	}

	for _, method := range decl.Type.Methods {
		err := compileMethod(thisCtx, method, service)
		if err != nil {
			return nil, err
		}
	}
	return service, nil
}

func compileMessage(ctx context.Context, decl *lang.StructDecl) (*types.Message, error) {
	msg := &types.Message{
		Name: decl.Name,
	}

	for _, f := range decl.GetType().GetFields() {
		t := f.GetType()
		fieldType := &types.FieldType{}

		switch t.GetFullName() {
		case "mojo.core.Array":
			if len(t.GenericArguments) > 0 {
				fieldType.Name = t.GenericArguments[0].Name
				fieldType.ArrayType = true
			}
		case "mojo.core.Map":
		default:
			fieldType.Name = t.Name
		}

		msg.Fields = append(msg.Fields, &types.Field{
			Name: strcase.ToCamel(f.Name),
			Type: fieldType,
		})
	}

	return msg, nil
}

func compileMethod(ctx context.Context, method *lang.FunctionDecl, service *types.Service) error {
	thisCtx := context.WithType(ctx, method)

	m := &types.InterfaceMethod{
		Name:      method.Name,
		CamelName: strcase.ToCamel(method.Name),
		Bindings:  nil,
	}

	r := &lang.StructDecl{}
	r.Name = strcase.ToCamel(method.Name) + "Request"
	r.Type = &lang.StructType{
		Fields: method.Signature.Parameters,
	}

	var err error
	m.RequestType, err = compileMessage(ctx, r)
	if err != nil {
		return err
	}

	registerType := func(t *lang.NominalType) {
		if !t.IsScalar() && !t.IsMapType() && !t.IsArrayType() && !t.IsUnionType() {
			RegisterMessagePackage(t.Name, GetGoPackage(t.PackageName))

			if len(t.GenericArguments) == 0 {
				if t.TypeDeclaration != nil && t.TypeDeclaration.GetEnumDecl() != nil {
					service.ImportEnums = append(service.ImportEnums, t.Name)
				} else {
					service.ImportStructs = append(service.ImportStructs, t.Name)
				}

				if path, ok := GoPackageImport(ctx, t.PackageName).(string); ok {
					service.ImportPaths = append(service.ImportPaths, path)
				}
			}
		}
	}

	for _, field := range r.Type.Fields {
		t := field.Type
		if t.IsMapType() {
			for _, gt := range t.GenericArguments {
				registerType(gt)
			}
		} else if !t.IsScalar() {
			registerType(t)
		}
	}

	result := method.GetSignature().GetResult()
	var decl *lang.StructDecl
	if result != nil {
		decl = result.GetTypeDeclaration().GetStructDecl()

		if decl.GetFullName() == "mojo.core.Array" {
			pagination, _ := lang.GetBoolAttribute(method.Attributes, "pagination")
			if pagination {
				decl = protobuf.GeneratePaginationResponse(method)
			} else {
				decl, err = compiler.CompileArrayToStruct(result)
				if err != nil {
					return err
				}
			}
		}
	} else {
		decl = &lang.StructDecl{
			PackageName:    "mojo.core",
			SourceFileName: "mojo/core/null.mojo",
			Name:           "Null",
		}
		result = &lang.NominalType{
			PackageName:     "mojo.core",
			Name:            "Null",
			TypeDeclaration: lang.NewStructTypeDeclaration(decl),
		}
	}

	m.ResponseType, err = compileMessage(thisCtx, decl)
	if err != nil {
		return err
	}

	if result != nil {
		if v, e := lang.GetStringAttribute(result.Attributes, "http.body"); e == nil && len(v) > 0 {
			httpResponse := &types.HTTPResponse{
				BodyField: nil,
				Headers:   make(map[string]string),
			}

			for _, attribute := range result.Attributes {
				if attribute.SameName("http.header") {
					if len(attribute.Arguments) == 2 {
						key := attribute.Arguments[0].GetStringLiteralExpr().Value
						value := attribute.Arguments[1].GetStringLiteralExpr().Value
						httpResponse.Headers[key] = value
					}
				}
			}

			for _, field := range m.ResponseType.Fields {
				if strings.ToLower(v) == strings.ToLower(field.Name) {
					httpResponse.BodyField = field
				}
			}
			m.HTTPResponse = httpResponse
		}
	}

	if len(decl.PackageName) > 0 && decl.PackageName != service.FullPkgName {
		RegisterMessagePackage(decl.Name, GetGoPackage(decl.PackageName))

		if result != nil && len(result.GenericArguments) == 0 {
			if result.TypeDeclaration != nil && result.TypeDeclaration.GetEnumDecl() != nil {
				service.ImportEnums = append(service.ImportEnums, decl.Name)
			} else {
				service.ImportStructs = append(service.ImportStructs, decl.Name)
			}
			service.ImportStructs = append(service.ImportStructs, decl.Name)

			if path, ok := GoPackageImport(ctx, decl.PackageName).(string); ok {
				service.ImportPaths = append(service.ImportPaths, path)
			}
		}
	}

	for _, attribute := range method.Attributes {
		if attribute.PackageName != "http" {
			continue
		}

		switch attribute.Name {
		case "get", "post", "put", "patch", "delete", "options", "head", "trace":
			if len(attribute.Arguments) > 0 {
				if value := attribute.Arguments[0].GetStringLiteralExpr(); value != nil {
					if len(value.Value) == 0 {
						continue
					}
					binding, err := compileBindings(attribute.Name, value.Value, method)
					if err != nil {
						continue
					}

					m.Bindings = append(m.Bindings, binding)
				}
			}
		}
	}

	service.Interface.Methods = append(service.Interface.Methods, m)
	return nil
}

func compileBindings(methodName string, path string, method *lang.FunctionDecl) (*types.HTTPBinding, error) {
	p, pathParams := apicompiler.CompilePath(path)

	binding := &types.HTTPBinding{}
	binding.Verb = methodName
	binding.Path = p

	for _, param := range method.Signature.Parameters {
		err := compileBindingParameter(param, pathParams, nil, binding)
		if err != nil {
			return nil, err
		}
	}

	return binding, nil
}

func compileStructDecl(decl *lang.StructDecl, pathParams map[string]bool, enclosing *types.HTTPParameter, binding *types.HTTPBinding) error {
	if decl == nil {
		return errors.New("structDecl is empty")
	}

	for _, inherit := range decl.Type.Inherits {
		compileStructDecl(inherit.GetTypeDeclaration().GetStructDecl(), pathParams, enclosing, binding)
	}

	for _, field := range decl.Type.Fields {
		compileBindingParameter(field, pathParams, enclosing, binding)
	}

	return nil
}

func compileBindingParameter(decl *lang.ValueDecl, pathParams map[string]bool, enclosing *types.HTTPParameter, binding *types.HTTPBinding) error {
	compiler := &apicompiler.NominalTypeCompiler{}

	schema, err := compiler.Compile(decl.Type)
	if err != nil {
		return err
	}

	param := &types.HTTPParameter{
		Field: compileBindingField(schema, context.Components(compiler.Context).GetSchemas()),
	}

	if param.Field == nil {
		return err
	}

	param.Field.Name = decl.Name

	if enclosing != nil {
		param.Field.Enclosing = enclosing.Field
		param.Field.FullName = enclosing.Field.FullName + "." + decl.Name
	} else {
		param.Field.FullName = decl.Name
	}

	if v, e := lang.GetBoolAttribute(decl.Type.Attributes, "unwrap"); e == nil && v {
		param.Field.Unwrap = true

		structDecl := decl.Type.GetTypeDeclaration().GetStructDecl()
		if len(decl.Type.GenericArguments) == 0 && structDecl != nil {
			return compileStructDecl(structDecl, pathParams, param, binding)
		}
	}

	//TODO update sync with openapi, add request body
	if pathParams[param.Field.Name] || pathParams[param.Field.FullName] {
		param.Location = "path"
	} else {
		param.Location = "query"
	}

	binding.Params = append(binding.Params, param)

	if fieldType := param.Field.GetType(); fieldType != nil && fieldType.Message != nil {
		for pathParam := range pathParams {
			if strings.Contains(pathParam, ".") {
				for _, field := range fieldType.Message.Fields {
					fullName := param.Field.Name + "." + field.Name
					if pathParam == fullName {
						field.FullName = fullName
						field.Enclosing = param.Field
						binding.Params = append(binding.Params, &types.HTTPParameter{
							Field:    field,
							Location: "path",
						})
					}
				}
			}
		}
	}

	return nil
}

func compileBindingField(schema *openapi.Schema, index map[string]*openapi.Schema) *types.Field {
	if schema == nil {
		// error
		return nil
	}

	field := &types.Field{}
	switch schema.Type {
	case openapi.Schema_TYPE_BOOLEAN:
		field.Type = &types.FieldType{
			Name: "bool",
		}
	case openapi.Schema_TYPE_INTEGER:
		if len(schema.Format) > 0 {
			switch schema.Format {
			case "int8", "int16":
				field.Type = &types.FieldType{
					Name: "int32",
				}
			default:
				field.Type = &types.FieldType{
					Name: schema.Format,
				}
			}
		} else {
			field.Type = &types.FieldType{
				Name: "int64",
			}
		}
	case openapi.Schema_TYPE_NUMBER:
		field.Type = &types.FieldType{
			Name: "float64",
		}
	case openapi.Schema_TYPE_STRING:
		if len(schema.Enum) > 0 {
			field.Type = &types.FieldType{
				Name: "string",
				Enum: &types.Enum{Name: schema.Title},
			}
		} else {
			field.Type = &types.FieldType{
				Name: "string",
			}
		}
	case openapi.Schema_TYPE_ARRAY:
		s := schema.Items.GetSchema()
		f := compileBindingField(s, index)
		if f != nil {
			field.Type = &types.FieldType{
				Name:      f.Type.Name,
				StarExpr:  f.Type.StarExpr,
				ArrayType: true,
			}
		}
	case openapi.Schema_TYPE_OBJECT:
		if schema.AdditionalProperties != nil { // map
			field.Type = &types.FieldType{
				Map: &types.Map{
					KeyType: &types.FieldType{
						Name: "string",
					},
					ValueType: compileBindingField(schema.GetAdditionalProperties().GetSchema(), index).GetType(),
				},
			}
		} else {
			field.Type = &types.FieldType{
				Name: schema.Title,
				Message: &types.Message{
					Name: schema.Title,
				},
				StarExpr: true,
			}

			for name, item := range schema.Properties {
				f := compileBindingField(item.GetSchemaOf(index), index)
				f.Name = name
				field.Type.Message.Fields = append(field.Type.Message.Fields, f)
			}
		}
	}

	return field
}
