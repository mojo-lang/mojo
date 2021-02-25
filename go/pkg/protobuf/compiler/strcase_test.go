package compiler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSnake(t *testing.T) {
	assert.Equal(t, "uint8", ToSnake("UInt8"))
	assert.Equal(t, "uint", ToSnake("UInt"))
}
