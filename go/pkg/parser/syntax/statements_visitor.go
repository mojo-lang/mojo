package syntax

import (
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type StatementsVisitor struct {
	*BaseMojoParserVisitor
}

func NewStatementsVisitor() *StatementsVisitor {
	visitor := &StatementsVisitor{}
	return visitor
}

func (s *StatementsVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	statementCtxes := ctx.AllStatement()

	var statements []*lang.Statement
	for _, statementCtx := range statementCtxes {
		if statement, ok := statementCtx.Accept(s).(*lang.Statement); ok {
			if statement != nil {
				statements = append(statements, statement)
			}
		}
	}
	return statements
}

func (s *StatementsVisitor) VisitStatement(ctx *StatementContext) interface{} {
	declCtx := ctx.Declaration()
	if declCtx != nil {
		return declCtx.Accept(s)
	}

	expressionCtx := ctx.Expression()
	if expressionCtx != nil {
		visitor := NewExpressionVisitor()
		if expression, ok := expressionCtx.Accept(visitor).(*lang.Expression); ok {
			return lang.NewExpressionStatement(expression)
		} else {
			fmt.Print("===> error")
		}
	}

	return nil
}

func (s *StatementsVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	document := GetDocument(ctx.Document())
	attributes := GetAttributes(ctx.Attributes())

	packageCtx := ctx.PackageDeclaration()
	if packageCtx != nil {
		visitor := NewPackageDeclarationVisitor()
		if decl, ok := packageCtx.Accept(visitor).(*lang.PackageDecl); ok {
			decl.Document = document
			return lang.NewPackageDeclStatement(decl)
		} else {
			fmt.Print("===> error")
		}
	}

	importCtx := ctx.ImportDeclaration()
	if importCtx != nil {
		visitor := NewImportDeclarationVisitor()
		if decl, ok := packageCtx.Accept(visitor).(*lang.ImportDecl); ok {
			return lang.NewImportDeclStatement(decl)
		} else {
			fmt.Print("===> error")
		}
	}

	enumCtx := ctx.EnumDeclaration()
	if enumCtx != nil {
		visitor := NewEnumDeclarationVisitor()
		if decl, ok := enumCtx.Accept(visitor).(*lang.EnumDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			return lang.NewEnumDeclStatement(decl)
		} else {
			fmt.Print("===> error")
		}
	}

	structCtx := ctx.StructDeclaration()
	if structCtx != nil {
		visitor := NewStructDeclarationVisitor()
		if decl, ok := structCtx.Accept(visitor).(*lang.StructDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			return lang.NewStructDeclStatement(decl)
		} else {
			fmt.Print("===> error")
		}
	}

	interfaceCtx := ctx.InterfaceDeclaration()
	if interfaceCtx != nil {
		visitor := NewInterfaceDeclarationVisitor()
		if decl, ok := interfaceCtx.Accept(visitor).(*lang.InterfaceDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			return lang.NewInterfaceDeclStatement(decl)
		} else {
			fmt.Print("===> error")
		}
	}

	typeAliasCtx := ctx.TypeAliasDeclaration()
	if typeAliasCtx != nil {
		visitor := NewTypeAliasDeclarationVisitor()
		if decl, ok := typeAliasCtx.Accept(visitor).(*lang.TypeAliasDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			return lang.NewTypeAliasDeclStatement(decl)
		} else {
			fmt.Print("===> error")
		}
	}

	functionCtx := ctx.FunctionDeclaration()
	if functionCtx != nil {
		visitor := NewFuncDeclarationVisitor()
		if decl, ok := functionCtx.Accept(visitor).(*lang.FunctionDecl); ok {
			decl.Document = document
			decl.Attributes = attributes
			return lang.NewFunctionDeclStatement(decl)
		} else {
			fmt.Print("===> error")
		}
	}

	//variableDeclaration := ctx.VariableDeclaration()
	//if variableDeclaration != nil {
	//	visitor := NewValueDeclarationVisitor()
	//	if decl, ok := variableDeclaration.Accept(visitor).(*lang.ValueDecl); ok {
	//		return lang.New
	//	}
	//}

	return nil
}
