package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_VisitParameterTypeList1(t *testing.T) {
	p1 := `void aX(void);`

	file, err := New(nil).ParseString(p1)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.NotEmpty(t, file.Statements)
	decl := file.Statements[0].GetDeclaration().GetFunctionDecl()
	assert.Equal(t, "aX", decl.GetName())
	assert.Equal(t, "void", decl.GetSignature().GetResultType().GetName())
	assert.Equal(t, "void", decl.GetSignature().GetParameters()[0].GetType().GetName())
}

func TestParser_VisitParameterTypeList2(t *testing.T) {
	p2 := `int a1(int param1);`

	file, err := New(nil).ParseString(p2)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.NotEmpty(t, file.Statements)
	decl := file.Statements[0].GetDeclaration().GetFunctionDecl()
	assert.Equal(t, "a1", decl.GetName())
	assert.Equal(t, "int", decl.GetSignature().GetResultType().GetName())
	assert.Equal(t, "param1", decl.GetSignature().GetParameters()[0].GetName())
	assert.Equal(t, "int", decl.GetSignature().GetParameters()[0].GetType().GetName())
}

func TestParser_VisitParameterTypeList3(t *testing.T) {
	p3 := `int a2(int param1, param2, ...);`

	file, err := New(nil).ParseString(p3)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.NotEmpty(t, file.Statements)
	decl := file.Statements[0].GetDeclaration().GetFunctionDecl()
	assert.Equal(t, "a2", decl.GetName())
	assert.Equal(t, "int", decl.GetSignature().GetResultType().GetName())
	assert.Equal(t, 3, len(decl.GetSignature().GetParameters()))
	assert.Equal(t, "param1", decl.GetSignature().GetParameters()[0].GetName())
	assert.Equal(t, "int", decl.GetSignature().GetParameters()[0].GetType().GetName())
	// assert.Equal(t, "param2", decl.GetSignature().GetParameters()[1].GetName())
	assert.Equal(t, "...", decl.GetSignature().GetParameters()[2].GetName())
}

func TestParser_VisitParameterTypeList4(t *testing.T) {
	p4 := `int* a2(int param1, char ** param2, ...);`

	file, err := New(nil).ParseString(p4)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.NotEmpty(t, file.Statements)
	decl := file.Statements[0].GetDeclaration().GetFunctionDecl()
	assert.Equal(t, "a2", decl.GetName())
	assert.Equal(t, "int", decl.GetSignature().GetResultType().GetName())
	assert.Equal(t, 3, len(decl.GetSignature().GetParameters()))
	assert.Equal(t, "param1", decl.GetSignature().GetParameters()[0].GetName())
	assert.Equal(t, "int", decl.GetSignature().GetParameters()[0].GetType().GetName())
	assert.Equal(t, "param2", decl.GetSignature().GetParameters()[1].GetName())
	assert.Equal(t, "char", decl.GetSignature().GetParameters()[1].GetType().GetName())
	assert.Equal(t, "...", decl.GetSignature().GetParameters()[2].GetName())
}
