package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type StatementsVisitor struct {
	*BaseMojoParserVisitor

	FreeFloatingDocument *lang.Document
}

func NewStatementsVisitor() *StatementsVisitor {
	visitor := &StatementsVisitor{}
	return visitor
}

func (s *StatementsVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	statementCtxes := ctx.AllStatement()

	var statements []*lang.Statement
	var document *lang.Document
	for _, statementCtx := range statementCtxes {
		obj := statementCtx.Accept(s)
		if statement, ok := obj.(*lang.Statement); ok && statement != nil {
			if document != nil {
				lang.SetStartPosition(statement, &lang.Position{LeadingComments: lang.NewComments(document)})
				document = nil
			}

			statements = append(statements, statement)
		}
		if doc, ok := obj.(*lang.Document); ok && doc != nil {
			document = doc
		}
	}

	if document != nil {
		s.FreeFloatingDocument = document
	}
	return statements
}

func (s *StatementsVisitor) VisitStatement(ctx *StatementContext) interface{} {
	if declCtx := ctx.Declaration(); declCtx != nil {
		return declCtx.Accept(s)
	}

	if expressionCtx := ctx.Expression(); expressionCtx != nil {
		if expression, ok := expressionCtx.Accept(NewExpressionVisitor()).(*lang.Expression); ok && expression != nil {
			return lang.NewExpressionStatement(expression)
		}
	}

	if documentCtx := ctx.FloatingStatement(); documentCtx != nil {
		if document := GetFreeFloatingDocument(documentCtx); document != nil {
			return document
		}
	}

	return nil
}

func (s *StatementsVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	document := GetDocument(ctx.Document())
	attributes := GetAttributes(ctx.Attributes())

	var startPosition *lang.Position
	if document != nil {
		startPosition = document.StartPosition
	}
	if startPosition == nil && len(attributes) > 0 {
		startPosition = attributes[0].StartPosition
	}

	if packageCtx := ctx.PackageDeclaration(); packageCtx != nil {
		if decl, ok := packageCtx.Accept(NewPackageDeclarationVisitor()).(*lang.PackageDecl); ok {
			decl.Document = document
			decl.SetStartPosition(startPosition)
			return lang.NewPackageDeclStatement(decl)
		}
	}

	if importCtx := ctx.ImportDeclaration(); importCtx != nil {
		if decl, ok := importCtx.Accept(NewImportDeclarationVisitor()).(*lang.ImportDecl); ok {
			decl.Document = document
			decl.SetStartPosition(startPosition)
			return lang.NewImportDeclStatement(decl)
		}
	}

	if enumCtx := ctx.EnumDeclaration(); enumCtx != nil {
		if decl, ok := enumCtx.Accept(NewEnumDeclarationVisitor()).(*lang.EnumDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			decl.SetStartPosition(startPosition)
			return lang.NewEnumDeclStatement(decl)
		}
	}

	if structCtx := ctx.StructDeclaration(); structCtx != nil {
		if decl, ok := structCtx.Accept(NewStructDeclarationVisitor()).(*lang.StructDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			decl.SetStartPosition(startPosition)
			return lang.NewStructDeclStatement(decl)
		}
	}

	if interfaceCtx := ctx.InterfaceDeclaration(); interfaceCtx != nil {
		if decl, ok := interfaceCtx.Accept(NewInterfaceDeclarationVisitor()).(*lang.InterfaceDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			decl.SetStartPosition(startPosition)
			return lang.NewInterfaceDeclStatement(decl)
		}
	}

	if typeAliasCtx := ctx.TypeAliasDeclaration(); typeAliasCtx != nil {
		if decl, ok := typeAliasCtx.Accept(NewTypeAliasDeclarationVisitor()).(*lang.TypeAliasDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			decl.SetStartPosition(startPosition)
			return lang.NewTypeAliasDeclStatement(decl)
		}
	}

	if functionCtx := ctx.FunctionDeclaration(); functionCtx != nil {
		if decl, ok := functionCtx.Accept(NewFuncDeclarationVisitor()).(*lang.FunctionDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			decl.SetStartPosition(startPosition)
			return lang.NewFunctionDeclStatement(decl)
		}
	}

	if variableDeclaration := ctx.VariableDeclaration(); variableDeclaration != nil {
		if decl, ok := variableDeclaration.Accept(NewValueDeclarationVisitor()).(*lang.VariableDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			decl.SetStartPosition(startPosition)
			return lang.NewVariableDeclStatement(decl)
		}
	}

	if attributeDeclaration := ctx.AttributeDeclaration(); attributeDeclaration != nil {
		if decl, ok := attributeDeclaration.Accept(NewAttributeDeclarationVisitor()).(*lang.AttributeDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			decl.SetStartPosition(startPosition)
			return lang.NewAttributeDeclStatement(decl)
		}
	}

	if attributeAliasDeclaration := ctx.AttributeAliasDeclaration(); attributeAliasDeclaration != nil {
		if decl, ok := attributeAliasDeclaration.Accept(NewAttributeAliasDeclarationVisitor()).(*lang.AttributeAliasDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			decl.SetStartPosition(startPosition)
			return lang.NewAttributeAliasDeclStatement(decl)
		}
	}

	return nil
}
