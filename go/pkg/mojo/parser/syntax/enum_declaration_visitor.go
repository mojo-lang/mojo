package syntax

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type EnumDeclarationVisitor struct {
	*BaseMojoParserVisitor

	followingDocument *lang.Document
}

func NewEnumDeclarationVisitor() *EnumDeclarationVisitor {
	return &EnumDeclarationVisitor{}
}

func (e *EnumDeclarationVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	if nameCtx := ctx.EnumName(); nameCtx != nil {
		if name, ok := nameCtx.Accept(NewTypeNameVisitor()).(string); ok {
			decl := &lang.EnumDecl{
				StartPosition:   GetPosition(ctx.GetStart()),
				EndPosition:     GetPosition(ctx.GetStop()),
				KeywordPosition: GetPosition(ctx.KEYWORD_ENUM().GetSymbol()),
				Name:            name,
				NamePosition:    GetPosition(nameCtx.GetStart()),
			}

			e.followingDocument = nil
			if bodyCtx := ctx.EnumBody(); bodyCtx != nil {
				if t, ok := bodyCtx.Accept(e).(*lang.EnumType); ok {
					decl.Type = t

					inheritances := GetTypeInheritances(ctx.TypeInheritanceClause())
					if len(inheritances) > 0 {
						inheritances[0].Document = e.followingDocument
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
	if followingDocumentCtx := ctx.FollowingDocument(); followingDocumentCtx != nil {
		e.followingDocument = GetFollowingDocument(followingDocumentCtx)
	}

	if membersCtx := ctx.EnumMembers(); membersCtx != nil {
		return membersCtx.Accept(e)
	}

	return nil
}

func (e *EnumDeclarationVisitor) VisitEnumMembers(ctx *EnumMembersContext) interface{} {
	t := &lang.EnumType{}
	t.StartPosition = GetPosition(ctx.GetStart())
	t.EndPosition = GetPosition(ctx.GetStop())

	documentCtxes := ctx.AllEosWithDocument()
	memberCtxes := ctx.AllEnumMember()
	var freeDocument *lang.Document
	for i, memberCtx := range memberCtxes {
		member := memberCtx.Accept(NewValueDeclarationVisitor())
		var document *lang.Document
		if i < len(documentCtxes) {
			document = GetEosDocument(documentCtxes[i])
		}

		if value, ok := member.(*lang.ValueDecl); ok {
			if freeDocument != nil {
				value.SetStartPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})
				freeDocument = nil
			}
			value.Document = document
			t.Enumerators = append(t.Enumerators, value)
		} else if doc, ok := member.(*lang.Document); ok {
			freeDocument = doc
		}
	}

	if freeDocument != nil {
		t.SetEndPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})
	}

	return t
}
