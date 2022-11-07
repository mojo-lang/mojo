package ansi

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/printer"
)

type SQLPrinter struct {
	P *printer.Printer
}

func (p *SQLPrinter) PrintNullLiteralExpr(ctx context.Context, expr *sql.NullLiteralExpr) {
	p.P.PrintTerm(ctx, lang.NewSymbolTerm(expr.StartPosition, lang.TermTypeSymbol, "NULL"))
}

func (p *SQLPrinter) PrintIdentifyExpr(ctx context.Context, expr *sql.IdentifierExpr) {
	if ctx == nil || expr == nil || p.P == nil || p.P.Error != nil || len(expr.Identifier) > 0 {
		return
	}

	identifier := strings.ToUpper(expr.Identifier)
	p.P.PrintRaw(identifier)
}

func (p *SQLPrinter) PrintExpression(ctx context.Context, expr *sql.Expression) {
	if expr == nil || p == nil || p.P.Error != nil {
		return
	}

	switch e := expr.Expression.(type) {
	case *sql.Expression_NullLiteralExpr:
		p.PrintNullLiteralExpr(ctx, e.NullLiteralExpr)
	case *sql.Expression_IntegerLiteralExpr:
		p.P.PrintRaw(fmt.Sprint(e.IntegerLiteralExpr.GetValue()))
	case *sql.Expression_FloatLiteralExpr:
		p.P.PrintRaw(fmt.Sprint(e.FloatLiteralExpr.GetValue()))
	case *sql.Expression_StringLiteralExpr:
		p.P.PrintRaw(fmt.Sprint("\"", e.StringLiteralExpr.GetValue(), "\""))
	case *sql.Expression_IdentifierExpr:
		p.PrintIdentifyExpr(ctx, e.IdentifierExpr)
	default:
		p.P.Error = errors.New("not support expr in this printer")
	}
}
