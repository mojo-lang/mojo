package lang

func (x *NumericSuffixLiteralExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *NumericSuffixLiteralExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *NumericSuffixLiteralExpr) GetIntegerLiteralExpr() *IntegerLiteralExpr {
	return x.GetArgument().GetIntegerLiteralExpr()
}

func (x *NumericSuffixLiteralExpr) GetFloatLiteralExpr() *FloatLiteralExpr {
	return x.GetArgument().GetFloatLiteralExpr()
}

func (x *NumericSuffixLiteralExpr) SetCallee(operator *Operator) *NumericSuffixLiteralExpr {
	return x.SetOperator(operator)
}

func (x *NumericSuffixLiteralExpr) SetOperator(operator *Operator) *NumericSuffixLiteralExpr {
	if x != nil && operator != nil {
		x.Callee = &NumericSuffixLiteralExpr_Operator{Operator: operator}
	}
	return x
}
