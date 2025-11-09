package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimestamp_Format(t *testing.T) {
	ts := &Timestamp{Seconds: 1656489818}
	str := ts.Format()

	assert.NotEmpty(t, str)
	assert.Equal(t, "2022-06-29T16:03:38.000+08:00", str)
}
