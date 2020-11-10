package parser

type TypeNameVisitor struct {
	*BaseMojoParserVisitor
}

func NewTypeNameVisitor() *TypeNameVisitor {
	return &TypeNameVisitor{}
}

func (t *TypeNameVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return ctx.TypeName().GetText()
}

func (t *TypeNameVisitor) VisitEnumName (ctx *EnumNameContext) interface{} {
	return ctx.TypeName().GetText()
}
