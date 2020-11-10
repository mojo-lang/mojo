package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

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

	decl.Document = GetDocument(ctx.Document())
	decl.Attributes = GetAttributes(ctx.Attributes())

	nameCtx := ctx.EnumName()
	if nameCtx != nil {
		visitor := NewTypeNameVisitor()
		decl.Name = nameCtx.Accept(visitor).(string)
	}

	bodyCtx := ctx.EnumBody()
	if bodyCtx != nil {
		decl.Type = bodyCtx.Accept(e).(*lang.EnumType)
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

		value := memberCtx.Accept(visitor).(*lang.ValueDecl)
		if value != nil {
			t.Enumerators = append(t.Enumerators, value)
		}
	}

	return t
}
