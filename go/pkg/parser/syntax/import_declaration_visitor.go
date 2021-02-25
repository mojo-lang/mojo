package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type ImportDeclarationVisitor struct {
	*BaseMojoParserVisitor

	ImportDecl *lang.ImportDecl
}

func NewImportDeclarationVisitor() *ImportDeclarationVisitor {
	visitor := &ImportDeclarationVisitor{}
	visitor.ImportDecl = &lang.ImportDecl{}
	return visitor
}

func (i *ImportDeclarationVisitor) VisitImportDeclaration(ctx *ImportDeclarationContext) interface{} {
	return nil
}

func (i *ImportDeclarationVisitor) VisitImportPath(ctx *ImportPathContext) interface{} {
	return nil
}

func (i *ImportDeclarationVisitor) VisitImportValueAsClause(ctx *ImportValueAsClauseContext) interface{} {
	return nil
}

func (i *ImportDeclarationVisitor) VisitImportTypeClause(ctx *ImportTypeClauseContext) interface{} {
	return nil
}

func (i *ImportDeclarationVisitor) VisitImportTypeAsClause(ctx *ImportTypeAsClauseContext) interface{} {
	return nil
}

func (i *ImportDeclarationVisitor) VisitImportGroupClause(ctx *ImportGroupClauseContext) interface{} {
	return nil
}

func (i *ImportDeclarationVisitor) VisitImportGroup(ctx *ImportGroupContext) interface{} {
	return nil
}

func (i *ImportDeclarationVisitor) VisitImportValue(ctx *ImportValueContext) interface{} {
	return nil
}

func (i *ImportDeclarationVisitor) VisitImportType(ctx *ImportTypeContext) interface{} {
	return nil
}
