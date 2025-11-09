package longrunning

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
)

const operationMetadata = `{"name":"130f36bf-547f-4994-b33c-fd7fb289701f","metadata":{"@type":"mojo.core.Checksum","value":"SHA256:a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"}}`

const operationError = `{
    "name": "130f36bf-547f-4994-b33c-fd7fb289701f",
    "done": true,
    "metadata": {
        "@type": "mojo.core.Checksum",
        "value": "SHA256:a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"
    },
    "error": {
        "code": "404"
        "message": "something wrong"
    }
}`

const operationValue = `{
    "name": "130f36bf-547f-4994-b33c-fd7fb289701f",
    "done": true,
    "metadata": {
        "@type": "mojo.core.Checksum",
        "value": "SHA256:a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"
    },
    "response": {
        "@type": "mojo.core.Checksum",
        "value": "SHA256:a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"
    }
}`

func TestOperationCodec_Decode(t *testing.T) {
	operation := &Operation{}
	jsoniter.ConfigFastest.Unmarshal([]byte(operationMetadata), operation)

	assert.Equal(t, "130f36bf-547f-4994-b33c-fd7fb289701f", operation.Name)
	assert.Equal(t, core.Checksum_ALGORITHM_SHA256, operation.Metadata.Get().(*core.Checksum).Algorithm)
}

func TestOperationCodec_Decode_Error(t *testing.T) {
	operation := &Operation{}
	jsoniter.ConfigFastest.Unmarshal([]byte(operationError), operation)

	assert.Equal(t, "130f36bf-547f-4994-b33c-fd7fb289701f", operation.Name)
	assert.Equal(t, core.Checksum_ALGORITHM_SHA256, operation.Metadata.Get().(*core.Checksum).Algorithm)
}

func TestOperationCodec_Decode_Value(t *testing.T) {
	operation := &Operation{}
	jsoniter.ConfigFastest.Unmarshal([]byte(operationValue), operation)

	assert.Equal(t, "130f36bf-547f-4994-b33c-fd7fb289701f", operation.Name)
	assert.Equal(t, core.Checksum_ALGORITHM_SHA256, operation.Metadata.Get().(*core.Checksum).Algorithm)
}

func TestOperationCodec_Encode(t *testing.T) {
	operation := &Operation{
		Name:     "130f36bf-547f-4994-b33c-fd7fb289701f",
		Metadata: core.NewAny(core.NewChecksum(core.Checksum_ALGORITHM_SHA256, []byte("Hello World"))),
		Done:     false,
	}

	str, err := jsoniter.ConfigFastest.MarshalToString(operation)
	assert.NoError(t, err)
	assert.Equal(t, operationMetadata, str)
}
