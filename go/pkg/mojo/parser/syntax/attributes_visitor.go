package syntax

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type AttributesVisitor struct {
	*BaseMojoParserVisitor
}

func NewAttributesVisitor() *AttributesVisitor {
	return &AttributesVisitor{}
}

func GetAttributes(ctx IAttributesContext) []*lang.Attribute {
	if ctx != nil {
		visitor := NewAttributesVisitor()
		if attributes, ok := ctx.Accept(visitor).([]*lang.Attribute); ok {
			return attributes
		} else {
			logs.Warnw("failed to parse attributes", "code", ctx.GetText())
		}
	}
	return []*lang.Attribute{}
}

func (a *AttributesVisitor) VisitAttributes(ctx *AttributesContext) interface{} {
	var attributes []*lang.Attribute

	attrs := ctx.AllAttribute()
	for _, attributeCtx := range attrs {
		if attributeCtx != nil {
			visitor := NewAttributeVisitor()
			if attribute, ok := attributeCtx.Accept(visitor).(*lang.Attribute); ok {
				attributes = append(attributes, attribute)
			} else {
				logs.Warnw("failed to parse attribute", "code", ctx.GetText())
			}
		}
	}

	return attributes
}
