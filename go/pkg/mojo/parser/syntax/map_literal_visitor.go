package syntax

import (
	"fmt"
	"strconv"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

func GetMapLiteral(ctx IMapLiteralContext) *lang.MapLiteralExpr {
	if ctx != nil {
		visitor := NewMapLiteralVisitor()
		if expr, ok := ctx.Accept(visitor).(*lang.Expression); ok {
			return expr.GetMapLiteralExpr()
		} else {
			fmt.Print("===> error")
		}
	}
	return nil
}

type MapLiteralVisitor struct {
	*BaseMojoParserVisitor
}

func NewMapLiteralVisitor() *MapLiteralVisitor {
	visitor := &MapLiteralVisitor{}
	return visitor
}

func (e *MapLiteralVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	if ctx != nil {
		if context := ctx.MapLiteralItems(); context != nil {
			if expr, ok := context.Accept(e).(*lang.MapLiteralExpr); ok {
				expr.StartPosition = GetPosition(ctx.GetStart())
				expr.EndPosition = GetPosition(ctx.GetStop())
				return lang.NewMapLiteralExpression(expr)
			}
		}
	}
	return nil
}

func (e *MapLiteralVisitor) VisitMapLiteralItems(ctx *MapLiteralItemsContext) interface{} {
	if ctx == nil {
		return nil
	}

	expr := &lang.MapLiteralExpr{}
	for _, item := range ctx.AllMapLiteralItem() {
		if entry, ok := item.Accept(e).(*lang.MapLiteralExpr_Entry); ok {
			expr.Entries = append(expr.Entries, entry)
		}
	}

	return expr
}

func (e *MapLiteralVisitor) VisitMapLiteralItem(ctx *MapLiteralItemContext) interface{} {
	key := ""
	numeric := false
	if stringLiteralCtx := ctx.StringLiteral(); stringLiteralCtx != nil {
		if expr, ok := stringLiteralCtx.Accept(NewExpressionVisitor()).(*lang.Expression); ok {
			key = expr.GetStringLiteralExpr().Value
		}
	}
	if integerLiteralCtx := ctx.IntegerLiteral(); integerLiteralCtx != nil {
		if expr, ok := integerLiteralCtx.Accept(NewExpressionVisitor()).(*lang.Expression); ok {
			integerLiteral := expr.GetIntegerLiteralExpr()
			key = strconv.FormatUint(integerLiteral.Value, 10)
			if integerLiteral.IsNegative {
				key = "-" + key
			}
			numeric = true
		}
	}

	entry := &lang.MapLiteralExpr_Entry{
		Key:     key,
		Numeric: numeric,
		Value:   GetExpression(ctx.Expression()),
	}

	return entry
}
