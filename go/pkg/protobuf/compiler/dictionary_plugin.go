package compiler

import (
    "errors"
    "fmt"
    "github.com/golang/protobuf/protoc-gen-go/descriptor"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    desc "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type MapPlugin struct {
}

func init() {
    p := plugins["mojo.core.Map"]
    if p == nil {
        p = make([]Plugin, 0)
    }
    p = append(p, &MapPlugin{})

    plugins["mojo.core.Map"] = p
}

func (p *MapPlugin) Compile(ctx context.Context, t *lang.NominalType) (string, string, error) {
    if t.Name != "Map" {
        return "", "", errors.New("")
    }

    if len(t.GenericArguments) != 2 {
        return "", "", errors.New("")
    }

    s := &lang.StructDecl{}
    fieldName := ""
    if decl := context.ValueDecl(ctx); decl != nil {
        fieldName = decl.Name
    }
    s.Name = strcase.ToCamel(fieldName) + "Entry"

    keyType := t.GenericArguments[0]
    keyType.Attributes = lang.SetIntegerAttribute(keyType.Attributes, core.NumberAttributeName, 1)

    valType := t.GenericArguments[1]
    valType.Attributes = lang.SetIntegerAttribute(valType.Attributes, core.NumberAttributeName, 2)

    if len(valType.GenericArguments) > 0 {
        structCtx := context.WithValues(ctx, "register_struct", true)
        _, name, err := CompileNominalType(structCtx, valType)
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

    msgDescriptor := desc.NewMessageDescriptor(context.FileDescriptor(ctx))
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

    message := context.MessageDescriptor(ctx)
    if message != nil {
        message.AddInnerMessage(msgDescriptor)
    }

    return "struct", s.Name, nil
}
