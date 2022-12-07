package parser

import (
	"context"
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

func TestConstRenamer_ParseNominalType1(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "const char* const",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cConstAttributionName))
	assert.True(t, nominal.HasAttribute(cPointerConstAttributeName))
	assert.Equal(t, "char*", nominal.Name)
}

func TestConstRenamer_ParseNominalType2(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "const char*",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cConstAttributionName))
	assert.Equal(t, "char*", nominal.Name)
}

func TestConstRenamer_ParseNominalType3(t *testing.T) {
	nominal := &lang.NominalType{
		Name: " char* const",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cPointerConstAttributeName))
	assert.Equal(t, "char*", nominal.Name)
}

func TestConstRenamer_ParseNominalType4(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "char const* const",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cConstAttributionName))
	assert.True(t, nominal.HasAttribute(cPointerConstAttributeName))
	assert.Equal(t, "char *", nominal.Name)
}

func TestConstRenamer_ParseNominalType5(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "char const*",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cConstAttributionName))
	assert.Equal(t, "char *", nominal.Name)
}

func TestConstRenamer_ParseNominalType6(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "char const",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cConstAttributionName))
	assert.Equal(t, "char", nominal.Name)
}

func TestConstRenamer_ParseNominalType7(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "const char **",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cConstAttributionName))
	assert.Equal(t, "char **", nominal.Name)
}

func TestConstRenamer_ParseNominalType8(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "char * const *",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.False(t, nominal.HasAttribute(cConstAttributionName))
	assert.Equal(t, "char * const *", nominal.Name)
}

func TestConstRenamer_ParseNominalType9(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "const char * const *",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cConstAttributionName))
	assert.Equal(t, "char * const *", nominal.Name)
}

func TestConstRenamer_ParseNominalType10(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "char ** const",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cPointerConstAttributeName))
	assert.Equal(t, "char **", nominal.Name)
}

func TestConstRenamer_ParseNominalType11(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "const char ** const",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cConstAttributionName))
	assert.True(t, nominal.HasAttribute(cPointerConstAttributeName))
	assert.Equal(t, "char **", nominal.Name)
}

func TestConstRenamer_ParseNominalType12(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "char * const * const",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cPointerConstAttributeName))
	assert.Equal(t, "char * const *", nominal.Name)
}

func TestConstRenamer_ParseNominalType13(t *testing.T) {
	nominal := &lang.NominalType{
		Name: "const char * const * const",
	}

	err := NewConstRenamer(nil).ParseNominalType(context.Background(), nominal)
	assert.NoError(t, err)
	assert.True(t, nominal.HasAttribute(cConstAttributionName))
	assert.True(t, nominal.HasAttribute(cPointerConstAttributeName))
	assert.Equal(t, "char * const *", nominal.Name)
}
