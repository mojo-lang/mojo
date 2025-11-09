package lang

func NewBinaryExpr(operator *Operator, left *Expression, right *Expression) *BinaryExpr {
	return (&BinaryExpr{LeftArgument: left, RightArgument: right}).SetOperator(operator)
}

func (x *BinaryExpr) SetCallee(operator *Operator) *BinaryExpr {
	return x.SetOperator(operator)
}

func (x *BinaryExpr) SetOperator(operator *Operator) *BinaryExpr {
	if x != nil && operator != nil {
		x.Callee = &BinaryExpr_Operator{Operator: operator}
	}
	return x
}
