package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type ArrayLiteralVisitor struct {
	*BaseMojoParserVisitor

	Expression *lang.ArrayLiteralExpr
}

func NewArrayLiteralVisitor() *ArrayLiteralVisitor {
	visitor := &ArrayLiteralVisitor{}
	return visitor
}

func (e *ArrayLiteralVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return nil
}

func (e *ArrayLiteralVisitor) VisitArrayLiteralItems(ctx *ArrayLiteralItemsContext) interface{} {
	return nil
}
