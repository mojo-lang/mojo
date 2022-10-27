package format

import (
	"errors"
	"fmt"
	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"strings"
)

func (p *Printer) PrintExpression(ctx context.Context, expr *sql.Expression) *Printer {
	if expr == nil || p == nil || p.Error != nil {
		return p
	}

	switch e := expr.Expression.(type) {
	case *sql.Expression_NullLiteralExpr:
		p.PrintTerm(ctx, lang.NewSymbolTerm(e.NullLiteralExpr.StartPosition, lang.TermTypeSymbol, "null"))
	case *sql.Expression_IntegerLiteralExpr:
		p.PrintTerm(ctx, lang.NewSymbolTerm(e.IntegerLiteralExpr.StartPosition,
			lang.TermTypeSymbol,
			fmt.Sprint(e.IntegerLiteralExpr.EvalValue())))
	case *sql.Expression_FloatLiteralExpr:
		p.PrintTerm(ctx, lang.NewSymbolTerm(e.FloatLiteralExpr.StartPosition,
			lang.TermTypeSymbol,
			fmt.Sprint(e.FloatLiteralExpr.EvalValue())))
	case *sql.Expression_StringLiteralExpr:
		p.PrintTerm(ctx, lang.NewSymbolTerm(e.StringLiteralExpr.StartPosition,
			lang.TermTypeSymbol,
			fmt.Sprint("\"", e.StringLiteralExpr.EvalValue(), "\"")))
	case *sql.Expression_IdentifierExpr:
		p.PrintIdentifyExpr(ctx, e.IdentifierExpr)
	default:
		p.Error = errors.New("not support expr in this printer")
	}

	return p
}

func (p *Printer) PrintIdentifyExpr(ctx context.Context, expr *sql.IdentifierExpr) *Printer {
	if ctx == nil || expr == nil || p == nil || p.Error != nil || len(expr.Identifier) > 0 {
		return p
	}

	identifier := strings.ToUpper(expr.Identifier)
	p.PrintRaw(identifier)
	return p
}
