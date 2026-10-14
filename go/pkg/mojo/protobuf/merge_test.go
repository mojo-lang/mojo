package protobuf

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func mergeTestDescriptor(t *testing.T) protoreflect.MessageDescriptor {
	t.Helper()
	field := func(name string, number int32, kind descriptorpb.FieldDescriptorProto_Type, typeName string, repeated bool) *descriptorpb.FieldDescriptorProto {
		f := &descriptorpb.FieldDescriptorProto{Name: proto.String(name), Number: proto.Int32(number), Type: kind.Enum()}
		if typeName != "" {
			f.TypeName = proto.String(typeName)
		}
		if repeated {
			f.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
		}
		return f
	}
	child := &descriptorpb.DescriptorProto{Name: proto.String("Child"), Field: []*descriptorpb.FieldDescriptorProto{
		field("name", 1, descriptorpb.FieldDescriptorProto_TYPE_STRING, "", false),
		field("blob", 2, descriptorpb.FieldDescriptorProto_TYPE_BYTES, "", false),
	}}
	record := &descriptorpb.DescriptorProto{Name: proto.String("Record"), Field: []*descriptorpb.FieldDescriptorProto{
		field("name", 1, descriptorpb.FieldDescriptorProto_TYPE_STRING, "", false),
		field("blob", 2, descriptorpb.FieldDescriptorProto_TYPE_BYTES, "", false),
		field("child", 3, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, ".merge.test.Child", false),
		field("names", 4, descriptorpb.FieldDescriptorProto_TYPE_STRING, "", true),
		field("blobs", 5, descriptorpb.FieldDescriptorProto_TYPE_BYTES, "", true),
		field("children", 6, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, ".merge.test.Child", true),
	}}
	for i, entry := range []struct {
		name, field string
		kind        descriptorpb.FieldDescriptorProto_Type
		typeName    string
	}{
		{"NamesMapEntry", "names_map", descriptorpb.FieldDescriptorProto_TYPE_STRING, ""},
		{"BlobsMapEntry", "blobs_map", descriptorpb.FieldDescriptorProto_TYPE_BYTES, ""},
		{"ChildrenMapEntry", "children_map", descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, ".merge.test.Child"},
	} {
		record.NestedType = append(record.NestedType, &descriptorpb.DescriptorProto{Name: proto.String(entry.name), Options: &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)}, Field: []*descriptorpb.FieldDescriptorProto{
			field("key", 1, descriptorpb.FieldDescriptorProto_TYPE_STRING, "", false),
			field("value", 2, entry.kind, entry.typeName, false),
		}})
		record.Field = append(record.Field, field(entry.field, int32(i+7), descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, ".merge.test.Record."+entry.name, true))
	}
	file, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{Name: proto.String("merge_test.proto"), Package: proto.String("merge.test"), Syntax: proto.String("proto3"), MessageType: []*descriptorpb.DescriptorProto{child, record}}, nil)
	require.NoError(t, err)
	return file.Messages().ByName("Record")
}

func TestMergeReflectiveMatchesProtobufAndCopiesMutableValues(t *testing.T) {
	desc := mergeTestDescriptor(t)
	dst, src := dynamicpb.NewMessage(desc), dynamicpb.NewMessage(desc)
	require.NoError(t, protojson.Unmarshal([]byte("{\"name\":\"old\",\"child\":{\"name\":\"old\",\"blob\":\"AQ==\"},\"names\":[\"old\"],\"blobs\":[\"AQ==\"],\"namesMap\":{\"keep\":\"yes\",\"replace\":\"old\"},\"childrenMap\":{\"replace\":{\"name\":\"old\",\"blob\":\"AQ==\"}}}"), dst))
	require.NoError(t, protojson.Unmarshal([]byte("{\"name\":\"new\",\"blob\":\"Ag==\",\"child\":{\"name\":\"new\"},\"names\":[\"new\"],\"blobs\":[\"Ag==\"],\"children\":[{\"blob\":\"Ag==\"}],\"namesMap\":{\"replace\":\"new\"},\"blobsMap\":{\"key\":\"Ag==\"},\"childrenMap\":{\"replace\":{\"name\":\"new\",\"blob\":\"Ag==\"}}}"), src))
	dst.SetUnknown([]byte{0xa0, 0x06, 0x01})
	src.SetUnknown([]byte{0xa0, 0x06, 0x02})
	expected := proto.Clone(dst)
	proto.Merge(expected, src)
	Merge(dst, src)
	require.True(t, proto.Equal(expected, dst), "merge result must match the documented protobuf merge contract")
	require.Equal(t, protoreflect.RawFields{0xa0, 0x06, 0x01, 0xa0, 0x06, 0x02}, dst.GetUnknown())
	fields := desc.Fields()
	src.Get(fields.ByName("blob")).Bytes()[0] = 9
	src.Get(fields.ByName("blobs")).List().Get(0).Bytes()[0] = 9
	child := src.Get(fields.ByName("children")).List().Get(0).Message()
	child.Get(child.Descriptor().Fields().ByName("blob")).Bytes()[0] = 9
	src.Get(fields.ByName("blobs_map")).Map().Get(protoreflect.ValueOfString("key").MapKey()).Bytes()[0] = 9
	child = src.Get(fields.ByName("children_map")).Map().Get(protoreflect.ValueOfString("replace").MapKey()).Message()
	child.Get(child.Descriptor().Fields().ByName("blob")).Bytes()[0] = 9
	src.GetUnknown()[2] = 9
	require.True(t, proto.Equal(expected, dst), "destination must not alias mutable source data")
}

func TestMergeGeneratedMessagesAndDescriptorMismatch(t *testing.T) {
	dst := &structpb.ListValue{Values: []*structpb.Value{structpb.NewStringValue("old")}}
	src := &structpb.ListValue{Values: []*structpb.Value{structpb.NewStringValue("new")}}
	Merge(dst, src)
	require.Len(t, dst.Values, 2)
	src.Values[0].Kind = &structpb.Value_StringValue{StringValue: "changed"}
	require.Equal(t, "new", dst.Values[1].GetStringValue())
	require.Panics(t, func() { Merge(wrapperspb.String("a"), wrapperspb.Int32(1)) })
	desc := mergeTestDescriptor(t)
	empty := dynamicpb.NewMessage(desc)
	Merge(empty, dynamicpb.NewMessage(desc))
	require.False(t, empty.Has(desc.Fields().ByName("name")))
	var invalid *dynamicpb.Message
	require.Panics(t, func() { Merge(invalid, empty) })
}
