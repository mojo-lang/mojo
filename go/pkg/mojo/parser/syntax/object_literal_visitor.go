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
		if expr, ok := ctx.Accept(visitor).(*lang.Expression); ok {
			return expr.GetObjectLiteralExpr()
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

	//TODO need to be check the validation of the object expression
	//if len(e.Expression.Fields) > 0 {
	//	fieldName := e.Expression.Fields[0].Name
	//	_, err := fieldName.EvalIdentifier()
	//	if err != nil { // NOT IDENTIFIER, change to map
	//		map := &lang.MapLiteralExpr{
	//			StartPosition: e.Expression.StartPosition,
	//			EndPosition:   e.Expression.EndPosition,
	//			Kind:          e.Expression.Kind,
	//			Implicit:      e.Expression.Implicit,
	//		}
	//
	//		for _, field := range e.Expression.Fields {
	//			entry := &lang.MapLiteralExpr_Entry{
	//				Key:   field.Name,
	//				Value: field.Value,
	//			}
	//			map.Entries = append(map.Entries, entry)
	//		}
	//
	//		return lang.NewMapLiteralExpr(map)
	//	}
	//}

	return lang.NewObjectLiteralExpression(e.Expression)
}

func (e *ObjectLiteralVisitor) VisitObjectLiteralItem(ctx *ObjectLiteralItemContext) interface{} {
	key := GetPathIdentifier(ctx.PathIdentifier())
	if len(key) > 0 {
		field := &lang.ObjectLiteralExpr_Field{
			//FIXME deal with the multi path segments situation
			Name:  key[0],
			Value: GetExpression(ctx.Expression()),
		}
		e.Expression.Fields = append(e.Expression.Fields, field)
	}

	return nil
}
