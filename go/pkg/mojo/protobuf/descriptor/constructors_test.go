package descriptor

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestDescriptorWrappingDoesNotDuplicateProtoChildren(t *testing.T) {
	enumProto := &descriptorpb.EnumDescriptorProto{Name: proto.String("Status"), Value: []*descriptorpb.EnumValueDescriptorProto{{Name: proto.String("UNKNOWN"), Number: proto.Int32(0)}}}
	messageProto := &descriptorpb.DescriptorProto{Name: proto.String("Item"),
		Field:      []*descriptorpb.FieldDescriptorProto{{Name: proto.String("name"), Number: proto.Int32(1), Type: descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(), OneofIndex: proto.Int32(0)}},
		NestedType: []*descriptorpb.DescriptorProto{{Name: proto.String("Nested")}},
		EnumType:   []*descriptorpb.EnumDescriptorProto{enumProto},
		OneofDecl:  []*descriptorpb.OneofDescriptorProto{{Name: proto.String("choice")}},
	}
	serviceProto := &descriptorpb.ServiceDescriptorProto{Name: proto.String("Store"), Method: []*descriptorpb.MethodDescriptorProto{{Name: proto.String("Get")}}}
	fileProto := &descriptorpb.FileDescriptorProto{Name: proto.String("sample.proto"), Package: proto.String("sample"),
		MessageType: []*descriptorpb.DescriptorProto{messageProto}, EnumType: []*descriptorpb.EnumDescriptorProto{proto.Clone(enumProto).(*descriptorpb.EnumDescriptorProto)},
		Service: []*descriptorpb.ServiceDescriptorProto{serviceProto},
	}
	before := proto.Clone(fileProto)
	for i := 0; i < 2; i++ {
		file := NewFileFrom(fileProto)
		require.True(t, proto.Equal(before, fileProto), "wrapping must not change the supplied proto")
		require.Len(t, file.Messages, 1)
		message := file.Messages[0]
		require.Len(t, message.Fields, 1)
		require.Len(t, message.Messages, 1)
		require.Len(t, message.Enums, 1)
		require.Len(t, message.Oneofs, 1)
		require.Same(t, message, message.Messages[0].Parent)
		require.Same(t, message, message.Enums[0].Parent)
		require.Same(t, message.Oneofs[0], message.Fields[0].Oneof)
		require.Same(t, message.Fields[0], message.Oneofs[0].Fields[0])
		require.Len(t, file.Enums[0].Values, 1)
		require.True(t, file.Enums[0].IsValueExist("UNKNOWN"))
		require.Nil(t, file.Enums[0].GetValue("missing"))
		require.Len(t, file.Services[0].Methods, 1)
		require.True(t, file.Services[0].IsMethodExist("Get"))
		require.Nil(t, file.Services[0].GetMethod("missing"))
	}
}

func TestPackagesIndexFilterAndMethodTypes(t *testing.T) {
	packages := NewPackages()
	file := NewFileWithName("sample.proto", "sample")
	input := NewMessage(file).SetName("Request")
	output := NewMessage(file).SetName("Response")
	service := NewService(file).SetName("Store")
	enum := NewEnum(file).SetName("Status").AppendValueWith("UNKNOWN", 0)
	file.Messages = []*Message{input, output}
	file.Services = []*Service{service}
	file.Enums = []*Enum{enum}
	file.Packages = packages
	require.Same(t, packages, packages.AddFile(file).AddFile(file).AddFile(nil))
	require.Len(t, packages.FilesByPath, 1)
	require.Same(t, input, packages.GetMessage("sample.Request"))
	require.Same(t, service, packages.ServicesByName["sample.Store"])
	child := NewFileWithName("child.proto", "sample.child")
	sibling := NewFileWithName("other.proto", "sampleOther")
	packages.AddFile(child).AddFile(sibling)
	require.ElementsMatch(t, []*File{file, child}, packages.Filter("sample", false))
	require.Equal(t, []*File{file}, packages.Filter("sample", true))
	method := NewMethod(service).SetName("Get").SetInput(input).SetOutput(output)
	service.AppendMethod(method)
	require.Same(t, input, method.GetInput())
	require.Same(t, output, method.GetOutput())
	wrapped := NewMethodFrom(service, method.Proto)
	require.Same(t, input, wrapped.GetInput())
	require.Same(t, output, wrapped.GetOutput())
	wrapped.Input = nil
	wrapped.Output = nil
	require.Same(t, input, wrapped.GetInput())
	require.Same(t, output, wrapped.GetOutput())
	require.Nil(t, packages.GetMessage("missing"))
}
