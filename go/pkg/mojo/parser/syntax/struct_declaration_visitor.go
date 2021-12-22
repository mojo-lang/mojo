package syntax

import (
	"fmt"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type StructDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewStructDeclarationVisitor() *StructDeclarationVisitor {
	visitor := &StructDeclarationVisitor{}
	return visitor
}

func (s *StructDeclarationVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	structName := ctx.StructName()
	if structName != nil {
		if name, ok := structName.Accept(s).(string); ok {
			decl := &lang.StructDecl{
				StartPosition:     nil,
				EndPosition:       nil,
				Name:              name,
				GenericParameters: GetGenericParameters(ctx.GenericParameterClause()),
				EnclosingType:     nil,
			}

			enclosingType := &lang.NominalType{
				Name: name,
				TypeDeclaration: lang.NewStructTypeDeclaration(decl), // may cause reference circle
			}

			structType := ctx.StructType()
			if structType != nil {
				if len(structType.GetText()) > 0 {
					if d, ok := structType.Accept(s).(*lang.StructDecl); ok {
						decl.Type = d.Type
						decl.EnumDecls = d.EnumDecls
						decl.StructDecls = d.StructDecls
						decl.TypeAliasDecls = d.TypeAliasDecls

						for _, enumDecl := range decl.EnumDecls {
							enumDecl.EnclosingType = enclosingType
						}
						for _, structDecl := range decl.StructDecls {
							structDecl.EnclosingType = enclosingType

							for _, eDecl := range structDecl.EnumDecls {
								eDecl.EnclosingType.EnclosingType = enclosingType
							}
							for _, sDecl := range structDecl.StructDecls {
								sDecl.EnclosingType.EnclosingType = enclosingType
							}
							for _, aDecl := range structDecl.TypeAliasDecls {
								aDecl.EnclosingType.EnclosingType = enclosingType
							}
						}
						for _, typeAliasDecl := range decl.TypeAliasDecls {
							typeAliasDecl.EnclosingType = enclosingType
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
	body := ctx.StructBody()
	if body != nil {
		if d, ok := body.Accept(s).(*lang.StructDecl); ok {
			decl = d
		}
	}

	inheritances := ctx.TypeInheritanceClause()
	if inheritances != nil {
		visitor := NewTypeInheritancesVisitor()
		if t, ok := inheritances.Accept(visitor).([]*lang.NominalType); ok {
			if decl == nil {
				decl = &lang.StructDecl{Type: &lang.StructType{}}
			}
			if decl.Type == nil {
				decl.Type = &lang.StructType{}
			}
			decl.Type.Inherits = t
		} else {
			fmt.Print("===> error")
		}
	}

	if decl != nil {
		return decl
	}

	// remove the type info
	return nil
}

func (s *StructDeclarationVisitor) VisitStructName(ctx *StructNameContext) interface{} {
	return ctx.GetText()
}

func (s *StructDeclarationVisitor) VisitStructBody(ctx *StructBodyContext) interface{} {
	members := ctx.StructMembers()
	if members != nil {
		return members.Accept(s)
	}

	return &lang.StructDecl{}
}

func (s *StructDeclarationVisitor) VisitStructMembers(ctx *StructMembersContext) interface{} {
	members := ctx.AllStructMember()
	documents := ctx.AllEosWithDocument()
	if len(members) > 0 {
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
		for i, member := range members {
			if len(documents) > 0 {
				document = GetEosDocument(documents[i])
			}
			switch m := member.Accept(s).(type) {
			case *lang.StructDecl:
				decl.StructDecls = append(decl.StructDecls, m)
			case *lang.EnumDecl:
				decl.EnumDecls = append(decl.EnumDecls, m)
			case *lang.TypeAliasDecl:
				decl.TypeAliasDecls = append(decl.TypeAliasDecls, m)
				m.Document = appendDocumentLine(m.Document, document)
			case *lang.ValueDecl:
				decl.Type.Fields = append(decl.Type.Fields, m)
				m.Document = appendDocumentLine(m.Document, document)
			default:
			}
		}
		return decl
	}

	return nil
}

func (s *StructDeclarationVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	document := GetDocument(ctx.Document())
	attributes := GetAttributes(ctx.Attributes())

	structCtx := ctx.StructDeclaration()
	if structCtx != nil {
		visitor := NewStructDeclarationVisitor()
		if structDecl, ok := structCtx.Accept(visitor).(*lang.StructDecl); ok {
			structDecl.Document = document
			structDecl.Attributes = attributes
			return structDecl
		} else {
			fmt.Print("===> error")
		}
	}

	enumCtx := ctx.EnumDeclaration()
	if enumCtx != nil {
		visitor := NewEnumDeclarationVisitor()
		if enumDecl, ok := enumCtx.Accept(visitor).(*lang.EnumDecl); ok {
			enumDecl.Document = document
			enumDecl.Attributes = attributes
			return enumDecl
		} else {
			fmt.Print("===> error")
		}
	}

	typeAliasCtx := ctx.TypeAliasDeclaration()
	if typeAliasCtx != nil {
		visitor := NewTypeAliasDeclarationVisitor()
		if typeAlias, ok := typeAliasCtx.Accept(visitor).(*lang.TypeAliasDecl); ok {
			typeAlias.Document = document
			typeAlias.Attributes = attributes
			return typeAlias
		} else {
			fmt.Print("===> error")
		}
	}

	memberCtx := ctx.StructMemberDeclaration()
	if memberCtx != nil {
		visitor := NewValueDeclarationVisitor()
		if memberDecl, ok := memberCtx.Accept(visitor).(*lang.ValueDecl); ok {
			memberDecl.Document = document
			memberDecl.Attributes = attributes
			return memberDecl
		} else {
			fmt.Print("===> error")
		}
	}

	return nil
}
