package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type IdentifierVisitor struct {
	*BaseProtobuf2Visitor
}

func NewIdentifierVisitor() *IdentifierVisitor {
	visitor := &IdentifierVisitor{}
	return visitor
}

func (v *IdentifierVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return ctx.GetText()
}

func (v *IdentifierVisitor) VisitFullIdent(ctx *FullIdentContext) interface{} {
	idents := ctx.AllIdent()
	var names []string
	for _, id := range idents {
		names = append(names, id.GetText())
	}

	if len(names) > 0 {
		return lang.NewIdentifier(names...)
	}

	return nil
}

func (v *IdentifierVisitor) VisitType_(ctx *Type_Context) interface{} {
	if message := ctx.MessageType(); message != nil {
		return message.Accept(v)
	}
	if enum := ctx.EnumType(); enum != nil {
		return enum.Accept(v)
	}
	return &lang.NominalType{Name: ctx.GetText()}
}

func (v *IdentifierVisitor) VisitKeyType(ctx *KeyTypeContext) interface{} {
	return &lang.NominalType{Name: ctx.GetText()}
}

func (v *IdentifierVisitor) VisitMessageType(ctx *MessageTypeContext) interface{} {
	if name := ctx.MessageName(); name != nil {
		typ := &lang.NominalType{
			Name: name.GetText(),
		}

		// in the syntax parse stage, can't to figure out which ident is type or package name
		// should be process in the semantic parse stage
		idents := ctx.AllIdent()
		cur := typ
		for i := len(idents) - 1; i >= 0; i-- {
			t := &lang.NominalType{Name: idents[i].GetText()}
			cur.Enclosing = t
			cur = t
		}
		return typ
	}
	return nil
}

func (v *IdentifierVisitor) VisitEnumType(ctx *EnumTypeContext) interface{} {
	if name := ctx.EnumName(); name != nil {
		typ := &lang.NominalType{
			Name: name.GetText(),
		}

		// in the syntax parse stage, can't to figure out which ident is type or package name
		// should be process in the semantic parse stage
		idents := ctx.AllIdent()
		cur := typ
		for i := len(idents) - 1; i >= 0; i-- {
			t := &lang.NominalType{Name: idents[i].GetText()}
			cur.Enclosing = t
			cur = t
		}
		return typ
	}
	return nil
}
