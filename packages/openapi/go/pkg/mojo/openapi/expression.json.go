package openapi

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

func init() {
	core.RegisterJSONTypeDecoder("openapi.Expression", &ExpressionStringCodec{})
	core.RegisterJSONTypeEncoder("openapi.Expression", &ExpressionStringCodec{})
}

type ExpressionStringCodec struct {
}

func (codec *ExpressionStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	e := (*Expression)(ptr)
	if any.ValueType() == jsoniter.StringValue {
		if err := e.Parse(any.ToString()); err != nil {
			iter.ReportError("ExpressionStringCodec.Decode", err.Error())
		}
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
