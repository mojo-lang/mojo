package syntax

import (
	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"
)

type LimitSmtVisitor struct {
	*BaseSQLiteParserVisitor
}

func NewLimitSmtVisitor() *LimitSmtVisitor {
	visitor := &LimitSmtVisitor{}
	return visitor
}

//goland:noinspection GoSnakeCaseUsage
func (v *LimitSmtVisitor) VisitLimit_stmt(ctx *Limit_stmtContext) interface{} {
	allExprs := ctx.AllExpr()
	if len(allExprs) > 0 {
		limited := &sql.LimitClause{}

		if expr, ok := allExprs[0].Accept(NewExprVisitor()).(*sql.Expression); ok {
			if ctx.COMMA() != nil {
				limited.SetOffsetExpr(expr)
			} else {
				limited.Rows = expr
			}

			if len(allExprs) > 1 {
				if expr, ok = allExprs[1].Accept(NewExprVisitor()).(*sql.Expression); ok {
					if ctx.COMMA() != nil {
						limited.Rows = expr
					} else if ctx.OFFSET_() != nil {
						offset := &sql.OffsetClause{
							Value: expr,
						}
						limited.SetOffsetClause(offset)
					}
				}
			}
			return limited
		}
	}
	return nil
}
