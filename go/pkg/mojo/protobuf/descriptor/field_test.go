package descriptor

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/mojo"
	"github.com/stretchr/testify/assert"
)

func newField() *Field {
	file := NewFile()
	message := NewMessage(file)
	return NewField(message, "foo")
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
