package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArchitecture(t *testing.T) {
	ar, err := ParseArchitecture("AMD64")
	assert.NoError(t, err)
	assert.Equal(t, Architecture_ARCHITECTURE_AMD64, ar)
}
