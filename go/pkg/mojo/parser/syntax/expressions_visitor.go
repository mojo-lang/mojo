package syntax

import (
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

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
        if expression, ok := expressionCtx.Accept(visitor).(*lang.Expression); ok {
            expressions = append(expressions, expression)
        } else {
            logs.Warnw("failed to parse expression", "code", expressionCtx.GetText())
        }
    }
    return expressions
}
