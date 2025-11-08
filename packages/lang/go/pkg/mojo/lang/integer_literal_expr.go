package lang

func NewIntegerLiteralExpr(value int64) *IntegerLiteralExpr {
	isNegative := value < 0
	if isNegative {
		value = -value
	}
	return &IntegerLiteralExpr{
		Kind:       0,
		IsNegative: isNegative,
		Value:      uint64(value),
	}
}

func (x *IntegerLiteralExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *IntegerLiteralExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *IntegerLiteralExpr) EvalValue() int64 {
	if x != nil {
		if x.IsNegative {
			return -int64(x.Value)
		}
		return int64(x.Value)
	} else {
		return 0
	}
}

func (x *IntegerLiteralExpr) SetNegative() *IntegerLiteralExpr {
	if x != nil {
		x.IsNegative = true
	}
	return x
}
