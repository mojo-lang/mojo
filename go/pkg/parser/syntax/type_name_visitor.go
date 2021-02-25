package syntax

import "fmt"

type TypeNameVisitor struct {
	*BaseMojoParserVisitor
}

func NewTypeNameVisitor() *TypeNameVisitor {
	return &TypeNameVisitor{}
}

func GetTypeName(ctx ITypeNameContext) string {
	visitor := NewTypeNameVisitor()
	if name, ok := ctx.Accept(visitor).(string); ok {
		return name
	} else {
		fmt.Print("===> error")
		return ""
	}
}

func (t *TypeNameVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return ctx.GetText()
}

func (t *TypeNameVisitor) VisitEnumName(ctx *EnumNameContext) interface{} {
	return ctx.TypeName().GetText()
}
