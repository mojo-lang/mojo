package compiler

import (
    "errors"
    "fmt"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "strings"

    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/compiler/transformer"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type ArrayPlugin struct {
}

func init() {
    p := plugins["mojo.core.Array"]
    if p == nil {
        p = make([]Plugin, 0)
    }
    p = append(p, &ArrayPlugin{})

    plugins["mojo.core.Array"] = p
}

func CompileArrayToStruct(t *lang.NominalType) (*lang.StructDecl, error) {
    if t.Name != "Array" {
        return nil, errors.New("")
    }

    if len(t.GenericArguments) != 1 {
        return nil, errors.New("")
    }

    t.Attributes = lang.SetIntegerAttribute(t.Attributes, core.NumberAttributeName, 1)

    s := &lang.StructDecl{}
    s.Type = &lang.StructType{
        Fields: []*lang.ValueDecl{{
            Name: "vals",
            Type: t,
        }},
    }

    val := t.GenericArguments[0]

    if name, _ := lang.GetStringAttribute(t.Attributes, lang.OriginalTypeAliasName); len(name) > 0 {
        if strings.Contains(name, "<") {
            nominalType, _ := lang.ParseNominalTypeFullName(name)
            s.Name = transformer.GenericTypeNamer{}.Name(nominalType)
        } else {
            s.Name = name
        }
    } else {
        s.Name = transformer.Plural(strcase.ToCamel(val.Name))
    }

    return s, nil
}

func (p *ArrayPlugin) Compile(ctx context.Context, t *lang.NominalType) (string, string, error) {
    s, err := CompileArrayToStruct(t)
    if err != nil {
        return "", "", err
    }

    err = CompileStruct(ctx, s, descriptor.NewMessageDescriptor(context.FileDescriptor(ctx)))
    if err != nil {
        return "", "", errors.New(fmt.Sprintf("failed to compile the arry field in %s.%s: %s",
            "", //ctx.Message.Name,
            "", //ctx.FieldName,
            err.Error()))
    }

    //ctx.File.Message.AddInnerMessage(context.Message)
    return "struct", s.Name, nil
}
