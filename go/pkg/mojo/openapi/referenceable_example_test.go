package openapi

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReferenceableExampleResolution(t *testing.T) {
	value := &Example{Summary: "inline"}
	inline := NewReferenceableExample(value)
	assert.Same(t, value, inline.GetExampleOf(nil))
	url, err := core.NewUrl("#" + ExampleReferenceRoot + "Thing")
	require.NoError(t, err)
	reference := &Reference{Ref: url}
	wrapper := NewReferencedExample(reference)
	assert.Same(t, reference, wrapper.GetReference())
	assert.Nil(t, wrapper.GetExampleOf(nil))
	assert.Nil(t, wrapper.GetExampleOf(map[string]*Example{}))
	assert.Same(t, value, wrapper.GetExampleOf(map[string]*Example{"Thing": value}))
	assert.Nil(t, (*ReferenceableExample)(nil).GetExampleOf(nil))
	assert.Nil(t, (&ReferenceableExample{}).GetExampleOf(nil))
	assert.Nil(t, NewReferencedExample(&Reference{}).GetExampleOf(nil))
	wrapper.SetExample(value)
	assert.Nil(t, wrapper.GetReference())
	assert.Same(t, value, wrapper.GetExample())
	wrapper.SetReference(reference)
	assert.Nil(t, wrapper.GetExample())
	assert.Same(t, reference, wrapper.GetReference())
	assert.NotPanics(t, func() {
		(*ReferenceableExample)(nil).SetExample(value)
		(*ReferenceableExample)(nil).SetReference(reference)
	})
	assert.Nil(t, NewReferenceableExample(nil).GetExample())
	assert.Nil(t, NewReferencedExample(nil).GetReference())
}
