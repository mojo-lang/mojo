package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/protobuf3/parser/syntax/testdata"
)

func TestParser_ParseString_AddressBook(t *testing.T) {
	file, err := New(nil).ParseString(context.Empty(), testdata.AddressBookProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)

	assert.Equal(t, 4, len(file.Attributes))
	assert.Equal(t, 3, len(file.Statements))

	assert.Equal(t, "tutorial", file.Statements[0].GetDeclaration().GetPackageDecl().GetName())

	decl := file.Statements[1].GetDeclaration().GetStructDecl()
	fields := decl.GetType().GetFields()
	assert.Equal(t, "Person", decl.Name)
	assert.Equal(t, 4, len(fields))
	assert.Equal(t, 1, len(decl.EnumDecls))
	assert.Equal(t, 1, len(decl.StructDecls))

	decl = file.Statements[2].GetDeclaration().GetStructDecl()
	fields = decl.GetType().GetFields()
	assert.Equal(t, "AddressBook", decl.Name)
	assert.Equal(t, 1, len(fields))
	assert.Equal(t, "people", fields[0].Name)
	assert.Equal(t, 1, len(fields[0].GetType().GetAttributes()))
}

func TestParser_ParseString_Conformance(t *testing.T) {
	file, err := New(nil).ParseString(context.Empty(), testdata.ConformanceProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseString_Example(t *testing.T) {
	file, err := New(nil).ParseString(context.Empty(), testdata.ExampleProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)

	assert.Equal(t, 2, len(file.Attributes))
	assert.Equal(t, 5, len(file.Statements))

	decl := file.Statements[3].GetDeclaration().GetStructDecl()
	fields := decl.GetType().GetFields()

	assert.Equal(t, 2, len(decl.Attributes))
	assert.Equal(t, 3, len(decl.Attributes[0].Arguments))
	argument := decl.Attributes[0].Arguments[0]
	assert.NotNil(t, 2, argument.GetIntegerLiteralExpr().EvalValue())

	assert.Equal(t, 2, len(decl.Attributes[1].Arguments))
	assert.Equal(t, 0, len(fields))
}

func TestParser_ParseString_HelloWorld(t *testing.T) {
	file, err := New(nil).ParseString(context.Empty(), testdata.HelloWorldProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseString_HelloWorldReserved(t *testing.T) {
	file, err := New(nil).ParseString(context.Empty(), testdata.HelloWorldReservedProto)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}

func TestParser_ParseString_HelloWorldReserved2(t *testing.T) {
	file, err := New(nil).ParseFile(context.Empty(), "./testdata/helloworldreserved.proto")
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
