package document

import (
	"math"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttribute_SetBoolProperty(t *testing.T) {
	attribute := &Attribute{}
	attribute.SetBoolProperty("foo", true)

	assert.True(t, attribute.HasProperty("foo"))
	v, err := attribute.GetBoolProperty("foo")

	assert.NoError(t, err)
	assert.True(t, v)
}

func TestAttribute_SetIntProperty(t *testing.T) {
	for _, value := range []int64{0, -1, math.MinInt64, math.MaxInt64} {
		attribute := &Attribute{}
		assert.Same(t, attribute, attribute.SetIntProperty("number", value))
		assert.True(t, attribute.HasProperty("number"))
		got, err := attribute.GetIntProperty("number")
		require.NoError(t, err)
		assert.Equal(t, value, got)
	}
}

func TestAttributeProperties(t *testing.T) {
	for _, tt := range []struct {
		name         string
		newAttribute func() *Attribute
	}{
		{"constructor", NewAttribute},
		{"zero value", func() *Attribute { return &Attribute{} }},
	} {
		t.Run(tt.name, func(t *testing.T) {
			attribute := tt.newAttribute()
			assert.Same(t, attribute, attribute.SetStringProperty("string", "first"))
			assert.Same(t, attribute, attribute.SetIntProperty("int", -42))
			assert.Same(t, attribute, attribute.SetBoolProperty("bool", true))
			integer, err := attribute.GetIntProperty("int")
			require.NoError(t, err)
			assert.Equal(t, int64(-42), integer)
			attribute.SetStringProperty("string", "").SetIntProperty("int", 0).SetBoolProperty("bool", false)
			str, err := attribute.GetStringProperty("string")
			require.NoError(t, err)
			assert.Empty(t, str)
			integer, err = attribute.GetIntProperty("int")
			require.NoError(t, err)
			assert.Zero(t, integer)
			boolean, err := attribute.GetBoolProperty("bool")
			require.NoError(t, err)
			assert.False(t, boolean)
			for _, key := range []string{"string", "int", "bool"} {
				assert.True(t, attribute.HasProperty(key), "zero values are present properties")
			}
			require.Len(t, attribute.Properties, 3)
		})
	}
}

func TestAttributeMissingProperties(t *testing.T) {
	for _, tt := range []struct {
		name      string
		attribute *Attribute
	}{
		{"nil receiver", nil},
		{"nil map", &Attribute{}},
		{"missing key", NewAttribute()},
		{"nil value", &Attribute{Properties: map[string]*core.Value{"missing": nil}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.False(t, tt.attribute.HasProperty("missing"))
			str, err := tt.attribute.GetStringProperty("missing")
			assert.Empty(t, str)
			assert.True(t, core.IsNotFoundError(err), "got %v", err)
			integer, err := tt.attribute.GetIntProperty("missing")
			assert.Zero(t, integer)
			assert.True(t, core.IsNotFoundError(err), "got %v", err)
			boolean, err := tt.attribute.GetBoolProperty("missing")
			assert.False(t, boolean)
			assert.True(t, core.IsNotFoundError(err), "got %v", err)
		})
	}
}

func TestAttributePropertyConversions(t *testing.T) {
	attribute := NewAttribute().SetIntProperty("number", 42).SetStringProperty("integer", "-12").SetStringProperty("boolean", "true")
	str, err := attribute.GetStringProperty("number")
	require.NoError(t, err)
	assert.Equal(t, "42", str)
	integer, err := attribute.GetIntProperty("integer")
	require.NoError(t, err)
	assert.Equal(t, int64(-12), integer)
	boolean, err := attribute.GetBoolProperty("boolean")
	require.NoError(t, err)
	assert.True(t, boolean)
}

func TestAttributeNilSetters(t *testing.T) {
	var attribute *Attribute
	assert.Nil(t, attribute.SetStringProperty("key", "value"))
	assert.Nil(t, attribute.SetIntProperty("key", 1))
	assert.Nil(t, attribute.SetBoolProperty("key", true))
}

func TestNewAttributeIndependentProperties(t *testing.T) {
	first, second := NewAttribute(), NewAttribute()
	require.NotNil(t, first.Properties)
	require.NotNil(t, second.Properties)
	first.SetStringProperty("key", "value")
	assert.False(t, second.HasProperty("key"))
}

func TestAttribute_SetStringProperty(t *testing.T) {
	attribute := NewAttribute()
	attribute.SetStringProperty("foo", "bar")

	assert.True(t, attribute.HasProperty("foo"))
	v, err := attribute.GetStringProperty("foo")

	assert.NoError(t, err)
	assert.Equal(t, "bar", v)
}
