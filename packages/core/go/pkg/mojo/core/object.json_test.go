package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestObjectCodec_Decode(t *testing.T) {
	const Val = "{\"key\":\"value\"}"
	object := &Object{}
	err := jsoniter.ConfigFastest.Unmarshal([]byte(Val), object)
	assert.NoError(t, err)
	assert.Equal(t, "value", object.GetString("key"))
}

func TestObjectCodec_Decode_1(t *testing.T) {
	type Foo struct {
		Name   string  `json:"name"`
		Object *Object `json:"object,omitempty"`
	}

	foo := &Foo{}
	err := jsoniter.ConfigFastest.Unmarshal([]byte(`{"name":"foo","object":{"f1":"v1","f2":100}}`), foo)
	assert.NoError(t, err)
	assert.Equal(t, `v1`, foo.Object.GetString("f1"))

	// str, err := jsoniter.ConfigFastest.MarshalToString(&Foo{Name: "foo", Object: NewObject()})
	// assert.NoError(t, err)
	// assert.Equal(t, `{"name":"foo"}`, str)
}

func TestObjectCodec_Decode_Empty(t *testing.T) {
	const Val = "{}"
	object := &Object{}
	err := jsoniter.ConfigFastest.Unmarshal([]byte(Val), object)
	assert.NoError(t, err)
}

func TestObjectCodec_Encode(t *testing.T) {
	type Foo struct {
		Name   string  `json:"name"`
		Object *Object `json:"object,omitempty"`
	}

	str, err := jsoniter.ConfigFastest.MarshalToString(&Foo{Name: "foo"})
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"foo"}`, str)

	// str, err := jsoniter.ConfigFastest.MarshalToString(&Foo{Name: "foo", Object: NewObject()})
	// assert.NoError(t, err)
	// assert.Equal(t, `{"name":"foo"}`, str)
}
