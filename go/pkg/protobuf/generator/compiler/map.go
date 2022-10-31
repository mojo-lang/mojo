package compiler

import (
	"errors"
	"fmt"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

type Map struct {
}

func init() {
	RegisterNominalPlugin(&Map{})
}

func (p *Map) Name() string {
	return core.MapTypeFullName
}

func (p *Map) Compile(ctx context.Context, t *lang.NominalType) (string, string, error) {
	if t.GetFullName() != core.MapTypeFullName {
		return "", "", logs.NewErrorw("invalid Map type", "type", t.GetFullName())
	}

	if len(t.GenericArguments) != 2 {
		return "", "", logs.NewErrorw("invalid generic arguments size for Map type", "size", len(t.GenericArguments))
	}

	s := &lang.StructDecl{}
	fieldName := ""
	if decl := context.ValueDecl(ctx); decl != nil {
		fieldName = decl.Name
	}
	s.Name = strcase.ToCamel(fieldName) + "Entry"

	keyType := t.GenericArguments[0]
	keyType.Attributes = lang.SetIntegerAttribute(keyType.Attributes, core.NumberAttributeName, 1)
	keyLabel := "key"
	if label, _ := keyType.GetStringAttribute(core.LabelAttributeName); len(label) > 0 {
		keyLabel = label
	}

	valType := t.GenericArguments[1]
	valType.Attributes = lang.SetIntegerAttribute(valType.Attributes, core.NumberAttributeName, 2)
	valLabel := "value"
	if label, _ := valType.GetStringAttribute(core.LabelAttributeName); len(label) > 0 {
		valLabel = label
	}

	if len(valType.GenericArguments) > 0 {
		structCtx := context.WithValues(ctx, "register_struct", true)
		_, name, err := Nominal{}.Compile(structCtx, valType)
		if err != nil {
			return "", "", err
		}
		valType.Name = name
	}

	s.Type = &lang.StructType{
		Fields: []*lang.ValueDecl{{
			Name: keyLabel,
			Type: keyType,
		}, {
			Name: valLabel,
			Type: valType,
		}},
	}

	msgDescriptor := descriptor.NewMessage(context.FileDescriptor(ctx))
	err := Struct{}.Compile(ctx, s, msgDescriptor)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("failed to compile the map field entry in %s.%s: %s",
			"", // ctx.Message.Name,
			"", // ctx.FieldName,
			err.Error()))
	}

	msgDescriptor.SetMapEntry(true)
	message := context.MessageDescriptor(ctx)
	if message != nil {
		message.AppendMessage(msgDescriptor)
	}

	return "struct", s.Name, nil
}
