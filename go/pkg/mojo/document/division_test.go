package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTocDivision(t *testing.T) {
	division := NewTocDivision()
	assert.Equal(t, "toc", division.GetAttribute().GetIdentifier())
	require.Len(t, division.Content, 1)
	plain := division.Content[0].GetPlain()
	require.NotNil(t, plain)
	require.Len(t, plain.Inlines, 1)
	assert.Equal(t, "[TOC]", plain.Inlines[0].GetText().GetVal())
	division.Attribute.Identifier = "changed"
	assert.Equal(t, "toc", NewTocDivision().Attribute.Identifier)
}
