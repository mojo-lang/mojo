package openapi

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReferenceableParameterResolution(t *testing.T) {
	value := &Parameter{Name: "inline"}
	inline := NewReferenceableParameter(value)
	assert.Same(t, value, inline.GetParameterOf(nil))
	url, err := core.NewUrl("#" + ParameterReferenceRoot + "Thing")
	require.NoError(t, err)
	reference := &Reference{Ref: url}
	wrapper := NewReferencedParameter(reference)
	assert.Same(t, reference, wrapper.GetReference())
	assert.Nil(t, wrapper.GetParameterOf(nil))
	assert.Nil(t, wrapper.GetParameterOf(map[string]*Parameter{}))
	assert.Same(t, value, wrapper.GetParameterOf(map[string]*Parameter{"Thing": value}))
	assert.Nil(t, (*ReferenceableParameter)(nil).GetParameterOf(nil))
	assert.Nil(t, (&ReferenceableParameter{}).GetParameterOf(nil))
	assert.Nil(t, NewReferencedParameter(&Reference{}).GetParameterOf(nil))
	wrapper.SetParameter(value)
	assert.Nil(t, wrapper.GetReference())
	assert.Same(t, value, wrapper.GetParameter())
	wrapper.SetReference(reference)
	assert.Nil(t, wrapper.GetParameter())
	assert.Same(t, reference, wrapper.GetReference())
	assert.NotPanics(t, func() {
		(*ReferenceableParameter)(nil).SetParameter(value)
		(*ReferenceableParameter)(nil).SetReference(reference)
	})
	assert.Nil(t, NewReferenceableParameter(nil).GetParameter())
	assert.Nil(t, NewReferencedParameter(nil).GetReference())
}
