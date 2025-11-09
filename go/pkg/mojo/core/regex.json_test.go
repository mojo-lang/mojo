package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

const regexExpr = "^[A-Za-z0-9_.\u4E00-\u9FA5]{1,32}$"
const regexExprJson = `"^[A-Za-z0-9_.\u4E00-\u9FA5]{1,32}$"`

func TestRegex_Decode(t *testing.T) {
	regex := &Regex{}
	err := jsoniter.Unmarshal([]byte(regexExprJson), regex)

	assert.NoError(t, err)
	assert.Equal(t, regexExpr, regex.Expression)
}

func TestRegex_Encode(t *testing.T) {
	regex := &Regex{
		Expression: regexExpr,
	}
	str, err := jsoniter.MarshalToString(regex)

	assert.NoError(t, err)
	assert.NotEmpty(t, str)
	// assert.Equal(t, regexExprJson, str)
}
