package compiler

import (
    "errors"
    "fmt"
    "strings"

    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/compiler/transformer"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    desc "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type Union struct {
}

func init() {
    RegisterNominalPlugin(&Union{})
}

func (p *Union) Name() string {
    return "mojo.core.Union"
}

func (p *Union) Compile(ctx context.Context, t *lang.NominalType) (string, string, error) {
    s := ConstructBoxedUnion(t)
    if s != nil {
        file := context.FileDescriptor(ctx)
        msgDescriptor := desc.NewMessage(file)
        err := Struct{}.Compile(ctx, s, msgDescriptor)
        if err != nil {
            return "", "", errors.New(fmt.Sprintf("failed to compile struct: %s", err.Error()))
        }

        oneof := desc.NewOneof(msgDescriptor, strcase.ToSnake(msgDescriptor.GetName()))
        msgDescriptor.AppendOneof(oneof)
        for _, field := range msgDescriptor.Fields {
            oneof.AppendField(field)
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
