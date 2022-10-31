package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/go/generator/data"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
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
