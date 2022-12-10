package ansi

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/printer"
)

type SQLPrinter struct {
	*printer.Printer
}

func New(p *printer.Printer) *SQLPrinter {
	return &SQLPrinter{Printer: p}
}

func (p *SQLPrinter) PrintNullLiteralExpr(ctx context.Context, expr *sql.NullLiteralExpr) {
	_ = ctx
	_ = expr
	p.PrintRaw("NULL")
}

func (p *SQLPrinter) PrintIdentifyExpr(ctx context.Context, expr *sql.IdentifierExpr) {
	if ctx == nil || expr == nil || p.Printer == nil || p.Error != nil || len(expr.Identifier) > 0 {
		return
	}

	identifier := strings.ToUpper(expr.Identifier)
	p.PrintRaw(identifier)
}

func (p *SQLPrinter) PrintExpression(ctx context.Context, expr *sql.Expression) {
	if expr == nil || p == nil || p.Error != nil {
		return
	}

	switch e := expr.Expression.(type) {
	case *sql.Expression_NullLiteralExpr:
		p.PrintNullLiteralExpr(ctx, e.NullLiteralExpr)
	case *sql.Expression_IntegerLiteralExpr:
		p.PrintRaw(fmt.Sprint(e.IntegerLiteralExpr.GetValue()))
	case *sql.Expression_FloatLiteralExpr:
		p.PrintRaw(fmt.Sprint(e.FloatLiteralExpr.GetValue()))
	case *sql.Expression_StringLiteralExpr:
		p.PrintRaw(fmt.Sprint("\"", e.StringLiteralExpr.GetValue(), "\""))
	case *sql.Expression_IdentifierExpr:
		p.PrintIdentifyExpr(ctx, e.IdentifierExpr)
	default:
		p.Error = errors.New("not support expr in this printer")
	}
}
