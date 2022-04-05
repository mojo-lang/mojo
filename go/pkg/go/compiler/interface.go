package compiler

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/data"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/protobuf/precompiler"
)

type Interface struct {
    *data.Data
}

func (i *Interface) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
    thisCtx := context.WithType(ctx, decl)

    for _, method := range decl.Type.Methods {
        methodCtx := context.WithType(ctx, method)
        isPagination, _ := lang.GetBoolAttribute(method.Attributes, "pagination")
        if isPagination {
            pagination := &PaginationResult{Data: i.Data}
            if err := pagination.CompileMethod(thisCtx, method); err != nil {
                return err
            }
        } else {
            req, resp, err := precompiler.CompileMethod(thisCtx, method)
            if err != nil {
                return err
            }

            s := &Struct{Data: i.Data}
            if err = s.CompileStruct(methodCtx, req); err != nil {
                return err
            }
            if err = s.CompileStruct(methodCtx, resp); err != nil {
                return err
            }
        }
    }

    return nil
}
