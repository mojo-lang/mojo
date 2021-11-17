package syntax

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNominalTypeVisitor_VisitArrayType(t *testing.T) {
	const arrayType = `type Val{ val: [Int] }`

	parser := &Parser{}
	file, err := parser.ParseString(arrayType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Array", nominalType.Name)
	assert.Equal(t, 1, len(nominalType.GenericArguments))
	assert.Equal(t, "Int", nominalType.GenericArguments[0].Name)
}

func TestNominalTypeVisitor_VisitArrayType2(t *testing.T) {
	const arrayType = `type Val{ val: [mojo.lang.StructType] }`

	parser := &Parser{}
	file, err := parser.ParseString(arrayType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Array", nominalType.Name)
	assert.Equal(t, 1, len(nominalType.GenericArguments))
	assert.Equal(t, "mojo.lang", nominalType.GenericArguments[0].PackageName)
	assert.Equal(t, "StructType", nominalType.GenericArguments[0].Name)
}

func TestNominalTypeVisitor_VisitPrimeType_PackageName(t *testing.T) {
	const primeType = `type Val{ val: mojo.core.Int }`

	parser := &Parser{}
	file, err := parser.ParseString(primeType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Int", nominalType.Name)
	assert.Equal(t, "mojo.core", nominalType.PackageName)
}

func TestNominalTypeVisitor_VisitType_EnclosingName(t *testing.T) {
	const primeType = `type Val{ val: mojo.core.Url.Path }`

	parser := &Parser{}
	file, err := parser.ParseString(primeType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Path", nominalType.Name)
	assert.Equal(t, "Url", nominalType.EnclosingType.Name)
	assert.Equal(t, "mojo.core", nominalType.PackageName)
}

func TestNominalTypeVisitor_VisitMapType(t *testing.T) {
	const dictType = `type Val{ val: {String: Int} }`

	parser := &Parser{}
	file, err := parser.ParseString(dictType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Map", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
}

func TestNominalTypeVisitor_VisitUnion(t *testing.T) {
	const unionType = `type Val{ val: String | Int | [String] }`

	parser := &Parser{}
	file, err := parser.ParseString(unionType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Union", nominalType.Name)
	assert.Equal(t, 3, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
	assert.Equal(t, "Array", nominalType.GenericArguments[2].Name)
}

func TestNominalTypeVisitor_VisitTupleType(t *testing.T) {
	const tupleType = `type Val{ val: (String, Int) }`

	parser := &Parser{}
	file, err := parser.ParseString(tupleType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Tuple", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
}

func TestNominalTypeVisitor_VisitTupleTypeWithLabel(t *testing.T) {
	const tupleType = `type Val{ val: (str: String, integer: Int) }`

	parser := &Parser{}
	file, err := parser.ParseString(tupleType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Tuple", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)

	label, _ := lang.GetStringAttribute(nominalType.GenericArguments[0].Attributes, "label")
	assert.Equal(t, "str", label)

	label, _ = lang.GetStringAttribute(nominalType.GenericArguments[1].Attributes, "label")
	assert.Equal(t, "integer", label)
}

func TestNominalTypeVisitor_VisitQuestionType(t *testing.T) {
	const questionType = `type Val{ val: Int? }`

	parser := &Parser{}
	file, err := parser.ParseString(questionType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Int", nominalType.Name)
	v, _ := lang.GetBoolAttribute(nominalType.Attributes, "optional")
	assert.True(t, v)
}

func TestNominalTypeVisitor_VisitBangType(t *testing.T) {
	const bangType = `type Val{ val: Int! }`

	parser := &Parser{}
	file, err := parser.ParseString(bangType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Int", nominalType.Name)
	v, _ := lang.GetBoolAttribute(nominalType.Attributes, "required")
	assert.True(t, v)
}

func getNominalType(file *lang.SourceFile) *lang.NominalType {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			if structDecl := decl.GetStructDecl(); structDecl != nil {
				if t := structDecl.Type; t != nil {
					if len(t.Fields) > 0 {
						if field := t.Fields[0]; field != nil {
							return field.Type
						}
					}
				}
			}
		}
	}
	return nil
}
