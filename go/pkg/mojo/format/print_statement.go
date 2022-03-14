package format

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func (p *Printer) PrintStatement(ctx context.Context, decl *lang.Statement) {
    if decl == nil || p == nil || p.Error != nil {
        return
    }

    switch d := decl.Statement.(type) {
    case *lang.Statement_Declaration:
        p.PrintDeclaration(ctx, d.Declaration)
    case *lang.Statement_Expression:
        p.PrintExpression(ctx, d.Expression)
    }
}
