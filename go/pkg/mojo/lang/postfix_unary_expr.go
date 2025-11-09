package lang

func (x *PostfixUnaryExpr) SetCallee(operator *Operator) *PostfixUnaryExpr {
	return x.SetOperator(operator)
}

func (x *PostfixUnaryExpr) SetOperator(operator *Operator) *PostfixUnaryExpr {
	if x != nil && operator != nil {
		x.Callee = &PostfixUnaryExpr_Operator{Operator: operator}
	}
	return x
}
