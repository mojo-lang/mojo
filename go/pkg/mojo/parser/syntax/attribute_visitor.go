package syntax

import (
	"strconv"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
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
			StartPosition:    GetPosition(ctx.GetStart()),
			EndPosition:      GetPosition(ctx.GetStop()),
			PackageName:      "",
			Name:             core.NumberAttributeName,
			GenericArguments: nil,
			Arguments: []*lang.Argument{
				lang.NewIntegerLiteralArgument(&lang.IntegerLiteralExpr{
					StartPosition: nil,
					EndPosition:   nil,
					Value:         uint64(v),
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
	arguments := ctx.AttributeArguments()
	if arguments != nil {
		if as, ok := arguments.Accept(a).([]*lang.Argument); ok {
			return as
		} else {
			logs.Warnw("failed to parse attribute arguments", "code", arguments.GetText(), "source", arguments.GetStart().GetSource())
		}
	}

	return nil
}

func (a *AttributeVisitor) VisitAttributeArguments(ctx *AttributeArgumentsContext) interface{} {
	argumentCtxes := ctx.AllAttributeArgument()
	if argumentCtxes != nil {
		var arguments []*lang.Argument
		for _, argumentCtx := range argumentCtxes {
			if argument, ok := argumentCtx.Accept(a).(*lang.Argument); ok {
				arguments = append(arguments, argument)
			} else {
				logs.Warnw("failed to parse attribute argument", "code", argumentCtx.GetText(), "source", argumentCtx.GetStart().GetSource())
			}
		}
		return arguments
	}
	return nil
}

func (a *AttributeVisitor) VisitAttributeArgument(ctx *AttributeArgumentContext) interface{} {
	if expression := GetExpression(ctx.Expression()); expression != nil {
		argument := &lang.Argument{
			StartPosition: GetPosition(ctx.GetStart()),
			EndPosition:   GetPosition(ctx.GetStop()),
			Value:         expression,
		}
		if identifier := ctx.LabelIdentifier(); identifier != nil {
			argument.Label = identifier.GetText()
		}
		return argument
	}
	return nil
}
