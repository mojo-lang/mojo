package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const errorCodeStr = "system.408"

var errorCodeStruct = &ErrorCode{Domain: "system", Code: 408}

func TestErrorCode_Format(t *testing.T) {
	assert.Equal(t, errorCodeStr, errorCodeStruct.Format())
}

func TestErrorCode_Parse(t *testing.T) {
	ec := &ErrorCode{}
	err := ec.Parse(errorCodeStr)

	assert.NoError(t, err)
	assert.Equal(t, errorCodeStruct.Domain, ec.Domain)
	assert.Equal(t, errorCodeStruct.Code, ec.Code)
}
