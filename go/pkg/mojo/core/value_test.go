package core

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestNewValueScalarTypes(t *testing.T) {
	for _, tt := range []struct {
		input, want interface{}
		kind        ValueKind
	}{
		{nil, nil, ValueKind_VALUE_KIND_NULL}, {true, true, ValueKind_VALUE_KIND_BOOLEAN},
		{int(-7), int64(-7), ValueKind_VALUE_KIND_INTEGER}, {int8(-8), int64(-8), ValueKind_VALUE_KIND_INTEGER},
		{int16(-16), int64(-16), ValueKind_VALUE_KIND_INTEGER}, {int32(-32), int64(-32), ValueKind_VALUE_KIND_INTEGER},
		{int64(math.MinInt64), int64(math.MinInt64), ValueKind_VALUE_KIND_INTEGER},
		{int64(0), uint64(0), ValueKind_VALUE_KIND_INTEGER}, {uint(7), uint64(7), ValueKind_VALUE_KIND_INTEGER},
		{uint8(8), uint64(8), ValueKind_VALUE_KIND_INTEGER}, {uint16(16), uint64(16), ValueKind_VALUE_KIND_INTEGER},
		{uint32(32), uint64(32), ValueKind_VALUE_KIND_INTEGER}, {uint64(math.MaxUint64), uint64(math.MaxUint64), ValueKind_VALUE_KIND_INTEGER},
		{float32(1.25), float64(1.25), ValueKind_VALUE_KIND_NUMBER}, {float64(2.5), float64(2.5), ValueKind_VALUE_KIND_NUMBER},
		{"hello", "hello", ValueKind_VALUE_KIND_STRING}, {[]byte{0, 255}, []byte{0, 255}, ValueKind_VALUE_KIND_BYTES},
	} {
		t.Run(fmt.Sprintf("%T/%v", tt.input, tt.input), func(t *testing.T) {
			v, err := NewValue(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.kind, v.GetKind())
			assert.Equal(t, tt.want, v.ToInterface())
		})
	}
	for _, tt := range []struct {
		value float64
		want  string
	}{{math.NaN(), "NaN"}, {math.Inf(1), "Infinity"}, {math.Inf(-1), "-Infinity"}} {
		assert.Equal(t, tt.want, NewFloat64Value(tt.value).ToInterface())
	}
}

func TestNewValueCollectionsAndErrors(t *testing.T) {
	for _, tt := range []struct {
		name   string
		input  interface{}
		kind   ValueKind
		length int
	}{
		{"slice", []int{-1, 2}, ValueKind_VALUE_KIND_ARRAY, 2},
		{"array", [2]string{"a", "b"}, ValueKind_VALUE_KIND_ARRAY, 2},
		{"empty slice", []int{}, ValueKind_VALUE_KIND_ARRAY, 0},
		{"nil slice", []int(nil), ValueKind_VALUE_KIND_ARRAY, 0},
		{"typed map", map[string]int{"n": -1}, ValueKind_VALUE_KIND_OBJECT, 1},
		{"empty map", map[string]int{}, ValueKind_VALUE_KIND_OBJECT, 0},
		{"interface slice", []interface{}{true, "x"}, ValueKind_VALUE_KIND_ARRAY, 2},
		{"interface map", map[string]interface{}{"x": true}, ValueKind_VALUE_KIND_OBJECT, 1},
		{"value map", map[string]*Value{"x": NewBoolValue(true)}, ValueKind_VALUE_KIND_OBJECT, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			v, err := NewValue(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.kind, v.GetKind())
			if tt.kind == ValueKind_VALUE_KIND_ARRAY {
				assert.Len(t, v.GetValues(), tt.length)
			} else {
				assert.Len(t, v.GetObject().Vals, tt.length)
			}
		})
	}
	obj := NewObject().SetString("name", "test")
	values := &Values{Vals: []*Value{NewBoolValue(true)}}
	for _, input := range []interface{}{obj, values, NewStringValue("same")} {
		v, err := NewValue(input)
		require.NoError(t, err)
		switch x := input.(type) {
		case *Object:
			assert.Same(t, x, v.GetObject())
		case *Values:
			assert.Same(t, x, v.GetValuesVal())
		case *Value:
			assert.Same(t, x, v)
		}
	}
	for _, input := range []interface{}{string([]byte{255}), make(chan int), func() {}, []interface{}{make(chan int)}, map[string]interface{}{"bad": make(chan int)}, map[int]string{1: "x"}} {
		_, err := NewValue(input)
		assert.Error(t, err, "%T", input)
	}
	t.Run("typed nil pointer", func(t *testing.T) {
		var input *struct{ Name string }
		var v *Value
		require.NotPanics(t, func() { var err error; v, err = NewValue(input); require.NoError(t, err) })
		assert.Equal(t, ValueKind_VALUE_KIND_NULL, v.GetKind())
	})
}

func TestValueScalarAccessors(t *testing.T) {
	assert.True(t, NewStringValue("TRUE").GetBool())
	assert.False(t, NewStringValue("false").GetBool())
	assert.Equal(t, -12, NewStringValue("-12").GetInt())
	assert.Equal(t, int32(-12), NewStringValue("-12").GetInt32())
	assert.Zero(t, NewStringValue("invalid").GetInt64())
	assert.Equal(t, -12, NewIntValue(-12).GetIntVal())
	assert.Equal(t, int32(-12), NewIntValue(-12).GetInt32Val())
	assert.Equal(t, int64(math.MinInt64), NewInt64Value(math.MinInt64).GetInt64())
	assert.Equal(t, uint(12), NewUInt64Value(12).GetUint())
	assert.Equal(t, uint32(12), NewUInt64Value(12).GetUint32())
	assert.Equal(t, uint64(12), NewUInt64Value(12).GetUint64())
	assert.Equal(t, float32(1.25), NewStringValue("1.25").GetFloat32())
	assert.Equal(t, float64(1.25), NewStringValue("1.25").GetDouble())
	assert.Equal(t, float64(-12), NewIntValue(-12).GetFloat64())
	assert.Equal(t, float64(uint64(math.MaxUint64)), NewUInt64Value(math.MaxUint64).GetFloat64())
	assert.Zero(t, NewStringValue("invalid").GetFloat64())
	assert.Equal(t, "-12", NewIntValue(-12).GetString())
	assert.Equal(t, []byte{1, 2}, NewBytesValue([]byte{1, 2}).GetBytes())
	var nilValue *Value
	assert.Equal(t, ValueKind_VALUE_KIND_UNSPECIFIED, nilValue.GetKind())
	assert.Nil(t, nilValue.ToInterface())
	assert.False(t, nilValue.GetBool())
	assert.Zero(t, nilValue.GetInt64())
	assert.Zero(t, nilValue.GetFloat64())
	assert.Empty(t, nilValue.GetString())
	assert.Nil(t, nilValue.GetBytes())
	assert.Nil(t, nilValue.GetObject())
}

func TestValueArrayAccessors(t *testing.T) {
	obj := NewObject().SetString("name", "x")
	assert.Equal(t, []bool{true, false}, NewArrayValue(NewBoolValue(true), NewBoolValue(false)).GetBoolArray())
	assert.Equal(t, []int{-1, 2}, NewIntArrayValue(-1, 2).GetIntArray())
	assert.Equal(t, []int32{-1, 2}, NewInt32ArrayValue(-1, 2).GetInt32Array())
	assert.Equal(t, []int64{-1, 2}, NewInt64ArrayValue(-1, 2).GetInt64Array())
	assert.Equal(t, []uint{1, 2}, NewUInt32ArrayValue(1, 2).GetUintArray())
	assert.Equal(t, []uint32{1, 2}, NewUInt32ArrayValue(1, 2).GetUint32Array())
	assert.Equal(t, []uint64{1, 2}, NewUInt64ArrayValue(1, 2).GetUint64Array())
	assert.Equal(t, []float32{1.25, 2.5}, NewFloat32ArrayValue(1.25, 2.5).GetFloat32Array())
	assert.Equal(t, []float64{1.25, 2.5}, NewFloat64ArrayValue(1.25, 2.5).GetFloat64Array())
	assert.Equal(t, []string{"a", "b"}, NewStringArrayValue("a", "b").GetStringValues().Vals)
	assert.Equal(t, []*Object{obj}, NewObjectArrayValue(obj).GetObjectArray())
	assert.Nil(t, NewArrayValue(NewStringValue("a"), NewIntValue(1)).GetStringArray())
	assert.Nil(t, NewArrayValue().GetStringValues())
	assert.Nil(t, (*Value)(nil).GetObjectArray())
	assert.Empty(t, (*Value)(nil).GetIntArray())
}
