package util

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApplyTemplate(t *testing.T) {
	out, err := ApplyTemplate("abc",
		`{{range  $i, $e := .}}{{if $i}}{{if Last $i $}} and {{else}}, {{end}}{{end}}{{$e}}{{end}}.`,
		[]string{"one", "two", "three"},
		FuncMap)

	if assert.NoError(t, err) {
		str, err := ioutil.ReadAll(out)
		assert.NoError(t, err)

		assert.Equal(t, "one, two and three.", string(str))
	}
}

func TestApplyTemplateErrors(t *testing.T) {
	for _, tt := range []struct{ name, template, message string }{
		{"parse", "{{if}}", "cannot create templates"},
		{"execute", "{{.Missing}}", `attempting to execute templates "broken"`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			reader, err := ApplyTemplate("broken", tt.template, struct{}{}, FuncMap)
			require.ErrorContains(t, err, tt.message)
			require.Nil(t, reader)
		})
	}
}

func TestGoLocalName(t *testing.T) {
	require.Equal(t, "userName", GoLocalName("user_name"))
	require.Equal(t, "typeVar", GoLocalName("Type", "type"))
	require.Equal(t, "typeVarVar", GoLocalName("Type", "type", "typeVar"))
	require.Empty(t, GoLocalName(""))
}
