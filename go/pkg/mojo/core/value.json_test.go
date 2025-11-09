package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type ValueTag struct {
	Tag   string `json:"tag"`
	Value *Value `json:"value"`
}

func TestValuesCodec_Decode(t *testing.T) {
	json := `{"tag":"bool", "value": true}`
	vt := &ValueTag{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(json, vt)
	assert.NoError(t, err)
	assert.Equal(t, "bool", vt.Tag)
	assert.Equal(t, ValueKind_VALUE_KIND_BOOLEAN, vt.Value.GetKind())
}

func TestValuesCodec_Decode2(t *testing.T) {
	json := `{"tag":"integer", "value": 9223372036854775808}`
	vt := &ValueTag{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(json, vt)
	assert.NoError(t, err)
	assert.Equal(t, "integer", vt.Tag)
	assert.Equal(t, ValueKind_VALUE_KIND_INTEGER, vt.Value.GetKind())
	assert.Equal(t, uint64(9223372036854775808), vt.Value.GetUint64())
}

func TestValuesCodec_Decode3(t *testing.T) {
	json := `{"tag":"integer", "value": 18446744073709551615}`
	vt := &ValueTag{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(json, vt)
	assert.NoError(t, err)
	assert.Equal(t, "integer", vt.Tag)
	assert.Equal(t, ValueKind_VALUE_KIND_INTEGER, vt.Value.GetKind())
	assert.Equal(t, uint64(18446744073709551615), vt.Value.GetUint64())
}

func TestValuesCodec_Decode4(t *testing.T) {
	json := `{"tag":"integer", "value": -9223372036854775808}`
	vt := &ValueTag{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(json, vt)
	assert.NoError(t, err)
	assert.Equal(t, "integer", vt.Tag)
	assert.Equal(t, ValueKind_VALUE_KIND_INTEGER, vt.Value.GetKind())
	assert.Equal(t, int64(-9223372036854775808), vt.Value.GetInt64())
}

func TestValuesCodec_Decode5(t *testing.T) {
	json := `{"tag":"integer", "value": -9223372036854775809.91}`
	vt := &ValueTag{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(json, vt)
	assert.NoError(t, err)
	assert.Equal(t, "integer", vt.Tag)
	assert.Equal(t, ValueKind_VALUE_KIND_NUMBER, vt.Value.GetKind())
	assert.Equal(t, float64(-9223372036854776000), vt.Value.GetDouble())
}

func TestValuesCodec_Decode6(t *testing.T) {
	json := `{"tag":"integer", "value": 18446744073709551616}`
	vt := &ValueTag{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(json, vt)
	assert.NoError(t, err)
	assert.Equal(t, "integer", vt.Tag)
	assert.Equal(t, ValueKind_VALUE_KIND_NUMBER, vt.Value.GetKind())
	assert.Equal(t, float64(18446744073709552000), vt.Value.GetDouble())
}

func TestValuesCodec_Decode7(t *testing.T) {
	json := `{"tag":"integer", "value": -9223372036854775808.98765432}`
	vt := &ValueTag{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(json, vt)
	assert.NoError(t, err)
	assert.Equal(t, "integer", vt.Tag)
	assert.Equal(t, ValueKind_VALUE_KIND_INTEGER, vt.Value.GetKind())
	assert.Equal(t, int64(-9223372036854775808), vt.Value.GetInt64())
}

func TestValuesCodec_Decode8(t *testing.T) {
	json := `{"tag":"integer", "value": 0}`
	vt := &ValueTag{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(json, vt)
	assert.NoError(t, err)
	assert.Equal(t, "integer", vt.Tag)
	assert.Equal(t, ValueKind_VALUE_KIND_INTEGER, vt.Value.GetKind())
	assert.Equal(t, int64(0), vt.Value.GetInt64())
}
