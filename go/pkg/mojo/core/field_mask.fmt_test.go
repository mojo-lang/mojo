package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldMaskParser(t *testing.T) {
	fm := NewFieldMask(" field1, field2 { filed2-1 , field2-2 }, field3 ")

	assert.Equal(t, 4, len(fm.Paths))
	assert.Equal(t, "field1", fm.Paths[0])
	assert.Equal(t, "field2.filed2-1", fm.Paths[1])
	assert.Equal(t, "field2.field2-2", fm.Paths[2])
	assert.Equal(t, "field3", fm.Paths[3])
}
