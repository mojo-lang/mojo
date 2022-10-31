package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type NumericOperatorLiteralVisitor struct {
	*BaseMojoParserVisitor
}

func NewNumericOperatorLiteralVisitor() *NumericOperatorLiteralVisitor {
	visitor := &NumericOperatorLiteralVisitor{}
	return visitor
}

func (s *NumericOperatorLiteralVisitor) VisitNumericOperatorLiteral(ctx *NumericOperatorLiteralContext) interface{} {
	expr := &lang.NumericSuffixLiteralExpr{
		StartPosition: GetPosition(ctx.GetStart()),
		EndPosition:   GetPosition(ctx.GetStop()),
	}

	if operatorCtx := ctx.SuffixLiteralOperator(); operatorCtx != nil {
		expr.Callee = &lang.NumericSuffixLiteralExpr_Operator{
			Operator: &lang.Operator{
				StartPosition: GetPosition(operatorCtx.GetStart()),
				EndPosition:   GetPosition(operatorCtx.GetStop()),
				Symbol:        operatorCtx.GetText(),
			},
		}
		if literalCtx := ctx.NumericLiteral(); literalCtx != nil {
			if expression, ok := literalCtx.Accept(NewExpressionVisitor()).(*lang.Expression); ok {
				expr.Argument = expression
				return lang.NewNumericSuffixLiteralExpression(expr)
			}
		}
	}
	return nil
}
