package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_ParseFile1(t *testing.T) {
	file, err := New(nil).ParseFile("./testdata/func_call_with_var_args.h", nil)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseFile2(t *testing.T) {
	file, err := New(nil).ParseFile("./testdata/bt.c", nil)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseFile_CPP(t *testing.T) {
	file, err := New(nil).ParseFile("./testdata/zipper.hpp", nil)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseFile_Static(t *testing.T) {
	file, err := New(nil).ParseFile("./testdata/func_static_extern.h", nil)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, 2, len(file.GetStatements()))

	decl := file.Statements[0].GetDeclaration().GetFunctionDecl()
	assert.True(t, decl.HasAttribute("c_static"))

	decl = file.Statements[1].GetDeclaration().GetFunctionDecl()
	assert.True(t, decl.HasAttribute("c_extern"))
}
