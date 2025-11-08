package lang

func NewFloatLiteralExpr(value float64) *FloatLiteralExpr {
	return &FloatLiteralExpr{
		Kind:  0,
		Value: value,
	}
}

func (x *FloatLiteralExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *FloatLiteralExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *FloatLiteralExpr) EvalValue() float64 {
	if x != nil {
		if x.IsNegative {
			return -x.Value
		}
		return x.Value
	}
	return 0
}

func (x *FloatLiteralExpr) SetNegative() *FloatLiteralExpr {
	if x != nil {
		x.IsNegative = true
	}
	return x
}
