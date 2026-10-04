package openapi

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReferenceableResponseResolution(t *testing.T) {
	value := &Response{Description: "inline"}
	inline := NewReferenceableResponse(value)
	assert.Same(t, value, inline.GetResponseOf(nil))
	url, err := core.NewUrl("#" + ResponseReferenceRoot + "Thing")
	require.NoError(t, err)
	reference := &Reference{Ref: url}
	wrapper := NewReferencedResponse(reference)
	assert.Same(t, reference, wrapper.GetReference())
	assert.Nil(t, wrapper.GetResponseOf(nil))
	assert.Nil(t, wrapper.GetResponseOf(map[string]*Response{}))
	assert.Same(t, value, wrapper.GetResponseOf(map[string]*Response{"Thing": value}))
	assert.Nil(t, (*ReferenceableResponse)(nil).GetResponseOf(nil))
	assert.Nil(t, (&ReferenceableResponse{}).GetResponseOf(nil))
	assert.Nil(t, NewReferencedResponse(&Reference{}).GetResponseOf(nil))
	wrapper.SetResponse(value)
	assert.Nil(t, wrapper.GetReference())
	assert.Same(t, value, wrapper.GetResponse())
	wrapper.SetReference(reference)
	assert.Nil(t, wrapper.GetResponse())
	assert.Same(t, reference, wrapper.GetReference())
	assert.NotPanics(t, func() {
		(*ReferenceableResponse)(nil).SetResponse(value)
		(*ReferenceableResponse)(nil).SetReference(reference)
	})
	assert.Nil(t, NewReferenceableResponse(nil).GetResponse())
	assert.Nil(t, NewReferencedResponse(nil).GetReference())
}
