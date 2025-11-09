package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[string]*TypeName{
	"mojo.core.Checksum": {
		PackageName: "mojo.core",
		Name:        "Checksum",
	},
	"Array<mojo.core.Checksum>": {
		Name:              ArrayTypeName,
		GenericParameters: []*TypeName{{PackageName: "mojo.core", Name: "Checksum"}},
	},
	"Array<mojo.core.Array<Int32>>": {
		Name:              ArrayTypeName,
		GenericParameters: []*TypeName{{PackageName: "mojo.core", Name: ArrayTypeName, GenericParameters: []*TypeName{{Name: "Int32"}}}},
	},
	"Map<mojo.core.Array<Int32>, mojo.core.Checksum>": {
		Name: MapTypeName,
		GenericParameters: []*TypeName{
			{PackageName: "mojo.core", Name: ArrayTypeName, GenericParameters: []*TypeName{{Name: "Int32"}}},
			{PackageName: "mojo.core", Name: "Checksum"},
		},
	},
}

var errCases = []string{
	"Array<mojo.core.Checksum",
}

func TestTypeName_Parse(t *testing.T) {
	for k, v := range cases {
		tn, err := ParseTypeName(k)
		assert.NoError(t, err)
		assert.True(t, v.IsSame(tn))
	}
}

func TestTypeName_Parse_Error(t *testing.T) {
	for _, v := range errCases {
		_, err := ParseTypeName(v)
		assert.Error(t, err)
	}
}
