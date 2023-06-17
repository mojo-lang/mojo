package syntax

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type StructDeclarationVisitor struct {
	*BaseMojoParserVisitor

	name              string
	followingDocument *lang.Document
}

func NewStructDeclarationVisitor() *StructDeclarationVisitor {
	visitor := &StructDeclarationVisitor{}
	return visitor
}

func (s *StructDeclarationVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	s.name = ""
	if structName := ctx.StructName(); structName != nil {
		if name, ok := structName.Accept(s).(string); ok {
			s.name = name

			decl := &lang.StructDecl{
				StartPosition:     GetPosition(ctx.GetStart()),
				EndPosition:       GetPosition(ctx.GetStop()),
				KeywordPosition:   GetPosition(ctx.KEYWORD_TYPE().GetSymbol()),
				Name:              name,
				NamePosition:      GetPosition(structName.GetStart()),
				GenericParameters: GetGenericParameters(ctx.GenericParameterClause()),
				Enclosing:         nil,
			}

			enclosing := &lang.NominalType{
				Name: name,
				// TypeDeclaration: lang.NewStructTypeDeclaration(decl), // may cause reference circle
			}

			if structType := ctx.StructType(); structType != nil {
				if len(structType.GetText()) > 0 {
					if d, ok := structType.Accept(s).(*lang.StructDecl); ok {
						decl.Type = d.Type
						decl.EnumDecls = d.EnumDecls
						decl.StructDecls = d.StructDecls
						decl.TypeAliasDecls = d.TypeAliasDecls
						decl.SetEndPosition(d.EndPosition)

						for _, enumDecl := range decl.EnumDecls {
							enumDecl.Enclosing = enclosing
						}
						for _, structDecl := range decl.StructDecls {
							structDecl.Enclosing = enclosing

							for _, eDecl := range structDecl.EnumDecls {
								eDecl.Enclosing.Enclosing = enclosing
							}
							for _, sDecl := range structDecl.StructDecls {
								sDecl.Enclosing.Enclosing = enclosing
							}
							for _, aDecl := range structDecl.TypeAliasDecls {
								aDecl.Enclosing.Enclosing = enclosing
							}
						}
						for _, typeAliasDecl := range decl.TypeAliasDecls {
							typeAliasDecl.Enclosing = enclosing
						}
						if d.Type == nil {
							logs.Warnw("declaration an opaque struct", "struct", name)
						}
					} else {
						logs.Errorw("failed to parse struct type", "struct", name, "code", structType.GetText())
					}
				} else {
					logs.Warnw("declaration an opaque struct", "struct", name)
				}
			}
			return decl
		}
	}
	return nil
}

func (s *StructDeclarationVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	var decl *lang.StructDecl
	s.followingDocument = nil

	if body := ctx.StructBody(); body != nil {
		if d, ok := body.Accept(s).(*lang.StructDecl); ok {
			decl = d
		}
	}

	if inheritances := ctx.TypeInheritanceClause(); inheritances != nil {
		if types := GetTypeInheritances(inheritances); len(types) > 0 {
			if decl == nil {
				decl = &lang.StructDecl{Type: &lang.StructType{}}
			}
			if decl.Type == nil {
				decl.Type = &lang.StructType{}
			}
			decl.Type.Inherits = types
			decl.Type.Inherits[len(decl.Type.Inherits)-1].Document = s.followingDocument
		}
	}

	if decl != nil {
		if decl.Type != nil {
			decl.Type.SetStartPosition(GetPosition(ctx.GetStart()))
			decl.Type.SetEndPosition(GetPosition(ctx.GetStop()))
		}
		return decl
	}

	// remove the type info
	return nil
}

func (s *StructDeclarationVisitor) VisitStructName(ctx *StructNameContext) interface{} {
	return ctx.GetText()
}

func (s *StructDeclarationVisitor) VisitStructBody(ctx *StructBodyContext) interface{} {
	if followingDocumentCtx := ctx.FollowingDocument(); followingDocumentCtx != nil {
		s.followingDocument = GetFollowingDocument(followingDocumentCtx)
	}

	if members := ctx.StructMembers(); members != nil {
		return members.Accept(s)
	}
	return &lang.StructDecl{} // opaque struct
}

func (s *StructDeclarationVisitor) VisitStructMembers(ctx *StructMembersContext) interface{} {
	documents := ctx.AllEosWithDocument()
	if members := ctx.AllStructMember(); len(members) > 0 {
		decl := &lang.StructDecl{
			Type: &lang.StructType{},
		}
		appendDocumentLine := func(doc *lang.Document, following *lang.Document) *lang.Document {
			if doc != nil {
				if following != nil {
					for _, line := range following.Lines {
						doc.Lines = append(doc.Lines, line)
					}
				}
				return doc
			} else {
				return following
			}
		}
		var document *lang.Document
		var freeDocument *lang.Document
		for i, member := range members {
			if len(documents) > i {
				document = GetEosDocument(documents[i])
			}
			switch m := member.Accept(s).(type) {
			case *lang.StructDecl:
				if freeDocument != nil {
					m.SetStartPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})
					freeDocument = nil
				}
				decl.StructDecls = append(decl.StructDecls, m)
			case *lang.EnumDecl:
				if freeDocument != nil {
					m.SetStartPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})
					freeDocument = nil
				}
				decl.EnumDecls = append(decl.EnumDecls, m)
			case *lang.TypeAliasDecl:
				if freeDocument != nil {
					m.SetStartPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})
					freeDocument = nil
				}
				decl.TypeAliasDecls = append(decl.TypeAliasDecls, m)
				m.Document = appendDocumentLine(m.Document, document)
			case *lang.ValueDecl:
				if freeDocument != nil {
					m.SetStartPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})
					freeDocument = nil
				}
				decl.Type.Fields = append(decl.Type.Fields, m)
				m.Document = appendDocumentLine(m.Document, document)
			case *lang.Document:
				freeDocument = m
			}
		}

		if freeDocument != nil {
			decl.Type.SetEndPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})
		}
		return decl
	}

	return nil
}

func (s *StructDeclarationVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	document := GetDocument(ctx.Document())
	attributes := GetAttributes(ctx.Attributes())

	var startPosition *lang.Position
	if document != nil {
		startPosition = document.StartPosition
	}
	if startPosition == nil && len(attributes) > 0 {
		startPosition = attributes[0].StartPosition
	}

	if structCtx := ctx.StructDeclaration(); structCtx != nil {
		if structDecl, ok := structCtx.Accept(NewStructDeclarationVisitor()).(*lang.StructDecl); ok {
			structDecl.Document = document
			structDecl.Attributes = attributes
			structDecl.SetStartPosition(startPosition)
			return structDecl
		}
	}

	if enumCtx := ctx.EnumDeclaration(); enumCtx != nil {
		if enumDecl, ok := enumCtx.Accept(NewEnumDeclarationVisitor()).(*lang.EnumDecl); ok {
			enumDecl.Document = document
			enumDecl.Attributes = attributes
			enumDecl.SetStartPosition(startPosition)
			return enumDecl
		}
	}

	if typeAliasCtx := ctx.TypeAliasDeclaration(); typeAliasCtx != nil {
		if typeAlias, ok := typeAliasCtx.Accept(NewTypeAliasDeclarationVisitor()).(*lang.TypeAliasDecl); ok {
			typeAlias.Document = document
			typeAlias.Attributes = attributes
			typeAlias.SetStartPosition(startPosition)
			return typeAlias
		}
	}

	if memberCtx := ctx.StructMemberDeclaration(); memberCtx != nil {
		if memberDecl, ok := memberCtx.Accept(NewValueDeclarationVisitor()).(*lang.ValueDecl); ok {
			memberDecl.Document = document
			memberDecl.Attributes = attributes
			memberDecl.SetStartPosition(startPosition)
			return memberDecl
		}
	}

	if freeFloatingDocumentContext := ctx.FloatingStatement(); freeFloatingDocumentContext != nil {
		return GetFreeFloatingDocument(freeFloatingDocumentContext)
	}

	return nil
}
