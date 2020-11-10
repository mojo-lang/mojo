package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type FuncDeclarationVisitor struct {
	*BaseMojoParserVisitor
}

func NewFuncDeclarationVisitor() *FuncDeclarationVisitor {
	decl := &FuncDeclarationVisitor{}
	return decl
}

func (f *FuncDeclarationVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	decl := &lang.FuncDecl{}

	nameCtx := ctx.FunctionName()
	if nameCtx != nil {
		decl.Name = nameCtx.Accept(f).(string)
	}

	decl.GenericParameters = GetGenericParameters(ctx.GenericParameterClause())

	signatureCtx := ctx.FunctionSignature()
	if signatureCtx != nil {
		decl.Type = signatureCtx.Accept(f).(*lang.FuncType)
	}

	return decl
}

func (f *FuncDeclarationVisitor) VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{} {
	decl := &lang.FuncDecl{}
	
	decl.Document = GetDocument(ctx.Document())
	decl.Attributes = GetAttributes(ctx.Attributes())

	nameCtx := ctx.FunctionName()
	if nameCtx != nil {
		decl.Name = nameCtx.Accept(f).(string)
	}

	decl.GenericParameters = GetGenericParameters(ctx.GenericParameterClause())

	signatureCtx := ctx.FunctionSignature()
	if signatureCtx != nil {
		decl.Type = signatureCtx.Accept(f).(*lang.FuncType)
	}

	return decl
}

func (f *FuncDeclarationVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return ctx.GetText()
}

func (f *FuncDeclarationVisitor) VisitFunctionSignature(ctx *FunctionSignatureContext) interface{} {
	funcType := &lang.FuncType{}
	
	parameterCtx := ctx.FunctionParameterClause()
	if parameterCtx != nil {
		funcType.Parameters = parameterCtx.Accept(f).([]*lang.ValueDecl)
	}

	resultCtx := ctx.FunctionResult()
	if resultCtx != nil {
		funcType.ReturnType = resultCtx.Accept(f).(*lang.NominalType)
	}

	return funcType
}

func (f *FuncDeclarationVisitor) VisitFunctionParameterClause(ctx *FunctionParameterClauseContext) interface{} {
	paramCtx := ctx.FunctionParameters()
	if paramCtx != nil {
		return paramCtx.Accept(f)
	}

	return []*lang.ValueDecl{}
}

func (f *FuncDeclarationVisitor) VisitFunctionParameters(ctx *FunctionParametersContext) interface{} {
	paramCtxes := ctx.AllFunctionParameter()
	var parameters []*lang.ValueDecl
	for _, paramCtx := range paramCtxes {
		visitor := NewValueDeclarationVisitor()
		value := paramCtx.Accept(visitor).(*lang.ValueDecl)
		if value != nil {
			parameters = append(parameters, value)
		}
	}

	return parameters
}

func (f *FuncDeclarationVisitor) VisitFunctionResult(ctx *FunctionResultContext) interface{} {
	returnType := GetType(ctx.Type_())
	if returnType != nil {
		returnType.Attributes = GetAttributes(ctx.Attributes())
	}

	return returnType
}

func (f *FuncDeclarationVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return nil
}
