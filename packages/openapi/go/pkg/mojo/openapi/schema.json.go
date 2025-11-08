package openapi

import (
	"sort"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterFieldEncoder("openapi.Schema", "Properties", &SchemaPropertiesEncoder{})
}

type SchemaPropertiesEncoder struct {
}

func (codec *SchemaPropertiesEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	properties := (*map[string]*ReferenceableSchema)(ptr)

	var keys []string
	for key := range *properties {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	stream.WriteObjectStart()
	for i, key := range keys {
		if i > 0 {
			stream.WriteMore()
		}

		stream.WriteObjectField(key)
		stream.WriteVal((*properties)[key])
	}
	stream.WriteObjectEnd()
}

func (codec *SchemaPropertiesEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	properties := (*map[string]*ReferenceableSchema)(ptr)
	return properties == nil || len(*properties) == 0
}
