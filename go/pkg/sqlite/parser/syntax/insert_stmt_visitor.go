package syntax

type InsertSmtVisitor struct {
	*BaseSQLiteParserVisitor
}

func NewInsertSmtVisitor() *InsertSmtVisitor {
	visitor := &InsertSmtVisitor{}
	return visitor
}

//goland:noinspection GoSnakeCaseUsage
func (v *InsertSmtVisitor) VisitInsert_stmt(ctx *Insert_stmtContext) interface{} {
	return nil
}
