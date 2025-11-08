package mojo

import (
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GetOptionFullName(option protoreflect.ExtensionType) string {
	return string(option.TypeDescriptor().FullName())
}

var fileOptions []protoreflect.ExtensionType
var fileOptionsOnce sync.Once

func FileOptionsExtensions() []protoreflect.ExtensionType {
	fileOptionsOnce.Do(func() {
		for i := 0; i < len(file_mojo_mojo_proto_extTypes); i++ {
			if _, ok := file_mojo_mojo_proto_extTypes[i].ExtendedType.(*descriptorpb.FileOptions); ok {
				fileOptions = append(fileOptions, &file_mojo_mojo_proto_extTypes[i])
			}
		}
	})
	return fileOptions
}

var enumOptions []protoreflect.ExtensionType
var enumOptionsOnce sync.Once

func EnumOptionsExtensions() []protoreflect.ExtensionType {
	enumOptionsOnce.Do(func() {
		for i := 0; i < len(file_mojo_mojo_proto_extTypes); i++ {
			if _, ok := file_mojo_mojo_proto_extTypes[i].ExtendedType.(*descriptorpb.EnumOptions); ok {
				enumOptions = append(enumOptions, &file_mojo_mojo_proto_extTypes[i])
			}
		}
	})
	return enumOptions
}

var enumValueOptions []protoreflect.ExtensionType
var enumValueOptionsOnce sync.Once

func EnumValueOptionsExtensions() []protoreflect.ExtensionType {
	enumValueOptionsOnce.Do(func() {
		for i := 0; i < len(file_mojo_mojo_proto_extTypes); i++ {
			if _, ok := file_mojo_mojo_proto_extTypes[i].ExtendedType.(*descriptorpb.EnumOptions); ok {
				enumValueOptions = append(enumValueOptions, &file_mojo_mojo_proto_extTypes[i])
			}
		}
	})
	return enumValueOptions
}

var messageOptions []protoreflect.ExtensionType
var messageOptionsOnce sync.Once

func MessageOptionsExtensions() []protoreflect.ExtensionType {
	messageOptionsOnce.Do(func() {
		for i := 0; i < len(file_mojo_mojo_proto_extTypes); i++ {
			if _, ok := file_mojo_mojo_proto_extTypes[i].ExtendedType.(*descriptorpb.MessageOptions); ok {
				messageOptions = append(messageOptions, &file_mojo_mojo_proto_extTypes[i])
			}
		}
	})
	return messageOptions
}

var fieldOptions []protoreflect.ExtensionType
var fieldOptionsOnce sync.Once

func FieldOptionsExtensions() []protoreflect.ExtensionType {
	fieldOptionsOnce.Do(func() {
		for i := 0; i < len(file_mojo_mojo_proto_extTypes); i++ {
			if _, ok := file_mojo_mojo_proto_extTypes[i].ExtendedType.(*descriptorpb.FieldOptions); ok {
				fieldOptions = append(fieldOptions, &file_mojo_mojo_proto_extTypes[i])
			}
		}
	})
	return fieldOptions
}
