package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type ExpressionsVisitor struct {
	*BaseMojoParserVisitor
}

func NewExpressionsVisitor() *ExpressionsVisitor {
	return &ExpressionsVisitor{}
}

func (e *ExpressionsVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	var expressions []*lang.Expression
	expressionCtxes := ctx.AllExpression()
	for _, expressionCtx := range expressionCtxes {
		visitor := &ExpressionVisitor{}
		expression := expressionCtx.Accept(visitor).(*lang.Expression)

		expressions = append(expressions, expression)
	}
	return expressions
}
