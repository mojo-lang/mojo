package decompiler

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/packages/protobuf/go/pkg/mojo/protobuf"
	"github.com/pkg/errors"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/compiler"
	"github.com/mojo-lang/mojo/go/pkg/mojo/compiler/transformer"
)

func CompileMethod(ctx context.Context, method *lang.FunctionDecl) (request *lang.StructDecl, response *lang.StructDecl, err error) {
	if request, err = CompileMethodRequest(ctx, method); err != nil {
		return nil, nil, err
	} else if response, err = CompileMethodResponse(ctx, method); err != nil {
		return nil, nil, err
	} else {
		return request, response, nil
	}
}

func CompileMethodRequest(ctx context.Context, method *lang.FunctionDecl) (*lang.StructDecl, error) {
	service := context.InterfaceDecl(ctx)

	// special case for the request, do NOT generate new request type
	if parameters := method.Signature.GetParameters(); len(parameters) == 1 {
		param := parameters[0]
		if requestType, _ := param.GetBoolAttribute(protobuf.MethodRequestTypeAttributeName); requestType {
			return param.GetType().GetTypeDeclaration().GetStructDecl(), nil
		}
	}

	// generate the request message
	req := &lang.StructDecl{
		Name:           strcase.ToCamel(method.Name) + "Request",
		Implicit:       true,
		PackageName:    method.PackageName,
		SourceFileName: method.SourceFileName,
	}

	parameters := method.Signature.GetParameters()

	// add number attribute if there is no field has number attribute
	noneFieldNumber := true
	for _, param := range parameters {
		if param.HasAttribute(core.NumberAttributeName) {
			noneFieldNumber = false
			break
		}
	}
	if noneFieldNumber {
		for i, param := range parameters {
			param.SetIntegerAttribute(core.NumberAttributeName, int64(i+1))
		}
		logs.Infow("all the arguments has not set the number attribute, will auto set by sequence, warning for compatibility",
			"method", method.GetName(), "interface", service.GetName())
	}

	req.Type = &lang.StructType{
		Fields: parameters,
	}

	pagination, _ := lang.GetBoolAttribute(method.Attributes, core.PaginationAttributeName)
	if pagination {
		req.Type.MergeFields(compiler.PaginationRequestFields())
	}

	if query, _ := lang.GetBoolAttribute(method.Attributes, core.QueryAttributeName); query {
		req.Type.MergeFields(compiler.QueryRequestFields())

		pkg := context.Package(ctx)
		cp := pkg.ResolvedDependencies["mojo.core"]
		ids := []*lang.Identifier{
			cp.GetIdentifier("mojo.core.FieldMask"),
			cp.GetIdentifier("mojo.core.Ordering"),
		}

		service.ResolvedIdentifiers = lang.MergeDependencies(service.ResolvedIdentifiers, ids)

		sf := context.SourceFile(ctx)
		sf.ResolvedIdentifiers = lang.MergeDependencies(sf.ResolvedIdentifiers, ids)
	}

	return req, nil
}

func CompileMethodResponse(ctx context.Context, method *lang.FunctionDecl) (*lang.StructDecl, error) {
	result := method.Signature.GetResultType()
	pkg := context.Package(ctx)
	service := context.InterfaceDecl(ctx)
	if result == nil {
		if id := pkg.GetIdentifier(core.NullTypeFullName); id != nil {
			return id.Declaration.GetStructDecl(), nil
		} else {
			return nil, logs.NewErrorw("failed to find the identifier", "name", core.NullTypeFullName,
				"method", method.Name, "interface", service.Name, "package", pkg.FullName)
		}
	} else if result.GetFullName() == core.ArrayTypeFullName {
		pagination, _ := lang.GetBoolAttribute(method.Attributes, core.PaginationAttributeName)
		return compileArrayTypeResponse(ctx, method, pagination)
	} else if result.GetFullName() == core.MapTypeFullName {
		return compileMapTypeResponse(ctx, method)
	} else if result.IsScalar() {
		scalarTypeName := core.GetProperBoxedScalarName(result.GetFullName())
		if id := pkg.GetIdentifier(scalarTypeName); id != nil {
			return id.GetDeclaration().GetStructDecl(), nil
		} else {
			return nil, logs.NewErrorw("failed to find the identifier", "name", scalarTypeName,
				"method", method.Name, "interface", service.Name, "package", pkg.FullName)
		}
	} else {
		if result.GetFullName() == core.TupleTypeFullName {
			return compileTupleTypeResponse(ctx, method)
		} else {
			return result.TypeDeclaration.GetStructDecl(), nil
		}
	}
}

func compileArrayTypeResponse(ctx context.Context, method *lang.FunctionDecl, pagination bool) (*lang.StructDecl, error) {
	_ = ctx
	result := method.Signature.GetResultType()
	if result == nil || result.GetFullName() != core.ArrayTypeFullName {
		return nil, errors.New("invalid array type for the response in method")
	}

	result.Attributes = lang.SetIntegerAttribute(result.Attributes, core.NumberAttributeName, 1)
	val := result.GenericArguments[0]
	name := transformer.Plural(strcase.ToSnake(val.Name))

	// generate the response message
	resp := &lang.StructDecl{
		Implicit:       true,
		PackageName:    method.PackageName,
		SourceFileName: method.SourceFileName,
		Name:           strcase.ToCamel(method.Name) + "Response",
		Type: &lang.StructType{
			Fields: []*lang.ValueDecl{{
				Name: name,
				Type: result,
			}},
		},
	}

	if pagination {
		resp.Type.MergeFields(compiler.PaginationResponseFields())
	}

	return resp, nil
}

func compileMapTypeResponse(ctx context.Context, method *lang.FunctionDecl) (*lang.StructDecl, error) {
	_ = ctx
	result := method.Signature.GetResultType()
	if result == nil || result.GetFullName() != core.MapTypeFullName {
		return nil, errors.New("invalid map type for the response in method")
	}

	result.Attributes = lang.SetIntegerAttribute(result.Attributes, core.NumberAttributeName, 1)
	// val := result.GenericArguments[0]
	// name := transformer.Plural(strcase.ToSnake(val.Name))
	name := "vals"

	// generate the request message
	resp := &lang.StructDecl{
		Implicit:       true,
		PackageName:    method.PackageName,
		SourceFileName: method.SourceFileName,
		Name:           strcase.ToCamel(method.Name) + "Response",
		Type: &lang.StructType{
			Fields: []*lang.ValueDecl{{
				Name: name,
				Type: result,
			}},
		},
	}

	return resp, nil
}

func compileTupleTypeResponse(ctx context.Context, method *lang.FunctionDecl) (*lang.StructDecl, error) {
	_ = ctx
	result := method.Signature.GetResultType()
	if result == nil || result.GetFullName() != core.TupleTypeFullName {
		return nil, errors.New("invalid map type for the response in method")
	}

	name := strcase.ToCamel(method.Name) + "Response"
	result.Attributes = lang.SetStringAttribute(result.Attributes, core.AliasAttributeName, name)

	var fields []*lang.ValueDecl
	for i, element := range result.GenericArguments {
		number, _ := element.GetIntegerAttribute(core.NumberAttributeName)
		if number == 0 {
			number = int64(i + 1)
			element.SetIntegerAttribute(core.NumberAttributeName, number)
		}

		fields = append(fields, &lang.ValueDecl{
			Name: transformer.GenericArgumentLabel(element),
			Type: element,
		})
	}

	// generate the request message
	resp := &lang.StructDecl{
		Implicit:       true,
		PackageName:    method.PackageName,
		SourceFileName: method.SourceFileName,
		Name:           strcase.ToCamel(method.Name) + "Response",
		Type: &lang.StructType{
			Fields: fields,
		},
	}

	return resp, nil
}
