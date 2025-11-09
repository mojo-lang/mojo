package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLogLevel(t *testing.T) {
	assert.Equal(t, DebugLevel, ParseLogLevel("Debug"))
	assert.Equal(t, InfoLevel, ParseLogLevel("info"))
	assert.Equal(t, ErrorLevel, ParseLogLevel("ERROR"))
	assert.Equal(t, InfoLevel, ParseLogLevel("ERR"))
}
