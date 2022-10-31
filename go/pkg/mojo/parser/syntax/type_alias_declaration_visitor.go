package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type TypeAliasDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewTypeAliasDeclarationVisitor() *TypeAliasDeclarationVisitor {
	visitor := &TypeAliasDeclarationVisitor{}
	return visitor
}

func (s *TypeAliasDeclarationVisitor) VisitTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) interface{} {
	if nameCtx := ctx.TypeAliasName(); nameCtx != nil {
		if name, ok := nameCtx.Accept(s).(string); ok {
			decl := &lang.TypeAliasDecl{
				StartPosition:     GetPosition(ctx.GetStart()),
				EndPosition:       GetPosition(ctx.GetStop()),
				KeywordPosition:   GetPosition(ctx.KEYWORD_TYPE().GetSymbol()),
				Name:              name,
				NamePosition:      GetPosition(nameCtx.GetStart()),
				GenericParameters: GetGenericParameters(ctx.GenericParameterClause()),
			}

			if assignmentCtx := ctx.TypeAliasAssignment(); assignmentCtx != nil {
				if decl.Type, ok = assignmentCtx.Accept(s).(*lang.NominalType); ok {
					return decl
				}
			}
		}
	}
	return nil
}

func (s *TypeAliasDeclarationVisitor) VisitTypeAliasAssignment(ctx *TypeAliasAssignmentContext) interface{} {
	nominal := GetType(ctx.Type_())
	nominal.Attributes = GetAttributes(ctx.Attributes())
	nominal.Document = GetFollowingDocument(ctx.FollowingDocument())
	return nominal
}

func (s *TypeAliasDeclarationVisitor) VisitTypeAliasName(ctx *TypeAliasNameContext) interface{} {
	return ctx.GetText()
}
