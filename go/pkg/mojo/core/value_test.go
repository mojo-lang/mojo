package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArrayValue(t *testing.T) {
	val := NewInt32ArrayValue(1, 2, 3, 4)

	vi, ok := val.ToInterface().([]interface{})
	assert.True(t, ok)
	assert.Equal(t, 4, len(vi))
}

func TestValue_GetKind(t *testing.T) {
	val := NewInt32Value(32)
	assert.Equal(t, ValueKind_VALUE_KIND_INTEGER, val.GetKind())

	val = NewStringValue("foo")
	assert.Equal(t, ValueKind_VALUE_KIND_STRING, val.GetKind())

	val = NewStringArrayValue("foo")
	assert.Equal(t, ValueKind_VALUE_KIND_ARRAY, val.GetKind())
}

func TestValue_GetStringArray(t *testing.T) {
	val := NewInt32ArrayValue(1, 2, 3, 4)
	strs := val.GetStringArray()
	assert.Empty(t, strs)

	strs = NewStringArrayValue("foo", "bar", "baz").GetStringArray()
	assert.Equal(t, 3, len(strs))
	assert.Equal(t, "foo", strs[0])
	assert.Equal(t, "baz", strs[2])
}

func TestNewValue(t *testing.T) {
	p := &Platform{Os: OS_OS_LINUX, Architecture: Architecture_ARCHITECTURE_AMD64}
	val, err := NewValue(p)
	assert.NoError(t, err)
	assert.NotNil(t, val)
	assert.Equal(t, p.Format(), val.GetString())
}

func TestNewValue1(t *testing.T) {
	v := &BoolValue{Val: true}
	val, err := NewValue(v)
	assert.NoError(t, err)
	assert.NotNil(t, val)
	assert.Equal(t, true, val.GetBoolVal())
}

func TestNewValue2(t *testing.T) {
	v := &Error{Code: Aborted, Message: "aborted"}
	val, err := NewValue(v)
	assert.NoError(t, err)
	assert.NotNil(t, val)
	assert.Equal(t, v.Code.Format(), val.GetObject().GetString("code"))
}

func TestNewValue3(t *testing.T) {
	v := NewStringMap()
	v.Vals["foo"] = "bar"

	val, err := NewValue(v)
	assert.NoError(t, err)
	assert.NotNil(t, val)
	assert.Equal(t, "bar", val.GetObject().GetString("foo"))
}
