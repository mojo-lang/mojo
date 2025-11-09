package sql

func (x *LimitClause) SetOffsetExpr(offset *Expression) *LimitClause {
	if x != nil && offset != nil {
		x.Offset = &LimitClause_Expression{Expression: offset}
	}
	return x
}

func (x *LimitClause) SetOffsetClause(offset *OffsetClause) *LimitClause {
	if x != nil && offset != nil {
		x.Offset = &LimitClause_OffsetClause{OffsetClause: offset}
	}
	return x
}
