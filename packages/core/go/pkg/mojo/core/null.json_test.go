package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestNullCodec_Encode(t *testing.T) {
	str, err := jsoniter.ConfigFastest.Marshal(&Null{})
	assert.NoError(t, err)
	assert.Empty(t, str)
}
