package descriptor

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	FieldTypeBool   = "bool"
	FieldTypeInt32  = "int32"
	FieldTypeUInt32 = "uint32"
	FieldTypeInt64  = "int64"
	FieldTypeUInt64 = "uint64"
	FieldTypeFloat  = "float"
	FieldTypeDouble = "double"
	FieldTypeString = "string"
	FieldTypeBytes  = "bytes"

	FieldTypeFixed32 = "fixed32"
	FieldTypeFixed64 = "fixed64"

	FieldTypeSInt32   = "sint32"
	FieldTypeSInt64   = "sint64"
	FieldTypeSFixed32 = "sfixed32"
	FieldTypeSFixed64 = "sfixed64"
)

// A Field describes a message field.
type Field struct {
	Descriptor
	Proto *descriptorpb.FieldDescriptorProto

	Parent *Message // message in which this field is declared; nil if top-level extension

	Oneof *Oneof // containing oneof; nil if not part of an oneof

	Enum    *Enum    // type for enum fields; nil otherwise
	Message *Message // type for message or group fields; nil otherwise
}

var (
	enumType    = descriptorpb.FieldDescriptorProto_TYPE_ENUM
	messageType = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE

	repeated = descriptorpb.FieldDescriptorProto_LABEL_REPEATED
)

func NewField(parent *Message, name string) *Field {
	field := &Field{
		Descriptor: Descriptor{
			File: parent.File,
		},
		Proto: &descriptorpb.FieldDescriptorProto{
			Name: &name,
		},
		Parent: parent,
	}
	return field
}

func NewEnumField(parent *Message, name string, enum *Enum) *Field {
	field := &Field{
		Descriptor: Descriptor{
			File: parent.File,
		},
		Proto: &descriptorpb.FieldDescriptorProto{
			Name: &name,
			Type: &enumType,
		},
		Parent: parent,
		Enum:   enum,
	}
	return field
}

func NewMessageField(parent *Message, name string, message *Message) *Field {
	field := &Field{
		Descriptor: Descriptor{
			File: parent.File,
		},
		Proto: &descriptorpb.FieldDescriptorProto{
			Name: &name,
			Type: &messageType,
		},
		Parent:  parent,
		Message: message,
	}
	return field
}

func NewFieldFrom(parent *Message, proto *descriptorpb.FieldDescriptorProto) *Field {
	field := &Field{
		Descriptor: Descriptor{
			File: parent.File,
		},
		Proto:  proto,
		Parent: parent,
	}
	return field
}

func (m *Field) proto() *descriptorpb.FieldDescriptorProto {
	if m != nil {
		return m.Proto
	}
	return nil
}

func (m *Field) IsEnumType() bool {
	return m.proto().GetType() == descriptorpb.FieldDescriptorProto_TYPE_ENUM
}

func (m *Field) IsMessageType() bool {
	return m.proto().GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
}

func (m *Field) GetName() string {
	return m.proto().GetName()
}

func (m *Field) GetNumber() int32 {
	return m.proto().GetNumber()
}

func (m *Field) GetTypeName() string {
	if name, ok := fieldDescriptorProtoTypeName[*m.Proto.Type]; ok {
		return name
	}
	return m.proto().GetTypeName()
}

// GetOptions get the predefined options in the field options
func (m *Field) GetOptions() *descriptorpb.FieldOptions {
	return m.proto().GetOptions()
}

func (m *Field) HasOptions() bool {
	return m.GetOptions() != nil
}

func (m *Field) HasOption(extension protoreflect.ExtensionType) bool {
	return proto.HasExtension(m.Proto.Options, extension)
}

func (m *Field) HasExtension(extension protoreflect.ExtensionType) bool {
	return m.HasOption(extension)
}

func (m *Field) GetBoolOption(extension protoreflect.ExtensionType) bool {
	if v, ok := m.GetOption(extension).(bool); ok {
		return v
	}
	return false
}

func (m *Field) GetInt64Option(extension protoreflect.ExtensionType) int64 {
	if v, ok := m.GetOption(extension).(int64); ok {
		return v
	}
	return 0
}

func (m *Field) GetFloat64Option(extension protoreflect.ExtensionType) float64 {
	if v, ok := m.GetOption(extension).(float64); ok {
		return v
	}
	return 0
}

func (m *Field) GetStringOption(extension protoreflect.ExtensionType) string {
	if v, ok := m.GetOption(extension).(string); ok {
		return v
	}
	return ""
}

func (m *Field) GetOption(extension protoreflect.ExtensionType) interface{} {
	if m.Proto.GetOptions() != nil {
		return proto.GetExtension(m.Proto.Options, extension)
	}

	return nil
}

func (m *Field) SetPacked() *Field {
	if m != nil && m.Proto != nil {
		if m.Proto.Options == nil {
			m.Proto.Options = &descriptorpb.FieldOptions{}
		}
		packed := true
		m.Proto.Options.Packed = &packed
	}
	return m
}

func (m *Field) SetOption(extension protoreflect.ExtensionType, value interface{}) *Field {
	if !m.HasExtension(extension) {
		if m.Proto.Options == nil {
			m.Proto.Options = &descriptorpb.FieldOptions{}
		}
	}
	proto.SetExtension(m.Proto.Options, extension, value)
	return m
}

func (m *Field) SetBoolOption(extension protoreflect.ExtensionType, value bool) *Field {
	return m.SetOption(extension, value)
}

func (m *Field) SetStringOption(extension protoreflect.ExtensionType, value string) *Field {
	return m.SetOption(extension, value)
}

func (m *Field) SetInt64Option(extension protoreflect.ExtensionType, value int64) *Field {
	return m.SetOption(extension, value)
}

func (m *Field) SetFloat64Option(extension protoreflect.ExtensionType, value float64) *Field {
	return m.SetOption(extension, value)
}

func (m *Field) IsRepeated() bool {
	return m.proto().GetLabel() == repeated
}

func (m *Field) SetRepeated() *Field {
	if m != nil && m.Proto != nil {
		m.Proto.Label = &repeated
	}
	return m
}

func (m *Field) SetName(name string) *Field {
	if m != nil && m.Proto != nil {
		m.Proto.Name = &name
	}
	return m
}

func (m *Field) SetNumber(number int32) *Field {
	if m != nil && m.Proto != nil {
		m.Proto.Number = &number
	}
	return m
}

func (m *Field) SetType(t string, options ...string) *Field {
	if m != nil && m.Proto != nil {
		typ := protoType(t)
		m.Proto.Type = &typ
	}
	return m
}

func (m *Field) SetTypeName(name string, options ...string) *Field {
	if m != nil && m.Proto != nil {
		name = protoTypeName(name)
		m.Proto.TypeName = &name
	}
	return m
}

var fieldDescriptorProtoTypeName = map[descriptorpb.FieldDescriptorProto_Type]string{
	descriptorpb.FieldDescriptorProto_TYPE_BOOL:     FieldTypeBool,
	descriptorpb.FieldDescriptorProto_TYPE_INT32:    FieldTypeInt32,
	descriptorpb.FieldDescriptorProto_TYPE_UINT32:   FieldTypeUInt32,
	descriptorpb.FieldDescriptorProto_TYPE_INT64:    FieldTypeInt64,
	descriptorpb.FieldDescriptorProto_TYPE_UINT64:   FieldTypeUInt64,
	descriptorpb.FieldDescriptorProto_TYPE_FLOAT:    FieldTypeFloat,
	descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:   FieldTypeDouble,
	descriptorpb.FieldDescriptorProto_TYPE_STRING:   FieldTypeString,
	descriptorpb.FieldDescriptorProto_TYPE_BYTES:    FieldTypeBytes,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED64:  FieldTypeFixed64,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED32:  FieldTypeFixed32,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: FieldTypeSFixed32,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: FieldTypeSFixed64,
	descriptorpb.FieldDescriptorProto_TYPE_SINT32:   FieldTypeSInt32,
	descriptorpb.FieldDescriptorProto_TYPE_SINT64:   FieldTypeSInt64,
}

func protoType(t string, options ...string) descriptorpb.FieldDescriptorProto_Type {
	switch t {
	case core.UInt32TypeName, core.UInt32TypeFullName:
		if core.NewStringValues(options...).Contains(FixedAttributeName) {
			return descriptorpb.FieldDescriptorProto_TYPE_FIXED32
		}
	case core.UInt64TypeName, core.UInt64TypeFullName:
		if core.NewStringValues(options...).Contains(FixedAttributeName) {
			return descriptorpb.FieldDescriptorProto_TYPE_FIXED32
		}
	case core.Int32TypeName, core.Int32TypeFullName:
		values := core.NewStringValues(options...)
		if values.Contains(FixedAttributeName) {
			if values.Contains(SignedAttributeName) {
				return descriptorpb.FieldDescriptorProto_TYPE_SFIXED32
			} else {
				return descriptorpb.FieldDescriptorProto_TYPE_FIXED32
			}
		} else if values.Contains(SignedAttributeName) {
			return descriptorpb.FieldDescriptorProto_TYPE_SINT32
		}
	case core.Int64TypeName, core.Int64TypeFullName:
		values := core.NewStringValues(options...)
		if values.Contains(FixedAttributeName) {
			if values.Contains(SignedAttributeName) {
				return descriptorpb.FieldDescriptorProto_TYPE_SFIXED64
			} else {
				return descriptorpb.FieldDescriptorProto_TYPE_FIXED64
			}
		} else if values.Contains(SignedAttributeName) {
			return descriptorpb.FieldDescriptorProto_TYPE_SINT64
		}
	}

	switch t {
	case core.DoubleTypeName, core.Float64TypeName, core.DoubleTypeFullName, core.Float64TypeFullName:
		return descriptorpb.FieldDescriptorProto_TYPE_DOUBLE
	case core.FloatTypeName, core.Float32TypeName, core.FloatTypeFullName, core.Float32TypeFullName:
		return descriptorpb.FieldDescriptorProto_TYPE_FLOAT
	case core.Int64TypeName, core.IntTypeName, core.Int64TypeFullName, core.IntTypeFullName:
		return descriptorpb.FieldDescriptorProto_TYPE_INT64
	case core.UInt64TypeName, core.UIntTypeName, core.PositiveTypeName, core.NegativeTypeName, core.UInt64TypeFullName, core.UIntTypeFullName, core.PositiveTypeFullName, core.NegativeTypeFullName:
		return descriptorpb.FieldDescriptorProto_TYPE_UINT64
	case core.Int8TypeName, core.Int16TypeName, core.Int32TypeName, core.ByteTypeName, core.Int8TypeFullName, core.Int16TypeFullName, core.Int32TypeFullName, core.ByteTypeFullName:
		return descriptorpb.FieldDescriptorProto_TYPE_INT32
	case core.UInt8TypeName, core.UInt16TypeName, core.UInt32TypeName, core.UInt8TypeFullName, core.UInt16TypeFullName, core.UInt32TypeFullName:
		return descriptorpb.FieldDescriptorProto_TYPE_UINT32
	case core.BoolTypeName, core.BoolTypeFullName:
		return descriptorpb.FieldDescriptorProto_TYPE_BOOL
	case core.StringTypeName, core.StringTypeFullName:
		return descriptorpb.FieldDescriptorProto_TYPE_STRING
	case core.BytesTypeName, core.BytesTypeFullName:
		return descriptorpb.FieldDescriptorProto_TYPE_BYTES
	case "Enum":
		return descriptorpb.FieldDescriptorProto_TYPE_ENUM
	default:
		return descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
	}
}

func protoTypeName(typeName string, options ...string) string {
	switch typeName {
	case core.DoubleTypeName, core.Float64TypeName, core.FloatTypeName, core.Float32TypeName,
		core.DoubleTypeFullName, core.Float64TypeFullName, core.FloatTypeFullName, core.Float32TypeFullName,
		core.Int64TypeName, core.IntTypeName, core.NegativeTypeName, core.UInt64TypeName, core.UIntTypeName, core.PositiveTypeName,
		core.Int64TypeFullName, core.IntTypeFullName, core.NegativeTypeFullName, core.UInt64TypeFullName, core.UIntTypeFullName, core.PositiveTypeFullName,
		core.Int8TypeName, core.Int16TypeName, core.Int32TypeName, core.ByteTypeName,
		core.Int8TypeFullName, core.Int16TypeFullName, core.Int32TypeFullName, core.ByteTypeFullName,
		core.UInt8TypeName, core.UInt16TypeName, core.UInt32TypeName,
		core.UInt8TypeFullName, core.UInt16TypeFullName, core.UInt32TypeFullName,
		core.BoolTypeName, core.StringTypeName, core.BytesTypeName,
		core.BoolTypeFullName, core.StringTypeFullName, core.BytesTypeFullName:
		return fieldDescriptorProtoTypeName[protoType(typeName, options...)]
	default:
		return typeName
	}
}
