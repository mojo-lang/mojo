package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathItem_GetOperations(t *testing.T) {
	pathItem := &PathItem{Get: &Operation{Summary: "foo"}, Post: &Operation{Summary: "bar"}, Delete: &Operation{Summary: "baz"}}

	operations := pathItem.GetOperations()
	assert.Equal(t, 3, len(operations))
	assert.Equal(t, "GET", operations[0].Method)
	assert.Equal(t, "foo", operations[0].Operation.Summary)
	assert.Equal(t, "POST", operations[1].Method)
	assert.Equal(t, "bar", operations[1].Operation.Summary)
	assert.Equal(t, "DELETE", operations[2].Method)
	assert.Equal(t, "baz", operations[2].Operation.Summary)
}
