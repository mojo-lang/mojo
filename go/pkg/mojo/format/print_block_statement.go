package format

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

const ignoreBlockStartSymbol = "ignore-block-start-symbol"

func (p *Printer) PrintBlockStatement(ctx context.Context, stmt *lang.BlockStmt) *Printer {
	if stmt == nil || p == nil || p.Error != nil {
		return p
	}

	if !context.GetBool(ctx, ignoreBlockStartSymbol) {
		p.PrintRaw("{").BreakLine()
	}

	p.Indent()
	for _, statement := range stmt.Statements {
		p.PrintStatement(ctx, statement)
	}
	p.Outdent()
	p.PrintTerm(ctx, lang.NewSymbolTerm(stmt.EndPosition, lang.TermTypeEnd, "}"))

	return p
}
