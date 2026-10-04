package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestObject_From(t *testing.T) {
	image := &Image{
		Title: "test",
		Type:  "png",
		Properties: map[string]*Value{
			"p1": NewInt32Value(1234),
		},
	}

	object := &Object{}
	err := object.From(image)
	assert.NoError(t, err)

	assert.Equal(t, image.Title, object.GetString("title"))
	assert.Equal(t, int32(1234), object.GetValue("properties").GetObject().GetInt32("p1"))
}

func TestObject_From2(t *testing.T) {
	v := &Error{Code: Aborted, Message: "aborted"}

	object := &Object{}
	err := object.From(v)
	assert.NoError(t, err)

	assert.Equal(t, v.Code.Format(), object.GetString("code"))
}

func TestObject_From3(t *testing.T) {
	v := NewStringMap()
	v.Vals["Foo"] = "bar"

	object := &Object{}
	err := object.From(v)
	assert.NoError(t, err)

	assert.Equal(t, "bar", object.GetString("Foo"))
}

func TestObject_From4(t *testing.T) {
	type Foo struct {
		Name string `json:"name"`
	}

	type Test struct {
		Name string `json:"name"`
		Foo  *Foo   `json:"foo"`
	}

	v := &Test{
		Name: "name",
		Foo:  &Foo{},
	}

	object := &Object{}
	err := object.From(v)
	assert.NoError(t, err)
}

func TestObject_From5(t *testing.T) {
	type V struct {
		Name string
		ID   int
		Desc string
	}
	type KeyValuePair struct {
		Key    string `json:"key"`
		Source string `json:"src"`
		Value  *V     `json:"value"`
	}

	kv := KeyValuePair{}
	o := NewObject()
	err := o.From(kv.Value)
	assert.NoError(t, err)
	assert.Empty(t, o.Vals)
}

func TestObject_To(t *testing.T) {
	obj := NewObject().SetString("title", "test").
		SetString("type", "png").
		SetObject("properties", NewObject().SetInt32("p1", 1234))

	image := &Image{}
	err := obj.To(image)
	assert.NoError(t, err)

	assert.Equal(t, "test", image.Title)
	assert.Equal(t, int32(1234), image.Properties["p1"].GetInt32())
}

func TestObject_Delete(t *testing.T) {
	obj := NewObject().SetString("title", "test")
	obj.Delete("title")
	assert.True(t, obj.IsEmpty())
}

func TestObject_Merge(t *testing.T) {
	o1, _ := NewObjectFromKeyValues("x", 12, "y", "foo")
	o2, _ := NewObjectFromKeyValues("x", 16, "z", "bar")
	o1.Merge(o2)
	assert.Equal(t, 3, len(o1.Vals))
	assert.Equal(t, int64(16), o1.GetInt64("x"))
}

func TestMergeObjects(t *testing.T) {
	o1, _ := NewObjectFromKeyValues("x", 12, "y", "foo")
	o2, _ := NewObjectFromKeyValues("x", 16, "z", "bar")
	o3, _ := NewObjectFromKeyValues("a", "alpha", "b", "blue")

	o := MergeObjects(o1, o2, o3)
	assert.NotNil(t, o)
	assert.Equal(t, 5, len(o.Vals))
	assert.Equal(t, int64(16), o.GetInt64("x"))
}

func TestObjectConstructionAndConversion(t *testing.T) {
	obj, err := NewObjectFromKeyValues("name", "x", "count", -2, "name", "latest")
	require.NoError(t, err)
	assert.Equal(t, "latest", obj.GetString("name"))
	assert.Equal(t, int64(-2), obj.GetInt64("count"))
	assert.Equal(t, map[string]interface{}{"name": "latest", "count": int64(-2)}, obj.ToMapInterface())
	assert.Equal(t, obj.Vals, obj.ToMap())
	for _, kvs := range [][]interface{}{{"odd"}, {1, "x"}, {string([]byte{255}), "x"}, {"x", make(chan int)}} {
		_, err := NewObjectFromKeyValues(kvs...)
		assert.Error(t, err)
	}
	for _, input := range []interface{}{obj, NewObjectValue(obj), obj.Vals, map[string]interface{}{"name": "latest", "count": -2}} {
		copy, err := NewObjectFrom(input)
		require.NoError(t, err)
		assert.Equal(t, obj.ToMapInterface(), copy.ToMapInterface())
	}
	for _, input := range []interface{}{true, 1, "x", []int{1}, NewStringValue("x"), struct{ Bad chan int }{make(chan int)}} {
		_, err := NewObjectFrom(input)
		assert.Error(t, err, "%T", input)
	}
	var out struct {
		Name  string
		Count int
	}
	require.NoError(t, obj.To(&out))
	assert.Equal(t, "latest", out.Name)
	assert.Equal(t, -2, out.Count)
	assert.Error(t, obj.To(42))
	assert.NoError(t, (*Object)(nil).From(obj))
	assert.NoError(t, (*Object)(nil).To(&out))
	assert.Nil(t, (*Object)(nil).ToMap())
	assert.Nil(t, (&Object{}).ToMapInterface())
}

func TestObjectMergeAndKeyConversion(t *testing.T) {
	source := NewObject().SetInt("user_id", 7).SetString("user_name", "x")
	target := &Object{}
	require.NotPanics(t, func() { assert.Same(t, target, target.Merge(source)) })
	assert.Equal(t, int64(7), target.GetInt64("user_id"))
	target.Merge(NewObject().SetInt("user_id", 8))
	assert.Equal(t, int64(8), target.GetInt64("user_id"))
	assert.Same(t, target, target.Merge(nil))
	assert.Nil(t, (*Object)(nil).Merge(source))
	camel := source.ToLowerCamelKeys()
	assert.Equal(t, "x", camel.GetString("userName"))
	assert.Nil(t, camel.GetValue("user_name"))
	assert.Equal(t, "x", source.GetString("user_name"), "conversion leaves source keys unchanged")
	assert.Equal(t, source.ToMapInterface(), camel.ToSnakeKeys().ToMapInterface())
	assert.Nil(t, (*Object)(nil).ToLowerCamelKeys())
	assert.Nil(t, (*Object)(nil).ToSnakeKeys())
	assert.Nil(t, (*Object)(nil).Delete("x"))
	target.Delete("user_id").Delete("missing")
	assert.Nil(t, target.GetValue("user_id"))
}

func TestObjectTypedAccessors(t *testing.T) {
	t.Run("Bool", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetBool("key", true))
		assert.Equal(t, true, obj.GetBool("key"))
		var zero bool
		assert.Equal(t, zero, obj.GetBool("missing"))
		assert.Equal(t, zero, (*Object)(nil).GetBool("key"))
		assert.Nil(t, (*Object)(nil).SetBool("key", true))
	})
	t.Run("Int", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetInt("key", -1))
		assert.Equal(t, int64(-1), obj.GetInt64("key"))
		var zero int64
		assert.Equal(t, zero, obj.GetInt64("missing"))
		assert.Equal(t, zero, (*Object)(nil).GetInt64("key"))
		assert.Nil(t, (*Object)(nil).SetInt("key", -1))
	})
	t.Run("Int32", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetInt32("key", -2))
		assert.Equal(t, int32(-2), obj.GetInt32("key"))
		var zero int32
		assert.Equal(t, zero, obj.GetInt32("missing"))
		assert.Equal(t, zero, (*Object)(nil).GetInt32("key"))
		assert.Nil(t, (*Object)(nil).SetInt32("key", -2))
	})
	t.Run("Int64", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetInt64("key", -3))
		assert.Equal(t, int64(-3), obj.GetInt64("key"))
		var zero int64
		assert.Equal(t, zero, obj.GetInt64("missing"))
		assert.Equal(t, zero, (*Object)(nil).GetInt64("key"))
		assert.Nil(t, (*Object)(nil).SetInt64("key", -3))
	})
	t.Run("Uint32", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetUint32("key", 4))
		assert.Equal(t, uint64(4), obj.GetUint64("key"))
		var zero uint64
		assert.Equal(t, zero, obj.GetUint64("missing"))
		assert.Equal(t, zero, (*Object)(nil).GetUint64("key"))
		assert.Nil(t, (*Object)(nil).SetUint32("key", 4))
	})
	t.Run("Uint64", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetUint64("key", 5))
		assert.Equal(t, uint64(5), obj.GetUint64("key"))
		var zero uint64
		assert.Equal(t, zero, obj.GetUint64("missing"))
		assert.Equal(t, zero, (*Object)(nil).GetUint64("key"))
		assert.Nil(t, (*Object)(nil).SetUint64("key", 5))
	})
	t.Run("Float32", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetFloat32("key", 1.25))
		assert.Equal(t, float32(1.25), obj.GetFloat32("key"))
		var zero float32
		assert.Equal(t, zero, obj.GetFloat32("missing"))
		assert.Equal(t, zero, (*Object)(nil).GetFloat32("key"))
		assert.Nil(t, (*Object)(nil).SetFloat32("key", 1.25))
	})
	t.Run("Float64", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetFloat64("key", 2.5))
		assert.Equal(t, float64(2.5), obj.GetFloat64("key"))
		var zero float64
		assert.Equal(t, zero, obj.GetFloat64("missing"))
		assert.Equal(t, zero, (*Object)(nil).GetFloat64("key"))
		assert.Nil(t, (*Object)(nil).SetFloat64("key", 2.5))
	})
	t.Run("String", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetString("key", "x"))
		assert.Equal(t, "x", obj.GetString("key"))
		var zero string
		assert.Equal(t, zero, obj.GetString("missing"))
		assert.Equal(t, zero, (*Object)(nil).GetString("key"))
		assert.Nil(t, (*Object)(nil).SetString("key", "x"))
	})
	t.Run("Int32Array", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetInt32Array("key", -1, 2))
		assert.Equal(t, []int32{-1, 2}, obj.GetInt32Array("key"))
		assert.Empty(t, obj.GetInt32Array("missing"))
		assert.Nil(t, (*Object)(nil).SetInt32Array("key", -1, 2))
	})
	t.Run("Int64Array", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetInt64Array("key", -1, 2))
		assert.Equal(t, []int64{-1, 2}, obj.GetInt64Array("key"))
		assert.Empty(t, obj.GetInt64Array("missing"))
		assert.Nil(t, (*Object)(nil).SetInt64Array("key", -1, 2))
	})
	t.Run("Uint32Array", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetUint32Array("key", 1, 2))
		assert.Equal(t, []uint32{1, 2}, obj.GetUint32Array("key"))
		assert.Empty(t, obj.GetUint32Array("missing"))
		assert.Nil(t, (*Object)(nil).SetUint32Array("key", 1, 2))
	})
	t.Run("Uint64Array", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetUint64Array("key", 1, 2))
		assert.Equal(t, []uint64{1, 2}, obj.GetUint64Array("key"))
		assert.Empty(t, obj.GetUint64Array("missing"))
		assert.Nil(t, (*Object)(nil).SetUint64Array("key", 1, 2))
	})
	t.Run("Float32Array", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetFloat32Array("key", 1.25, 2.5))
		assert.Equal(t, []float32{1.25, 2.5}, obj.GetFloat32Array("key"))
		assert.Empty(t, obj.GetFloat32Array("missing"))
		assert.Nil(t, (*Object)(nil).SetFloat32Array("key", 1.25, 2.5))
	})
	t.Run("Float64Array", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetFloat64Array("key", 1.25, 2.5))
		assert.Equal(t, []float64{1.25, 2.5}, obj.GetFloat64Array("key"))
		assert.Empty(t, obj.GetFloat64Array("missing"))
		assert.Nil(t, (*Object)(nil).SetFloat64Array("key", 1.25, 2.5))
	})
	t.Run("StringArray", func(t *testing.T) {
		obj := &Object{}
		assert.Same(t, obj, obj.SetStringArray("key", "a", "b"))
		assert.Equal(t, []string{"a", "b"}, obj.GetStringArray("key"))
		assert.Empty(t, obj.GetStringArray("missing"))
		assert.Nil(t, (*Object)(nil).SetStringArray("key", "a", "b"))
	})
	child := NewObject().SetString("name", "child")
	obj := (&Object{}).SetObject("child", child).SetObjectArray("children", child)
	assert.Same(t, child, obj.GetObject("child"))
	assert.Equal(t, []*Object{child}, obj.GetObjectArray("children"))
	assert.Nil(t, (*Object)(nil).SetObject("child", child))
	assert.Nil(t, (*Object)(nil).SetObjectArray("children", child))
	assert.Nil(t, obj.GetObject("missing"))
	assert.Empty(t, obj.GetObjectArray("missing"))
	assert.Same(t, obj, obj.SetValue("bools", NewArrayValue(NewBoolValue(true))))
	assert.Equal(t, []bool{true}, obj.GetBoolArray("bools"))
	assert.Nil(t, (*Object)(nil).SetValue("x", NewBoolValue(true)))
}
