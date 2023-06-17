package syntax

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type InterfaceDeclarationVisitor struct {
	*BaseMojoParserVisitor

	name              string
	followingDocument *lang.Document
}

func NewInterfaceDeclarationVisitor() *InterfaceDeclarationVisitor {
	decl := &InterfaceDeclarationVisitor{}
	return decl
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{} {
	i.name = ""

	if nameContext := ctx.InterfaceName(); nameContext != nil {
		if name, ok := nameContext.Accept(i).(string); ok && len(name) > 0 {
			i.name = name

			decl := &lang.InterfaceDecl{
				StartPosition:     GetPosition(ctx.GetStart()),
				EndPosition:       GetPosition(ctx.GetStop()),
				KeywordPosition:   GetPosition(ctx.KEYWORD_INTERFACE().GetSymbol()),
				Name:              name,
				NamePosition:      GetPosition(nameContext.GetStart()),
				GenericParameters: GetGenericParameters(ctx.GenericParameterClause()),
			}

			if typeCtx := ctx.InterfaceType(); typeCtx != nil {
				if t, ok := typeCtx.Accept(i).(*lang.InterfaceType); ok && t != nil {
					decl.Type = t
				}
			}

			return decl
		}
	} else {
		// error happened
	}
	return nil
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceType(ctx *InterfaceTypeContext) interface{} {
	i.followingDocument = nil
	inherits := GetTypeInheritances(ctx.TypeInheritanceClause())

	if bodyCtx := ctx.InterfaceBody(); bodyCtx != nil {
		if t, ok := bodyCtx.Accept(i).(*lang.InterfaceType); ok && t != nil {
			if len(inherits) > 0 && i.followingDocument != nil {
				inherits[len(inherits)-1].Document = i.followingDocument
			}

			t.Inherits = inherits
			t.SetStartPosition(GetPosition(ctx.GetStart()))
			t.SetEndPosition(GetPosition(ctx.GetStop()))
			return t
		}
	}

	if len(inherits) > 0 {
		if i.followingDocument != nil {
			inherits[len(inherits)-1].Document = i.followingDocument
		}
		return &lang.InterfaceType{
			Inherits:      inherits,
			StartPosition: GetPosition(ctx.GetStart()),
			EndPosition:   GetPosition(ctx.GetStop()),
		}
	}

	logs.Warnw("declaration an opaque interface", "interface", i.name)
	return nil
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceName(ctx *InterfaceNameContext) interface{} {
	return ctx.GetText()
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceBody(ctx *InterfaceBodyContext) interface{} {
	if followingDocumentCtx := ctx.FollowingDocument(); followingDocumentCtx != nil {
		i.followingDocument = GetFollowingDocument(followingDocumentCtx)
	}

	if membersCtx := ctx.InterfaceMembers(); membersCtx != nil {
		return membersCtx.Accept(i)
	}

	return nil
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceMembers(ctx *InterfaceMembersContext) interface{} {
	var freeDocument *lang.Document

	interfaceType := &lang.InterfaceType{}
	allInterfaceMember := ctx.AllInterfaceMember()
	for _, memberCtx := range allInterfaceMember {
		member := memberCtx.Accept(i)
		if funcDecl, ok := member.(*lang.FunctionDecl); ok && funcDecl != nil {
			if freeDocument != nil {
				funcDecl.SetStartPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})
				freeDocument = nil
			}
			interfaceType.Methods = append(interfaceType.Methods, funcDecl)
		}
		if document, ok := member.(*lang.Document); ok {
			freeDocument = document
		}
	}

	interfaceType.SetEndPosition(&lang.Position{LeadingComments: lang.NewComments(freeDocument)})

	return interfaceType
}

func (i *InterfaceDeclarationVisitor) VisitInterfaceMember(ctx *InterfaceMemberContext) interface{} {
	if freeDocument := ctx.FloatingStatement(); freeDocument != nil {
		return GetFreeFloatingDocument(freeDocument)
	} else {
		document := GetDocument(ctx.Document())
		attributes := GetAttributes(ctx.Attributes())

		var startPosition *lang.Position
		if document != nil {
			startPosition = document.StartPosition
		}
		if startPosition == nil && len(attributes) > 0 {
			startPosition = attributes[0].StartPosition
		}

		if methodCtx := ctx.InterfaceMethodDeclaration(); methodCtx != nil {
			if funcDecl, ok := methodCtx.Accept(NewFuncDeclarationVisitor()).(*lang.FunctionDecl); ok {
				funcDecl.Document = document
				funcDecl.Attributes = attributes
				funcDecl.SetStartPosition(startPosition)
				return funcDecl
			} else {
				logs.Errorw("failed to parse the method in the interface", "interface", i.name)
			}
		}
	}

	return nil
}
