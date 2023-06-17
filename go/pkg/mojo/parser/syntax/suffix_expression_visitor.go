package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type SuffixExpressionVisitor struct {
	*BaseMojoParserVisitor
	PrimaryExpression *lang.Expression
}

func (e *SuffixExpressionVisitor) VisitSuffixExpression(ctx *SuffixExpressionContext) interface{} {
	if functionCall := ctx.FunctionCallSuffix(); functionCall != nil {
		if expr, ok := functionCall.Accept(e).(*lang.FunctionCallExpr); ok {
			expr.SetCallee(e.PrimaryExpression)
			return lang.NewFunctionCallExpression(expr)
		}
	}

	if explicitMember := ctx.ExplicitMemberSuffix(); explicitMember != nil {
		if expr, ok := explicitMember.Accept(e).(*lang.ExplicitMemberExpr); ok {
			expr.SetCallee(e.PrimaryExpression)
			return lang.NewExplicitMemberExpression(expr)
		}
	}

	if subscript := ctx.SubscriptSuffix(); subscript != nil {
		if expr, ok := subscript.Accept(e).(*lang.SubscriptExpr); ok {
			expr.SetCallee(e.PrimaryExpression)
			return lang.NewSubscriptExpression(expr)
		}
	}

	return nil
}

func (e *SuffixExpressionVisitor) VisitFunctionCallSuffix(ctx *FunctionCallSuffixContext) interface{} {
	if clauseCtx := ctx.FunctionCallArgumentClause(); clauseCtx != nil {
		if arguments, ok := clauseCtx.Accept(e).([]*lang.Argument); ok {
			return &lang.FunctionCallExpr{Arguments: arguments}
		}
	}
	return nil
}

func (e *SuffixExpressionVisitor) VisitFunctionCallArgumentClause(ctx *FunctionCallArgumentClauseContext) interface{} {
	if argumentsCtx := ctx.FunctionCallArguments(); argumentsCtx != nil {
		return argumentsCtx.Accept(e)
	}
	return nil
}

func (e *SuffixExpressionVisitor) VisitFunctionCallArguments(ctx *FunctionCallArgumentsContext) interface{} {
	argumentCtxes := ctx.AllFunctionCallArgument()
	var arguments []*lang.Argument
	for _, argumentCtx := range argumentCtxes {
		if argument, ok := argumentCtx.Accept(e).(*lang.Argument); ok {
			arguments = append(arguments, argument)
		}
	}
	return arguments
}

func (e *SuffixExpressionVisitor) VisitFunctionCallArgument(ctx *FunctionCallArgumentContext) interface{} {
	argument := &lang.Argument{}

	if labelCtx := ctx.LabelIdentifier(); labelCtx != nil {
		argument.Label = labelCtx.GetText()
	}

	if expressionCtx := ctx.Expression(); expressionCtx != nil {
		argument.Value = GetExpression(expressionCtx)
		return argument
	}
	return nil
}

func (e *SuffixExpressionVisitor) VisitExplicitMemberSuffix(ctx *ExplicitMemberSuffixContext) interface{} {
	if identifier := ctx.Identifier(); identifier != nil {
		return &lang.ExplicitMemberExpr{Member: identifier.GetText()}
	}

	if decimal := ctx.DECIMAL_LITERAL(); decimal != nil {
		return &lang.ExplicitMemberExpr{Member: decimal.GetText()}
	}

	return nil
}

func (e *SuffixExpressionVisitor) VisitSubscriptSuffix(ctx *SubscriptSuffixContext) interface{} {
	if argumentsCtx := ctx.FunctionCallArguments(); argumentsCtx != nil {
		if arguments, ok := argumentsCtx.Accept(e).([]*lang.Argument); ok {
			return &lang.SubscriptExpr{Arguments: arguments}
		}
	}
	return nil
}
