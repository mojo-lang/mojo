package core

import (
	"errors"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.Id", &IdCodec{})
	RegisterJSONTypeEncoder("core.Id", &IdCodec{})
}

type IdCodec struct {
}

func NewIdCodec() *IdCodec {
	return &IdCodec{}
}

func (codec *IdCodec) DecodeAny(any jsoniter.Any) (*Id, error) {
	switch any.ValueType() {
	case jsoniter.NumberValue:
		return NewIntId(any.ToUint64()), nil
	case jsoniter.StringValue:
		return ParseId(any.ToString())
	}
	return nil, errors.New("type is invalid for ID")
}

func (codec *IdCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	id := (*Id)(ptr)

	switch any.ValueType() {
	case jsoniter.NumberValue:
		id.SetInt(any.ToUint64())
	case jsoniter.StringValue:
		value := any.ToString()
		uuid, err := ParseUuid(value)
		if err != nil {
			id.SetUuid(uuid)
		} else {
			id.SetString(value)
		}
	default:
		iter.ReportError("Id Decode", "invalid value type")
	}
}

func (codec *IdCodec) IsEmpty(ptr unsafe.Pointer) bool {
	id := (*Id)(ptr)
	return id == nil || id.Id == nil
}

func (codec *IdCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	id := (*Id)(ptr)
	stream.WriteVal(id.Format())
}
