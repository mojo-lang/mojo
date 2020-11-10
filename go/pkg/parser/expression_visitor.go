package parser

import (
	"github.com/mojo-lang/lang/go/pkg/lang"
	"strconv"
)

type ExpressionVisitor struct {
	*BaseMojoParserVisitor

	Expression *lang.Expression
}

func NewExpressionVisitor() *ExpressionVisitor {
	visitor := &ExpressionVisitor{}
	return visitor
}

func GetExpression(ctx IExpressionContext) *lang.Expression {
	if ctx != nil {
		visitor := NewExpressionVisitor()
		return ctx.Accept(visitor).(*lang.Expression)
	}
	return nil
}

func (e *ExpressionVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	prefixCtx := ctx.PrefixExpression()
	if prefixCtx != nil {
		prefixCtx.Accept(e)
	}

	binaryCtx := ctx.BinaryExpressions()
	if binaryCtx != nil {
	}

	return e.Expression
}

func (e *ExpressionVisitor) VisitPrefixExpression(ctx *PrefixExpressionContext) interface{} {
	postfixCtx := ctx.PostfixExpression()
	if postfixCtx != nil {
		postfixCtx.Accept(e)
	}

	return true
}

func (e *ExpressionVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	literalCtx := ctx.LiteralExpression()
	if literalCtx != nil {
		literalCtx.Accept(e)
	}

	return true
}

func (e *ExpressionVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	primaryCtx := ctx.PrimaryExpression()
	if primaryCtx != nil {
		primaryCtx.Accept(e)
	}

	return true
}

func (e *ExpressionVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	literalCtx := ctx.Literal()
	if literalCtx != nil {
		return literalCtx.Accept(e)
	}

	arrayCtx := ctx.ArrayLiteral()
	if arrayCtx != nil {
		visitor := NewArrayLiteralVisitor()
		return arrayCtx.Accept(visitor)
	}

	objectCtx := ctx.ObjectLiteral()
	if objectCtx != nil {
		visitor := NewObjectLiteralVisitor()
		return objectCtx.Accept(visitor)
	}

	return nil
}

func (e *ExpressionVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	nullCtx := ctx.NullLiteral()
	if nullCtx != nil {
		return lang.NewNullLiteralExpr(&lang.NullLiteralExpr{
			StartPosition: nil,
			EndPosition:   nil,
		})
	}

	boolCtx := ctx.BoolLiteral()
	if boolCtx != nil {
		return lang.NewBoolLiteralExpr(&lang.BoolLiteralExpr{
			StartPosition: nil,
			EndPosition:   nil,
			Kind:          0,
			Implicit:      false,
			Value:         boolCtx.GetText() == "true",
		})
	}

	numericCtx := ctx.NumericLiteral()
	if numericCtx != nil {
		return numericCtx.Accept(e)
	}

	strCtx := ctx.StringLiteral()
	if strCtx != nil {
		return strCtx.Accept(e)
	}

	return nil
}

func (e *ExpressionVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	e.Expression = nil
	isNegative := false
	negatePrefix := ctx.NegatePrefixOperator()
	if negatePrefix != nil {
		isNegative = true
	}

	integerCtx := ctx.IntegerLiteral()
	if integerCtx != nil {
		v := integerCtx.Accept(e).(*lang.IntegerLiteralExpr)
		if v != nil {
			v.IsNegative = isNegative
			return lang.NewIntegerLiteralExpr(v)
		}
	}

	floatCtx := ctx.FloatLiteral()
	if floatCtx != nil {
		v, err := strconv.ParseFloat(floatCtx.GetText(), 64)
		if err != nil {
			return lang.NewFloatLiteralExpr(&lang.FloatLiteralExpr{
				StartPosition: nil,
				EndPosition:   nil,
				Kind:          0,
				Implicit:      false,
				IsNegative:    isNegative,
				Value:         v,
			})
		}
	}

	return nil
}

func (e *ExpressionVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	decimal := ctx.DecimalLiteral()
	if decimal != nil {
		v, err := strconv.ParseUint(decimal.GetText(), 10, 64)
		if err != nil {
			//FIXME log error
			return nil
		}

		return &lang.IntegerLiteralExpr{
			StartPosition: nil,
			EndPosition:   nil,
			Kind:          0,
			Implicit:      false,
			IsNegative:    false,
			Value:         int64(v),
		}
	}

	return nil
}

func (e *ExpressionVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	static := ctx.StaticStringLiteral()
	if static != nil {
		return lang.NewStringLiteralExpr(&lang.StringLiteralExpr{
			//StartPosition:        nil,
			//EndPosition:          nil,
			//Kind:                 0,
			Implicit: false,
			Value:    static.GetText(),
		})
	}

	return nil
}
