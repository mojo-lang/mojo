package syntax

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type ExpressionVisitor struct {
	*BaseMojoParserVisitor
}

func NewExpressionVisitor() *ExpressionVisitor {
	return &ExpressionVisitor{}
}

func GetExpression(ctx IExpressionContext) *lang.Expression {
	if ctx != nil {
		visitor := NewExpressionVisitor()
		switch expr := ctx.Accept(visitor).(type) {
		case *lang.Expression:
			return expr
		case *lang.ArrayLiteralExpr:
			return lang.NewArrayLiteralExpression(expr)
		case *lang.ObjectLiteralExpr:
			return lang.NewObjectLiteralExpression(expr)
		case *lang.MapLiteralExpr:
			return lang.NewMapLiteralExpression(expr)
		case *lang.IdentifierExpr:
			return lang.NewIdentifierExpression(expr)
		default:
			fmt.Print("===> error")
		}
	}
	return nil
}

func (e *ExpressionVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	if prefixCtx := ctx.PrefixExpression(); prefixCtx != nil {
		if expr, ok := prefixCtx.Accept(e).(*lang.Expression); ok {
			if binaryCtx := ctx.BinaryExpressions(); binaryCtx != nil {
				if binaryExprs, ok := binaryCtx.Accept(e).([]*lang.BinaryExpr); ok {
					return BinaryExprParser{}.Parse(expr, binaryExprs)
				}
			}
			return expr
		}
	}

	return nil
}

func (e *ExpressionVisitor) VisitBinaryExpressions(ctx *BinaryExpressionsContext) interface{} {
	binaryCtxes := ctx.AllBinaryExpression()
	var binaries []*lang.BinaryExpr
	for _, binaryCtx := range binaryCtxes {
		if expr, ok := binaryCtx.Accept(e).(*lang.BinaryExpr); ok {
			binaries = append(binaries, expr)
		}
	}
	return binaries
}

func (e *ExpressionVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	if prefixExprCtx := ctx.PrefixExpression(); prefixExprCtx != nil {
		if binaryOperator := ctx.BinaryOperator(); binaryOperator != nil {
			if expression, ok := prefixExprCtx.Accept(e).(*lang.Expression); ok {
				return &lang.BinaryExpr{
					StartPosition: GetPosition(ctx.GetStart()),
					EndPosition:   GetPosition(ctx.GetStop()),
					Callee: &lang.BinaryExpr_Operator{
						Operator: &lang.Operator{
							StartPosition: GetPosition(binaryOperator.GetStart()),
							EndPosition:   GetPosition(binaryOperator.GetStop()),
							Symbol:        binaryOperator.GetText(),
						},
					},
					LeftArgument:  nil,
					RightArgument: expression,
				}
			}
		} else if conditionalOperator := ctx.ConditionalOperator(); conditionalOperator != nil {
			conditionalExpr := &lang.ConditionalExpr{
				EndPosition: GetPosition(ctx.GetStop()),
			}

			if expr, ok := conditionalOperator.Accept(e).(*lang.Expression); ok {
				conditionalExpr.ThenBranch = expr
			} else {
				return nil
			}

			if expression, ok := prefixExprCtx.Accept(e).(*lang.Expression); ok {
				conditionalExpr.ElseBranch = expression
				return &lang.BinaryExpr{
					StartPosition: GetPosition(ctx.GetStart()),
					EndPosition:   GetPosition(ctx.GetStop()),
					Callee: &lang.BinaryExpr_Operator{
						Operator: &lang.Operator{
							StartPosition: GetPosition(conditionalOperator.GetStart()),
							EndPosition:   GetPosition(conditionalOperator.GetStop()),
							Symbol:        "?",
						},
					},
					LeftArgument:  nil,
					RightArgument: lang.NewConditionalExpression(conditionalExpr),
				}
			}
		}
	}
	return nil
}

func (e *ExpressionVisitor) VisitConditionalOperator(ctx *ConditionalOperatorContext) interface{} {
	return GetExpression(ctx.Expression())
}

func (e *ExpressionVisitor) VisitPrefixExpression(ctx *PrefixExpressionContext) interface{} {
	if postfixCtx := ctx.PostfixExpression(); postfixCtx != nil {
		if expression, ok := postfixCtx.Accept(e).(*lang.Expression); ok {
			if operator := ctx.PrefixOperator(); operator != nil {
				if operator.GetText() == "-" {
					switch expression.Expression.(type) {
					case *lang.Expression_IntegerLiteralExpr:
						expression.GetIntegerLiteralExpr().SetNegative()
						return expression
					case *lang.Expression_FloatLiteralExpr:
						expression.GetFloatLiteralExpr().SetNegative()
						return expression
					}
				}
				return lang.NewPrefixUnaryExpression(&lang.PrefixUnaryExpr{
					StartPosition: GetPosition(ctx.GetStart()),
					EndPosition:   GetPosition(ctx.GetStop()),
					Callee: &lang.PrefixUnaryExpr_Operator{
						Operator: &lang.Operator{
							StartPosition: GetPosition(operator.GetStart()),
							EndPosition:   GetPosition(operator.GetStop()),
							Symbol:        operator.GetText(),
						},
					},
					Argument: expression,
				})
			}
			return expression
		}
	}
	return nil
}

func (e *ExpressionVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	if primaryCtx := ctx.PrimaryExpression(); primaryCtx != nil {
		if expression, ok := primaryCtx.Accept(e).(*lang.Expression); ok {
			suffixExprs := ctx.AllSuffixExpression()

			for _, suffixExpr := range suffixExprs {
				visitor := &SuffixExpressionVisitor{
					PrimaryExpression: expression,
				}
				var expr *lang.Expression
				if expr, ok = suffixExpr.Accept(visitor).(*lang.Expression); ok {
					expression = expr
				}
			}

			if operator := ctx.PostfixOperator(); operator != nil {
				return lang.NewPostfixUnaryExpression(&lang.PostfixUnaryExpr{
					StartPosition: GetPosition(ctx.GetStart()),
					EndPosition:   GetPosition(ctx.GetStop()),
					Callee: &lang.PostfixUnaryExpr_Operator{
						Operator: &lang.Operator{
							StartPosition: GetPosition(operator.GetStart()),
							EndPosition:   GetPosition(operator.GetStop()),
							Symbol:        operator.GetText(),
						},
					},
					Argument: expression,
				})
			}

			return expression
		}
	}

	return nil
}

func (e *ExpressionVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	if literalCtx := ctx.LiteralExpression(); literalCtx != nil {
		return literalCtx.Accept(e)
	}

	if identifierCtx := ctx.DeclarationIdentifier(); identifierCtx != nil {
		if identifier, ok := identifierCtx.Accept(e).(*lang.Identifier); ok {
			arguments := GetGenericArguments(ctx.GenericArgumentClause())
			enclosing := GetTypeIdentifier(ctx.TypeIdentifier())
			if enclosing != nil {
				identifier.Enclosing = enclosing.ToIdentifier()
			}

			return lang.NewIdentifierExpression(&lang.IdentifierExpr{
				StartPosition:    GetPosition(ctx.GetStart()),
				EndPosition:      GetPosition(ctx.GetStop()),
				Identifier:       identifier,
				GenericArguments: arguments,
			})
		}
	}

	if parenthesizedCtx := ctx.ParenthesizedExpression(); parenthesizedCtx != nil {
		return parenthesizedCtx.Accept(e)
	}

	if closureCtx := ctx.ClosureExpression(); closureCtx != nil {
		return closureCtx.Accept(&ClosureExprVisitor{})
	}

	if wildcardCtx := ctx.WildcardExpression(); wildcardCtx != nil {
		return wildcardCtx.Accept(e)
	}

	if tupleCtx := ctx.TupleLiteralExpression(); tupleCtx != nil {
		return tupleCtx.Accept(e)
	}

	return nil
}

func (e *ExpressionVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	literalCtx := ctx.Literal()
	if literalCtx != nil {
		return literalCtx.Accept(e)
	}

	arrayCtx := ctx.ArrayLiteral()
	if arrayCtx != nil {
		return arrayCtx.Accept(NewArrayLiteralVisitor())
	}

	objectCtx := ctx.ObjectLiteral()
	if objectCtx != nil {
		return objectCtx.Accept(NewObjectLiteralVisitor())
	}

	mapCtx := ctx.MapLiteral()
	if mapCtx != nil {
		return mapCtx.Accept(NewMapLiteralVisitor())
	}

	structLiteralCtx := ctx.StructLiteral()
	if structLiteralCtx != nil {
		return structLiteralCtx.Accept(NewStructLiteralVisitor())
	}

	stringOperatorLiteralCtx := ctx.StringOperatorLiteral()
	if stringOperatorLiteralCtx != nil {
		return stringOperatorLiteralCtx.Accept(NewStringOperatorLiteralVisitor())
	}

	numericOperatorLiteralCtx := ctx.NumericOperatorLiteral()
	if numericOperatorLiteralCtx != nil {
		return numericOperatorLiteralCtx.Accept(NewNumericOperatorLiteralVisitor())
	}

	return nil
}

func (e *ExpressionVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	nullCtx := ctx.NullLiteral()
	if nullCtx != nil {
		return lang.NewNullLiteralExpression(&lang.NullLiteralExpr{
			StartPosition: GetPosition(nullCtx.GetStart()),
			EndPosition:   GetPosition(nullCtx.GetStop()),
		})
	}

	boolCtx := ctx.BoolLiteral()
	if boolCtx != nil {
		return lang.NewBoolLiteralExpression(&lang.BoolLiteralExpr{
			StartPosition: GetPosition(boolCtx.GetStart()),
			EndPosition:   GetPosition(boolCtx.GetStop()),
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
	isNegative := false
	negatePrefix := ctx.NegatePrefixOperator()
	if negatePrefix != nil {
		isNegative = true
	}

	integerCtx := ctx.IntegerLiteral()
	if integerCtx != nil {
		v := integerCtx.Accept(e).(*lang.IntegerLiteralExpr)
		if v != nil {
			v.StartPosition = GetPosition(ctx.GetStart())
			v.EndPosition = GetPosition(ctx.GetStop())
			v.IsNegative = isNegative
			return lang.NewIntegerLiteralExpression(v)
		}
	}

	floatCtx := ctx.FLOAT_LITERAL()
	if floatCtx != nil {
		if v, err := strconv.ParseFloat(floatCtx.GetText(), 64); err == nil {
			return lang.NewFloatLiteralExpression(&lang.FloatLiteralExpr{
				StartPosition: GetPosition(ctx.GetStart()),
				EndPosition:   GetPosition(ctx.GetStop()),
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
	decimal := ctx.DECIMAL_LITERAL()
	if decimal != nil {
		v, err := strconv.ParseUint(decimal.GetText(), 10, 64)
		if err != nil {
			// FIXME log error
			return nil
		}

		return &lang.IntegerLiteralExpr{
			StartPosition: GetPosition(ctx.GetStart()),
			EndPosition:   GetPosition(ctx.GetStop()),
			Kind:          0,
			Implicit:      false,
			IsNegative:    false,
			Value:         v,
		}
	}

	return nil
}

func (e *ExpressionVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	static := ctx.STATIC_STRING_LITERAL()
	if static != nil {
		value := static.GetText()
		if strings.HasPrefix(value, "'") {
			value = value[1 : len(value)-1]
		} else if strings.HasPrefix(value, "\"") {
			value = value[1 : len(value)-1]
		}

		return lang.NewStringLiteralExpression(&lang.StringLiteralExpr{
			StartPosition: GetPosition(ctx.GetStart()),
			EndPosition:   GetPosition(ctx.GetStop()),
			Kind:          0,
			Implicit:      false,
			Value:         value,
		})
	}

	return nil
}

func (e *ExpressionVisitor) VisitDeclarationIdentifier(ctx *DeclarationIdentifierContext) interface{} {
	identifier := ctx.GetText()
	if len(identifier) > 0 {
		return &lang.Identifier{
			StartPosition: GetPosition(ctx.GetStart()),
			EndPosition:   GetPosition(ctx.GetStop()),
			Name:          identifier,
		}
	}
	return nil
}

func (e *ExpressionVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	if exprCtx := ctx.Expression(); exprCtx != nil {
		if expr, ok := exprCtx.Accept(e).(*lang.Expression); ok {
			return lang.NewParenthesizedExpression(&lang.ParenthesizedExpr{
				StartPosition: GetPosition(ctx.GetStart()),
				EndPosition:   GetPosition(ctx.GetStop()),
				Kind:          0,
				Expression:    expr,
			})
		}
	}
	return nil
}

func (e *ExpressionVisitor) VisitWildcardExpression(ctx *WildcardExpressionContext) interface{} {
	if ctx.UNDERSCORE().GetText() == "_" {
		return lang.NewWildcardExpression(&lang.WildcardExpr{
			StartPosition: GetPosition(ctx.GetStart()),
			EndPosition:   GetPosition(ctx.GetStop()),
		})
	}
	return nil
}

func (e *ExpressionVisitor) VisitTupleExpression(ctx *TupleLiteralExpressionContext) interface{} {
	elementCtxes := ctx.AllTupleElement()
	tupleExpr := &lang.TupleExpr{
		StartPosition: GetPosition(ctx.GetStart()),
		EndPosition:   GetPosition(ctx.GetStop()),
	}

	for _, elementCtx := range elementCtxes {
		if element, ok := elementCtx.Accept(e).(*lang.Argument); ok {
			tupleExpr.Elements = append(tupleExpr.Elements, element)
			if len(element.Label) > 0 {
				tupleExpr.HasElementLabels = true
			}
		}
	}

	if len(tupleExpr.Elements) > 0 {
		return lang.NewTupleExpression(tupleExpr)
	}
	return nil
}

func (e *ExpressionVisitor) VisitTupleElement(ctx *TupleElementContext) interface{} {
	element := &lang.Argument{}
	if exprCtx := ctx.Expression(); exprCtx != nil {
		element.Value = GetExpression(exprCtx)

		element.StartPosition = GetPosition(ctx.GetStart())
		element.EndPosition = GetPosition(ctx.GetStop())
	}

	if identifierCtx := ctx.LabelIdentifier(); identifierCtx != nil {
		element.Label = identifierCtx.GetText()
	}

	return element
}
