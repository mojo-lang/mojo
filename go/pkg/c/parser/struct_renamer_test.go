package parser

import (
	"context"
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

func TestStructRenamer_ParseNominalType1(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "struct foo",
	}

	err := NewStructRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cStructAttributionName))
	assert.Equal(t, "foo", nominal.Name)
}
