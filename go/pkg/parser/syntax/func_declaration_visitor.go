package syntax

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type FuncDeclarationVisitor struct {
	*BaseMojoParserVisitor

	Decl *lang.FunctionDecl
}

func NewFuncDeclarationVisitor() *FuncDeclarationVisitor {
	decl := &FuncDeclarationVisitor{}
	return decl
}

func (f *FuncDeclarationVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	decl := &lang.FunctionDecl{}

	nameCtx := ctx.FunctionName()
	if nameCtx != nil {
		decl.Name = nameCtx.Accept(f).(string)
	}

	decl.GenericParameters = GetGenericParameters(ctx.GenericParameterClause())

	signatureCtx := ctx.FunctionSignature()
	if signatureCtx != nil {
		decl.Signature = signatureCtx.Accept(f).(*lang.FunctionDecl_Signature)
	}

	return decl
}

func (f *FuncDeclarationVisitor) VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{} {
	f.Decl = &lang.FunctionDecl{}

	nameCtx := ctx.FunctionName()
	if nameCtx != nil {
		f.Decl.Name = nameCtx.Accept(f).(string)
	} else {
		logs.Errorw("failed to parse method name in the interface")
		return nil
	}

	f.Decl.GenericParameters = GetGenericParameters(ctx.GenericParameterClause())

	signatureCtx := ctx.FunctionSignature()
	if signatureCtx != nil {
		f.Decl.Signature = signatureCtx.Accept(f).(*lang.FunctionDecl_Signature)
	} else {
		logs.Errorw("failed to parse method signature in the interface", "name", f.Decl.Name)
		return nil
	}

	return f.Decl
}

func (f *FuncDeclarationVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return ctx.GetText()
}

func (f *FuncDeclarationVisitor) VisitFunctionSignature(ctx *FunctionSignatureContext) interface{} {
	funcType := &lang.FunctionDecl_Signature{}

	parameterCtx := ctx.FunctionParameterClause()
	if parameterCtx != nil {
		funcType.Parameters = parameterCtx.Accept(f).([]*lang.ValueDecl)
	}

	document := GetFollowingDocument(ctx.FollowingDocument())
	if document != nil {
		if len(funcType.Parameters) > 0 {
			funcType.Parameters[len(funcType.Parameters)-1].Document = document
		} else {
			logs.Warnw("should not add following document after function signature without parameter")
		}
	}

	resultCtx := ctx.FunctionResult()
	if resultCtx != nil {
		funcType.Result = resultCtx.Accept(f).(*lang.NominalType)
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
	documents := ctx.AllEovWithDocument()
	var parameters []*lang.ValueDecl
	var document *lang.Document
	for i, paramCtx := range paramCtxes {
		if i < len(documents) {
			document = GetEovDocument(documents[i])
		} else {
			document = nil
		}

		visitor := NewValueDeclarationVisitor()
		if value, ok := paramCtx.Accept(visitor).(*lang.ValueDecl); ok {
			if document != nil {
				value.Document = document
			}
			parameters = append(parameters, value)
		} else {
			logs.Errorw("failed to parse the parameter in the function decl", "function", f.Decl.Name)
			return nil
		}
	}

	return parameters
}

func (f *FuncDeclarationVisitor) VisitFunctionResult(ctx *FunctionResultContext) interface{} {
	returnType := GetType(ctx.Type_())
	if returnType != nil {
		returnType.Document = GetFollowingDocument(ctx.FollowingDocument())
		returnType.Attributes = GetAttributes(ctx.Attributes())
	}

	return returnType
}

func (f *FuncDeclarationVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	documentCtx := ctx.FollowingDocument()
	if documentCtx != nil {
		document := GetFollowingDocument(documentCtx)
		if document != nil {
			if f.Decl.Signature != nil && len(f.Decl.Signature.Parameters) > 0 {
				parameters := f.Decl.Signature.Parameters
				parameters[len(parameters)-1].Document = document
			} else {
				logs.Warnw("should not add following document after function signature without parameter")
			}
		}
	}

	return nil
}
