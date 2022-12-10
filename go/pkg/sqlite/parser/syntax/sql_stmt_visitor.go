package syntax

import "github.com/mojo-lang/db/go/pkg/mojo/db/sql"

type SqlSmtVisitor struct {
	*BaseSQLiteParserVisitor
}

func NewSqlSmtVisitor() *SqlSmtVisitor {
	visitor := &SqlSmtVisitor{}
	return visitor
}

func (v *SqlSmtVisitor) VisitParse(ctx *ParseContext) interface{} {
	source := &sql.SourceFile{}

	allList := ctx.AllSql_stmt_list()
	for _, list := range allList {
		if statements, ok := list.Accept(v).([]*sql.Statement); ok {
			source.Statements = append(source.Statements, statements...)
		}
	}
	return source
}

//goland:noinspection GoSnakeCaseUsage
func (v *SqlSmtVisitor) VisitSql_stmt_list(ctx *Sql_stmt_listContext) interface{} {
	allStatements := ctx.AllSql_stmt()

	var stmts []*sql.Statement
	for _, statement := range allStatements {
		if stmt, ok := statement.Accept(v).(*sql.Statement); ok {
			stmts = append(stmts, stmt)
		}
	}
	return stmts
}

//goland:noinspection GoSnakeCaseUsage
func (v *SqlSmtVisitor) VisitSql_stmt(ctx *Sql_stmtContext) interface{} {
	if selectStmt := ctx.Select_stmt(); selectStmt != nil {
		return selectStmt.Accept(NewSelectSmtVisitor())
	}
	if updateStmt := ctx.Update_stmt(); updateStmt != nil {
		return updateStmt.Accept(NewUpdateSmtVisitor())
	}
	if updateStmtLimited := ctx.Update_stmt_limited(); updateStmtLimited != nil {
		return updateStmtLimited.Accept(NewUpdateSmtVisitor())
	}
	if insertStmt := ctx.Insert_stmt(); insertStmt != nil {
		return insertStmt.Accept(NewInsertSmtVisitor())
	}
	if deleteStmt := ctx.Delete_stmt_limited(); deleteStmt != nil {
		return deleteStmt.Accept(NewDeleteSmtVisitor())
	}
	if createTableStmt := ctx.Create_table_stmt(); createTableStmt != nil {
		return createTableStmt.Accept(NewCreateTableSmtVisitor())
	}
	if alterTableStmt := ctx.Alter_table_stmt(); alterTableStmt != nil {
		return alterTableStmt.Accept(NewAlterTableSmtVisitor())
	}
	return nil
}
