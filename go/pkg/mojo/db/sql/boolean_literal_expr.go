package sql

func NewBooleanLiteralExpr(value Boolean) *BooleanLiteralExpr {
	return &BooleanLiteralExpr{
		Kind:  0,
		Value: value,
	}
}
