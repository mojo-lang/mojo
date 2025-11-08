package descriptor

import (
    "strings"

    "google.golang.org/protobuf/types/descriptorpb"
)

// Message represents a protocol buffer message.
type Message struct {
    Descriptor
    Proto *descriptorpb.DescriptorProto

    FullName string

    Parent *Message // The containing message, if any.

    Enums    []*Enum    // Inner Enums, if any.
    Messages []*Message // Inner Messages, if any.

    Fields []*Field // message field declarations
    Oneofs []*Oneof // message oneof declarations
}

func NewMessage(file *File) *Message {
    return &Message{
        Descriptor: Descriptor{
            File: file,
        },
        Proto: &descriptorpb.DescriptorProto{},
    }
}

func NewMessageFrom(file *File, proto *descriptorpb.DescriptorProto) *Message {
    message := &Message{
        Descriptor: Descriptor{
            File: file,
        },
        Proto: proto,
    }
    //
    //var loc Location
    //if parent != nil {
    //    loc = parent.Location.AppendPath(DescriptorProto_NestedType_field_number, desc.Index())
    //} else {
    //    loc = f.location.AppendPath(FileDescriptorProto_MessageType_field_number, desc.Index())
    //}

    for _, enum := range proto.EnumType {
        message.AppendInnerEnum(NewEnumFrom(file, enum))
    }
    for _, msg := range proto.NestedType {
        message.AppendMessage(NewMessageFrom(file, msg))
    }
    for _, field := range proto.Field {
        message.AppendField(NewFieldFrom(message, field))
    }
    for _, oneof := range proto.OneofDecl {
        message.AppendOneof(NewOneofFrom(message, oneof))
    }

    // Resolve local references between fields and oneofs.
    for _, field := range message.Fields {
        if index := field.Proto.OneofIndex; index != nil {
            message.Oneofs[*index].AppendField(field)
        }
    }

    return message
}

func (m *Message) proto() *descriptorpb.DescriptorProto {
    if m != nil {
        return m.Proto
    }
    return nil
}

func (m *Message) HasMessage() bool {
    return m != nil && len(m.Messages) > 0
}

func (m *Message) GetMessage(name string) *Message {
    for _, msg := range m.Messages {
        if msg.GetName() == name {
            return msg
        }
    }
    return nil
}

func (m *Message) IsMessageExist(name string) bool {
    return m.GetMessage(name) != nil
}

func (m *Message) HasEnum() bool {
    return m != nil && len(m.Enums) > 0
}

func (m *Message) GetEnum(name string) *Enum {
    for _, e := range m.Enums {
        if e.GetName() == name {
            return e
        }
    }
    return nil
}

func (m *Message) IsEnumExist(name string) bool {
    return m.GetEnum(name) != nil
}

func (m *Message) HasField() bool {
    return m != nil && len(m.Fields) > 0
}

func (m *Message) GetField(name string) *Field {
    for _, f := range m.Fields {
        if f.GetName() == name {
            return f
        }
    }
    return nil
}

func (m *Message) IsFieldExist(name string) bool {
    return m.GetField(name) != nil
}

func (m *Message) HasOneof() bool {
    return m != nil && len(m.Fields) > 0
}

func (m *Message) GetOneof(name string) *Oneof {
    for _, o := range m.Oneofs {
        if o.GetName() == name {
            return o
        }
    }
    return nil
}

func (m *Message) IsOneofExist(name string) bool {
    return m.GetOneof(name) != nil
}

func (m *Message) GetName() string {
    return m.proto().GetName()
}

func (m *Message) GetFullName() string {
    if m != nil {
        return m.FullName
    }
    return ""
}

func (m *Message) GetPackageName() string {
    if m != nil && m.File != nil {
        return m.File.GetPackageName()
    }
    return ""
}

func (m *Message) GetOneofs() []*Oneof {
    if m != nil {
        return m.Oneofs
    }
    return nil
}

func (m *Message) GetLastOneof() *Oneof {
    if m != nil && len(m.Oneofs) > 0 {
        return m.Oneofs[len(m.Oneofs)-1]
    }
    return nil
}

func (m *Message) AppendMessage(msg *Message) *Message {
    if m != nil && m.Proto != nil {
        msg.Parent = m
        m.Messages = append(m.Messages, msg)
        m.Proto.NestedType = append(m.Proto.NestedType, msg.Proto)
    }
    return m
}

func (m *Message) AppendInnerEnum(enum *Enum) *Message {
    if m != nil && m.Proto != nil {
        enum.Parent = m
        m.Enums = append(m.Enums, enum)
        m.Proto.EnumType = append(m.Proto.EnumType, enum.Proto)
    }
    return m
}

func (m *Message) AppendField(field *Field) *Message {
    if m != nil && m.Proto != nil {
        m.Fields = append(m.Fields, field)
        m.Proto.Field = append(m.Proto.Field, field.Proto)
    }
    return m
}

func (m *Message) AppendOneof(oneof *Oneof) *Message {
    if m != nil && m.Proto != nil {
        m.Oneofs = append(m.Oneofs, oneof)
        m.Proto.OneofDecl = append(m.Proto.OneofDecl, oneof.Proto)
    }
    return m
}

func (m *Message) AppendOneofWith(name string) *Message {
    if m != nil && m.Proto != nil {
        m.AppendOneof(NewOneof(m, name))
    }
    return m
}

func concatFullName(pkg string, name string) string {
    if len(name) > 0 {
        if len(pkg) > 0 {
            return strings.Join([]string{pkg, name}, ".")
        }
        return name
    }
    return ""
}

func (m *Message) SetName(name string) *Message {
    if m != nil && m.Proto != nil {
        m.Proto.Name = &name
        m.FullName = concatFullName(m.File.GetPackageName(), name)
    }
    return m
}

func (m *Message) IsDeprecated() bool {
    return m.proto().GetOptions().GetDeprecated()
}

func (m *Message) IsMapEntry() bool {
    return m.proto().GetOptions().GetMapEntry()
}

func (m *Message) SetDeprecated(value bool) *Message {
    if m != nil && m.Proto != nil {
        if m.Proto.Options == nil {
            m.Proto.Options = &descriptorpb.MessageOptions{}
        }
        m.Proto.Options.Deprecated = &value
    }
    return m
}

func (m *Message) SetMapEntry(value bool) *Message {
    if m != nil && m.Proto != nil {
        if m.Proto.Options == nil {
            m.Proto.Options = &descriptorpb.MessageOptions{}
        }
        m.Proto.Options.MapEntry = &value
    }
    return m
}
