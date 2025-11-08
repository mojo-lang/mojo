package test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

// to avoid cyclic core package reference, move the test to separate folder

func TestLogJSON(t *testing.T) {
	const logFileName = "./test.log"
	l := logs.New(&logs.Config{
		Output: "file",
		File: logs.FileSinkConfig{
			Path: logFileName,
		},
	})

	value := core.NewStringValue("foo")
	values := core.NewStringArrayValue("foo", "bar", "baz")
	mapVals := core.NewMapValue(map[string]*core.Value{"foo": core.NewBoolValue(true), "bar": core.NewInt32Value(100)})
	l.Infow("info", "value", value, "values", values, "map", mapVals)

	content, err := os.ReadFile(logFileName)
	assert.NoError(t, err)
	assert.NotEmpty(t, content)
	assert.True(t, strings.Contains(string(content), "{\"value\": \"foo\", \"values\": [\"foo\",\"bar\",\"baz\"], \"map\": {\"foo\":true,\"bar\":100}}"))

	_ = os.Remove(logFileName)
}
