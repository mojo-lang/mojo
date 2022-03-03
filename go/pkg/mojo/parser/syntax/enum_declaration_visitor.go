package syntax

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type EnumDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewEnumDeclarationVisitor() *EnumDeclarationVisitor {
	return &EnumDeclarationVisitor{}
}

func (e *EnumDeclarationVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	if nameCtx := ctx.EnumName(); nameCtx != nil {
		if name, ok := nameCtx.Accept(NewTypeNameVisitor()).(string); ok {
			decl := &lang.EnumDecl{}
			decl.Name = name

			if bodyCtx := ctx.EnumBody(); bodyCtx != nil {
				if t, ok := bodyCtx.Accept(e).(*lang.EnumType); ok {
					decl.Type = t

					inheritances := GetTypeInheritances(ctx.TypeInheritanceClause())
					if len(inheritances) > 0 {
						decl.Type.UnderlyingType = inheritances[0]
					}

					return decl
				}
				logs.Errorw("failed to parse enum decl body", "name", name, "body", bodyCtx.GetText())
			} else {
				logs.Errorw("failed to parse enum decl without body", "name", name)
			}
		}
	}

	return nil
}

func (e *EnumDeclarationVisitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	membersCtx := ctx.EnumMembers()
	if membersCtx != nil {
		return membersCtx.Accept(e)
	}

	return nil
}

func (e *EnumDeclarationVisitor) VisitEnumMembers(ctx *EnumMembersContext) interface{} {
	t := &lang.EnumType{}
	t.StartPosition = GetPosition(ctx.GetStart())
	t.EndPosition = GetPosition(ctx.GetStop())

	memberCtxes := ctx.AllEnumMember()
	var freeDocument *lang.Document
	for _, memberCtx := range memberCtxes {
		member := memberCtx.Accept(NewValueDeclarationVisitor())
		if value, ok := member.(*lang.ValueDecl); ok {
			if freeDocument != nil {
				value.SetStartPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})
				freeDocument = nil
			}
			t.Enumerators = append(t.Enumerators, value)
		}
		if document, ok := member.(*lang.Document); ok {
			freeDocument = document
		}
	}

	if freeDocument != nil {
		t.SetEndPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})
	}

	return t
}
