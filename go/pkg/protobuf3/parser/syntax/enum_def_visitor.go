package syntax

import (
	"strconv"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type EnumDefVisitor struct {
	*BaseProtobuf3Visitor
}

func NewEnumDefVisitor() *EnumDefVisitor {
	visitor := &EnumDefVisitor{}
	return visitor
}

func (v *EnumDefVisitor) VisitEnumDef(ctx *EnumDefContext) interface{} {
	if name := ctx.EnumName(); name != nil {
		if body := ctx.EnumBody(); body != nil {
			if decl, ok := body.Accept(v).(*lang.EnumDecl); ok {
				decl.Name = name.GetText()
				return decl
			}
		}
	}
	return nil
}

func (v *EnumDefVisitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	elements := ctx.AllEnumElement()

	decl := &lang.EnumDecl{Type: &lang.EnumType{}}
	for _, element := range elements {
		ele := element.Accept(v)
		if valueDecl, ok := ele.(*lang.ValueDecl); ok {
			decl.Type.Enumerators = append(decl.Type.Enumerators, valueDecl)
		}
	}
	return decl
}

func (v *EnumDefVisitor) VisitEnumElement(ctx *EnumElementContext) interface{} {
	if field := ctx.EnumField(); field != nil {
		return field.Accept(v)
	}
	return nil
}

func (v *EnumDefVisitor) VisitEnumField(ctx *EnumFieldContext) interface{} {
	decl := &lang.ValueDecl{Name: ctx.Ident().GetText()}

	lit := ctx.IntLit().GetText()
	value, _ := strconv.ParseInt(lit, 10, 64)
	if ctx.MINUS() != nil {
		value = -value
	}

	decl.Initializer = &lang.Initializer{
		Value: lang.NewIntegerLiteralExpressionFrom(value),
	}

	if options := ctx.EnumValueOptions(); options != nil {
		if attrs, ok := options.Accept(v).([]*lang.Attribute); ok {
			decl.Attributes = attrs
		}
	}

	return decl
}

func (v *EnumDefVisitor) VisitEnumValueOptions(ctx *EnumValueOptionsContext) interface{} {
	allOptions := ctx.AllEnumValueOption()
	var attrs []*lang.Attribute
	for _, option := range allOptions {
		if attr, ok := option.Accept(v).(*lang.Attribute); ok {
			attrs = append(attrs, attr)
		}
	}
	return attrs
}

func (v *EnumDefVisitor) VisitEnumValueOption(ctx *EnumValueOptionContext) interface{} {
	if name := ctx.OptionName(); name != nil {
		if attribute, ok := name.Accept(NewOptionStatementVisitor()).(*lang.Attribute); ok {
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
