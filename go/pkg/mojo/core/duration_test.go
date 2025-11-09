package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuration_FromSeconds(t *testing.T) {
	duration := &Duration{}
	duration.FromSeconds(1.001)
	assert.Equal(t, int64(1), duration.Seconds)
	assert.Equal(t, int32(1000000), duration.Nanoseconds)
}

func TestDuration_ToSeconds(t *testing.T) {
	duration := Duration{Nanoseconds: 1000, Seconds: 10}
	assert.Equal(t, 10.000001, duration.ToSeconds())
}
