package descriptor

import "google.golang.org/protobuf/types/descriptorpb"

// An Oneof describes a message oneof.
type Oneof struct {
    Descriptor
    Proto *descriptorpb.OneofDescriptorProto

    Parent *Message // message in which this oneof is declared

    Fields []*Field // fields that are part of this oneof
}

func NewOneof(parent *Message, name string) *Oneof {
    return &Oneof{
        Descriptor: Descriptor{
            File: parent.File,
        },
        Proto: &descriptorpb.OneofDescriptorProto{
            Name: &name,
        },
        Parent: parent,
    }
}

func NewOneofFrom(parent *Message, proto *descriptorpb.OneofDescriptorProto) *Oneof {
    return &Oneof{
        Descriptor: Descriptor{
            File: parent.File,
        },
        Proto:  proto,
        Parent: parent,
    }
}

func (o *Oneof) proto() *descriptorpb.OneofDescriptorProto {
    if o != nil {
        return o.Proto
    }
    return nil
}

func (o *Oneof) GetName() string {
    return o.proto().GetName()
}

func (o *Oneof) AppendField(field *Field) *Oneof {
    if o != nil {
        field.Oneof = o
        o.Fields = append(o.Fields, field)
    }
    return o
}
