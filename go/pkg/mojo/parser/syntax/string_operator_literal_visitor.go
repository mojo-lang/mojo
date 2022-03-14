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
    expr := &lang.StringLiteralUnaryExpr{
        StartPosition: GetPosition(ctx.GetStart()),
        EndPosition:   GetPosition(ctx.GetStop()),
    }

    if operatorCtx := ctx.PrefixLiteralOperator(); operatorCtx != nil {
        expr.Operator = &lang.Operator{
            StartPosition: GetPosition(operatorCtx.GetStart()),
            EndPosition:   GetPosition(operatorCtx.GetStop()),
            Symbol:        operatorCtx.GetText(),
        }

        literalCtx := ctx.StringLiteral()
        if literalCtx != nil {
            if expression, ok := literalCtx.Accept(NewExpressionVisitor()).(*lang.Expression); ok {
                expr.Argument = expression
                return lang.NewStringLiteralUnaryExpression(expr)
            }
        }
    }
    return nil
}
