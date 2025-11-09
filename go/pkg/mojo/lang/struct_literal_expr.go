package lang

func (x *StructLiteralExpr) SetCallee(expr *Expression) *StructLiteralExpr {
	if x != nil && expr != nil {
		x.Callee = &StructLiteralExpr_Expression{Expression: expr}
	}
	return x
}
