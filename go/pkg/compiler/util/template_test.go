package util

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
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
