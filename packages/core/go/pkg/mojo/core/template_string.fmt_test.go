package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const urlTemplate = "/maps/tiles/v1/layers/{layer}/tiles/{id.level}/{id.x}/{id.y}{.format}"
const urlTemplateWithTemplate = "/maps/{{tiles}}/v1/layers/{layer}/tiles/{id.level}/{id.x}/{id.y}{.format}"

func TestTemplateString_Parse(t *testing.T) {
	ts := &TemplateString{}
	err := ts.Parse(urlTemplate)
	assert.NoError(t, err)
	assert.Equal(t, "/maps/tiles/v1/layers/", ts.Segments[0].Content)
	assert.Equal(t, "layer", ts.Segments[1].Content)
	assert.True(t, ts.Segments[1].Templated)
}

func TestTemplateString_Parse2(t *testing.T) {
	ts := &TemplateString{}
	err := ts.Parse(urlTemplateWithTemplate)
	assert.NoError(t, err)
	assert.Equal(t, "/maps/{{tiles}}/v1/layers/", ts.Segments[0].Content)
	assert.Equal(t, "layer", ts.Segments[1].Content)
	assert.True(t, ts.Segments[1].Templated)
}

func TestTemplateString_Parse3(t *testing.T) {
	ts := &TemplateString{}
	err := ts.Parse(`/maps/tiles/v1/layers/{layerId}`)
	assert.NoError(t, err)
	assert.Equal(t, "/maps/tiles/v1/layers/", ts.Segments[0].Content)
	assert.Equal(t, "layerId", ts.Segments[1].Content)
	assert.True(t, ts.Segments[1].Templated)
}

func TestTemplateString_Format(t *testing.T) {
	ts := &TemplateString{}
	err := ts.Parse(urlTemplate)
	assert.NoError(t, err)
	assert.Equal(t, urlTemplate, ts.Format())
}

func TestTemplateString_Format2(t *testing.T) {
	ts := &TemplateString{}
	err := ts.Parse(urlTemplateWithTemplate)
	assert.NoError(t, err)
	assert.Equal(t, urlTemplateWithTemplate, ts.Format())
}
