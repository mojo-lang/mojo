package sql

func NewStringLiteralExpr(value string) *StringLiteralExpr {
	return &StringLiteralExpr{
		Kind:  0,
		Value: value,
	}
}
