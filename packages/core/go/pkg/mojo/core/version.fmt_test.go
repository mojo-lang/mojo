package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion_Parse(t *testing.T) {
	const V = "0.0.0-20211119132822-836b7b2cb4ed"
	version, err := ParseVersion(V)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), version.Major)
	assert.Equal(t, 1, len(version.PreReleases))
	assert.Equal(t, "20211119132822-836b7b2cb4ed", version.PreReleases[0])
}

func TestVersion_Parse_1(t *testing.T) {
	const V = "v0"
	version, err := ParseVersion(V)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), version.Major)
	assert.Equal(t, int32(1), version.Level)
}

func TestVersion_Parse_2(t *testing.T) {
	const V = "v0.0"
	version, err := ParseVersion(V)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), version.Major)
	assert.Equal(t, int32(2), version.Level)
}

func TestVersion_Parse_3(t *testing.T) {
	const V = "3.0.3"
	version, err := ParseVersion(V)
	assert.NoError(t, err)
	assert.Equal(t, uint64(3), version.Major)
	assert.Equal(t, int32(3), version.Level)
}
