package sql

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
