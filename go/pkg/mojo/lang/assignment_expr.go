package lang

func (x *AssignmentExpr) SetCallee(operator *Operator) *AssignmentExpr {
	return x.SetOperator(operator)
}

func (x *AssignmentExpr) SetOperator(operator *Operator) *AssignmentExpr {
	if x != nil && operator != nil {
		x.Callee = &AssignmentExpr_Operator{Operator: operator}
	}
	return x
}
