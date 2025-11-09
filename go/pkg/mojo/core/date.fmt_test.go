package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDate_Format(t *testing.T) {
	date := &Date{
		Year:  2024,
		Month: 1,
		Day:   9,
	}
	assert.Equal(t, "2024-01-09", date.Format())
}

func TestDate_Parse(t *testing.T) {
	date := &Date{}
	err := date.Parse("2024-01-09")
	assert.NoError(t, err)
	assert.Equal(t, int32(2024), date.Year)
	assert.Equal(t, int32(1), date.Month)
	assert.Equal(t, int32(9), date.Day)
}
