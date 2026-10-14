package descriptor

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/descriptorpb"
)

func newField() *Field {
	file := NewFile()
	message := NewMessage(file)
	return NewField(message, "foo")
}

func TestFieldNumericEncodingOptions(t *testing.T) {
	for _, tt := range []struct {
		name     string
		options  []string
		kind     descriptorpb.FieldDescriptorProto_Type
		typeName string
	}{
		{"Int32", nil, descriptorpb.FieldDescriptorProto_TYPE_INT32, "int32"},
		{"Int64", []string{SignedAttributeName}, descriptorpb.FieldDescriptorProto_TYPE_SINT64, "sint64"},
		{"Int32", []string{FixedAttributeName}, descriptorpb.FieldDescriptorProto_TYPE_FIXED32, "fixed32"},
		{"Int64", []string{FixedAttributeName, SignedAttributeName}, descriptorpb.FieldDescriptorProto_TYPE_SFIXED64, "sfixed64"},
		{"Int32", []string{FixedAttributeName, SignedAttributeName}, descriptorpb.FieldDescriptorProto_TYPE_SFIXED32, "sfixed32"},
		{"UInt64", []string{FixedAttributeName}, descriptorpb.FieldDescriptorProto_TYPE_FIXED64, "fixed64"},
		{"UInt32", []string{FixedAttributeName}, descriptorpb.FieldDescriptorProto_TYPE_FIXED32, "fixed32"},
	} {
		t.Run(tt.typeName, func(t *testing.T) {
			field := newField().SetType(tt.name, tt.options...).SetTypeName(tt.name, tt.options...)
			require.Equal(t, tt.kind, field.Proto.GetType())
			require.Equal(t, tt.typeName, field.GetTypeName())
			require.Equal(t, tt.typeName, field.Proto.GetTypeName())
		})
	}
}

func TestNewEnumField(t *testing.T) {
	file := NewFile()
	message := NewMessage(file)
	enum := NewEnum(file)
	field := NewEnumField(message, "foo", enum)
	if assert.NotNil(t, field) {
		assert.Equal(t, "foo", field.GetName())
	}
}

func TestNewMessageField(t *testing.T) {
	file := NewFile()
	message := NewMessage(file)
	fm := NewMessage(file)
	field := NewMessageField(message, "foo", fm)
	if assert.NotNil(t, field) {
		assert.Equal(t, "foo", field.GetName())
	}
}

func TestField_SetBoolOption(t *testing.T) {
	field := newField()

	field.SetOption(mojo.E_Alias, "alias")
	assert.Equal(t, "alias", field.GetStringOption(mojo.E_Alias))

	field.SetBoolOption(mojo.E_DbIgnore, true)
	assert.True(t, field.GetBoolOption(mojo.E_DbIgnore))
}

func TestField_SetNumber(t *testing.T) {
	field := newField()
	field.SetNumber(14)
	assert.Equal(t, int32(14), field.GetNumber())
}
