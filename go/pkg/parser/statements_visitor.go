package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

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

func (s *StatementsVisitor) VisitStatement(ctx *StatementContext) interface{}  {
	declCtx := ctx.Declaration()
	if declCtx != nil {
		return declCtx.Accept(s)
	}

	expressionCtx := ctx.Expression()
	if expressionCtx != nil {
		visitor := NewExpressionVisitor()
		expression := expressionCtx.Accept(visitor).(*lang.Expression)
		return lang.NewExpressionStatement(expression)
	}

	return nil
}

func (s *StatementsVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	packageCtx := ctx.PackageDeclaration()
	if packageCtx != nil {
		visitor := NewPackageDeclarationVisitor()
		decl := packageCtx.Accept(visitor).(*lang.PackageDecl)
		if decl != nil {
			return lang.NewPackageDeclStatement(decl)
		} else {
			return nil
		}
	}

	importCtx := ctx.ImportDeclaration()
	if importCtx != nil {
		visitor := NewImportDeclarationVisitor()
		decl := packageCtx.Accept(visitor).(*lang.ImportDecl)
		if decl != nil {
			return lang.NewImportDeclStatement(decl)
		} else {
			return nil
		}
	}

	enumCtx := ctx.EnumDeclaration()
	if enumCtx != nil {
		visitor := NewEnumDeclarationVisitor()
		decl := enumCtx.Accept(visitor).(*lang.EnumDecl)
		if decl != nil {
			return lang.NewEnumDeclStatement(decl)
		} else {
			return nil
		}
	}

	structCtx := ctx.StructDeclaration()
	if structCtx != nil {
		visitor := NewStructDeclarationVisitor()
		decl := structCtx.Accept(visitor).(*lang.StructDecl)
		if decl != nil {
			return lang.NewStructDeclStatement(decl)
		} else {
			return nil
		}
	}

	interfaceCtx := ctx.InterfaceDeclaration()
	if interfaceCtx != nil {
		visitor := NewInterfaceDeclarationVisitor()
		decl := interfaceCtx.Accept(visitor).(*lang.InterfaceDecl)
		if decl != nil {
			return lang.NewInterfaceDeclStatement(decl)
		} else {
			return nil
		}
	}

	//typeAliasCtx := ctx.TypeAliasDeclaration()
	//if typeAliasCtx != nil {
	//	visitor := NewTypeAliasDeclarationVisitor()
	//	decl := typeAliasCtx.Accept(visitor).(*lang.TypeAliasDecl)
	//	if decl != nil {
	//		return lang.NewTypeAliasDeclStatement(decl)
	//	} else {
	//		return nil
	//	}
	//}

	functionCtx := ctx.FunctionDeclaration()
	if functionCtx != nil {
		visitor := NewFuncDeclarationVisitor()
		decl := functionCtx.Accept(visitor).(*lang.FuncDecl)
		if decl != nil {
			return lang.NewFunctionDeclStatement(decl)
		} else {
			return nil
		}
	}

	return nil
}
