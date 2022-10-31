package syntax

import (
	"strconv"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type ConstantVisitor struct {
	*BaseProtobuf2Visitor
}

func NewConstantVisitor() *ConstantVisitor {
	visitor := &ConstantVisitor{}
	return visitor
}

// VisitConstant returns the *lang.Expression
func (v *ConstantVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	if fullIdent := ctx.FullIdent(); fullIdent != nil {
		if expr, ok := fullIdent.Accept(NewIdentifierVisitor()).(*lang.IdentifierExpr); ok {
			return lang.NewIdentifierExpression(expr)
		}
	}

	if intLit := ctx.IntLit(); intLit != nil {
		if expr, ok := intLit.Accept(v).(*lang.IntegerLiteralExpr); ok {
			if ctx.MINUS() != nil {
				expr.IsNegative = true
			}
			return lang.NewIntegerLiteralExpression(expr)
		}
	}

	if floatLit := ctx.FloatLit(); floatLit != nil {
		if expr, ok := floatLit.Accept(v).(*lang.FloatLiteralExpr); ok {
			if ctx.MINUS() != nil {
				expr.IsNegative = true
			}
			return lang.NewFloatLiteralExpression(expr)
		}
	}

	if strLit := ctx.StrLit(); strLit != nil {
		if expr, ok := strLit.Accept(v).(*lang.StringLiteralExpr); ok {
			return lang.NewStringLiteralExpression(expr)
		}
	}

	if boolLit := ctx.BoolLit(); boolLit != nil {
		if expr, ok := boolLit.Accept(v).(*lang.BoolLiteralExpr); ok {
			return lang.NewBoolLiteralExpression(expr)
		}
	}

	if blockLit := ctx.BlockLit(); blockLit != nil {
		if expr, ok := blockLit.Accept(v).(*lang.ObjectLiteralExpr); ok {
			return lang.NewObjectLiteralExpression(expr)
		}
	}

	return nil
}

func (v *ConstantVisitor) VisitIntLit(ctx *IntLitContext) interface{} {
	lit := ctx.GetText()
	value, _ := strconv.ParseInt(lit, 10, 64)
	return lang.NewIntegerLiteralExpr(value)
}

func (v *ConstantVisitor) VisitFloatLit(ctx *FloatLitContext) interface{} {
	lit := ctx.GetText()
	value, _ := strconv.ParseFloat(lit, 64)
	return lang.NewFloatLiteralExpr(value)
}

func (v *ConstantVisitor) VisitStrLit(ctx *StrLitContext) interface{} {
	lit := ctx.GetText()
	if len(lit) > 0 {
		switch lit[0] {
		case '\'':
			return lang.NewStringLiteralExpr(core.RemoveSingleQuote(lit))
		case '"':
			return lang.NewStringLiteralExpr(core.RemoveDoubleQuote(lit))
		}
	} else {
		return lang.NewStringLiteralExpr("")
	}
	return nil
}

func (v *ConstantVisitor) VisitBoolLit(ctx *BoolLitContext) interface{} {
	lit := ctx.GetText()
	return lang.NewBoolLiteralExpr(lit == "true")
}

func (v *ConstantVisitor) VisitBlockLit(ctx *BlockLitContext) interface{} {
	keys := ctx.AllIdent()
	values := ctx.AllConstant()

	object := &lang.ObjectLiteralExpr{}
	for i := 0; i < len(keys); i++ {
		if value, ok := values[i].Accept(v).(*lang.Expression); ok {
			object.Fields = append(object.Fields, &lang.ObjectLiteralExpr_Field{
				StartPosition: nil,
				EndPosition:   nil,
				Name:          keys[i].GetText(),
				Value:         value,
			})
		}
	}

	return object
}
