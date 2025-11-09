package lang

func (x *StringPrefixLiteralExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *StringPrefixLiteralExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *StringPrefixLiteralExpr) GetStringLiteralExpr() *StringLiteralExpr {
	return x.GetArgument().GetStringLiteralExpr()
}

func (x *StringPrefixLiteralExpr) SetCallee(operator *Operator) *StringPrefixLiteralExpr {
	return x.SetOperator(operator)
}

func (x *StringPrefixLiteralExpr) SetOperator(operator *Operator) *StringPrefixLiteralExpr {
	if x != nil && operator != nil {
		x.Callee = &StringPrefixLiteralExpr_Operator{Operator: operator}
	}
	return x
}
