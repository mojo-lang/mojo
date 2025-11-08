package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMediaType(t *testing.T) {
	str := `multipart/form-data;boundary="boundary"`

	mt, err := ParseMediaType(str)
	assert.NoError(t, err)

	assert.Equal(t, "multipart", mt.Type)
	assert.Equal(t, "form-data", mt.Subtype)
	assert.Equal(t, "boundary", mt.Parameter.Key)
	assert.Equal(t, `boundary`, mt.Parameter.Value.GetString())
}
