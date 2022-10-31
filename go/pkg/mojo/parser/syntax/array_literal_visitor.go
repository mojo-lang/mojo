package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type ArrayLiteralVisitor struct {
	*BaseMojoParserVisitor
}

func NewArrayLiteralVisitor() *ArrayLiteralVisitor {
	visitor := &ArrayLiteralVisitor{}
	return visitor
}

func (e *ArrayLiteralVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	if context := ctx.ArrayLiteralItems(); context != nil {
		if expr, ok := context.Accept(e).(*lang.Expression); ok {
			array := expr.GetArrayLiteralExpr()
			array.StartPosition = GetPosition(ctx.GetStart())
			array.EndPosition = GetPosition(ctx.GetStop())
			return expr
		}
	}
	return nil
}

func (e *ArrayLiteralVisitor) VisitArrayLiteralItems(ctx *ArrayLiteralItemsContext) interface{} {
	itemCtxes := ctx.AllArrayLiteralItem()
	var expressions []*lang.Expression
	for _, itemCtx := range itemCtxes {
		if expression, ok := itemCtx.Accept(e).(*lang.Expression); ok {
			expressions = append(expressions, expression)
		}
	}

	if len(expressions) > 0 {
		return lang.NewArrayLiteralExpression(&lang.ArrayLiteralExpr{
			Kind:     0,
			Implicit: false,
			Elements: expressions,
		})
	}

	return nil
}

func (e *ArrayLiteralVisitor) VisitArrayLiteralItem(ctx *ArrayLiteralItemContext) interface{} {
	return GetExpression(ctx.Expression())
}
