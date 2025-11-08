package lang

func (x *PrefixUnaryExpr) SetCallee(operator *Operator) *PrefixUnaryExpr {
	return x.SetOperator(operator)
}

func (x *PrefixUnaryExpr) SetOperator(operator *Operator) *PrefixUnaryExpr {
	if x != nil && operator != nil {
		x.Callee = &PrefixUnaryExpr_Operator{Operator: operator}
	}
	return x
}
