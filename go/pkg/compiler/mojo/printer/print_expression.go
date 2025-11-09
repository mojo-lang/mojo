package printer

import (
	"errors"
	"fmt"

	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

func (p *Printer) PrintExpression(ctx context.Context, expr *lang.Expression) *Printer {
	if expr == nil || p.GetError() != nil {
		return p
	}

	switch e := expr.Expression.(type) {
	case *lang.Expression_NullLiteralExpr:
		p.PrintTerm(ctx, lang.NewSymbolTerm(e.NullLiteralExpr.StartPosition, lang.TermTypeSymbol, "null"))
	case *lang.Expression_IntegerLiteralExpr:
		p.PrintTerm(ctx, lang.NewSymbolTerm(e.IntegerLiteralExpr.StartPosition,
			lang.TermTypeSymbol,
			fmt.Sprint(e.IntegerLiteralExpr.EvalValue())))
	case *lang.Expression_FloatLiteralExpr:
		p.PrintTerm(ctx, lang.NewSymbolTerm(e.FloatLiteralExpr.StartPosition,
			lang.TermTypeSymbol,
			fmt.Sprint(e.FloatLiteralExpr.EvalValue())))
	case *lang.Expression_StringLiteralExpr:
		p.PrintTerm(ctx, lang.NewSymbolTerm(e.StringLiteralExpr.StartPosition,
			lang.TermTypeSymbol,
			fmt.Sprint("\"", e.StringLiteralExpr.EvalValue(), "\"")))
	case *lang.Expression_ArrayLiteralExpr:
		p.PrintArrayLiteralExpr(ctx, e.ArrayLiteralExpr)
	case *lang.Expression_MapLiteralExpr:
		p.PrintMapLiteralExpr(ctx, e.MapLiteralExpr)
	case *lang.Expression_ObjectLiteralExpr:
		p.PrintObjectLiteralExpr(ctx, e.ObjectLiteralExpr)
	case *lang.Expression_IdentifierExpr:
		p.PrintIdentifyExpr(ctx, e.IdentifierExpr)
	default:
		p.SetError(errors.New("not support expr in this printer"))
	}

	return p
}

func (p *Printer) PrintArrayLiteralExpr(ctx context.Context, expr *lang.ArrayLiteralExpr) *Printer {
	if expr == nil || p.GetError() != nil {
		return p
	}

	p.PrintRaw("[")
	for i, element := range expr.Elements {
		if i > 0 {
			p.PrintRaw(", ")
		}
		p.PrintExpression(ctx, element)
	}
	p.PrintRaw("]")

	return p
}

func (p *Printer) PrintMapLiteralExpr(ctx context.Context, expr *lang.MapLiteralExpr) *Printer {
	if expr == nil || p.GetError() != nil {
		return p
	}

	p.PrintRaw("{")
	p.Indent()
	for _, entry := range expr.Entries {
		if entry.Numeric {
			p.PrintLine(entry.Key)
		} else {
			p.PrintLine("\"", entry.Key, "\"")
		}

		p.PrintRaw(": ").PrintExpression(ctx, entry.Value).BreakLine()
	}
	p.Outdent()
	p.PrintLine("}")

	return p
}

func (p *Printer) PrintObjectLiteralExpr(ctx context.Context, expr *lang.ObjectLiteralExpr) *Printer {
	if expr == nil || p.GetError() != nil {
		return p
	}

	p.PrintRaw("{")
	p.Indent()
	for _, field := range expr.Fields {
		p.PrintLine(field.Name).PrintRaw(": ").PrintExpression(ctx, field.Value).BreakLine()
	}
	p.Outdent()
	p.PrintLine("}")

	return p
}

func (p *Printer) PrintIdentifyExpr(ctx context.Context, expr *lang.IdentifierExpr) *Printer {
	if expr == nil || p.GetError() != nil || expr.Identifier == nil {
		return p
	}

	identifier := expr.Identifier
	if len(identifier.PackageName) > 0 {
		p.PrintRaw(identifier.PackageName, ".")
	}

	enclosingNames := identifier.EnclosingNames()
	if len(enclosingNames) > 0 {
		for _, name := range enclosingNames {
			p.PrintRaw(name, ".")
		}
	}

	p.PrintRaw(identifier.Name)
	p.PrintGenericArguments(ctx, expr.GenericArguments)
	return p
}
