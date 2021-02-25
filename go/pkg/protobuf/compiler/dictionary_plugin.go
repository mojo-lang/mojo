package compiler

import (
	"errors"
	"fmt"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	desc "github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
)

type DictionaryPlugin struct {
}

func init() {
	p := plugins["mojo.core.Dictionary"]
	if p == nil {
		p = make([]Plugin, 0)
	}
	p = append(p, &DictionaryPlugin{})

	plugins["mojo.core.Dictionary"] = p
}

func (p *DictionaryPlugin) Compile(ctx *Context, t *lang.NominalType) (string, string, error) {
	if t.Name != "Dictionary" {
		return "", "", errors.New("")
	}

	if len(t.GenericArguments) != 2 {
		return "", "", errors.New("")
	}

	s := &lang.StructDecl{}
	fieldName := ""
	if decl, ok := ctx.GetCurrent().(*lang.ValueDecl); ok {
		fieldName = decl.Name
	}
	s.Name = strcase.ToCamel(fieldName) + "Entry"

	keyType := t.GenericArguments[0]
	keyType.Attributes = lang.SetIntegerAttribute(keyType.Attributes, "number", 1)

	valType := t.GenericArguments[1]
	valType.Attributes = lang.SetIntegerAttribute(valType.Attributes, "number", 2)

	if len(valType.GenericArguments) > 0 {
		ctx.SetOption("register_struct", true)
		_, name, err := CompileNominalType(ctx, valType)
		ctx.DeleteOption("register_struct")

		if err != nil {
			return "", "", err
		}
		valType.Name = name
	}

	s.Type = &lang.StructType{
		Fields: []*lang.ValueDecl{{
			Name: "key",
			Type: keyType,
		}, {
			Name: "value",
			Type: valType,
		}},
	}

	msgDescriptor := desc.NewMessageDescriptor(ctx.GetFileDescriptor())
	err := CompileStruct(ctx, s, msgDescriptor)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("failed to compile the map field entry in %s.%s: %s",
			"", //ctx.Message.Name,
			"", //ctx.FieldName,
			err.Error()))
	}

	mapEntry := true
	msgDescriptor.Options = &descriptor.MessageOptions{
		MapEntry: &mapEntry,
	}

	message := ctx.GetParentMessageDescriptor()
	if message != nil {
		message.AddInnerMessage(msgDescriptor)
	}

	return "struct", s.Name, nil
}
