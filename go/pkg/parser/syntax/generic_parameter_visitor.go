package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type GenericParameterVisitor struct {
	*BaseMojoParserVisitor
}

func NewGenericParameterVisitor() *GenericParameterVisitor {
	return &GenericParameterVisitor{}
}

func (a *GenericParameterVisitor) VisitGenericParameter(ctx *GenericParameterContext) interface{} {
	parameter := &lang.GenericParameter{}

	ellipsis := ctx.ELLIPSIS()
	if ellipsis != nil {
		attribute := &lang.Attribute{Name: "ellipsis"}
		parameter.Attributes = append(parameter.Attributes, attribute)
	}

	annotationCtx := ctx.TypeAnnotation()
	if annotationCtx != nil {
		parameter.Constraint = GetTypeAnnotation(annotationCtx)
	}

	nameCtx := ctx.TypeName()
	if nameCtx != nil {
		parameter.Name = GetTypeName(nameCtx)
		return parameter
	}

	return nil
}
