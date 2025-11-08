package lang

func (x *ExplicitMemberExpr) SetCallee(expr *Expression) *ExplicitMemberExpr {
	if x != nil && expr != nil {
		x.Callee = &ExplicitMemberExpr_Expression{Expression: expr}
	}
	return x
}
