package lang

func NewBoolLiteralExpr(value bool) *BoolLiteralExpr {
	return &BoolLiteralExpr{
		Kind:  0,
		Value: value,
	}
}

func (x *BoolLiteralExpr) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *BoolLiteralExpr) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *BoolLiteralExpr) EvalValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}
