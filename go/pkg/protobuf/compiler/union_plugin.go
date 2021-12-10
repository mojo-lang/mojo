package compiler

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/compiler/transformer"
	"github.com/mojo-lang/mojo/go/pkg/context"
	desc "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
	"strings"
)

type UnionPlugin struct {
}

func init() {
	p := plugins["mojo.core.Union"]
	if p == nil {
		p = make([]Plugin, 0)
	}
	p = append(p, &UnionPlugin{})

	plugins["mojo.core.Union"] = p
}

func (p *UnionPlugin) Compile(ctx context.Context, t *lang.NominalType) (string, string, error) {
	s := ConstructBoxedUnion(t)
	if s != nil {
		file := context.FileDescriptor(ctx)
		msgDescriptor := desc.NewMessageDescriptor(file)
		err := CompileStruct(ctx, s, msgDescriptor)
		if err != nil {
			return "", "", errors.New(fmt.Sprintf("failed to compile struct: %s", err.Error()))
		}

		oneOfName := strcase.ToSnake(*msgDescriptor.Name)
		msgDescriptor.OneofDecl = append(msgDescriptor.OneofDecl, &descriptor.OneofDescriptorProto{
			Name: &oneOfName,
		})
		index := int32(0)
		for _, field := range msgDescriptor.Field {
			field.OneofIndex = &index
		}

		return "struct", s.Name, nil
	}

	return "", "", nil
}

func ConstructBoxedUnion(t *lang.NominalType) *lang.StructDecl {
	originName, err := lang.GetStringAttribute(t.Attributes, lang.OriginalTypeAliasName)
	if err != nil { // directly union declaration
		return nil
	} else {
		if strings.Contains(originName, "<") {
			nominalType, _ := lang.ParseNominalTypeFullName(originName)
			originName = transformer.GenericTypeNamer{}.Name(nominalType)
		}

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
				label = generateLabel(argument, labelFormat)
			}

			s.Type.Fields = append(s.Type.Fields, &lang.ValueDecl{
				Name: label,
				Type: argument,
			})
		}
		return s
	}
}

func generateLabel(argument *lang.NominalType, labelFormat string) string {
	if labelFormat == "{}" {
		return strcase.ToSnake(argument.Name)
	} else {
		fullName := argument.GetFullName()
		if argument.IsScalar() {
			return strcase.ToSnake(argument.Name) + "_val"
		} else if fullName == "mojo.core.Array" {
			name := argument.GenericArguments[0].Name
			name = strcase.ToSnake(name + "_val")
			return transformer.Plural(name)
		} else if fullName == "mojo.core.Map" {
			name := argument.GenericArguments[1].Name
			return strcase.ToSnake(name + "_map_vals")
		}
		return strcase.ToSnake(argument.Name) + "_val"
	}
}
