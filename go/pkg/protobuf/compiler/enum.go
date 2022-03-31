package compiler

import (
    "errors"

    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    desc "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type Enum struct {
}

func (e Enum) Compile(ctx context.Context, decl *lang.EnumDecl, enumDescriptor *desc.Enum) error {
    thisCtx := context.WithDescriptor(context.WithType(ctx, decl), enumDescriptor)
    enumDescriptor.SetName(decl.Name)

    for i, e := range decl.Type.Enumerators {
        name := decl.Name + "_" + e.Name
        name = strcase.ToScreamingSnake(name)

        number, err := lang.GetIntegerAttribute(e.Attributes, core.NumberAttributeName)
        if err != nil {
            number = int64(i)
        } else {
            if number < 0 {
                return errors.New("number attribute value must be positive")
            }
        }

        enumDescriptor.AppendValueWith(name, int32(number))
    }

    message := context.MessageDescriptor(thisCtx)
    file := context.FileDescriptor(thisCtx)
    if message == nil && file != nil {
        if register, ok := ctx.Value("register_enum").(bool); !ok || register {
            file.Enums = append(file.Enums, enumDescriptor)
        }
    }

    return nil
}
