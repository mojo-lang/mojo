package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fullNameCase struct {
	FullName       string
	PackageName    string
	Name           string
	EnclosingNames []string
	Arguments      []*NominalType
}

func testCase(t *testing.T, c fullNameCase) {
	nominalType, err := ParseNominalTypeFullName(c.FullName)
	assert.NoError(t, err)
	assert.NotNil(t, nominalType)
	assert.Equal(t, c.PackageName, nominalType.PackageName)
	assert.Equal(t, c.EnclosingNames, nominalType.GetEnclosingNames())
	assert.Equal(t, c.Name, nominalType.Name)

	assert.Equal(t, len(c.Arguments), len(nominalType.GenericArguments))
	for i, argument := range nominalType.GenericArguments {
		testNominalType(t, c.Arguments[i], argument)
	}
}

func testNominalType(t *testing.T, expect *NominalType, actual *NominalType) {
	assert.NotNil(t, expect)
	assert.NotNil(t, actual)
	assert.Equal(t, expect.Name, actual.Name)
	assert.Equal(t, expect.PackageName, actual.PackageName)
	assert.Equal(t, len(expect.GenericArguments), len(actual.GenericArguments))

	for i, argument := range expect.GenericArguments {
		testNominalType(t, argument, actual.GenericArguments[i])
	}

	if expect.Enclosing != nil {
		testNominalType(t, expect.Enclosing, actual.Enclosing)
	}
}

func TestParseNominalTypeFullName(t *testing.T) {
	cases := []fullNameCase{
		{
			FullName:       "a.b.c.Test.Inner.Inner2",
			PackageName:    "a.b.c",
			Name:           "Inner2",
			EnclosingNames: []string{"Test", "Inner"},
		},
		{
			FullName:       "a.Test.Inner",
			PackageName:    "a",
			Name:           "Inner",
			EnclosingNames: []string{"Test"},
		},
		{
			FullName:       "a.b.c.Test",
			PackageName:    "a.b.c",
			Name:           "Test",
			EnclosingNames: nil,
		},
		{
			FullName:       "Test",
			PackageName:    "",
			Name:           "Test",
			EnclosingNames: nil,
		},
	}

	for _, c := range cases {
		testCase(t, c)
	}
}

func TestParseGenericNominalTypeFullName(t *testing.T) {
	cases := []fullNameCase{
		{
			FullName:       "a.b.c.Test<A1<B1<C1,C2>,B2>,A2<a.B2>>",
			PackageName:    "a.b.c",
			Name:           "Test",
			EnclosingNames: nil,
			Arguments: []*NominalType{
				{
					Name: "A1",
					GenericArguments: []*NominalType{
						{
							Name: "B1",
							GenericArguments: []*NominalType{
								{Name: "C1"},
								{Name: "C2"},
							}},
						{
							Name: "B2",
						},
					}},
				{
					Name: "A2",
					GenericArguments: []*NominalType{
						{
							Name:        "B2",
							PackageName: "a"},
					}},
			},
		},
		{
			FullName:       "a.b.c.Test< A1 <B1  <C1,  C2  >,  B2>,  A2<  a.B2>  >",
			PackageName:    "a.b.c",
			Name:           "Test",
			EnclosingNames: nil,
			Arguments: []*NominalType{
				{
					Name: "A1",
					GenericArguments: []*NominalType{
						{
							Name: "B1",
							GenericArguments: []*NominalType{
								{Name: "C1"},
								{Name: "C2"},
							}},
						{
							Name: "B2",
						},
					}},
				{
					Name: "A2",
					GenericArguments: []*NominalType{
						{
							Name:        "B2",
							PackageName: "a"},
					}},
			},
		},
		{
			FullName:       "a.b.c.Test.Inner.Inner2<A1, A2>",
			PackageName:    "a.b.c",
			Name:           "Inner2",
			EnclosingNames: []string{"Test", "Inner"},
			Arguments: []*NominalType{
				{Name: "A1"},
				{Name: "A2"},
			},
		},
		{
			FullName:       "a.Test.Inner",
			PackageName:    "a",
			Name:           "Inner",
			EnclosingNames: []string{"Test"},
		},
		{
			FullName:       "Test",
			PackageName:    "",
			Name:           "Test",
			EnclosingNames: nil,
		},
	}

	for _, c := range cases {
		testCase(t, c)
	}
}
