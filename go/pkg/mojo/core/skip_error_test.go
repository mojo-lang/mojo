package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSkipError(t *testing.T) {
	err := &SkipError{}
	assert.True(t, IsSkipError(err))
}
