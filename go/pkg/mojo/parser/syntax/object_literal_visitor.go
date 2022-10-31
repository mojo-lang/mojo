package syntax

import (
	"fmt"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

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

type ObjectLiteralVisitor struct {
	*BaseMojoParserVisitor
}

func NewObjectLiteralVisitor() *ObjectLiteralVisitor {
	visitor := &ObjectLiteralVisitor{}
	return visitor
}

func (e *ObjectLiteralVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	if ctx != nil {
		if context := ctx.ObjectLiteralItems(); context != nil {
			if expr, ok := context.Accept(e).(*lang.ObjectLiteralExpr); ok {
				expr.StartPosition = GetPosition(ctx.GetStart())
				expr.EndPosition = GetPosition(ctx.GetStop())
				return lang.NewObjectLiteralExpression(expr)
			}
		}
	}
	return nil
}

func (e *ObjectLiteralVisitor) VisitObjectLiteralItems(ctx *ObjectLiteralItemsContext) interface{} {
	if ctx == nil {
		return nil
	}

	expr := &lang.ObjectLiteralExpr{}
	for _, item := range ctx.AllObjectLiteralItem() {
		if field, ok := item.Accept(e).(*lang.ObjectLiteralExpr_Field); ok {
			expr.Fields = append(expr.Fields, field)
		}
	}

	// TODO need to be check the validation of the object expression
	// if len(e.Expression.Fields) > 0 {
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
	// }

	return expr
}

func (e *ObjectLiteralVisitor) VisitObjectLiteralItem(ctx *ObjectLiteralItemContext) interface{} {
	key := GetPathIdentifier(ctx.PathIdentifier())
	if len(key) > 0 {
		return &lang.ObjectLiteralExpr_Field{
			StartPosition: GetPosition(ctx.GetStart()),
			EndPosition:   GetPosition(ctx.GetStop()),
			Name:          key[0], // FIXME deal with the multi path segments situation
			Value:         GetExpression(ctx.Expression()),
		}
	}
	return nil
}
