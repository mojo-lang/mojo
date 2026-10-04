package openapi

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReferenceableLinkResolution(t *testing.T) {
	value := &Link{OperationId: "inline"}
	inline := NewReferenceableLink(value)
	assert.Same(t, value, inline.GetLinkOf(nil))
	url, err := core.NewUrl("#" + LinkReferenceRoot + "Thing")
	require.NoError(t, err)
	reference := &Reference{Ref: url}
	wrapper := NewReferencedLink(reference)
	assert.Same(t, reference, wrapper.GetReference())
	assert.Nil(t, wrapper.GetLinkOf(nil))
	assert.Nil(t, wrapper.GetLinkOf(map[string]*Link{}))
	assert.Same(t, value, wrapper.GetLinkOf(map[string]*Link{"Thing": value}))
	assert.Nil(t, (*ReferenceableLink)(nil).GetLinkOf(nil))
	assert.Nil(t, (&ReferenceableLink{}).GetLinkOf(nil))
	assert.Nil(t, NewReferencedLink(&Reference{}).GetLinkOf(nil))
	wrapper.SetLink(value)
	assert.Nil(t, wrapper.GetReference())
	assert.Same(t, value, wrapper.GetLink())
	wrapper.SetReference(reference)
	assert.Nil(t, wrapper.GetLink())
	assert.Same(t, reference, wrapper.GetReference())
	assert.NotPanics(t, func() {
		(*ReferenceableLink)(nil).SetLink(value)
		(*ReferenceableLink)(nil).SetReference(reference)
	})
	assert.Nil(t, NewReferenceableLink(nil).GetLink())
	assert.Nil(t, NewReferencedLink(nil).GetReference())
}
