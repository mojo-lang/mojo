package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGenericTypeName(t *testing.T) {
	assert.True(t, IsGenericTypeName("a.b.C<T>"))
	assert.True(t, IsGenericTypeName("a.b.C<T<K,V<U>>>"))
	assert.False(t, IsGenericTypeName("a.b.C"))
}
