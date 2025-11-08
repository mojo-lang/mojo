package mojo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldOptionsExtensions(t *testing.T) {
	assert.Less(t, 9, len(FieldOptionsExtensions()))
}

func TestEnumValueOptionsExtensions(t *testing.T) {
	assert.NotEmpty(t, len(EnumValueOptionsExtensions()))
}
