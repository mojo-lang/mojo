package _go

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/builder"
	desc "github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	path2 "path"
)

type Builder struct {
	builder.Builder
	Output string
	Files  []*desc.FileDescriptor
}

func (b Builder) Build() error {
	if !b.APIEnabled {
		logs.Infow("disable generation, skip to build go.")
		return nil
	}

	logs.Infow("begin to build go.", "pwd", b.PWD, "path", b.Path)


	compiler := _go.NewCompiler(b.GetAbsolutePath(), b.Package, b.Files)
	files, err := compiler.Compile()
	if err != nil {
		logs.Errorw("failed to compile go", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	output := path2.Join(b.GetAbsolutePath(), "go")
	if len(b.Output) > 0 {
		output = b.Output
	}
	generator := _go.NewGenerator(files, compiler.Data)
	err = generator.Generate(output)
	if err != nil {
		logs.Errorw("generate go failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
		return err
	}

	return nil
}
