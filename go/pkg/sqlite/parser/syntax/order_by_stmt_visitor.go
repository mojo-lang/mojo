package syntax

import (
	"strings"

	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"
)

type OrderBySmtVisitor struct {
	*BaseSQLiteParserVisitor
}

func NewOrderBySmtVisitor() *OrderBySmtVisitor {
	visitor := &OrderBySmtVisitor{}
	return visitor
}

//goland:noinspection GoSnakeCaseUsage
func (v *OrderBySmtVisitor) VisitOrder_by_stmt(ctx *Order_by_stmtContext) interface{} {
	orderBy := &sql.OrderByClause{}

	allTerms := ctx.AllOrdering_term()
	for _, term := range allTerms {
		if ordering, ok := term.Accept(v).(*sql.Ordering); ok {
			orderBy.Orderings = append(orderBy.Orderings, ordering)
		}
	}
	return orderBy
}

//goland:noinspection GoSnakeCaseUsage
func (v *OrderBySmtVisitor) VisitOrdering_term(ctx *Ordering_termContext) interface{} {
	ordering := &sql.Ordering{}
	if expr, ok := ctx.Expr().Accept(NewExprVisitor()).(*sql.Expression); ok {
		ordering.Expression = expr
	}

	if collation := ctx.Collation_name(); collation != nil {
		ordering.Collation = collation.GetText()
	}

	if order := ctx.Asc_desc(); order != nil {
		ordering.Order, _ = sql.ParseOrder(strings.ToLower(order.GetText()))
	}

	if ctx.NULLS_() != nil {
		if ctx.FIRST_() != nil {
			ordering.NullsFirst = true
		} else if ctx.LAST_() != nil {
			ordering.NullsLast = true
		}
	}

	return ordering
}
