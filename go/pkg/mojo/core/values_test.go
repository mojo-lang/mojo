package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues_ToSlice(t *testing.T) {
	values, err := NewValues(1, 2, 3, 4, 5)
	assert.NoError(t, err)
	assert.Equal(t, 5, values.Len())
	assert.Equal(t, uint64(1), values.ToSlice()[0].(uint64))
}
