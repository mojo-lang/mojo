package context

import (
    "context"
    "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

const DescriptorKey = "@descriptor"
const FileDescriptorKey = "@descriptor/File"
const MessageDescriptorKey = "@descriptor/Message"
const ServiceDescriptorKey = "@descriptor/Service"
const EnumDescriptorKey = "@descriptor/Enum"
const FieldDescriptorKey = "@descriptor/Field"

func WithDescriptor(ctx context.Context, value interface{}) context.Context {
    key := ""
    switch value.(type) {
    case *descriptor.FileDescriptor:
        key = FileDescriptorKey
    case *descriptor.MessageDescriptor:
        key = MessageDescriptorKey
    case *descriptor.FieldDescriptorProto:
        key = FieldDescriptorKey
    case *descriptor.EnumDescriptor:
        key = EnumDescriptorKey
    case *descriptor.ServiceDescriptor:
        key = ServiceDescriptorKey
    }
    return WithValues(ctx, DescriptorKey, value, key, value)
}

func FileDescriptor(ctx context.Context) *descriptor.FileDescriptor {
    if file, ok := ctx.Value(FileDescriptorKey).(*descriptor.FileDescriptor); ok {
        return file
    }
    return nil
}

func MessageDescriptor(ctx context.Context) *descriptor.MessageDescriptor {
    if msg, ok := ctx.Value(MessageDescriptorKey).(*descriptor.MessageDescriptor); ok {
        return msg
    }
    return nil
}

func FieldDescriptor(ctx context.Context) *descriptor.FieldDescriptorProto {
    if field, ok := ctx.Value(FieldDescriptorKey).(*descriptor.FieldDescriptorProto); ok {
        return field
    }
    return nil
}

func EnumDescriptor(ctx context.Context) *descriptor.EnumDescriptor {
    if e, ok := ctx.Value(EnumDescriptorKey).(*descriptor.EnumDescriptor); ok {
        return e
    }
    return nil
}

func ServiceDescriptor(ctx context.Context) *descriptor.ServiceDescriptor {
    if service, ok := ctx.Value(ServiceDescriptorKey).(*descriptor.ServiceDescriptor); ok {
        return service
    }
    return nil
}

func DescriptorValue(ctx context.Context) interface{} {
    return ctx.Value(DescriptorKey)
}

func DescriptorValues(ctx context.Context) []interface{} {
    return Values(ctx, DescriptorKey)
}

func PreviousDescriptorValue(ctx context.Context, index int) interface{} {
    return PreviousValue(ctx, DescriptorKey, index)
}
