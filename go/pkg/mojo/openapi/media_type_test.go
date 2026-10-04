package openapi

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func exampleSchema() *ReferenceableSchema {
	return NewReferenceableSchema(&Schema{Type: Schema_TYPE_STRING, Enum: []*core.Value{core.NewStringValue("generated")}})
}

func TestMediaTypeExamples(t *testing.T) {
	for _, mediaType := range []string{core.ApplicationJson, core.ApplicationWwwFormUrlencoded, "text/plain"} {
		t.Run(mediaType, func(t *testing.T) {
			media := &MediaType{Schema: exampleSchema()}
			media.SupplementExample(nil, mediaType)
			require.NotNil(t, media.Example)
			assert.Equal(t, "generated", media.Example.GetString())
			explicit := core.NewStringValue("explicit")
			media.Example = explicit
			media.SupplementExample(nil, mediaType)
			assert.Same(t, explicit, media.Example)
			media.GenerateExample(nil, mediaType)
			assert.Equal(t, "generated", media.Example.GetString())
		})
	}
	for _, media := range []*MediaType{
		{Schema: exampleSchema(), Examples: map[string]*ReferenceableExample{"saved": NewReferenceableExample(&Example{Value: core.NewStringValue("saved")})}},
		{Schema: NewReferenceableSchema(&Schema{Type: Schema_TYPE_STRING, Example: core.NewStringValue("schema example")})},
		{},
	} {
		media.SupplementExample(nil, core.ApplicationJson)
		assert.Nil(t, media.Example)
	}
	url, err := core.NewUrl("#/components/schemas/Name")
	require.NoError(t, err)
	media := &MediaType{Schema: NewReferencedSchema(&Reference{Ref: url})}
	media.SupplementExample(map[string]*Schema{"Name": {Type: Schema_TYPE_STRING, Example: core.NewStringValue("referenced")}}, core.ApplicationJson)
	assert.Equal(t, "referenced", media.Example.GetString())
	assert.NotPanics(t, func() {
		(*MediaType)(nil).GenerateExample(nil, "")
		(*MediaType)(nil).SupplementExample(nil, "")
		(&MediaType{}).GenerateExample(nil, "")
	})
}
