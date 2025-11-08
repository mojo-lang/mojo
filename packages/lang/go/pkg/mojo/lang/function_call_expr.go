package lang

func (x *FunctionCallExpr) SetCallee(expr *Expression) *FunctionCallExpr {
	if x != nil && expr != nil {
		x.Callee = &FunctionCallExpr_Expression{Expression: expr}
	}
	return x
}
