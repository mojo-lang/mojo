package protobuf

// ProtoMessage name is message
type ProtoMessage struct {
	Enums    []*ProtoEnum
	Messages []*ProtoMessage

	Document []string
	Options  map[string]interface{}

	Name   string
	Fields []*ProtoMessageField
}

// ProtoMessageField name is field
type ProtoMessageField struct {
	Document []string
	Options  map[string]interface{}

	Name   string
	Number int32

	Type string

	IsOneOf     bool
	OneOfFields []*ProtoMessageField
}
