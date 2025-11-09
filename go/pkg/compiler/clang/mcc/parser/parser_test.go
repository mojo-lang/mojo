package parser

import (
	"os"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestParser_ParseJSON(t *testing.T) {
	file := "./testdata/mz_strm_zstd.c.mcc.json"
	data, err := os.ReadFile(file)
	assert.NoError(t, err)

	parser := &Parser{SrcPath: strings.TrimSuffix(file, ".mcc.json")}
	sf, err := parser.ParseJSON(data)
	assert.NoError(t, err)
	assert.NotNil(t, sf)

	bytes, err := jsoniter.ConfigDefault.MarshalIndent(sf, "", "    ")
	assert.NoError(t, err)
	assert.NotEmpty(t, bytes)
}
