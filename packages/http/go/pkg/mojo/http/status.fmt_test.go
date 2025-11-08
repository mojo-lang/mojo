package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus_Parse1(t *testing.T) {
	s, err := ParseStatus("404")
	assert.NoError(t, err)
	assert.Equal(t, int32(404), s.Code)
}

func TestStatus_Parse2(t *testing.T) {
	s, err := ParseStatus("404 NOT_FOUND")
	assert.NoError(t, err)
	assert.Equal(t, int32(404), s.Code)
	assert.Equal(t, "NOT_FOUND", s.Reason)
}

func TestStatus_Parse3(t *testing.T) {
	s, err := ParseStatus("404    ")
	assert.NoError(t, err)
	assert.Equal(t, int32(404), s.Code)
}

func TestStatus_Parse4(t *testing.T) {
	s, err := ParseStatus("404    NOT_FOUND")
	assert.NoError(t, err)
	assert.Equal(t, int32(404), s.Code)
	assert.Equal(t, "NOT_FOUND", s.Reason)
}
