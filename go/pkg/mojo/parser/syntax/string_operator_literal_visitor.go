package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type StringOperatorLiteralVisitor struct {
	*BaseMojoParserVisitor
}

func NewStringOperatorLiteralVisitor() *StringOperatorLiteralVisitor {
	visitor := &StringOperatorLiteralVisitor{}
	return visitor
}

func (s *StringOperatorLiteralVisitor) VisitStringOperatorLiteral(ctx *StringOperatorLiteralContext) interface{} {
	var expression *lang.Expression
	if prefixOperator := ctx.PrefixLiteralOperator(); prefixOperator != nil {
		expr := &lang.StringPrefixLiteralExpr{
			StartPosition: GetPosition(ctx.GetStart()),
			EndPosition:   GetPosition(ctx.GetStop()),
		}

		expr.SetOperator(&lang.Operator{
			StartPosition: GetPosition(prefixOperator.GetStart()),
			EndPosition:   GetPosition(prefixOperator.GetStop()),
			Symbol:        prefixOperator.GetText(),
		})

		literalCtx := ctx.StringLiteral()
		if literalCtx != nil {
			if ex, ok := literalCtx.Accept(NewExpressionVisitor()).(*lang.Expression); ok {
				expr.Argument = ex
				expression = lang.NewStringPrefixLiteralExpression(expr)
			}
		}
	}

	if suffixOperator := ctx.SuffixLiteralOperator(); suffixOperator != nil {
		expr := &lang.StringSuffixLiteralExpr{
			StartPosition: GetPosition(ctx.GetStart()),
			EndPosition:   GetPosition(ctx.GetStop()),
		}

		expr.SetOperator(&lang.Operator{
			StartPosition: GetPosition(suffixOperator.GetStart()),
			EndPosition:   GetPosition(suffixOperator.GetStop()),
			Symbol:        suffixOperator.GetText(),
		})

		if expression != nil {
			expr.Argument = expression
			expression = lang.NewStringSuffixLiteralExpression(expr)
		} else {
			literalCtx := ctx.StringLiteral()
			if literalCtx != nil {
				if ex, ok := literalCtx.Accept(NewExpressionVisitor()).(*lang.Expression); ok {
					expr.Argument = ex
					expression = lang.NewStringSuffixLiteralExpression(expr)
				}
			}
		}
	}

	return expression
}
