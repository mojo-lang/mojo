package decompiler

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

func CompileNominalType(ctx context.Context, t *lang.NominalType) (*lang.StructDecl, error) {
	if t != nil {
		switch t.GetFullName() {
		case core.ArrayTypeFullName:
			return CompileArray(ctx, t)
		case core.MapTypeFullName:
			return CompileMap(ctx, t)
		case core.TupleTypeFullName:
			return CompileUnion(ctx, t)
		}
	}
	return nil, nil
}
