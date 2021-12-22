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
	decl := &lang.TypeAliasDecl{}

	typeAliasName := ctx.TypeAliasName()
	if typeAliasName != nil {
		var ok bool
		if decl.Name, ok = typeAliasName.Accept(s).(string); ok {
			decl.GenericParameters = GetGenericParameters(ctx.GenericParameterClause())

			assignmentCtx := ctx.TypeAliasAssignment()
			if assignmentCtx != nil {
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
	return nominal
}

func (s *TypeAliasDeclarationVisitor) VisitTypeAliasName(ctx *TypeAliasNameContext) interface{} {
	return ctx.GetText()
}
