package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex_ToRegexp(t *testing.T) {
	m := NewRegex(".+").ToRegexp().MatchString("test")
	assert.True(t, m)
}

func TestRegex_ToRegexp2(t *testing.T) {
	m := NewRegex("^[\u4E00-\u9FA5A-Za-z0-9]{1,50}$").ToRegexp().MatchString("test")
	assert.True(t, m)
}
