package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDeleteTime(t *testing.T) {
	dt, err := ParseDeleteTime("2023-05-05T01:01:01.123")
	assert.NoError(t, err)
	assert.Equal(t, 2023, dt.ToTime().Year())
}
