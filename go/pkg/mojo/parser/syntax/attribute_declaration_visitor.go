package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type AttributeDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewAttributeDeclarationVisitor() *AttributeDeclarationVisitor {
	visitor := &AttributeDeclarationVisitor{}
	return visitor
}

func (a *AttributeDeclarationVisitor) VisitAttributeDeclaration(ctx *AttributeDeclarationContext) interface{} {
	if nameCtx := ctx.AttributeName(); nameCtx != nil {
		if name, ok := nameCtx.Accept(a).(string); ok {
			decl := &lang.AttributeDecl{
				StartPosition:     GetPosition(ctx.GetStart()),
				EndPosition:       GetPosition(ctx.GetStop()),
				KeywordPosition:   GetPosition(ctx.KEYWORD_ATTRIBUTE().GetSymbol()),
				Name:              name,
				NamePosition:      GetPosition(nameCtx.GetStart()),
				GenericParameters: GetGenericParameters(ctx.GenericParameterClause()),
			}

			if nominalType := GetTypeAnnotation(ctx.TypeAnnotation()); nominalType != nil {
				nominalType.Document = GetFollowingDocument(ctx.FollowingDocument())
				decl.Type = &lang.AttributeDecl_NominalType{NominalType: nominalType}
			} else if structBodyCtx := ctx.StructBody(); structBodyCtx != nil {
				v := NewStructDeclarationVisitor()
				if t, ok := structBodyCtx.Accept(v).(*lang.StructDecl); ok && t != nil && t.Type != nil {
					decl.Type = &lang.AttributeDecl_StructType{StructType: t.Type}
				}
			}

			return decl
		}
	}

	return nil
}

func (a *AttributeDeclarationVisitor) VisitAttributeName(ctx *AttributeNameContext) interface{} {
	return ctx.GetText()
}
