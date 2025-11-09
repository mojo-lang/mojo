package decompiler

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/db/sql"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
)

type ExpressionDecompiler interface {
	CompileExpression(ctx context.Context, expr *lang.Expression) (*sql.Expression, error)

	CompileNullLiteralExpr(ctx context.Context, expr *lang.NullLiteralExpr) (*sql.Expression, error)
	CompileBoolLiteralExpr(ctx context.Context, expr *lang.BoolLiteralExpr) (*sql.Expression, error)
	CompileIntegerLiteralExpr(ctx context.Context, expr *lang.IntegerLiteralExpr) (*sql.Expression, error)
	CompileFloatLiteralExpr(ctx context.Context, expr *lang.FloatLiteralExpr) (*sql.Expression, error)
	CompileStringLiteralExpr(ctx context.Context, expr *lang.StringLiteralExpr) (*sql.Expression, error)

	CompileStringPrefixLiteralExpr(ctx context.Context, expr *lang.StringPrefixLiteralExpr) (*sql.Expression, error)

	CompileIdentifierExpr(ctx context.Context, expr *lang.IdentifierExpr) (*sql.Expression, error)

	CompilePrefixUnaryExpr(ctx context.Context, expr *lang.PrefixUnaryExpr) (*sql.Expression, error)
	CompileBinaryExpr(ctx context.Context, expr *lang.BinaryExpr) (*sql.Expression, error)
	CompileFunctionCallExpr(ctx context.Context, expr *lang.FunctionCallExpr) (*sql.Expression, error)

	CompileParenthesizedExpr(ctx context.Context, expr *lang.ParenthesizedExpr) (*sql.Expression, error)
}
