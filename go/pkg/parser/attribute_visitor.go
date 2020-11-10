package parser

import (
	"github.com/mojo-lang/lang/go/pkg/lang"
	"strconv"
)

type AttributeVisitor struct {
	*BaseMojoParserVisitor

	Attribute *lang.Attribute
}

func NewAttributeVisitor() *AttributeVisitor {
	visitor := &AttributeVisitor{}
	visitor.Attribute = &lang.Attribute{}
	return visitor
}

func (a *AttributeVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	attributeName := ctx.AttributeName()
	if attributeName == nil {
		// error
		return nil
	}
	attributeName.Accept(a)

	argumentClause := ctx.AttributeArgumentClause()
	if argumentClause != nil {
		argumentClause.Accept(a)
	}

	return a.Attribute
}

func (a *AttributeVisitor) VisitAttributeName(ctx *AttributeNameContext) interface{} {
	identifier := ctx.AttributeNameIdentifier()
	if identifier != nil {
		a.Attribute.Name = identifier.GetText()
		//a.Attribute.StartPosition = GetPosition(identifier.GetSymbol())
	} else {
		a.Attribute.Name = "number"
		number := ctx.DecimalLiteral()
		v, err := strconv.ParseInt(number.GetText(), 10, 32)
		if err != nil {
			// err
			return false
		}
		if v < 0 {
			return false
		}
		a.Attribute.Expressions = []*lang.Expression{
			lang.NewIntegerLiteralExpr(&lang.IntegerLiteralExpr{
				StartPosition:        nil,
				EndPosition:          nil,
				Value:                v,
			}),
		}

		a.Attribute.StartPosition = GetPosition(number.GetSymbol())
	}

	return true
}

func (a *AttributeVisitor) VisitAttributeArgumentClause(ctx *AttributeArgumentClauseContext) interface{} {
	expressionList := ctx.Expressions()
	if expressionList != nil {
		visitor := &ExpressionsVisitor{}
		expressions := expressionList.Accept(visitor).([]*lang.Expression)

		for _, expression := range expressions {
			a.Attribute.Expressions = append(a.Attribute.Expressions, expression)
		}
	}

	return true
}
