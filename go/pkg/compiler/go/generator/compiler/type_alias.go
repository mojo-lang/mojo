package compiler

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/go/generator/data"
)

type TypeAlias struct {
	*data.Data
}

func (s *TypeAlias) CompileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	pkg := context.Package(ctx)
	thisCtx := context.WithType(ctx, decl)

	if pkg.GetFullName() != decl.GetPackageName() {
		return nil
	}

	formatComp := &FormatJson{Data: s.Data}
	if err := formatComp.CompileTypeAlias(thisCtx, decl); err != nil {
		return err
	}

	return nil
}
