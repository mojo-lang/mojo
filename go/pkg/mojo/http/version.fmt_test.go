package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion_Format(t *testing.T) {
	assert.Equal(t, "HTTP/1.1", (&Version{Major: 1, Minor: 1}).Format())
	assert.Equal(t, "HTTP/1.0", (&Version{Major: 1}).Format())
	assert.Equal(t, "HTTP/2", (&Version{Major: 2}).Format())
}

func TestVersion_Parse(t *testing.T) {
	v, err := ParseVersion("HTTP/1.1")
	assert.NoError(t, err)
	assert.Equal(t, int32(1), v.Major)
	assert.Equal(t, int32(1), v.Minor)

	v, err = ParseVersion("1.1")
	assert.NoError(t, err)
	assert.Equal(t, int32(1), v.Major)
	assert.Equal(t, int32(1), v.Minor)

	v, err = ParseVersion("HTTP/1.0")
	assert.NoError(t, err)
	assert.Equal(t, int32(1), v.Major)
	assert.Equal(t, int32(0), v.Minor)

	v, err = ParseVersion("1.0")
	assert.NoError(t, err)
	assert.Equal(t, int32(1), v.Major)
	assert.Equal(t, int32(0), v.Minor)

	v, err = ParseVersion("HTTP/2")
	assert.NoError(t, err)
	assert.Equal(t, int32(2), v.Major)
	assert.Equal(t, int32(0), v.Minor)

	v, err = ParseVersion("2")
	assert.NoError(t, err)
	assert.Equal(t, int32(2), v.Major)
	assert.Equal(t, int32(0), v.Minor)
}
