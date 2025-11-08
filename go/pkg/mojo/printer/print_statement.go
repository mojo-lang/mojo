package printer

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintStatement(ctx context.Context, decl *lang.Statement) *Printer {
	if decl == nil || p.GetError() != nil {
		return p
	}

	switch d := decl.Statement.(type) {
	case *lang.Statement_Declaration:
		p.PrintDeclaration(ctx, d.Declaration)
	case *lang.Statement_Expression:
		p.PrintExpression(ctx, d.Expression)
	}
	return p
}
