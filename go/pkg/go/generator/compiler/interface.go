package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/go/generator/data"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/generator/precompiler"
)

type Interface struct {
	*data.Data
}

func (i *Interface) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
	thisCtx := context.WithType(ctx, decl)

	for _, method := range decl.Type.Methods {
		methodCtx := context.WithType(ctx, method)
		isPagination, _ := lang.GetBoolAttribute(method.Attributes, core.PaginationAttributeName)
		result := method.GetSignature().GetResult().GetType()
		if isPagination || result.GetFullName() == core.ArrayTypeFullName {
			ar := &ArrayResponse{Data: i.Data}
			if err := ar.CompileMethod(thisCtx, method); err != nil {
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

			if result.GetFullName() == core.MapTypeFullName {
				bm := &BoxedMap{Data: i.Data}
				if err = bm.CompileStruct(thisCtx, resp); err != nil {
					return err
				}
			} else if result.GetFullName() == core.UnionTypeFullName {
				bu := &BoxedUnion{Data: i.Data}
				if err = bu.CompileStruct(thisCtx, resp); err != nil {
					return err
				}
			} else {
				if err = s.CompileStruct(methodCtx, resp); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
