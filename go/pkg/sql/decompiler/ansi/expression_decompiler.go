package ansi

import (
	"context"

	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type ExpressionDecompiler struct {
}

func (d *ExpressionDecompiler) CompileExpression(ctx context.Context, expr *lang.Expression) (*sql.Expression, error) {
	switch expression := expr.Expression.(type) {
	case *lang.Expression_NullLiteralExpr:
		return d.CompileNullLiteralExpr(ctx, expression.NullLiteralExpr)
	case *lang.Expression_BoolLiteralExpr:
		return d.CompileBoolLiteralExpr(ctx, expression.BoolLiteralExpr)
	case *lang.Expression_IntegerLiteralExpr:
		return d.CompileIntegerLiteralExpr(ctx, expression.IntegerLiteralExpr)
	case *lang.Expression_FloatLiteralExpr:
		return d.CompileFloatLiteralExpr(ctx, expression.FloatLiteralExpr)
	case *lang.Expression_StringLiteralExpr:
		return d.CompileStringLiteralExpr(ctx, expression.StringLiteralExpr)
	case *lang.Expression_StringPrefixLiteralExpr:
		return d.CompileStringPrefixLiteralExpr(ctx, expression.StringPrefixLiteralExpr)
	case *lang.Expression_IdentifierExpr:
		return d.CompileIdentifierExpr(ctx, expression.IdentifierExpr)
	case *lang.Expression_PrefixUnaryExpr:
		return d.CompilePrefixUnaryExpr(ctx, expression.PrefixUnaryExpr)
	case *lang.Expression_BinaryExpr:
		return d.CompileBinaryExpr(ctx, expression.BinaryExpr)
	case *lang.Expression_FunctionCallExpr:
		return d.CompileFunctionCallExpr(ctx, expression.FunctionCallExpr)
	case *lang.Expression_ParenthesizedExpr:
		return d.CompileParenthesizedExpr(ctx, expression.ParenthesizedExpr)
	}
	return nil, nil
}

func (d *ExpressionDecompiler) CompileNullLiteralExpr(ctx context.Context, expr *lang.NullLiteralExpr) (*sql.Expression, error) {
	_ = ctx

	return sql.NewNullLiteralExpressionFrom(), nil
}

func (d *ExpressionDecompiler) CompileBoolLiteralExpr(ctx context.Context, expr *lang.BoolLiteralExpr) (*sql.Expression, error) {
	_ = ctx
	var bv sql.Boolean
	if expr.GetValue() {
		bv = sql.Boolean_BOOLEAN_TRUE
	} else {
		bv = sql.Boolean_BOOLEAN_FALSE
	}
	return sql.NewBooleanLiteralExpressionFrom(bv), nil
}

func (d *ExpressionDecompiler) CompileIntegerLiteralExpr(ctx context.Context, expr *lang.IntegerLiteralExpr) (*sql.Expression, error) {
	return nil, nil
}

func (d *ExpressionDecompiler) CompileFloatLiteralExpr(ctx context.Context, expr *lang.FloatLiteralExpr) (*sql.Expression, error) {
	return nil, nil
}

func (d *ExpressionDecompiler) CompileStringLiteralExpr(ctx context.Context, expr *lang.StringLiteralExpr) (*sql.Expression, error) {
	return nil, nil
}

func (d *ExpressionDecompiler) CompileStringPrefixLiteralExpr(ctx context.Context, expr *lang.StringPrefixLiteralExpr) (*sql.Expression, error) {
	return nil, nil
}

func (d *ExpressionDecompiler) CompileIdentifierExpr(ctx context.Context, expr *lang.IdentifierExpr) (*sql.Expression, error) {
	return nil, nil
}

func (d *ExpressionDecompiler) CompilePrefixUnaryExpr(ctx context.Context, expr *lang.PrefixUnaryExpr) (*sql.Expression, error) {
	return nil, nil
}

func (d *ExpressionDecompiler) CompileBinaryExpr(ctx context.Context, expr *lang.BinaryExpr) (*sql.Expression, error) {
	return nil, nil
}

func (d *ExpressionDecompiler) CompileFunctionCallExpr(ctx context.Context, expr *lang.FunctionCallExpr) (*sql.Expression, error) {
	return nil, nil
}

func (d *ExpressionDecompiler) CompileParenthesizedExpr(ctx context.Context, expr *lang.ParenthesizedExpr) (*sql.Expression, error) {
	return nil, nil
}
