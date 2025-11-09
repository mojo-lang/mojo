package core

import (
	jsoniter "github.com/json-iterator/go"
	"testing"

	"github.com/stretchr/testify/assert"
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
