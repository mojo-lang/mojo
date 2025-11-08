package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatJson(t *testing.T) {
	const input = `{"key": "test123", "value": {"foo": 1, "bar": "baz"}}`

	output, err := FormatJson([]byte(input))
	assert.NoError(t, err)

	const expect1 = `{
    "key": "test123",
    "value": {
        "foo": 1,
        "bar": "baz"
    }
}`
	const expect2 = `{
    "key": "test123",
    "value": {
        "bar": "baz",
        "foo": 1
    }
}`
	assert.True(t, string(output) == expect1 || string(output) == expect2, "output: %s", output)
}
