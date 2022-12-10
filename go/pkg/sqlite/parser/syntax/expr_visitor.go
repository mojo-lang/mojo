package syntax

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type ExprVisitor struct {
	*BaseSQLiteParserVisitor
}

func NewExprVisitor() *ExprVisitor {
	visitor := &ExprVisitor{}
	return visitor
}

func (v *ExprVisitor) VisitExpr(ctx *ExprContext) interface{} {
	if literal := ctx.Literal_value(); literal != nil {
		return ctx.Accept(v)
	}

	return nil
}

//goland:noinspection GoSnakeCaseUsage
func (v *ExprVisitor) VisitLiteral_value(ctx *Literal_valueContext) interface{} {
	if numeric := ctx.NUMERIC_LITERAL(); numeric != nil {
	}
	if str := ctx.STRING_LITERAL(); str != nil {
		return lang.NewStringLiteralExpressionFrom(core.RemoveSingleQuote(str.GetText()))
	}
	if blob := ctx.BLOB_LITERAL(); blob != nil {
		str := blob.GetText()
		if len(str) > 0 {
			str = str[1:] // remove the `X` prefix
		}
	}
	if ctx.NULL_() != nil {
		return lang.NewNullLiteralExpressionFrom()
	}
	if ctx.TRUE_() != nil {
		return lang.NewBoolLiteralExpressionFrom(true)
	}
	if ctx.FALSE_() != nil {
		return lang.NewBoolLiteralExpressionFrom(false)
	}
	if time := ctx.CURRENT_TIME_(); time != nil {
		return lang.NewIdentifierExpressionFrom(time.GetText())
	}
	if date := ctx.CURRENT_DATE_(); date != nil {
		return lang.NewIdentifierExpressionFrom(date.GetText())
	}
	if timestamp := ctx.CURRENT_TIMESTAMP_(); timestamp != nil {
		return lang.NewIdentifierExpressionFrom(timestamp.GetText())
	}
	return nil
}
