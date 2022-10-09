package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	data2 "github.com/mojo-lang/mojo/go/pkg/go/generator/data"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/generator/precompiler"
)

type ArrayResponse struct {
	*data2.Data
}

func (p *ArrayResponse) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
	for _, method := range decl.Type.Methods {
		if err := p.CompileMethod(ctx, method); err != nil {
			return err
		}
	}
	return nil
}

func (p *ArrayResponse) CompileMethod(ctx context.Context, method *lang.FunctionDecl) error {
	pagination, _ := lang.GetBoolAttribute(method.Attributes, core.PaginationAttributeName)
	if !pagination {
		result := method.GetSignature().GetResult().GetType()
		if result.GetFullName() != core.ArrayTypeFullName {
			return nil
		}
	}

	if decl, err := precompiler.CompileMethodResponse(ctx, method); err != nil {
		return err
	} else {
		pr := &data2.ArrayResponse{}
		pr.PackageName = decl.GetPackageName()
		pr.GoPackageName = lang.GetGoPackageName(decl.GetPackageName())

		if pkg, _ := lang.GetStringAttribute(decl.Attributes, "go_package_name"); len(pkg) > 0 {
			pr.GoPackageName = pkg
		}
		pr.Name = decl.Name
		pr.FullName = GetFullName(pr.EnclosingName, pr.Name)
		pr.FieldName = strcase.ToCamel(decl.Type.Fields[0].Name)

		p.Data.ArrayResponses = append(p.Data.ArrayResponses, pr)
	}

	return nil
}
