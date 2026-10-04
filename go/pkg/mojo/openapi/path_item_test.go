package openapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestPathItemSetAndMerge(t *testing.T) {
	methods := []string{"GET", "PUT", "POST", "DELETE", "OPTIONS", "HEAD", "PATCH", "TRACE"}
	item := &PathItem{Summary: "preserved", Ref: "old", Servers: []*Server{{Description: &CachedDocument{Cache: "first"}}}, Parameters: []*ReferenceableParameter{NewReferenceableParameter(&Parameter{Name: "first"})}}
	other := &PathItem{Summary: "new", Ref: "new", Description: &CachedDocument{Cache: "description"}, Servers: []*Server{{Description: &CachedDocument{Cache: "second"}}}, Parameters: []*ReferenceableParameter{NewReferenceableParameter(&Parameter{Name: "second"})}}
	for _, method := range methods {
		other.SetOperation(method, &Operation{OperationId: method})
	}
	other.SetOperation("unknown", &Operation{})
	operations := other.GetOperations()
	require.Len(t, operations, len(methods))
	for i, method := range methods {
		assert.Equal(t, method, operations[i].Method)
		assert.Equal(t, method, operations[i].Operation.OperationId)
	}
	assert.Same(t, item, item.Merge(other))
	assert.Equal(t, "new", item.Summary)
	assert.Equal(t, "new", item.Ref)
	assert.Same(t, other.Description, item.Description)
	assert.Len(t, item.GetOperations(), len(methods))
	require.Len(t, item.Servers, 2)
	assert.Equal(t, "first", item.Servers[0].Description.GetStringDocument())
	assert.Equal(t, "second", item.Servers[1].Description.GetStringDocument())
	assert.Len(t, item.Parameters, 2)
	item.Merge(&PathItem{})
	assert.Equal(t, "new", item.Summary)
	assert.NotPanics(t, func() { assert.Same(t, item, item.Merge(nil)) })
	assert.Same(t, item, (*PathItem)(nil).Merge(item))
	assert.Nil(t, (*PathItem)(nil).Merge(nil))
	assert.Empty(t, (*PathItem)(nil).GetOperations())
	assert.NotPanics(t, func() { (*PathItem)(nil).SetOperation("GET", &Operation{}) })
	item.SetOperation("get", nil)
	assert.Nil(t, item.Get)
}

func TestOrderedPathItemIterator(t *testing.T) {
	items := map[string]*PathItem{"/z": {}, "/a": {}, "/m": nil}
	var visited []string
	require.NoError(t, OrderedPathItemIterator(items, func(key string, value *PathItem) error {
		visited = append(visited, key)
		assert.Equal(t, items[key], value)
		return nil
	}))
	assert.Equal(t, []string{"/a", "/m", "/z"}, visited)
	want := errors.New("stop")
	visited = nil
	err := OrderedPathItemIterator(items, func(key string, value *PathItem) error { visited = append(visited, key); return want })
	assert.ErrorIs(t, err, want)
	assert.Equal(t, []string{"/a"}, visited)
	assert.NoError(t, OrderedPathItemIterator(nil, nil))
}
