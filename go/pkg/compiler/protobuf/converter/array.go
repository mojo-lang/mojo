package converter

import (
	"errors"
	"fmt"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/protobuf/decompiler"
)

type Array struct {
}

func init() {
	RegisterNominalPlugin(&Array{})
}

func (p *Array) Name() string {
	return core.ArrayTypeFullName
}

func (p *Array) ConvertTo(ctx context.Context, t *lang.NominalType) (string, string, error) {
	s, err := decompiler.CompileArray(ctx, t)
	if err != nil {
		return "", "", err
	}

	err = Struct{}.Compile(ctx, s, descriptor.NewMessage(context.FileDescriptor(ctx)))
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("failed to compile the arry field in %s.%s: %s",
			"", // ctx.Message.Name,
			"", // ctx.FieldName,
			err.Error()))
	}

	// ctx.File.Message.AddInnerMessage(context.Message)
	return "struct", s.Name, nil
}
