package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArgumentLiteralValuesAndTypeErrors(t *testing.T) {
	boolean := NewBoolLiteralArgument(&BoolLiteralExpr{Value: true})
	b, err := boolean.GetBool()
	require.NoError(t, err)
	require.True(t, b)
	integer := NewArgument(NewIntegerLiteralExpressionFrom(-42))
	n, err := integer.GetInteger()
	require.NoError(t, err)
	require.Equal(t, int64(-42), n)
	floating := NewFloatLiteralArgument(&FloatLiteralExpr{Value: 1.5})
	f, err := floating.GetFloat()
	require.NoError(t, err)
	require.Equal(t, 1.5, f)
	text := NewStringArgument("hello")
	s, err := text.GetString()
	require.NoError(t, err)
	require.Equal(t, "hello", s)
	_, err = text.GetInteger()
	require.Error(t, err)
	var absent *Argument
	_, err = absent.GetString()
	require.Error(t, err)
	require.Nil(t, absent.GetBoolLiteralExpr())
	require.Nil(t, absent.GetArrayLiteralExpr())
}

func TestArgumentArraysAndTerms(t *testing.T) {
	arg := NewArrayLiteralArgument(&ArrayLiteralExpr{Elements: []*Expression{NewStringLiteralExpressionFrom("one"), NewStringLiteralExpressionFrom("two")}})
	values, err := arg.GetStringArray()
	require.NoError(t, err)
	require.Equal(t, []string{"one", "two"}, values)
	integer := NewArgument(NewArrayLiteralExpressionFrom(NewIntegerLiteralExpressionFrom(1), NewIntegerLiteralExpressionFrom(2)))
	numbers, err := integer.GetIntegerArray()
	require.NoError(t, err)
	require.Equal(t, []int64{1, 2}, numbers)
	arg.SetLabel("names").SetStartPosition(&Position{Line: 2, Column: 3})
	terms := arg.GetTerms()
	require.Len(t, terms, 2)
	label := terms[0].(*Term)
	require.Equal(t, "names", label.Value)
	require.Equal(t, int64(8), label.EndPosition.Column)
	require.Same(t, arg.Value, terms[1])
}

func TestMergeIdentifierCollections(t *testing.T) {
	first := &Identifier{FullName: "sample.First"}
	old := &Identifier{FullName: "sample.Shared"}
	replacement := &Identifier{FullName: "sample.Shared", SourceFileName: "new.mojo"}
	for _, merge := range []func(...[]*Identifier) []*Identifier{MergeDependencies, MergeUnresolvedIdentifiers} {
		require.Empty(t, merge())
		result := merge([]*Identifier{first, old}, nil, []*Identifier{replacement})
		require.ElementsMatch(t, []*Identifier{first, replacement}, result)
	}
}
