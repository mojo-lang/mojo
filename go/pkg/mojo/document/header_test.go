package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHeader(t *testing.T) {
	texts := []*Inline{NewTextInline("first"), NewStrongInline(NewTextInline("second"))}
	header := NewHeader(2, texts...)
	assert.Equal(t, int64(2), header.Level)
	assert.Equal(t, texts, header.Text)
	require.NotNil(t, header.Attribute)
	assert.NotNil(t, header.Attribute.Properties)
	empty := NewHeader(0)
	assert.Zero(t, empty.Level)
	assert.Empty(t, empty.Text)
	header.Attribute.SetStringProperty("key", "value")
	assert.False(t, empty.Attribute.HasProperty("key"))
}
