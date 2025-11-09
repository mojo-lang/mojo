package syntax

type TypeNameVisitor struct {
	*BaseCVisitor
}

func NewTypeNameVisitor() *TypeNameVisitor {
	visitor := &TypeNameVisitor{}
	return visitor
}

func (v *TypeNameVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	_ = ctx
	return nil
}
