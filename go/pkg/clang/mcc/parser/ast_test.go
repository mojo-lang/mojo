package parser

import (
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestAstParser(t *testing.T) {
	data, err := os.ReadFile("./testdata/test.json")
	assert.NoError(t, err)

	term := &Term{}
	err = jsoniter.Config{
		IndentionStep:                 0,
		MarshalFloatWith6Digits:       false,
		EscapeHTML:                    true,
		SortMapKeys:                   false,
		UseNumber:                     false,
		DisallowUnknownFields:         true,
		TagKey:                        "",
		OnlyTaggedField:               false,
		ValidateJsonRawMessage:        false,
		ObjectFieldMustBeSimpleString: false,
		CaseSensitive:                 false,
	}.Froze().Unmarshal(data, term)
	assert.NoError(t, err)
}
