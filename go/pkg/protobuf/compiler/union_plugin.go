package compiler

import (
	"errors"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	desc "github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
)

type UnionPlugin struct {
	pluralize *pluralize.Client
}

func init() {
	p := plugins["mojo.core.Union"]
	if p == nil {
		p = make([]Plugin, 0)
	}
	p = append(p, &UnionPlugin{
		pluralize: pluralize.NewClient(),
	})

	plugins["mojo.core.Union"] = p
}

func (p *UnionPlugin) Compile(ctx *Context, t *lang.NominalType) (string, string, error) {
	s := ConstructBoxedUnion(t, p.pluralize)
	if s == nil {
		return "", "", nil
	} else {
		file := ctx.GetFileDescriptor()
		msgDescriptor := desc.NewMessageDescriptor(file)
		err := CompileStruct(ctx, s, msgDescriptor)
		if err != nil {
			return "", "", errors.New(fmt.Sprintf("failed to compile struct: %s", err.Error()))
		}

		oneOfName := ToSnake(*msgDescriptor.Name)
		msgDescriptor.OneofDecl = append(msgDescriptor.OneofDecl, &descriptor.OneofDescriptorProto{
			Name: &oneOfName,
		})
		index := int32(0)
		for _, field := range msgDescriptor.Field {
			field.OneofIndex = &index
		}

		return "struct", s.Name, nil
	}
}

func ConstructBoxedUnion(t *lang.NominalType, p *pluralize.Client) *lang.StructDecl {
	originName, err := lang.GetStringAttribute(t.Attributes, lang.OriginalTypeAliasName)
	if err != nil { // directly union declaration
		return nil
	} else {
		labelFormat, _ := lang.GetStringAttribute(t.Attributes, "label_format")

		// generate the message
		s := &lang.StructDecl{}
		alias, err := lang.GetStringAttribute(t.Attributes, "alias")
		if err == nil {
			s.Name = alias
		} else {
			s.Name = originName
		}
		s.Type = &lang.StructType{}
		for _, argument := range t.GenericArguments {
			label, err := lang.GetStringAttribute(argument.Attributes, "label")
			if err != nil {
				format, _ := lang.GetStringAttribute(argument.Attributes, "label_format")
				if len(format) > 0 {
					labelFormat = format
				}

				// auto generate
				label = generateLabel(argument, labelFormat, p)
			}

			s.Type.Fields = append(s.Type.Fields, &lang.ValueDecl{
				Name: label,
				Type: argument,
			})
		}
		return s
	}
}

func generateLabel(argument *lang.NominalType, labelFormat string, p *pluralize.Client) string {
	if labelFormat == "{}" {
		return ToSnake(argument.Name)
	} else {
		fullName := argument.GetFullName()
		if argument.IsScalar() {
			return ToSnake(argument.Name) + "_val"
		} else if fullName == "mojo.core.Array" {
			name := argument.GenericArguments[0].Name
			name = ToSnake(name + "_val")
			return p.Plural(name)
		} else if fullName == "mojo.core.Dictionary" {
			name := argument.GenericArguments[1].Name
			return ToSnake(name + "_map_vals")
		}
		return ToSnake(argument.Name) + "_val"
	}
}
