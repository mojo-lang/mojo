package precompiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

// this package will compile the mojo types to StructDel which can be compiled to protobuf

func CompileNominalType(ctx context.Context, t *lang.NominalType) (*lang.StructDecl, error) {
    if t != nil {
        switch t.GetFullName() {
        case core.ArrayTypeFullName:
            return CompileArrayToStruct(ctx, t)
        case core.MapTypeFullName:
            return CompileMapToStruct(ctx, t)
        case core.TupleTypeFullName:
            return CompileUnionToStruct(ctx, t)
        }
    }
    return nil, nil
}
