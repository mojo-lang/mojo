package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttribute_SetBoolProperty(t *testing.T) {
	attribute := NewAttribute()
	attribute.SetBoolProperty("foo", true)

	assert.True(t, attribute.HasProperty("foo"))
	v, err := attribute.GetBoolProperty("foo")

	assert.NoError(t, err)
	assert.True(t, v)
}

func TestAttribute_SetStringProperty(t *testing.T) {
	attribute := NewAttribute()
	attribute.SetStringProperty("foo", "bar")

	assert.True(t, attribute.HasProperty("foo"))
	v, err := attribute.GetStringProperty("foo")

	assert.NoError(t, err)
	assert.Equal(t, "bar", v)
}
