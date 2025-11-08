package sql

func NewNullLiteralExpression(expr *NullLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_NullLiteralExpr{NullLiteralExpr: expr},
	}
}

func NewNullLiteralExpressionFrom() *Expression {
	return NewNullLiteralExpression(NewNullLiteralExpr())
}

func NewBooleanLiteralExpression(expr *BooleanLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_BooleanLiteralExpr{BooleanLiteralExpr: expr},
	}
}

func NewBooleanLiteralExpressionFrom(value Boolean) *Expression {
	return NewBooleanLiteralExpression(NewBooleanLiteralExpr(value))
}

func NewIntegerLiteralExpression(expr *IntegerLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_IntegerLiteralExpr{IntegerLiteralExpr: expr},
	}
}

func NewIntegerLiteralExpressionFrom(value int64) *Expression {
	return NewIntegerLiteralExpression(NewIntegerLiteralExpr(value))
}

func NewFloatLiteralExpression(expr *FloatLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_FloatLiteralExpr{FloatLiteralExpr: expr},
	}
}

func NewFloatLiteralExpressionFrom(value float64) *Expression {
	return NewFloatLiteralExpression(NewFloatLiteralExpr(value))
}

func NewStringLiteralExpression(expr *StringLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_StringLiteralExpr{StringLiteralExpr: expr},
	}
}

func NewStringLiteralExpressionFrom(value string) *Expression {
	return NewStringLiteralExpression(NewStringLiteralExpr(value))
}

func NewIdentifierExpression(expr *IdentifierExpr) *Expression {
	return &Expression{
		Expression: &Expression_IdentifierExpr{IdentifierExpr: expr},
	}
}

func NewIdentifierExpressionFrom(name string) *Expression {
	return &Expression{
		Expression: &Expression_IdentifierExpr{IdentifierExpr: &IdentifierExpr{Identifier: name}},
	}
}
