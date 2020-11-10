package parser

type GenericArgumentVisitor struct {
	*BaseMojoParserVisitor
}

func NewGenericArgumentVisitor() *GenericArgumentVisitor {
	return &GenericArgumentVisitor{}
}

func (g *GenericArgumentsVisitor) VisitGenericArgument(ctx *GenericArgumentContext) interface{} {
	return nil
}