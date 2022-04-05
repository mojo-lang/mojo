package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/data"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/compiler"
    "github.com/mojo-lang/mojo/go/pkg/protobuf/precompiler"
)

type PaginationResult struct {
    *data.Data
}

func (p *PaginationResult) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
    for _, method := range decl.Type.Methods {
        if err := p.CompileMethod(ctx, method); err != nil {
            return err
        }
    }
    return nil
}

func (p *PaginationResult) CompileMethod(ctx context.Context, method *lang.FunctionDecl) error {
    pagination, _ := lang.GetBoolAttribute(method.Attributes, core.PaginationAttributeName)
    if !pagination {
        return nil
    }

    if decl, err := precompiler.CompileMethodResponse(ctx, method); err != nil {
        return err
    } else {
        pr := &data.PaginationResult{}
        pr.PackageName = decl.GetPackageName()
        pr.GoPackageName = compiler.GetGoPackage(decl.GetPackageName())

        if pkg, _ := lang.GetStringAttribute(decl.Attributes, "go_package_name"); len(pkg) > 0 {
            pr.GoPackageName = pkg
        }
        pr.Name = decl.Name
        pr.FullName = GetFullName(pr.EnclosingName, pr.Name)
        pr.FieldName = strcase.ToCamel(decl.Type.Fields[0].Name)
    }

    return nil
}
