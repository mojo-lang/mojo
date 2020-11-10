package parser

import "github.com/mojo-lang/lang/go/pkg/lang"

type ObjectLiteralVisitor struct {
	*BaseMojoParserVisitor

	Expression *lang.ObjectLiteralExpr
}

func NewObjectLiteralVisitor() *ObjectLiteralVisitor {
	visitor := &ObjectLiteralVisitor{}
	return visitor
}

func GetObjectLiteral(ctx IObjectLiteralContext) *lang.ObjectLiteralExpr {
	if ctx != nil {
		visitor := NewObjectLiteralVisitor()
		return ctx.Accept(visitor).(*lang.ObjectLiteralExpr)
	}
	return nil
}

func (e *ObjectLiteralVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return nil
}

func (e *ObjectLiteralVisitor) VisitObjectLiteralItems(ctx *ObjectLiteralItemsContext) interface{} {
	return nil
}
