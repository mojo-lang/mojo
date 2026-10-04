package openapi

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReferenceableRequestBodyResolution(t *testing.T) {
	value := &RequestBody{Description: &CachedDocument{Cache: "inline"}}
	inline := NewReferenceableRequestBody(value)
	assert.Same(t, value, inline.GetRequestBodyOf(nil))
	url, err := core.NewUrl("#" + RequestBodyReferenceRoot + "Thing")
	require.NoError(t, err)
	reference := &Reference{Ref: url}
	wrapper := NewReferencedRequestBody(reference)
	assert.Same(t, reference, wrapper.GetReference())
	assert.Nil(t, wrapper.GetRequestBodyOf(nil))
	assert.Nil(t, wrapper.GetRequestBodyOf(map[string]*RequestBody{}))
	assert.Same(t, value, wrapper.GetRequestBodyOf(map[string]*RequestBody{"Thing": value}))
	assert.Nil(t, (*ReferenceableRequestBody)(nil).GetRequestBodyOf(nil))
	assert.Nil(t, (&ReferenceableRequestBody{}).GetRequestBodyOf(nil))
	assert.Nil(t, NewReferencedRequestBody(&Reference{}).GetRequestBodyOf(nil))
	wrapper.SetRequestBody(value)
	assert.Nil(t, wrapper.GetReference())
	assert.Same(t, value, wrapper.GetRequestBody())
	wrapper.SetReference(reference)
	assert.Nil(t, wrapper.GetRequestBody())
	assert.Same(t, reference, wrapper.GetReference())
	assert.NotPanics(t, func() {
		(*ReferenceableRequestBody)(nil).SetRequestBody(value)
		(*ReferenceableRequestBody)(nil).SetReference(reference)
	})
	assert.Nil(t, NewReferenceableRequestBody(nil).GetRequestBody())
	assert.Nil(t, NewReferencedRequestBody(nil).GetReference())
}
