package converter

import (
	"errors"
	"fmt"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/packages/protobuf/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/decompiler"
)

type Union struct {
}

func init() {
	RegisterNominalPlugin(&Union{})
}

func (p *Union) Name() string {
	return core.UnionTypeFullName
}

func getOneofName(name string) string {
	n := strcase.ToSnake(name)
	if n == "value" {
		return "val"
	}
	return n
}

func (p *Union) ConvertTo(ctx context.Context, t *lang.NominalType) (string, string, error) {
	if s, err := decompiler.CompileUnion(ctx, t); err != nil {
		return "", "", err
	} else {
		if s == nil {
			return "struct", t.Name, nil
		}

		file := context.FileDescriptor(ctx)
		msgDescriptor := descriptor.NewMessage(file)
		err = Struct{}.Compile(ctx, s, msgDescriptor)
		if err != nil {
			return "", "", errors.New(fmt.Sprintf("failed to compile struct: %s", err.Error()))
		}

		oneof := descriptor.NewOneof(msgDescriptor, getOneofName(msgDescriptor.GetName()))
		msgDescriptor.AppendOneof(oneof)
		for _, field := range msgDescriptor.Fields {
			oneof.AppendField(field)
		}

		return "struct", s.Name, nil
	}
}
