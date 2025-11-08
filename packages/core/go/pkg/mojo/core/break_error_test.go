package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBreakError(t *testing.T) {
	err := &BreakError{}
	assert.True(t, IsBreakError(err))
}
