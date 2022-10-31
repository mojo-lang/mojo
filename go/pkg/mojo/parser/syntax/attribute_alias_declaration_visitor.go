package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type AttributeAliasDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewAttributeAliasDeclarationVisitor() *AttributeAliasDeclarationVisitor {
	visitor := &AttributeAliasDeclarationVisitor{}
	return visitor
}

func (a *AttributeAliasDeclarationVisitor) VisitAttributeAliasDeclaration(ctx *AttributeAliasDeclarationContext) interface{} {
	if nameCtx := ctx.AttributeName(); nameCtx != nil {
		if name, ok := nameCtx.Accept(a).(string); ok {
			decl := &lang.AttributeAliasDecl{
				StartPosition:     GetPosition(ctx.GetStart()),
				EndPosition:       GetPosition(ctx.GetStop()),
				KeywordPosition:   GetPosition(ctx.KEYWORD_ATTRIBUTE().GetSymbol()),
				Name:              name,
				NamePosition:      GetPosition(nameCtx.GetStart()),
				GenericParameters: GetGenericParameters(ctx.GenericParameterClause()),
			}

			if assignmentCtx := ctx.AttributeAliasAssignment(); assignmentCtx != nil {
				if attr, ok := assignmentCtx.Accept(a).(*lang.Attribute); ok {
					decl.Attribute = attr
				}
			}

			return decl
		}
	}

	return nil
}

func (a *AttributeAliasDeclarationVisitor) VisitAttributeAliasAssignment(ctx *AttributeAliasAssignmentContext) interface{} {
	if nameCtx := ctx.AttributeName(); nameCtx != nil {
		if name, ok := nameCtx.Accept(a).(string); ok {
			attr := &lang.Attribute{
				Name:             name,
				PackageName:      GetPackageIdentifier(ctx.PackageIdentifier()),
				GenericArguments: GetGenericArguments(ctx.GenericArgumentClause()),
				Document:         GetFollowingDocument(ctx.FollowingDocument()),
			}
			return attr
		}
	}
	return nil
}

func (a *AttributeAliasDeclarationVisitor) VisitAttributeName(ctx *AttributeNameContext) interface{} {
	return ctx.GetText()
}
