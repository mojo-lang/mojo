package syntax

type GenericArgumentVisitor struct {
	*BaseMojoParserVisitor
}

func NewGenericArgumentVisitor() *GenericArgumentVisitor {
	return &GenericArgumentVisitor{}
}

func (g *GenericArgumentVisitor) VisitGenericArgument(ctx *GenericArgumentContext) interface{} {
	nominalType := GetType(ctx.Type_())
	if nominalType != nil {
		nominalType.Attributes = GetAttributes(ctx.Attributes())
	}
	return nominalType
}
