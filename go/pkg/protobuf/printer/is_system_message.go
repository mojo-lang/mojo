package printer

import "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"

var systemMessages = map[string]bool{
	"Bool":    true,
	"Int8":    true,
	"Int16":   true,
	"Int32":   true,
	"Int64":   true,
	"UInt8":   true,
	"UInt16":  true,
	"UInt32":  true,
	"UInt64":  true,
	"Int":     true,
	"UInt":    true,
	"Float32": true,
	"Float64": true,
	"String":  true,
	"Bytes":   true,
}

func isSystemMessage(msg *descriptor.Message) bool {
	return systemMessages[msg.GetName()]
}
