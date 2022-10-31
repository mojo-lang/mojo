package syntax

import (
	"strings"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type OptionStatementVisitor struct {
	*BaseProtobuf3Visitor
}

func NewOptionStatementVisitor() *OptionStatementVisitor {
	visitor := &OptionStatementVisitor{}
	return visitor
}

func (v *OptionStatementVisitor) VisitOptionStatement(ctx *OptionStatementContext) interface{} {
	if name := ctx.OptionName(); name != nil {
		if attribute, ok := name.Accept(v).(*lang.Attribute); ok {
			if constant := ctx.Constant(); constant != nil {
				if expr, ok := constant.Accept(NewConstantVisitor()).(*lang.Expression); ok {
					attribute.Arguments = append(attribute.Arguments, &lang.Argument{
						Value: expr,
					})

					return attribute
				}
			}
		}
	}
	return nil
}

func (v *OptionStatementVisitor) VisitOptionName(ctx *OptionNameContext) interface{} {
	idents := ctx.AllFullIdent()
	if ctx.LP() != nil { // custom options
		attr := &lang.Attribute{}

		if identifier, ok := idents[0].Accept(NewIdentifierVisitor()).(*lang.Identifier); ok {
			names := identifier.FullNames()

			attr.Name = names[len(names)-1]
			if len(names) > 1 {
				attr.PackageName = strings.Join(names[:len(names)-1], ".")
			}
		}

		if len(idents) > 1 {
			if identifier, ok := idents[1].Accept(NewIdentifierVisitor()).(*lang.Identifier); ok {
				attr.Fields = identifier.FullNames()
			}
		}
		return attr
	} else { // system options
		if identifier, ok := idents[0].Accept(NewIdentifierVisitor()).(*lang.Identifier); ok {
			names := identifier.FullNames()

			attr := &lang.Attribute{
				Name: names[0],
			}
			if len(names) > 1 {
				attr.Fields = names[1:]
			}

			return attr
		}
	}

	return nil
}
