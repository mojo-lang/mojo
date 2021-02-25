package descriptor

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/core/go/pkg/mojo"
)

// MessageDescriptor represents a protocol buffer message.
type MessageDescriptor struct {
	Common
	*descriptor.DescriptorProto

	Parent   *MessageDescriptor   // The containing message, if any.
	Messages []*MessageDescriptor // Inner Messages, if any.
	Enums    []*EnumDescriptor    // Inner Enums, if any.

	Index int    // The Index into the container, whether the File or another message.
	Path  string // The SourceCodeInfo Path as comma-separated integers.
}

func NewMessageDescriptor(file *FileDescriptor) *MessageDescriptor {
	return &MessageDescriptor{
		Common: Common{
			File: file,
		},
		DescriptorProto: &descriptor.DescriptorProto{},
	}
}

// Wrap this Descriptor, recursively
func wrapThisMessageDescriptor(sl []*MessageDescriptor, desc *descriptor.DescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) []*MessageDescriptor {
	sl = append(sl, newMessageDescriptor(desc, parent, file, index))
	me := sl[len(sl)-1]
	for i, nested := range desc.NestedType {
		sl = wrapThisMessageDescriptor(sl, nested, me, file, i)
	}
	return sl
}

// Construct the MessageDescriptor
func newMessageDescriptor(desc *descriptor.DescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) *MessageDescriptor {
	d := &MessageDescriptor{
		Common:          Common{file, nil},
		DescriptorProto: desc,
		Parent:          parent,
		Index:           index,
	}
	if parent == nil {
		//d.Path = fmt.Sprintf("%d,%d", messagePath, Index)
	} else {
		//d.Path = fmt.Sprintf("%s,%d,%d", Parent.Path, messageMessagePath, Index)
	}

	// The only way to distinguish a group from a message is whether
	// the containing message has a TYPE_GROUP field that matches.
	//if Parent != nil {
	//	parts := d.TypeName()
	//	if File.Package != nil {
	//		parts = append([]string{*File.Package}, parts...)
	//	}
	//	exp := "." + strings.Join(parts, ".")
	//	for _, field := range Parent.Field {
	//		if field.GetType() == FIELD_TYPE_GROUP && field.GetTypeName() == exp {
	//			d.group = true
	//			break
	//		}
	//	}
	//}

	return d
}

func (m *MessageDescriptor) IsMessageExist(name string) bool {
	if m == nil {
		return false
	}
	for _, msg := range m.Messages {
		if name == *msg.Name {
			return true
		}
	}
	return false
}

func (m *MessageDescriptor) GetInnerMessage(name string) *MessageDescriptor {
	for _, msg := range m.Messages {
		if msg.Name != nil && *msg.Name == name {
			return msg
		}
	}
	return nil
}

func (m *MessageDescriptor) AddInnerMessage(msg *MessageDescriptor) {
	m.Messages = append(m.Messages, msg)
	m.NestedType = append(m.NestedType, msg.DescriptorProto)
}

func (m *MessageDescriptor) AddInnerEnum(enum *EnumDescriptor) {
	m.Enums = append(m.Enums, enum)
	m.EnumType = append(m.EnumType, enum.EnumDescriptorProto)
}

type FieldDescriptorProto = descriptor.FieldDescriptorProto
type FieldType = descriptor.FieldDescriptorProto_Type

const (
	// 0 is reserved for errors.
	// Order is weird for historical reasons.
	FIELD_TYPE_DOUBLE FieldType = 1
	FIELD_TYPE_FLOAT  FieldType = 2
	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT64 if
	// negative values are likely.
	FIELD_TYPE_INT64  FieldType = 3
	FIELD_TYPE_UINT64 FieldType = 4
	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT32 if
	// negative values are likely.
	FIELD_TYPE_INT32   FieldType = 5
	FIELD_TYPE_FIXED64 FieldType = 6
	FIELD_TYPE_FIXED32 FieldType = 7
	FIELD_TYPE_BOOL    FieldType = 8
	FIELD_TYPE_STRING  FieldType = 9
	// Tag-delimited aggregate.
	// Group type is deprecated and not supported in proto3. However, Proto3
	// implementations should still be able to parse the group wire format and
	// treat group fields as unknown fields.
	FIELD_TYPE_GROUP   FieldType = 10
	FIELD_TYPE_MESSAGE FieldType = 11
	// New in version 2.
	FIELD_TYPE_BYTES    FieldType = 12
	FIELD_TYPE_UINT32   FieldType = 13
	FIELD_TYPE_ENUM     FieldType = 14
	FIELD_TYPE_SFIXED32 FieldType = 15
	FIELD_TYPE_SFIXED64 FieldType = 16
	FIELD_TYPE_SINT32   FieldType = 17
	FIELD_TYPE_SINT64   FieldType = 18

	FIELD_LABEL_OPTIONAL descriptor.FieldDescriptorProto_Label = descriptor.FieldDescriptorProto_LABEL_OPTIONAL
	FIELD_LABEL_REQUIRED descriptor.FieldDescriptorProto_Label = descriptor.FieldDescriptorProto_LABEL_REQUIRED
	FIELD_LABEL_REPEATED descriptor.FieldDescriptorProto_Label = descriptor.FieldDescriptorProto_LABEL_REPEATED
)

var fieldDescriptorProtoTypeName = map[FieldType]string{
	FIELD_TYPE_STRING: "string",
	FIELD_TYPE_BOOL:   "bool",
	FIELD_TYPE_INT32:  "int32",
	FIELD_TYPE_UINT32: "uint32",
	FIELD_TYPE_INT64:  "int64",
	FIELD_TYPE_UINT64: "uint64",
	FIELD_TYPE_FLOAT:  "float",
	FIELD_TYPE_DOUBLE: "double",
	FIELD_TYPE_BYTES:  "bytes",
}

func GetFieldTypeName(field *descriptor.FieldDescriptorProto) string {
	t := fieldDescriptorProtoTypeName[*field.Type]
	if len(t) > 0 {
		return t
	} else {
		return field.GetTypeName()
	}
}

func JsonTagLowerCamelCase(field *descriptor.FieldDescriptorProto) {
	//if field.IsRepeated() || field.IsMessage() {
	//	return
	//}
	//if field.DefaultValue != nil {
	//	return
	//}

	value := strcase.ToLowerCamel(*field.Name)
	value += ",omitempty"

	CompileAliasFieldOptions(field)
	SetStringFieldOption(gogoproto.E_Jsontag, value)(field)
}

func CompileAliasFieldOptions(field *descriptor.FieldDescriptorProto) {
	value := GetStringFieldOption(field, mojo.E_Alias)
	if len(value) > 0 {
		SetStringFieldOption(gogoproto.E_Jsontag, value+",omitempty")(field)
	}
}

func SetStringFieldOption(extension *proto.ExtensionDesc, value string) func(field *descriptor.FieldDescriptorProto) {
	return func(field *descriptor.FieldDescriptorProto) {
		if FieldHasStringExtension(field, extension) {
			return
		}
		if field.Options == nil {
			field.Options = &descriptor.FieldOptions{}
		}
		if err := proto.SetExtension(field.Options, extension, &value); err != nil {
			panic(err)
		}
	}
}

func GetStringFieldOption(field *descriptor.FieldDescriptorProto, extension *proto.ExtensionDesc) string {
	if field.Options == nil {
		return ""
	}
	value, err := proto.GetExtension(field.Options, extension)
	if err != nil || value == nil {
		return ""
	}
	if v := value.(*string); v != nil {
		return *v
	}
	return ""
}

func FieldHasStringExtension(field *descriptor.FieldDescriptorProto, extension *proto.ExtensionDesc) bool {
	if field.Options == nil {
		return false
	}
	value, err := proto.GetExtension(field.Options, extension)
	if err != nil {
		return false
	}
	if value == nil {
		return false
	}
	if value.(*string) == nil {
		return false
	}
	return true
}
