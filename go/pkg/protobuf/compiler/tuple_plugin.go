package compiler

import (
	"errors"
	"fmt"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type TuplePlugin struct {
}

func init() {
	p := plugins["mojo.core.Tuple"]
	if p == nil {
		p = make([]Plugin, 0)
	}
	p = append(p, &TuplePlugin{})

	plugins["mojo.core.Tuple"] = p
}

func (p *TuplePlugin) Compile(ctx context.Context, t *lang.NominalType) (string, string, error) {
	if t.Name != "Tuple" {
		return "", "", errors.New(fmt.Sprintf("type invalid, need Tuple, but is %s", t.Name))
	}

	if len(t.GenericArguments) == 1 {
		argument := t.GenericArguments[0]
		_, err := lang.GetStringAttribute(argument.Attributes, "label")
		if err != nil { // has no label
			argument.Attributes = append(argument.Attributes, t.Attributes...)
			return CompileNominalType(ctx, t.GenericArguments[0])
		}
	}

	// generate the message
	s := &lang.StructDecl{}
	alias, err := lang.GetStringAttribute(t.Attributes, "alias")
	if err == nil {
		s.Name = alias
	}
	s.Type = &lang.StructType{}
	for _, argument := range t.GenericArguments {
		label, err := lang.GetStringAttribute(argument.Attributes, "label")
		if err != nil {
			return "", "", errors.New(fmt.Sprintf("failed to get the label attribute from %s", argument.Name))
		}

		s.Type.Fields = append(s.Type.Fields, &lang.ValueDecl{
			Name: label,
			Type: argument,
		})
	}

	file := context.FileDescriptor(ctx)
	descriptor := descriptor.NewMessageDescriptor(file)
	err = CompileStruct(ctx, s, descriptor)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("failed to compile struct: %s", err.Error()))
	}

	file.Messages = append(file.Messages, descriptor)
	return "struct", s.Name, nil
}
