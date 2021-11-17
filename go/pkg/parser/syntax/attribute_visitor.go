package syntax

import (
	"fmt"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"strconv"
)

type AttributeVisitor struct {
	*BaseMojoParserVisitor
}

func NewAttributeVisitor() *AttributeVisitor {
	visitor := &AttributeVisitor{}
	return visitor
}

func (a *AttributeVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	number := ctx.DECIMAL_LITERAL()
	if number != nil {
		v, err := strconv.ParseInt(number.GetText(), 10, 32)
		if err != nil {
			// err
			return false
		}
		if v < 0 {
			return false
		}

		return &lang.Attribute{
			StartPosition:    GetPosition(number.GetSymbol()),
			EndPosition:      nil,
			PackageName:      "",
			Name:             "number",
			GenericArguments: nil,
			Arguments: []*lang.Argument{
				lang.NewIntegerLiteralArgument(&lang.IntegerLiteralExpr{
					StartPosition: nil,
					EndPosition:   nil,
					Value:         v,
				}),
			},
		}
	}

	attributeIdentifier := ctx.AttributeIdentifier()
	if attributeIdentifier != nil {
		if attribute, ok := attributeIdentifier.Accept(a).(*lang.Attribute); ok {
			attribute.GenericArguments = GetGenericArguments(ctx.GenericArgumentClause())

			argumentClause := ctx.AttributeArgumentClause()
			if argumentClause != nil {
				if arguments, ok := argumentClause.Accept(a).([]*lang.Argument); ok {
					attribute.Arguments = arguments
				} else {
					logs.Warnw("failed to parse arguments in attribute", "attribute", attribute.Name, "code", argumentClause.GetText(), "source", attributeIdentifier.GetStart().GetSource())
				}
			}
			return attribute
		}
	}

	return nil
}

func (a *AttributeVisitor) VisitAttributeIdentifier(ctx *AttributeIdentifierContext) interface{} {
	identifier := ctx.AttributeName()
	if identifier != nil {
		return &lang.Attribute{
			Name:          identifier.GetText(),
			StartPosition: GetPosition(identifier.GetStart()),
			PackageName:   GetPackageIdentifier(ctx.PackageIdentifier()),
		}
	}
	return nil
}

func (a *AttributeVisitor) VisitAttributeArgumentClause(ctx *AttributeArgumentClauseContext) interface{} {
	expressionList := ctx.Expressions()
	if expressionList != nil {
		visitor := &ExpressionsVisitor{}
		if expressions, ok := expressionList.Accept(visitor).([]*lang.Expression); ok {
			var arguments []*lang.Argument
			for _, expr := range expressions {
				arguments = append(arguments, &lang.Argument{Value: expr})
			}
			return arguments
		} else {
			fmt.Print("===> error")
		}
	}

	return nil
}
