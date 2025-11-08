package sql

func (x *SelectStmt) SetSelectClause(sc *SelectClause) *SelectStmt {
	if x != nil && sc != nil {
		x.Select = &SelectStmt_SelectClause{SelectClause: sc}
	}
	return x
}

func (x *SelectStmt) SetCompoundSelect(cs *CompoundSelect) *SelectStmt {
	if x != nil && cs != nil {
		x.Select = &SelectStmt_CompoundSelect{CompoundSelect: cs}
	}
	return x
}
