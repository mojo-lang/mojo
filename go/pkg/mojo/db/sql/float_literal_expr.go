package sql

func NewFloatLiteralExpr(value float64) *FloatLiteralExpr {
	return &FloatLiteralExpr{
		Kind:  0,
		Value: value,
	}
}
