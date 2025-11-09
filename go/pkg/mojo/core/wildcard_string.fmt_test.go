package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWildcardString_Parse(t *testing.T) {
	const str = "*foo."
	ws, err := ParseWildcardString(str)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(ws.Wildcards))
	assert.Equal(t, WildcardString_Wildcard_TYPE_ZERO_OR_MORE_CHARACTERS, ws.Wildcards[0].Type)
	assert.Equal(t, "foo", ws.Wildcards[1].Content)
	assert.Equal(t, WildcardString_Wildcard_TYPE_SINGLE_CHARACTER, ws.Wildcards[2].Type)
}

func TestWildcardString_Parse1(t *testing.T) {
	const str = `*foo\.`
	ws, err := ParseWildcardString(str)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(ws.Wildcards))
	assert.Equal(t, WildcardString_Wildcard_TYPE_ZERO_OR_MORE_CHARACTERS, ws.Wildcards[0].Type)
	assert.Equal(t, "foo.", ws.Wildcards[1].Content)
}

func TestWildcardString_Parse2(t *testing.T) {
	const str = "*foo[abcd]"
	ws, err := ParseWildcardString(str)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(ws.Wildcards))
	assert.Equal(t, WildcardString_Wildcard_TYPE_ZERO_OR_MORE_CHARACTERS, ws.Wildcards[0].Type)
	assert.Equal(t, "foo", ws.Wildcards[1].Content)
	assert.Equal(t, WildcardString_Wildcard_TYPE_IN_CHARACTERS, ws.Wildcards[2].Type)
	assert.Equal(t, "abcd", ws.Wildcards[2].Content)
}

func TestWildcardString_Parse3(t *testing.T) {
	const str = "*foo[^abcd]"
	ws, err := ParseWildcardString(str)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(ws.Wildcards))
	assert.Equal(t, WildcardString_Wildcard_TYPE_ZERO_OR_MORE_CHARACTERS, ws.Wildcards[0].Type)
	assert.Equal(t, "foo", ws.Wildcards[1].Content)
	assert.Equal(t, WildcardString_Wildcard_TYPE_NOT_IN_CHARACTERS, ws.Wildcards[2].Type)
	assert.Equal(t, "abcd", ws.Wildcards[2].Content)
}

func TestWildcardString_Format(t *testing.T) {
	ws := &WildcardString{Wildcards: []*WildcardString_Wildcard{{
		Type: WildcardString_Wildcard_TYPE_ZERO_OR_MORE_CHARACTERS,
	}, {
		Content: "foo",
	}, {
		Type: WildcardString_Wildcard_TYPE_SINGLE_CHARACTER,
	}}}
	assert.Equal(t, "*foo.", ws.Format())
}
