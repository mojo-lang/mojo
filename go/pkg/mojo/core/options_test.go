package core

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptions_SetValue(t *testing.T) {
	options := NewOptions().SetValue("k1", 1).SetValue("k2", 2)
	assert.Equal(t, 2, len(options))
	assert.Equal(t, 4, len(options.KeyValues()))
}

func TestOptions_GetValue(t *testing.T) {
	options := NewOptions().SetValue("k1", 1).SetValue("k2", 2)
	assert.True(t, options.HasValue("k1"))
	assert.True(t, options.HasValue("k2"))
	assert.Equal(t, 1, options.GetValue("k1").(int))
}

func TestOptions_GetBool(t *testing.T) {
	options := NewOptions().SetValue("k1", true)
	assert.True(t, options.GetBool("k1"))
}

func TestOptions_GetInteger(t *testing.T) {
	options := NewOptions().SetValue("k1", 100)
	assert.Equal(t, int64(100), options.GetInteger("k1"))
}

func TestOptions_GetFloat(t *testing.T) {
	options := NewOptions().SetValue("k1", 0.1)
	assert.Equal(t, 0.1, options.GetFloat("k1"))

	options = NewOptions().SetValue("k1", 100)
	assert.Equal(t, float64(100), options.GetFloat("k1"))
}

func TestOptions_GetString(t *testing.T) {
	options := NewOptions().SetValue("k1", "foo")
	assert.Equal(t, "foo", options.GetString("k1"))
}

func TestOptions_SetValues(t *testing.T) {
	options := NewOptions("k1", 1, "k2", "k2")
	assert.Equal(t, 2, len(options))

	options.SetValues("k1", 2, "k3", 3)
	assert.Equal(t, 3, len(options))
}

func TestOptions_Merge(t *testing.T) {
	options := NewOptions("k1", 1, "k2", "k2").Merge(NewOptions("k1", 2, "k3", 3))
	assert.Equal(t, 3, len(options))
	assert.Equal(t, 2, options["k1"].(int))
}

func TestOptions_GetValueByPath(t *testing.T) {
	const OptionsJson = `
	{
		"k1": {
			"k11": {
				"k111": {
					"int": 12,
					"float": 12.3
				}
			}
		}
	}
	`

	options := &Options{}
	_ = jsoniter.ConfigFastest.Unmarshal([]byte(OptionsJson), options)

	out := options.GetValueByPath("k1.k11.k111.int")
	assert.Equal(t, float64(12), out.(float64))
}

func TestOptionsNumericTypes(t *testing.T) {
	for _, input := range []interface{}{int(-3), int8(-3), int16(-3), int32(-3), int64(-3), uint(3), uint8(3), uint16(3), uint32(3), uint64(3)} {
		t.Run(fmt.Sprintf("%T", input), func(t *testing.T) {
			options := NewOptions("n", input)
			want := int64(3)
			switch input.(type) {
			case int, int8, int16, int32, int64:
				want = -3
			}
			require.NotPanics(t, func() {
				assert.Equal(t, want, options.GetInteger("n"))
				assert.Equal(t, float64(want), options.GetFloat("n"))
			})
		})
	}
	assert.Equal(t, 1.25, NewOptions("n", float32(1.25)).GetFloat("n"))
	assert.Equal(t, 2.5, NewOptions("n", 2.5).GetFloat("n"))
}

func TestOptionsMissingValuesAndPaths(t *testing.T) {
	for _, options := range []Options{nil, {}, {"wrong": struct{}{}}} {
		for _, key := range []string{"missing", "wrong"} {
			assert.False(t, options.GetBool(key))
			assert.Empty(t, options.GetString(key))
			assert.Zero(t, options.GetInteger(key))
			assert.Zero(t, options.GetFloat(key))
		}
		assert.False(t, options.HasValue("missing"))
		assert.Nil(t, options.GetValue("missing"))
	}
	options := Options{"child": map[string]interface{}{"value": 7}, "scalar": 1}
	assert.Equal(t, 7, options.GetValueByPath("child.value"))
	assert.Equal(t, 1, options.GetValueByPath("scalar"))
	for _, path := range []string{"", "missing", "child.missing", "scalar.value"} {
		assert.Nil(t, options.GetValueByPath(path))
	}
	assert.Empty(t, Options(nil).KeyValues())
	assert.Equal(t, Options{"a": 1}, NewOptions("a", 1, "unpaired"))
	assert.Equal(t, Options{"a": 2, "b": 3}, NewOptions("a", 1).Merge(Options{"a": 2, "b": 3}))
}

func TestOptionsTo(t *testing.T) {
	var out struct {
		Name  string
		Count int
	}
	require.NoError(t, NewOptions("name", "x", "count", 3).To(&out))
	assert.Equal(t, "x", out.Name)
	assert.Equal(t, 3, out.Count)
	assert.NoError(t, Options(nil).To(&out))
	assert.Equal(t, 3, out.Count)
	assert.Error(t, NewOptions("bad", make(chan int)).To(&out))
	assert.Error(t, NewOptions("count", "invalid").To(&out))
	assert.Error(t, NewOptions("count", 3).To(out))
}
