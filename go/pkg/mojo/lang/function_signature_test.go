package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunctionSignatureParameterLifecycle(t *testing.T) {
	first := &ValueDecl{Name: "first"}
	second := &ValueDecl{Name: "second"}
	signature := &FunctionSignature{}
	require.NotPanics(t, func() { require.Same(t, signature, signature.SetParameters(first)) })
	require.Same(t, first, signature.ParameterDecl(0))
	require.Same(t, signature, signature.AppendParameter(second))
	require.Same(t, second, signature.ParameterDeclByName("second"))
	require.Nil(t, signature.ParameterDecl(-1))
	require.Nil(t, signature.ParameterDecl(2))
	require.Nil(t, signature.ParameterDeclByName("missing"))
	require.Nil(t, signature.ParameterDeclByName(""))
	signature.SetParameters(second)
	require.Equal(t, []*ValueDecl{second}, signature.GetParameters())
	signature.AppendParameter(nil)
	require.Len(t, signature.GetParameters(), 1)
	result := &NominalType{Name: "Result"}
	require.Same(t, result, NewFunctionSignature(result, first).GetResultType())
	var empty *FunctionSignature
	require.Nil(t, empty.SetParameters(first))
	require.Nil(t, empty.AppendParameter(first))
	require.Nil(t, empty.GetParameters())
	require.Nil(t, empty.GetResultType())
}

func TestFunctionSignatureTermsAndDocuments(t *testing.T) {
	param := &ValueDecl{Name: "input"}
	result := &NominalType{Name: "Result"}
	signature := NewFunctionSignature(result, param)
	terms := signature.GetTerms()
	require.Len(t, terms, 5)
	require.Equal(t, "(", terms[0].(*Term).Value)
	require.Same(t, param, terms[1])
	require.Equal(t, ")", terms[2].(*Term).Value)
	require.Equal(t, "->", terms[3].(*Term).Value)
	require.Same(t, result, terms[4])
	signature.Result.Implicit = true
	require.Len(t, signature.GetTerms(), 3)
	require.NotPanics(t, func() {
		require.False(t, signature.Parameter.HasFollowingDocument())
		require.False(t, signature.Result.HasFollowingDocument())
	})
	param.Document = &Document{Following: true}
	result.Document = &Document{Following: true}
	require.True(t, signature.Parameter.HasFollowingDocument())
	require.True(t, signature.Result.HasFollowingDocument())
}
