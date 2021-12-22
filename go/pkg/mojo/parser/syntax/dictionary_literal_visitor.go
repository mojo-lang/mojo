package syntax

import (
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"strconv"
)

type MapLiteralVisitor struct {
	*BaseMojoParserVisitor

	Expression *lang.MapLiteralExpr
}

func NewMapLiteralVisitor() *MapLiteralVisitor {
	visitor := &MapLiteralVisitor{
		Expression: &lang.MapLiteralExpr{},
	}
	return visitor
}

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

func (e *MapLiteralVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	if ctx != nil {
		context := ctx.MapLiteralItems()
		if context != nil {
			return context.Accept(e)
		}
	}
	return nil
}

func (e *MapLiteralVisitor) VisitMapLiteralItems(ctx *MapLiteralItemsContext) interface{} {
	if ctx == nil {
		return nil
	}

	for _, item := range ctx.AllMapLiteralItem() {
		item.Accept(e)
	}

	return lang.NewMapLiteralExpression(e.Expression)
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
			key = strconv.FormatInt(expr.GetIntegerLiteralExpr().Value, 10)
			numeric = true
		}
	}

	entry := &lang.MapLiteralExpr_Entry{
		Key:     key,
		Numeric: numeric,
		Value:   GetExpression(ctx.Expression()),
	}
	e.Expression.Entries = append(e.Expression.Entries, entry)

	return nil
}
