package syntax

import (
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type EnumDeclarationVisitor struct {
	*BaseMojoParserVisitor

	EnumDecl *lang.EnumDecl
}

func NewEnumDeclarationVisitor() *EnumDeclarationVisitor {
	return &EnumDeclarationVisitor{}
}

func (e *EnumDeclarationVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	decl := &lang.EnumDecl{}
	decl.Type = &lang.EnumType{}

	//decl.Document.SetDocument(GetDocument(ctx.Document()))
	//decl.Attributes = GetAttributes(ctx.Attributes())

	nameCtx := ctx.EnumName()
	if nameCtx != nil {
		visitor := NewTypeNameVisitor()
		if name, ok := nameCtx.Accept(visitor).(string); ok {
			decl.Name = name
		} else {
			fmt.Print("===> error")
		}
	}

	bodyCtx := ctx.EnumBody()
	if bodyCtx != nil {
		if t, ok := bodyCtx.Accept(e).(*lang.EnumType); ok {
			decl.Type = t
		} else {
			fmt.Print("===> error")
		}
	}

	inheritances := GetTypeInheritances(ctx.TypeInheritanceClause())
	if len(inheritances) > 0 {
		decl.Type.UnderlyingType = inheritances[0]
	}

	return decl
}

func (e *EnumDeclarationVisitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	membersCtx := ctx.EnumMembers()
	if membersCtx != nil {
		return membersCtx.Accept(e)
	}

	return &lang.EnumType{}
}

func (e *EnumDeclarationVisitor) VisitEnumMembers(ctx *EnumMembersContext) interface{} {
	memberCtxes := ctx.AllEnumMember()
	t := &lang.EnumType{}
	for _, memberCtx := range memberCtxes {
		visitor := NewValueDeclarationVisitor()

		if value, ok := memberCtx.Accept(visitor).(*lang.ValueDecl); ok {
			t.Enumerators = append(t.Enumerators, value)
		} else {
			fmt.Print("===> error")
		}
	}

	return t
}
