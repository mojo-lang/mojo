package openapi

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

const referenceableExample = `{
    "summary": "example",
    "value": {
        "description": "parameter's description document",
        "in": "path",
        "name": "id",
        "schema": {
            "type": "string"
        }
    }
}`

const referenceableExample2 = `{
    "$ref": "#/components/examples/foo.Bar"
}`

func TestReferenceableExampleCodec_Decode(t *testing.T) {
	param := &ReferenceableExample{}
	err := jsoniter.Unmarshal([]byte(referenceableExample), param)
	assert.NoError(t, err)
	assert.Equal(t, "id", param.GetExample().GetValue().GetObject().GetString("name"))
}

func TestReferenceableExampleCodec_Decode2(t *testing.T) {
	param := &ReferenceableExample{}
	err := jsoniter.Unmarshal([]byte(referenceableExample2), param)
	assert.NoError(t, err)
	assert.Equal(t, "#/components/examples/foo.Bar", param.GetReference().Ref.ToString())
}
