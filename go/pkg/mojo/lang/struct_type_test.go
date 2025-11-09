package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeFields(t *testing.T) {
	src := []*ValueDecl{
		{Name: "a", Type: &NominalType{Name: "String"}},
		{Name: "b", Type: &NominalType{Name: "String"}}}

	fields := MergeFields(src, []*ValueDecl{
		{Name: "c", Type: &NominalType{Name: "Int32"}},
		{Name: "b", Type: &NominalType{Name: "String"}}})

	assert.Len(t, fields, 3)
}
