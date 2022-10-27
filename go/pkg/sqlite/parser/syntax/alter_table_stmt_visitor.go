package syntax

type AlterTableSmtVisitor struct {
	*BaseSQLiteParserVisitor
}

func NewAlterTableSmtVisitor() *AlterTableSmtVisitor {
	visitor := &AlterTableSmtVisitor{}
	return visitor
}

func (v *AlterTableSmtVisitor) VisitAlter_table_stmt(ctx *Alter_table_stmtContext) interface{} {
	return nil
}
