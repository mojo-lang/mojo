package core

import (
	"fmt"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.Float32Value", &Float32ValueCodec{})
	RegisterJSONTypeEncoder("core.Float32Value", &Float32ValueCodec{})
	RegisterJSONTypeDecoder("core.Float64Value", &Float64ValueCodec{})
	RegisterJSONTypeEncoder("core.Float64Value", &Float64ValueCodec{})
}

type Float32ValueCodec struct {
}

func (codec *Float32ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadAny()
	if val.ValueType() == jsoniter.NumberValue {
		(*Float32Value)(ptr).Val = val.ToFloat32()
	} else {
		iter.ReportError("number", fmt.Sprintf("invalid float32 type value, original type: %d", val.ValueType()))
	}
}

func (codec *Float32ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteRaw(fmt.Sprint((*Float32Value)(ptr).Val))
}

func (codec *Float32ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*Float32Value)(ptr)
	return e == nil
}

type Float64ValueCodec struct {
}

func (codec *Float64ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadAny()
	if val.ValueType() == jsoniter.NumberValue {
		(*Float64Value)(ptr).Val = val.ToFloat64()
	} else {
		iter.ReportError("number", fmt.Sprintf("invalid float64 type value, original type: %d", val.ValueType()))
	}
}

func (codec *Float64ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteRaw(fmt.Sprint((*Float64Value)(ptr).Val))
}

func (codec *Float64ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*Float64Value)(ptr)
	return e == nil
}
