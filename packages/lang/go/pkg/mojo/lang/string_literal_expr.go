package lang

func NewStringLiteralExpr(value string) *StringLiteralExpr {
	return &StringLiteralExpr{
		Kind:  0,
		Value: value,
	}
}

func NewImplicitStringLiteralExpr(value string) *StringLiteralExpr {
	return &StringLiteralExpr{
		Kind:     0,
		Implicit: true,
		Value:    value,
	}
}

func (x *StringLiteralExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *StringLiteralExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *StringLiteralExpr) EvalValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}
