package syntax

type FunctionDefinitionVisitor struct {
	*BaseCVisitor
}

func NewFunctionDefinitionVisitor() *FunctionDefinitionVisitor {
	visitor := &FunctionDefinitionVisitor{}
	return visitor
}

func (v *FunctionDefinitionVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return nil
}
