package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTemplateStringPatternRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		input    string
		segments []string
	}{
		{"/articles/{category}/{id:[0-9]+}", []string{"category", "id:[0-9]+"}},
		{"/files/{name:.+}", []string{"name:.+"}},
		{`/{id.level:\d{2,4}}/{name:[a-z]{3}}`, []string{`id.level:\d{2,4}`, "name:[a-z]{3}"}},
		{"/files/{name:(?:images|docs)/.+}", []string{"name:(?:images|docs)/.+"}},
		{"/{{base}}/{id:[0-9]{2}}{.format}", []string{"id:[0-9]{2}", ".format"}},
		{"/files/{name:.+", nil},
	} {
		t.Run(tc.input, func(t *testing.T) {
			template := NewTemplateString(tc.input)
			require.NotNil(t, template)
			require.Equal(t, tc.input, template.Format())
			var segments []string
			for _, s := range template.Segments {
				if s.Templated {
					segments = append(segments, s.Content)
				}
			}
			require.Equal(t, tc.segments, segments)
		})
	}
}

func TestTemplateStringApplyPatternByName(t *testing.T) {
	template := NewTemplateString("/{{base}}/{id:[0-9]{2}}/{name:.+}")
	result, err := template.Apply(map[string]interface{}{"id": int32(42), "name": "dir/file.txt"})
	require.NoError(t, err)
	require.Equal(t, "/{{base}}/42/dir/file.txt", result)
}
