package _go

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/go"
	desc "github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	"github.com/mojo-lang/mojo/go/pkg/wand/builder"
	path2 "path"
)

type Builder struct {
	builder.Builder
	Output string
	Files  []*desc.FileDescriptor
}

func (b Builder) Build() error {
	if b.DisableGeneration {
		logs.Infow("disable generation, skip to build go.")
		return nil
	}

	logs.Infow("begin to build go.", "pwd", b.PWD, "path", b.Path)

	compiler := _go.NewCompiler(path2.Join(b.PWD, b.Path), b.Package, b.Files)
	files, err := compiler.Compile()
	if err != nil {
		logs.Errorw("failed to compile go", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	output := path2.Join(b.Path, "go")
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
