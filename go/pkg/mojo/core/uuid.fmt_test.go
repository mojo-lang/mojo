package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUuidParseFormat(t *testing.T) {
	const first = "00112233-4455-4677-8899-aabbccddeeff"
	u, err := ParseUuid(first)
	require.NoError(t, err)
	require.Len(t, u.Val, 16)
	assert.Equal(t, first, u.Format())
	const second = "ffeeddcc-bbaa-4988-b766-554433221100"
	require.NoError(t, u.Parse(second))
	assert.Len(t, u.Val, 16, "parsing replaces the previous UUID")
	assert.Equal(t, second, u.Format())
	assert.Error(t, u.Parse("invalid"))
	assert.Equal(t, second, u.Format())
	_, err = ParseUuid("")
	assert.Error(t, err)
}
