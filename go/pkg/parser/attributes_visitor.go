package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type AttributesVisitor struct {
	*BaseMojoParserVisitor
}

func NewAttributesVisitor() *AttributesVisitor {
	return &AttributesVisitor{}
}

func GetAttributes(ctx IAttributesContext) []*lang.Attribute {
	if ctx != nil {
		visitor := NewAttributesVisitor()
		return ctx.Accept(visitor).([]*lang.Attribute)
	}
	return []*lang.Attribute{}
}

func (a *AttributesVisitor) VisitAttributes(ctx *AttributesContext) interface{} {
	var attributes []*lang.Attribute

	attrs := ctx.AllAttribute()
	for _, attributeCtx := range attrs {
		if attributeCtx != nil {
			visitor := NewAttributeVisitor()
			attributeCtx.Accept(visitor)

			attributes = append(attributes, visitor.Attribute)
		}
	}

	return attributes
}
