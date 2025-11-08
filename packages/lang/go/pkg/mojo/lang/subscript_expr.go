package lang

func (x *SubscriptExpr) SetCallee(expr *Expression) *SubscriptExpr {
	if x != nil && expr != nil {
		x.Callee = &SubscriptExpr_Expression{Expression: expr}
	}
	return x
}
