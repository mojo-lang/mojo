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
	case *descriptor.File:
		key = FileDescriptorKey
	case *descriptor.Message:
		key = MessageDescriptorKey
	case *descriptor.Field:
		key = FieldDescriptorKey
	case *descriptor.Enum:
		key = EnumDescriptorKey
	case *descriptor.Service:
		key = ServiceDescriptorKey
	}
	return WithValues(ctx, DescriptorKey, value, key, value)
}

func FileDescriptor(ctx context.Context) *descriptor.File {
	if file, ok := ctx.Value(FileDescriptorKey).(*descriptor.File); ok {
		return file
	}
	return nil
}

func MessageDescriptor(ctx context.Context) *descriptor.Message {
	if msg, ok := ctx.Value(MessageDescriptorKey).(*descriptor.Message); ok {
		return msg
	}
	return nil
}

func FieldDescriptor(ctx context.Context) *descriptor.Field {
	if field, ok := ctx.Value(FieldDescriptorKey).(*descriptor.Field); ok {
		return field
	}
	return nil
}

func EnumDescriptor(ctx context.Context) *descriptor.Enum {
	if e, ok := ctx.Value(EnumDescriptorKey).(*descriptor.Enum); ok {
		return e
	}
	return nil
}

func ServiceDescriptor(ctx context.Context) *descriptor.Service {
	if service, ok := ctx.Value(ServiceDescriptorKey).(*descriptor.Service); ok {
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
