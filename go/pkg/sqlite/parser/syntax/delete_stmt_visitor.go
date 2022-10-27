package syntax

type DeleteSmtVisitor struct {
	*BaseSQLiteParserVisitor
}

func NewDeleteSmtVisitor() *DeleteSmtVisitor {
	visitor := &DeleteSmtVisitor{}
	return visitor
}

func (v *DeleteSmtVisitor) VisitDelete_stmt(ctx *Delete_stmtContext) interface{} {
	return nil
}

func (v *DeleteSmtVisitor) VisitDelete_stmt_limited(ctx *Delete_stmt_limitedContext) interface{} {
	return nil
}
