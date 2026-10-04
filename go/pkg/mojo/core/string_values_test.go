package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringValuesOperations(t *testing.T) {
	values := NewStringValues("b", "a", "b", "", "a")
	assert.Same(t, values, values.Unique())
	assert.Equal(t, []string{"b", "a", ""}, values.Vals)
	assert.True(t, values.Contains("a"))
	assert.False(t, values.Contains("A"))
	assert.Same(t, values, values.Append("c", "d"))
	assert.Equal(t, []string{"b", "a", "", "c", "d"}, values.ToArray())
	assert.Same(t, values, values.Append())
}

func TestStringValuesMatching(t *testing.T) {
	values := NewStringValues("alpha", "beta", "alpine", "alpha")
	assert.True(t, values.Matched("^al"))
	assert.False(t, values.Matched("^z"))
	assert.False(t, values.Matched("["))
	assert.Equal(t, []string{"alpha", "alpine", "alpha"}, values.Matches("^al").Vals)
	assert.Empty(t, values.Matches("[").Vals)
	assert.Empty(t, values.Matches("^z").Vals)
	assert.Equal(t, []string{"alpha", "beta", "alpine", "alpha"}, values.Vals)
}

func TestStringValuesEmptyAndNil(t *testing.T) {
	for _, values := range []*StringValues{nil, {}, NewStringValues()} {
		assert.Empty(t, values.ToArray())
		assert.False(t, values.Contains(""))
		assert.NotPanics(t, func() {
			assert.False(t, values.Matched(".*"))
			assert.Empty(t, values.Matches(".*").Vals)
		})
		assert.Same(t, values, values.Unique())
	}
	assert.Nil(t, (*StringValues)(nil).Append("x"))
}
