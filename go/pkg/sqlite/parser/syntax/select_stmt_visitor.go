package syntax

import (
	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"
)

type SelectSmtVisitor struct {
	*BaseSQLiteParserVisitor
}

func NewSelectSmtVisitor() *SelectSmtVisitor {
	visitor := &SelectSmtVisitor{}
	return visitor
}

//goland:noinspection GoSnakeCaseUsage
func (v *SelectSmtVisitor) VisitSelect_stmt(ctx *Select_stmtContext) interface{} {
	var selectStmts []*sql.SelectStmt
	cores := ctx.AllSelect_core()
	for _, core := range cores {
		if stmt, ok := core.Accept(v).(*sql.SelectStmt); ok {
			selectStmts = append(selectStmts, stmt)
		} else {
			// error
		}
	}

	var selectStmt *sql.SelectStmt
	if len(selectStmts) == 1 {
		selectStmt = selectStmts[0]
	} else if len(selectStmts) > 1 { // TODO
	}

	return &sql.Statement{Statement: &sql.Statement_SelectStmt{SelectStmt: selectStmt}}
}

//goland:noinspection GoSnakeCaseUsage
func (v *SelectSmtVisitor) VisitCommon_table_stmt(ctx *Common_table_stmtContext) interface{} {
	_ = ctx
	return nil
}

//goland:noinspection GoSnakeCaseUsage
func (v *SelectSmtVisitor) VisitSelect_core(ctx *Select_coreContext) interface{} {
	selectStmt := &sql.SelectStmt{}

	if ctx.SELECT_() != nil {
		selectClause := &sql.SelectClause{}

		if ctx.ALL_() != nil {
			selectClause.Quantifier = sql.DataSetQuantifier_DATA_SET_QUANTIFIER_ALL
		}
		if ctx.DISTINCT_() != nil {
			selectClause.Quantifier = sql.DataSetQuantifier_DATA_SET_QUANTIFIER_DISTINCT
		}

		if columns := ctx.AllResult_column(); columns != nil {
			for _, column := range columns {
				if c, ok := column.Accept(v).(*sql.Column); ok {
					selectClause.Columns = append(selectClause.Columns, c)
				}
			}
		}
		selectStmt.SetSelectClause(selectClause)
	}

	if ctx.VALUES_() != nil {
		clause := &sql.ValuesClause{}
		if exprs := ctx.AllExpr(); exprs != nil {
			for _, expr := range exprs {
				if expression, ok := expr.Accept(NewExprVisitor()).(*sql.Expression); ok {
					clause.Values = append(clause.Values, expression)
				}
			}
		}
		selectStmt.Values = clause
	}

	if ctx.FROM_() != nil {
	}
	if ctx.WHERE_() != nil {
	}
	if ctx.GROUP_() != nil && ctx.BY_() != nil {
	}
	if ctx.HAVING_() != nil {
	}
	if ctx.WINDOW_() != nil {
	}

	return selectStmt
}

//goland:noinspection GoSnakeCaseUsage
func (v *SelectSmtVisitor) VisitResult_column(ctx *Result_columnContext) interface{} {
	_ = ctx
	return nil
}

//goland:noinspection GoSnakeCaseUsage
func (v *SelectSmtVisitor) VisitCommon_table_expression(ctx *Common_table_expressionContext) interface{} {
	_ = ctx
	return nil
}
