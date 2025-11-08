package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const colorStr = "#fff0f5"

var color = &Color{
	Red:   255,
	Green: 240,
	Blue:  245,
	Alpha: nil,
}

const colorAlphaStr = "#fff0f51a"

var colorAlpha = Rgba(255, 240, 245, 0.1)

func TestColor_Format(t *testing.T) {
	assert.Equal(t, colorStr, color.Format())
	assert.Equal(t, colorAlphaStr, colorAlpha.Format())
}

func TestColor_Parse(t *testing.T) {
	c := NewColor(0, 0, 0)
	err := c.Parse(colorStr)

	assert.NoError(t, err)
	assert.Equal(t, color.Red, c.Red)
	assert.Equal(t, color.Green, c.Green)
	assert.Equal(t, color.Blue, c.Blue)
}

func TestColor_ParseAlpha(t *testing.T) {
	c := NewColor(0, 0, 0)
	err := c.Parse(colorAlphaStr)

	assert.NoError(t, err)
	assert.Equal(t, colorAlpha.Red, c.Red)
	assert.Equal(t, colorAlpha.Green, c.Green)
	assert.Equal(t, colorAlpha.Blue, c.Blue)
	assert.Equal(t, colorAlpha.GetAlphaValue(), c.GetAlphaValue())
}
