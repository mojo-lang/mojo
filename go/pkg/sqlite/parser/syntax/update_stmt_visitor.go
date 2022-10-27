package syntax

type UpdateSmtVisitor struct {
	*BaseSQLiteParserVisitor
}

func NewUpdateSmtVisitor() *UpdateSmtVisitor {
	visitor := &UpdateSmtVisitor{}
	return visitor
}

func (v *UpdateSmtVisitor) VisitUpdate_stmt(ctx *Update_stmtContext) interface{} {
	return nil
}

func (v *UpdateSmtVisitor) VisitUpdate_stmt_limited(ctx *Update_stmt_limitedContext) interface{} {
	return nil
}
