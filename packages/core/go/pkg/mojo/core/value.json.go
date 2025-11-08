package core

import (
	"encoding/base64"
	"errors"
	"math"
	"strings"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.Value", &ValueCodec{})
	RegisterJSONTypeEncoder("core.Value", &ValueCodec{})
}

const Base64Prefix = "b64."

type ValueCodec struct {
}

func NewValueCodec() *ValueCodec {
	return &ValueCodec{}
}

func (codec *ValueCodec) DecodeAny(any jsoniter.Any) (*Value, error) {
	switch any.ValueType() {
	case jsoniter.NilValue:
		return &Value{}, nil
	case jsoniter.BoolValue:
		return NewBoolValue(any.ToBool()), nil
	case jsoniter.NumberValue:
		floatVal := any.ToFloat64()
		intVal := any.ToInt64() // not support uint64 (the highest bit is 1)
		uintVal := any.ToUint64()

		if intVal == 0 && uintVal > 0 { // > int64.max
			return NewUInt64Value(uintVal), nil
		} else if floatVal != float64(intVal) {
			return NewFloat64Value(floatVal), nil
		} else {
			return NewInt64Value(intVal), nil
		}
	case jsoniter.StringValue:
		str := any.ToString()
		if strings.HasPrefix(str, Base64Prefix) {
			bs, err := base64.StdEncoding.DecodeString(str[len(Base64Prefix):])
			if err != nil {
				return nil, err
			}
			return NewBytesValue(bs), nil
		}
		switch str {
		case "NaN":
			return NewFloat64Value(math.NaN()), nil
		case "Infinity":
			return NewFloat64Value(math.Inf(1)), nil
		case "-Infinity":
			return NewFloat64Value(math.Inf(-1)), nil
		default:
			return NewStringValue(str), nil
		}
	case jsoniter.ObjectValue:
		val := make(map[string]*Value)
		any.ToVal(&val)
		return NewMapValue(val), nil
	case jsoniter.ArrayValue:
		val := make([]*Value, 0)
		any.ToVal(&val)
		return NewArrayValue(val...), nil
	}
	return nil, errors.New("type is invalid")
}

func (codec *ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	v, _ := codec.DecodeAny(any)
	(*Value)(ptr).Val = v.Val
}

func (codec *ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	value := (*Value)(ptr)
	return value == nil || value.Val == nil
}

func (codec *ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	value := (*Value)(ptr)
	switch val := value.Val.(type) {
	case *Value_BoolVal:
		stream.WriteBool(val.BoolVal)
	case *Value_NegativeVal:
		stream.WriteInt64(value.GetInt64())
	case *Value_PositiveVal:
		stream.WriteUint64(val.PositiveVal)
	case *Value_DoubleVal:
		stream.WriteFloat64Lossy(val.DoubleVal)
	case *Value_StringVal:
		stream.WriteString(val.StringVal)
	case *Value_BytesVal:
		stream.WriteString(Base64Prefix + base64.StdEncoding.EncodeToString(val.BytesVal))
	case *Value_ValuesVal:
		stream.WriteVal(val.ValuesVal.Vals)
	case *Value_ObjectVal:
		stream.WriteVal(val.ObjectVal.Vals)
	default:
		stream.WriteNil()
	}
}
