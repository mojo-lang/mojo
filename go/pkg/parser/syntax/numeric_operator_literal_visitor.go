package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type NumericOperatorLiteralVisitor struct {
	*BaseMojoParserVisitor
}

func NewNumericOperatorLiteralVisitor() *NumericOperatorLiteralVisitor {
	visitor := &NumericOperatorLiteralVisitor{}
	return visitor
}

func (s *NumericOperatorLiteralVisitor) VisitNumericOperatorLiteral (ctx *NumericOperatorLiteralContext) interface{} {
	expr := &lang.NumericLiteralUnaryExpr{}

	operatorCtx := ctx.PostfixLiteralOperator()
	if operatorCtx != nil {
		expr.Operator = operatorCtx.GetText()

		literalCtx := ctx.NumericLiteral()
		if literalCtx != nil {
			if expression, ok := literalCtx.Accept(NewExpressionVisitor()).(*lang.Expression); ok {
				expr.Expression = expression
				return lang.NewNumericLiteralUnaryExpression(expr)
			}
		}
	}
	return nil
}
