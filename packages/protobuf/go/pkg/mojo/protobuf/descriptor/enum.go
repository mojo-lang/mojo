package descriptor

import (
    "google.golang.org/protobuf/types/descriptorpb"
    "strings"
)

// Enum describes an enum. If it's at top level, its Parent will be nil.
// Otherwise, it will be the descriptor of the message in which it is defined.
type Enum struct {
    Descriptor
    Proto *descriptorpb.EnumDescriptorProto

    FullName string

    Parent *Message     // The containing message, if any.
    Values []*EnumValue // enum value declarations
}

// NewEnum Construct the Enum
func NewEnum(file *File) *Enum {
    enum := &Enum{
        Descriptor: Descriptor{
            File: file,
        },
        Proto: &descriptorpb.EnumDescriptorProto{},
    }

    return enum
}

func NewEnumFrom(file *File, proto *descriptorpb.EnumDescriptorProto) *Enum {
    enum := &Enum{
        Descriptor: Descriptor{
            File: file,
        },
        Proto: proto,
    }

    for _, value := range proto.Value {
        enum.AppendValue(NewEnumValueFrom(enum, value))
    }

    return enum
}

func (m *Enum) proto() *descriptorpb.EnumDescriptorProto {
    if m != nil {
        return m.Proto
    }
    return nil
}

func (m *Enum) AppendValueWith(name string, number int32) *Enum {
    if m != nil && m.Proto != nil {
        m.AppendValue(NewEnumValue(m, name, number))
    }
    return m
}

func (m *Enum) AppendValue(value *EnumValue) *Enum {
    if m != nil && m.Proto != nil {
        m.Values = append(m.Values, value)
        m.Proto.Value = append(m.Proto.Value, value.Proto)
    }
    return m
}

func (m *Enum) IsDeprecated() bool {
    return m.proto().GetOptions().GetDeprecated()
}

func (m *Enum) GetName() string {
    return m.proto().GetName()
}

func (m *Enum) GetFullName() string {
    if m != nil {
        if len(m.FullName) == 0 {
            m.FullName = strings.Join([]string{m.GetPackageName(), m.GetName()}, ".")
        }
        return m.FullName
    }
    return ""
}

func (m *Enum) GetPackageName() string {
    if m != nil && m.File != nil {
        return m.File.GetPackageName()
    }
    return ""
}

func (m *Enum) SetName(name string) *Enum {
    if m != nil && m.Proto != nil {
        m.Proto.Name = &name
        m.FullName = concatFullName(m.File.GetPackageName(), name)
    }
    return m
}

func (m *Enum) HasValue() bool {
    return m != nil && len(m.Values) > 0
}

func (m *Enum) GetValue(name string) *EnumValue {
    for _, v := range m.Values {
        if v.GetName() == name {
            return v
        }
    }
    return nil
}

func (m *Enum) IsValueExist(name string) bool {
    return m.GetValue(name) != nil
}
