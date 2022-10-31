package syntax

import (
	"fmt"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type GenericArgumentsVisitor struct {
	*BaseMojoParserVisitor
}

func NewGenericArgumentsVisitor() *GenericArgumentsVisitor {
	return &GenericArgumentsVisitor{}
}

func GetGenericArguments(ctx IGenericArgumentClauseContext) []*lang.NominalType {
	if ctx != nil {
		visitor := NewGenericArgumentsVisitor()
		if t, ok := ctx.Accept(visitor).([]*lang.NominalType); ok {
			return t
		} else {
			fmt.Print("===> error")
		}
	}
	return []*lang.NominalType{}
}

func (a *GenericArgumentsVisitor) VisitGenericArgumentClause(ctx *GenericArgumentClauseContext) interface{} {
	argumentList := ctx.GenericArguments()
	if argumentList != nil {
		return argumentList.Accept(a)
	}

	return []*lang.NominalType{}
}

func (a *GenericArgumentsVisitor) VisitGenericArguments(ctx *GenericArgumentsContext) interface{} {
	var arguments []*lang.NominalType
	argumentCtxes := ctx.AllGenericArgument()
	for _, argumentCtx := range argumentCtxes {
		visitor := NewGenericArgumentVisitor()
		if argument, ok := argumentCtx.Accept(visitor).(*lang.NominalType); ok {
			arguments = append(arguments, argument)
		} else {
			fmt.Print("===> error")
		}
	}
	return arguments
}
