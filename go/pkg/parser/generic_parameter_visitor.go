package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type GenericParameterVisitor struct {
	*BaseMojoParserVisitor
}

func NewGenericParameterVisitor() *GenericParameterVisitor {
	return &GenericParameterVisitor{}
}

//TODO add impliment
func (a *GenericParameterVisitor) VisitGenericParameter(ctx *GenericParameterContext) interface{} {
	return &lang.GenericParameter{}
}
