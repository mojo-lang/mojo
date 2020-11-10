package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type GenericParametersVisitor struct {
	*BaseMojoParserVisitor
}

func NewGenericParametersVisitor() *GenericParametersVisitor {
	return &GenericParametersVisitor{}
}

func GetGenericParameters(ctx IGenericParameterClauseContext) []*lang.GenericParameter {
	if ctx != nil {
		visitor := NewGenericParametersVisitor()
		return ctx.Accept(visitor).([]*lang.GenericParameter)
	}
	return []*lang.GenericParameter{}
}

func (a *GenericParametersVisitor) VisitGenericParameterClause(ctx *GenericParameterClauseContext) interface{} {
	parameterList := ctx.GenericParameters()
	if parameterList != nil {
		return parameterList.Accept(a)
	}

	return []*lang.GenericParameter{}
}

func (a *GenericParametersVisitor) VisitGenericParameters(ctx *GenericParametersContext) interface{} {
	var params []*lang.GenericParameter
	parameters := ctx.AllGenericParameter()
	for _, parameterCtx := range parameters {
		visitor := NewGenericParameterVisitor()
		param := parameterCtx.Accept(visitor).(*lang.GenericParameter)

		if param != nil {
			params = append(params, param)
		}
	}
	return params
}
