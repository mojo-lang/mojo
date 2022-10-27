package syntax

type CreateTableSmtVisitor struct {
	*BaseSQLiteParserVisitor
}

func NewCreateTableSmtVisitor() *CreateTableSmtVisitor {
	visitor := &CreateTableSmtVisitor{}
	return visitor
}

func (v *CreateTableSmtVisitor) VisitCreate_table_stmt(ctx *Create_table_stmtContext) interface{} {
	return nil
}
