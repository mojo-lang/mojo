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
	expr := &lang.StructLiteralExpr{
		StartPosition: GetPosition(ctx.GetStart()),
		EndPosition:   GetPosition(ctx.GetStop()),
	}

	identifier := GetTypeIdentifier(ctx.TypeIdentifier())
	if identifier == nil {
		logs.Error("failed to get TypeIdentifier for StructLiteral")
		return nil
	}

	expr.SetCallee(lang.NewIdentifierExpression(lang.NewIdentifierExprFromType(identifier)))

	object := GetObjectLiteral(ctx.ObjectLiteral())
	if object != nil {
		expr.Argument = object
		return lang.NewStructLiteralExpression(expr)
	}

	return nil
}
