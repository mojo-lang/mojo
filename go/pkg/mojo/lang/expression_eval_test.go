package lang

import (
	"errors"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/require"
)

func TestExpressionArraysRejectMixedElementTypes(t *testing.T) {
	floats := NewArgument(NewArrayLiteralExpressionFrom(NewFloatLiteralExpressionFrom(1.25), NewFloatLiteralExpressionFrom(-2.5)))
	values, err := floats.GetFloatArray()
	require.NoError(t, err)
	require.Equal(t, []float64{1.25, -2.5}, values)
	elements, err := floats.GetArray()
	require.NoError(t, err)
	require.Len(t, elements, 2)
	mixed := NewArrayLiteralExpressionFrom(NewIntegerLiteralExpressionFrom(1), NewStringLiteralExpressionFrom("wrong"))
	_, err = mixed.EvalIntegerArrayLiteral()
	require.Error(t, err)
	_, err = mixed.EvalFloatArrayLiteral()
	require.Error(t, err)
	_, err = mixed.EvalStringArrayLiteral()
	require.Error(t, err)
	for _, expression := range []*Expression{nil, NewNullLiteralExpressionFrom()} {
		_, err = expression.EvalBoolLiteral()
		require.Error(t, err)
		_, err = expression.EvalIntegerLiteral()
		require.Error(t, err)
		_, err = expression.EvalFloatLiteral()
		require.Error(t, err)
		_, err = expression.EvalStringLiteral()
		require.Error(t, err)
		_, err = expression.EvalRangeLiteral()
		require.Error(t, err)
		_, err = expression.EvalArrayLiteral()
		require.Error(t, err)
		_, err = expression.EvalIntegerArrayLiteral()
		require.Error(t, err)
		_, err = expression.EvalFloatArrayLiteral()
		require.Error(t, err)
		_, err = expression.EvalStringArrayLiteral()
		require.Error(t, err)
		_, err = expression.EvalIdentifier()
		require.Error(t, err)
		require.Error(t, expression.EvalStringMapLiteral(func(string, *Expression) error { t.Fatal("invalid map must not invoke callback"); return nil }))
	}
	var absent *ArrayLiteralExpr
	_, err = absent.EvalIntegerArray()
	require.Error(t, err)
	_, err = absent.EvalFloatArray()
	require.Error(t, err)
	_, err = absent.EvalStringArray()
	require.Error(t, err)
	rangeValue := &core.IntRange{}
	got, err := NewRangeLiteralExpression(&RangeLiteralExpr{Value: rangeValue}).EvalRangeLiteral()
	require.NoError(t, err)
	require.Same(t, rangeValue, got)
	name, err := NewIdentifierExpressionFrom("sample", "Item").EvalIdentifier()
	require.NoError(t, err)
	require.Equal(t, "Item", name)
}

func TestExpressionMapIterationPreservesOrderAndPropagatesErrors(t *testing.T) {
	argument := NewMapLiteralArgument(&MapLiteralExpr{Entries: []*MapLiteralExpr_Entry{
		{Key: "first", Value: NewIntegerLiteralExpressionFrom(1)}, {Key: "second", Value: NewIntegerLiteralExpressionFrom(2)},
	}})
	var keys []string
	require.NoError(t, argument.IterateStringMap(func(key string, value *Expression) error {
		keys = append(keys, key)
		n, err := value.EvalIntegerLiteral()
		require.NoError(t, err)
		require.Equal(t, int64(len(keys)), n)
		return nil
	}))
	require.Equal(t, []string{"first", "second"}, keys)
	for _, stop := range []error{core.NewBreakError(), errors.New("callback failed")} {
		calls := 0
		err := argument.IterateStringMap(func(string, *Expression) error { calls++; return stop })
		require.Equal(t, 1, calls)
		if core.IsBreakError(stop) {
			require.NoError(t, err)
		} else {
			require.ErrorIs(t, err, stop)
		}
	}
}
