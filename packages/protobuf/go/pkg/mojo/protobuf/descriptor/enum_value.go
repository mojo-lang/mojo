package descriptor

import "google.golang.org/protobuf/types/descriptorpb"

type EnumValue struct {
    Descriptor
    Proto *descriptorpb.EnumValueDescriptorProto

    Parent *Enum // enum in which this value is declared
}

func NewEnumValue(enum *Enum, name string, number int32) *EnumValue {
    return &EnumValue{
        Descriptor: Descriptor{
            File: enum.File,
        },
        Proto: &descriptorpb.EnumValueDescriptorProto{
            Name:   &name,
            Number: &number,
        },
        Parent: enum,
    }
}

func NewEnumValueFrom(enum *Enum, proto *descriptorpb.EnumValueDescriptorProto) *EnumValue {
    return &EnumValue{
        Descriptor: Descriptor{
            File: enum.File,
        },
        Proto:  proto,
        Parent: enum,
    }
}

func (m *EnumValue) proto() *descriptorpb.EnumValueDescriptorProto {
    if m != nil {
        return m.Proto
    }
    return nil
}

func (m *EnumValue) IsDeprecated() bool {
    return m.proto().Options.GetDeprecated()
}

func (m *EnumValue) GetName() string {
    return m.proto().GetName()
}

func (m *EnumValue) GetNumber() int32 {
    return m.proto().GetNumber()
}
