package openapi

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

const referenceableParameter = `{
    "description": "parameter's description document",
    "in": "path",
    "name": "id",
    "schema": {
        "type": "string"
    }
}`

const referenceableParameter2 = `{
    "$ref": "#/components/schemas/foo.Bar"
}`

func TestReferenceableParameterCodec_Decode(t *testing.T) {
	param := &ReferenceableParameter{}
	err := jsoniter.Unmarshal([]byte(referenceableParameter), param)
	assert.NoError(t, err)
	assert.Equal(t, "id", param.GetParameter().Name)
}

func TestReferenceableParameterCodec_Decode2(t *testing.T) {
	param := &ReferenceableParameter{}
	err := jsoniter.Unmarshal([]byte(referenceableParameter2), param)
	assert.NoError(t, err)
	assert.Equal(t, "#/components/schemas/foo.Bar", param.GetReference().Ref.ToString())
}
