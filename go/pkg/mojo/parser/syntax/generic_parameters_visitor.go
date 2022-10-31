package syntax

import (
	"fmt"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type GenericParametersVisitor struct {
	*BaseMojoParserVisitor
}

func NewGenericParametersVisitor() *GenericParametersVisitor {
	return &GenericParametersVisitor{}
}

func GetGenericParameters(ctx IGenericParameterClauseContext) []*lang.GenericParameter {
	if ctx != nil {
		visitor := NewGenericParametersVisitor()
		if parameters, ok := ctx.Accept(visitor).([]*lang.GenericParameter); ok {
			return parameters
		} else {
			fmt.Print("===> error")
		}
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
		if param, ok := parameterCtx.Accept(visitor).(*lang.GenericParameter); ok {
			params = append(params, param)
		} else {
			fmt.Print("===> error")
		}
	}
	return params
}
