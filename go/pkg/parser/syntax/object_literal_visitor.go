package syntax

import (
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type ObjectLiteralVisitor struct {
	*BaseMojoParserVisitor

	Expression *lang.ObjectLiteralExpr
}

func NewObjectLiteralVisitor() *ObjectLiteralVisitor {
	visitor := &ObjectLiteralVisitor{
		Expression: &lang.ObjectLiteralExpr{},
	}
	return visitor
}

func GetObjectLiteral(ctx IObjectLiteralContext) *lang.ObjectLiteralExpr {
	if ctx != nil {
		visitor := NewObjectLiteralVisitor()
		if expr, ok := ctx.Accept(visitor).(*lang.ObjectLiteralExpr); ok {
			return expr
		} else {
			fmt.Print("===> error")
		}
	}
	return nil
}

func (e *ObjectLiteralVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	if ctx != nil {
		context := ctx.ObjectLiteralItems()
		if context != nil {
			return context.Accept(e)
		}
	}
	return nil
}

func (e *ObjectLiteralVisitor) VisitObjectLiteralItems(ctx *ObjectLiteralItemsContext) interface{} {
	if ctx == nil {
		return nil
	}

	for _, item := range ctx.AllObjectLiteralItem() {
		item.Accept(e)
	}

	return e.Expression
}

func (e *ObjectLiteralVisitor) VisitObjectLiteralItem(ctx *ObjectLiteralItemContext) interface{} {
	field := &lang.ObjectLiteralExpr_Field{
		Name:  GetExpression(ctx.Expression(0)),
		Value: GetExpression(ctx.Expression(1)),
	}

	e.Expression.Fields = append(e.Expression.Fields, field)
	return nil
}
