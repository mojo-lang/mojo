package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValues_ToSlice(t *testing.T) {
	values, err := NewValues(1, 2, 3, 4, 5)
	assert.NoError(t, err)
	assert.Equal(t, 5, values.Len())
	assert.Equal(t, uint64(1), values.ToSlice()[0].(uint64))
}

func TestValuesAppendAndRead(t *testing.T) {
	obj := NewObject().SetString("name", "x")
	for _, tt := range []struct {
		name   string
		append func(*Values) *Values
		want   interface{}
	}{
		{"bool", func(v *Values) *Values { return v.AppendBool(true) }, true},
		{"int", func(v *Values) *Values { return v.AppendInt(-1) }, int64(-1)},
		{"int32", func(v *Values) *Values { return v.AppendInt32(-2) }, int64(-2)},
		{"int64", func(v *Values) *Values { return v.AppendInt64(-3) }, int64(-3)},
		{"uint32", func(v *Values) *Values { return v.AppendUint32(4) }, uint64(4)},
		{"uint64", func(v *Values) *Values { return v.AppendUint64(5) }, uint64(5)},
		{"float32", func(v *Values) *Values { return v.AppendFloat32(1.25) }, float64(1.25)},
		{"float64", func(v *Values) *Values { return v.AppendFloat64(2.5) }, float64(2.5)},
		{"string", func(v *Values) *Values { return v.AppendString("x") }, "x"},
		{"object", func(v *Values) *Values { return v.AppendObject(obj) }, map[string]interface{}{"name": "x"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			v := &Values{}
			assert.Same(t, v, tt.append(v))
			require.Equal(t, 1, v.Len())
			assert.Equal(t, tt.want, v.GetValue(0).ToInterface())
			assert.Nil(t, tt.append(nil))
		})
	}
	v := &Values{}
	assert.Same(t, v, v.AppendValues(NewStringValue("a"), NewBoolValue(true), NewInt32Value(-2), NewUInt64Value(3)))
	assert.Equal(t, "a", v.GetString(0))
	assert.True(t, v.GetBool(1))
	assert.Equal(t, int32(-2), v.GetInt32(2))
	assert.Equal(t, int64(-2), v.GetInt64(2))
	assert.Equal(t, uint64(3), v.GetUint64(3))
	assert.Equal(t, []interface{}{"a", true, int64(-2), uint64(3)}, v.ToSlice())
	for _, index := range []int{-1, 4, 100} {
		assert.Nil(t, v.GetValue(index))
		assert.Empty(t, v.GetString(index))
		assert.False(t, v.GetBool(index))
		assert.Zero(t, v.GetInt64(index))
	}
	var absent *Values
	assert.Zero(t, absent.Len())
	assert.Empty(t, absent.ToSlice())
	assert.Nil(t, absent.AppendValue(NewBoolValue(true)))
	assert.Nil(t, absent.AppendValues(NewBoolValue(true)))
	assert.Nil(t, absent.GetValue(0))
	empty, err := NewValues()
	require.NoError(t, err)
	assert.Empty(t, empty.ToSlice())
	_, err = NewValues("valid", make(chan int))
	assert.Error(t, err)
}

func TestValuesNestedArrays(t *testing.T) {
	obj := NewObject().SetString("name", "x")
	v := (&Values{}).AppendValues(NewInt32ArrayValue(-1), NewInt64ArrayValue(-2), NewUInt32ArrayValue(3), NewUInt64ArrayValue(4), NewStringArrayValue("x"), NewObjectArrayValue(obj))
	assert.Equal(t, []int32{-1}, v.GetInt32Array(0))
	assert.Equal(t, []int64{-2}, v.GetInt64Array(1))
	assert.Equal(t, []uint32{3}, v.GetUint32Array(2))
	assert.Equal(t, []uint64{4}, v.GetUint64Array(3))
	assert.Equal(t, []string{"x"}, v.GetStringArray(4))
	assert.Equal(t, []*Object{obj}, v.GetObjectArray(5))
	assert.Empty(t, v.GetInt32Array(-1))
	assert.Empty(t, v.GetStringArray(6))
}
