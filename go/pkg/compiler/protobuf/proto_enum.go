package protobuf

// ProtoEnum is
type ProtoEnum struct {
	Document []string
	Options  map[string]interface{}

	Name   string
	Values []*ProtoEnumValue
}

// ProtoEnumValue is
type ProtoEnumValue struct {
	Document []string
	Options  map[string]interface{}

	Name   string
	Number int32
}
