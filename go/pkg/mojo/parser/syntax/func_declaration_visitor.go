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
	f.Decl = decl

	if nameCtx := ctx.FunctionName(); nameCtx != nil {
		if name, ok := nameCtx.Accept(f).(string); ok && len(name) > 0 {
			decl.Name = name
			decl.GenericParameters = GetGenericParameters(ctx.GenericParameterClause())

			if signatureCtx := ctx.FunctionSignature(); signatureCtx != nil {
				if signature, ok := signatureCtx.Accept(f).(*lang.FunctionDecl_Signature); ok && signature != nil {
					decl.Signature = signature
					return decl
				}
				logs.Errorw("failed to parse function signature", "name", decl.Name, "signature", signatureCtx.GetText())
			} else {
				logs.Errorw("failed to parse function without signature", "name", decl.Name)
			}
		}
		logs.Errorw("failed to parse function name", "code", nameCtx.GetText())
	} else {
		logs.Errorw("failed to parse function without name")
	}

	return nil
}

func (f *FuncDeclarationVisitor) VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{} {
	if nameCtx := ctx.FunctionName(); nameCtx != nil {
		if name, ok := nameCtx.Accept(f).(string); ok && len(name) > 0 {
			decl := &lang.FunctionDecl{}
			f.Decl = decl

			decl.Name = name
			decl.GenericParameters = GetGenericParameters(ctx.GenericParameterClause())
			decl.StartPosition = GetPosition(ctx.GetStart())
			decl.EndPosition = GetPosition(ctx.GetStop())

			if signatureCtx := ctx.FunctionSignature(); signatureCtx != nil {
				if signature, ok := signatureCtx.Accept(f).(*lang.FunctionDecl_Signature); ok && signature != nil {
					decl.Signature = signature
					return decl
				}
				logs.Errorw("failed to parse interface method signature", "name", decl.Name, "signature", signatureCtx.GetText())
			} else {
				logs.Errorw("failed to parse interface method without signature", "name", decl.Name)
			}
		}
		logs.Errorw("failed to parse interface method name", "code", nameCtx.GetText())
	} else {
		logs.Errorw("failed to parse interface method without name")
	}

	return nil
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
	for i, paramCtx := range paramCtxes {
		var document *lang.Document
		if i < len(documents) {
			document = GetEovDocument(documents[i])
		}

		if value, ok := paramCtx.Accept(NewValueDeclarationVisitor()).(*lang.ValueDecl); ok {
			if document != nil {
				value.Document = document
				value.SetEndPosition(document.EndPosition)
			}
			parameters = append(parameters, value)
		} else {
			logs.Errorw("failed to parse the parameter in the function decl", "function", f.Decl.Name, "code", paramCtx.GetText())
			return nil
		}
	}

	return parameters
}

func (f *FuncDeclarationVisitor) VisitFunctionResult(ctx *FunctionResultContext) interface{} {
	returnType := GetType(ctx.Type_())
	if returnType != nil {
		returnType.StartPosition = GetPosition(ctx.GetStart())
		returnType.EndPosition = GetPosition(ctx.GetStop())
		returnType.Document = GetFollowingDocument(ctx.FollowingDocument())
		returnType.Attributes = GetAttributes(ctx.Attributes())
	}

	return returnType
}

func (f *FuncDeclarationVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	if documentCtx := ctx.FollowingDocument(); documentCtx != nil {
		if document := GetFollowingDocument(documentCtx); document != nil {
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
