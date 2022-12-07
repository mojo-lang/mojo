package parser

import (
	"context"
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

func TestPointerRenamer_ParseNominalType1(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "char* *",
	}

	err := NewPointerRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cDoublePointerAttributionName))
	assert.Equal(t, "char", nominal.Name)
}

func TestPointerRenamer_ParseNominalType2(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "char* * const",
	}

	err := NewPointerRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.Equal(t, "char* * const", nominal.Name)
}
