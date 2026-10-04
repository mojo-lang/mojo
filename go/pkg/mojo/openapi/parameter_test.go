package openapi

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
)

func TestParameter_SupplementExample(t *testing.T) {
	param := &Parameter{
		Name: "test",
		In:   0,
		Schema: NewReferenceableSchema(&Schema{
			Example: core.NewBoolValue(true),
			Type:    Schema_TYPE_BOOLEAN,
		}),
	}

	param.SupplementExample(nil)
	assert.Nil(t, param.Example)
}

func TestParameterExamplePreservation(t *testing.T) {
	explicit := core.NewStringValue("explicit")
	parameter := &Parameter{Schema: exampleSchema(), Example: explicit}
	parameter.SupplementExample(nil)
	assert.Same(t, explicit, parameter.Example)
	parameter.GenerateExample(nil)
	assert.Equal(t, "generated", parameter.Example.GetString())
	parameter.Example = nil
	parameter.Examples = map[string]*ReferenceableExample{"saved": NewReferenceableExample(&Example{Value: explicit})}
	parameter.SupplementExample(nil)
	assert.Nil(t, parameter.Example)
	assert.NotPanics(t, func() {
		(*Parameter)(nil).GenerateExample(nil)
		(*Parameter)(nil).SupplementExample(nil)
		(&Parameter{}).GenerateExample(nil)
		(&Parameter{}).SupplementExample(nil)
	})
}
