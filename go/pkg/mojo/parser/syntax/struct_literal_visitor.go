package syntax

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type StructLiteralVisitor struct {
	*BaseMojoParserVisitor
}

func NewStructLiteralVisitor() *StructLiteralVisitor {
	visitor := &StructLiteralVisitor{}
	return visitor
}

func (s *StructLiteralVisitor) VisitStructLiteral(ctx *StructLiteralContext) interface{} {
	expr := &lang.StructLiteralExpr{}

	identifier := GetTypeIdentifier(ctx.TypeIdentifier())
	if identifier == nil {
		logs.Error("failed to get TypeIdentifier for StructLiteral")
		return nil
	}

	expr.Callee = lang.NewIdentifierExpression(lang.NewTypeIdentifierExpr(identifier))

	object := GetObjectLiteral(ctx.ObjectLiteral())
	if object != nil {
		expr.Value = object
		return lang.NewStructLiteralExpression(expr)
	}

	return nil
}
