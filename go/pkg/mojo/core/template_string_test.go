package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTemplateString(t *testing.T) {
	ts := NewTemplateString(urlTemplate)
	assert.Equal(t, "/maps/tiles/v1/layers/", ts.Segments[0].Content)
	assert.Equal(t, "layer", ts.Segments[1].Content)
	assert.True(t, ts.Segments[1].Templated)
}

func TestTemplateString_Apply(t *testing.T) {
	ts := NewTemplateString("/path/to/{service}/{format}")
	ts.Apply(map[string]interface{}{
		"service": "foo",
		"format":  "json",
	})

	assert.Equal(t, "/path/to/foo/json", ts.Format())
}

func TestTemplateString_ApplyStrings(t *testing.T) {
	ts := NewTemplateString("/path/to/{service}/{format}")
	ts.ApplyStrings(map[string]string{
		"service": "foo",
		"format":  "json",
	})

	assert.Equal(t, "/path/to/foo/json", ts.Format())
}
