package openapi

import (
	"testing"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
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
