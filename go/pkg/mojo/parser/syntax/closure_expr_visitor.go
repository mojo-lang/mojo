package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type ClosureExprVisitor struct {
	*BaseMojoParserVisitor
}

func (c *ClosureExprVisitor) VisitClosureExpression(ctx *ClosureExpressionContext) interface{} {
	closure := &lang.ClosureExpr{}

	if parametersCtx := ctx.ClosureParameters(); parametersCtx != nil {
		if params, ok := parametersCtx.Accept(c).([]*lang.ValueDecl); ok {
			closure.Signature = &lang.FunctionSignature{
				Parameter: &lang.FunctionSignature_Parameter{
					StartPosition: GetPosition(parametersCtx.GetStart()),
					EndPosition:   GetPosition(parametersCtx.GetStop()),
					Decls:         params,
				},
			}
		}
	}

	if resultCtx := ctx.Type_(); resultCtx != nil {
		if result := GetType(resultCtx); result != nil {
			closure.Signature.Result = &lang.FunctionSignature_Result{
				StartPosition: GetPosition(ctx.RIGHT_ARROW().GetSymbol()),
				EndPosition:   GetPosition(resultCtx.GetStop()),
				Type:          result,
			}
		}
	}

	if statementsCtx := ctx.Statements(); statementsCtx != nil {
		visitor := NewStatementsVisitor()
		if statements, ok := statementsCtx.Accept(visitor).([]*lang.Statement); ok {
			closure.Body = &lang.BlockStmt{
				StartPosition: GetPosition(ctx.LCURLY().GetSymbol()),
				EndPosition:   GetPosition(ctx.RCURLY().GetSymbol()),
				Statements:    statements,
			}
		}

		if visitor.FreeFloatingDocument != nil {
			closure.SetEndPosition(&lang.Position{LeadingComments: lang.NewComments(visitor.FreeFloatingDocument)})
		}
	}

	if closure.Body != nil {
		return lang.NewClosureExpression(closure)
	}
	return nil
}

func (c *ClosureExprVisitor) VisitClosureParameters(ctx *ClosureParametersContext) interface{} {
	parameterCtxes := ctx.AllClosureParameter()

	var parameters []*lang.ValueDecl
	for _, parameterCtx := range parameterCtxes {
		if decl, ok := parameterCtx.Accept(c).(*lang.ValueDecl); ok {
			parameters = append(parameters, decl)
		}
	}
	return parameters
}

func (c *ClosureExprVisitor) VisitClosureParameter(ctx *ClosureParameterContext) interface{} {
	if funcParamCtx := ctx.FunctionParameter(); funcParamCtx != nil {
		return funcParamCtx.Accept(NewValueDeclarationVisitor())
	}

	if identifier := ctx.LabelIdentifier(); identifier != nil {
		return &lang.ValueDecl{
			StartPosition: GetPosition(identifier.GetStart()),
			EndPosition:   GetPosition(identifier.GetStop()),
			Name:          identifier.GetText(),
		}
	}
	return nil
}
