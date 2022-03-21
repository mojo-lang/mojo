package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/protobuf"
)

type PaginationResult BoxedArray

func (p *PaginationResult) CompileInterface(decl *lang.InterfaceDecl) error {
    for _, method := range decl.Type.Methods {
        if err := p.CompileMethod(method); err != nil {
            return err
        }
    }
    return nil
}

func (p *PaginationResult) CompileMethod(method *lang.FunctionDecl) error {
    pagination, _ := lang.GetBoolAttribute(method.Attributes, core.PaginationAttributeName)
    if !pagination {
        return nil
    }

    if decl := protobuf.GenerateArrayTypeResponse(method, true); decl != nil {
        p.PackageName = decl.GetPackageName()
        p.GoPackageName = GetGoPackage(decl.GetPackageName())

        if pkg, _ := lang.GetStringAttribute(decl.Attributes, "go_package_name"); len(pkg) > 0 {
            p.GoPackageName = pkg
        }
        p.Name = decl.Name
        p.FullName = GetFullName(p.EnclosingName, p.Name)
        p.FieldName = strcase.ToCamel(decl.Type.Fields[0].Name)
    }

    return nil
}

func (p *PaginationResult) GetPackageName() string {
    return p.PackageName
}

func (p *PaginationResult) GetFullName() string {
    return p.FullName
}
