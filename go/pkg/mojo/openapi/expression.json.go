package openapi

import (
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

func init() {
	core.RegisterJSONTypeDecoder("openapi.Expression", &ExpressionStringCodec{})
	core.RegisterJSONTypeEncoder("openapi.Expression", &ExpressionStringCodec{})
}

type ExpressionStringCodec struct {
}

func (codec *ExpressionStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}
	e := (*Expression)(ptr)
	switch any.ValueType() {
	case jsoniter.StringValue:
		if err := e.Parse(any.ToString()); err != nil {
			iter.ReportError("ExpressionStringCodec.Decode", err.Error())
		}
	case jsoniter.NilValue:
		e.Type, e.Source = Expression_TYPE_UNSPECIFIED, nil
	default:
		iter.ReportError("ExpressionStringCodec.Decode", "expected a string or null")
	}
}

func (codec *ExpressionStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	e := (*Expression)(ptr)
	stream.WriteString(e.Format())
}

func (codec *ExpressionStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	e := (*Expression)(ptr)
	return e == nil || e.Type == Expression_TYPE_UNSPECIFIED
}
