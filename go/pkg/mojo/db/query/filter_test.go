package query

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func binaryFilter(op string, left, right *lang.Expression) *lang.Expression {
	return lang.NewBinaryExpression(lang.NewBinaryExpr(&lang.Operator{Symbol: op}, left, right))
}

func TestGenerateExpressionQueryEmptyAndOperators(t *testing.T) {
	for _, expr := range []*lang.Expression{nil, {}} {
		require.NotPanics(t, func() {
			query, args, err := GenerateExpressionQuery("", expr, nil)
			require.NoError(t, err)
			require.Empty(t, query)
			require.Empty(t, args)
		})
	}
	for _, tt := range []struct {
		op, want string
		right    *lang.Expression
	}{
		{"==", "age = 20", lang.NewIntegerLiteralExpressionFrom(20)},
		{"!=", "age IS NOT NULL", lang.NewNullLiteralExpressionFrom()},
		{"==", "age IS NULL", lang.NewNullLiteralExpressionFrom()},
		{"in", "age IN (10, 20)", lang.NewArrayLiteralExpressionFrom(lang.NewIntegerLiteralExpressionFrom(10), lang.NewIntegerLiteralExpressionFrom(20))},
		{"!in", "age NOT IN (10, 20)", lang.NewArrayLiteralExpressionFrom(lang.NewIntegerLiteralExpressionFrom(10), lang.NewIntegerLiteralExpressionFrom(20))},
		{"!in", "age NOT BETWEEN 10 AND 20", binaryFilter("..=", lang.NewIntegerLiteralExpressionFrom(10), lang.NewIntegerLiteralExpressionFrom(20))},
	} {
		t.Run(tt.want, func(t *testing.T) {
			got, _, err := GenerateExpressionQuery("sqlite", binaryFilter(tt.op, lang.NewIdentifierExpressionFrom("age"), tt.right), nil)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestGenerateExpressionQueryRejectsUnsupportedArguments(t *testing.T) {
	for _, expr := range []*lang.Expression{
		binaryFilter("+", lang.NewIntegerLiteralExpressionFrom(1), lang.NewIntegerLiteralExpressionFrom(2)),
		lang.NewArrayLiteralExpressionFrom(),
		binaryFilter("in", lang.NewIdentifierExpressionFrom("age"), lang.NewIntegerLiteralExpressionFrom(10)),
	} {
		require.NotPanics(t, func() {
			_, _, err := GenerateExpressionQuery("sqlite", expr, nil)
			require.Error(t, err)
		})
	}
}
