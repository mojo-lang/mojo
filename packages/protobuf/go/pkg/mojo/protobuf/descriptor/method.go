package descriptor

import "google.golang.org/protobuf/types/descriptorpb"

// A Method describes a method in a service.
type Method struct {
    Descriptor
    Proto *descriptorpb.MethodDescriptorProto

    Parent *Service // service in which this method is declared

    Input  *Message
    Output *Message
}

func NewMethod(parent *Service) *Method {
    return &Method{
        Descriptor: Descriptor{
            File: parent.File,
        },
        Proto:  &descriptorpb.MethodDescriptorProto{},
        Parent: parent,
        Input:  nil,
        Output: nil,
    }
}

func NewMethodFrom(parent *Service, proto *descriptorpb.MethodDescriptorProto) *Method {
    method := &Method{
        Descriptor: Descriptor{
            File: parent.File,
        },
        Proto:  proto,
        Parent: parent,
        Input:  nil,
        Output: nil,
    }

    if packages := method.File.GetPackages(); packages != nil {
        if proto.InputType != nil {
            method.Input = packages.GetMessage(*proto.InputType)
        }
        if proto.OutputType != nil {
            method.Output = packages.GetMessage(*proto.OutputType)
        }
    }

    return method
}

func (m *Method) proto() *descriptorpb.MethodDescriptorProto {
    if m != nil {
        return m.Proto
    }
    return nil
}

func (m *Method) GetName() string {
    return m.proto().GetName()
}

func (m *Method) SetName(name string) *Method {
    if m != nil && m.Proto != nil {
        m.Proto.Name = &name
    }
    return m
}

func (m *Method) GetInput() *Message {
    if m != nil {
        if m.Input == nil {
            if packages := m.File.GetPackages(); packages != nil {
                m.Input = packages.GetMessage(m.Proto.GetInputType())
            }
        }
        return m.Input
    }
    return nil
}

func (m *Method) GetOutput() *Message {
    if m != nil {
        if m.Output == nil {
            if packages := m.File.GetPackages(); packages != nil {
                m.Output = packages.GetMessage(m.Proto.GetOutputType())
            }
        }
        return m.Output
    }
    return nil
}

func (m *Method) SetInput(input *Message) *Method {
    if m != nil && m.Proto != nil {
        fullName := input.GetFullName()
        m.Input = input
        m.Proto.InputType = &fullName
    }
    return m
}

func (m *Method) SetOutput(output *Message) *Method {
    if m != nil && m.Proto != nil {
        fullName := output.GetFullName()
        m.Output = output
        m.Proto.OutputType = &fullName
    }
    return m
}
