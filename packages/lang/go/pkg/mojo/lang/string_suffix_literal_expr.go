package lang

func (x *StringSuffixLiteralExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *StringSuffixLiteralExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *StringSuffixLiteralExpr) GetStringLiteralExpr() *StringLiteralExpr {
	return x.GetArgument().GetStringLiteralExpr()
}

func (x *StringSuffixLiteralExpr) SetCallee(operator *Operator) *StringSuffixLiteralExpr {
	return x.SetOperator(operator)
}

func (x *StringSuffixLiteralExpr) SetOperator(operator *Operator) *StringSuffixLiteralExpr {
	if x != nil && operator != nil {
		x.Callee = &StringSuffixLiteralExpr_Operator{Operator: operator}
	}
	return x
}
