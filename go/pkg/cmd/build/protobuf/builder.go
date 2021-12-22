package protobuf

import (
	path2 "path"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
	"github.com/mojo-lang/mojo/go/pkg/protobuf"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type Builder struct {
	builder.Builder
	Output string
}

func (b Builder) Build() ([]*descriptor.FileDescriptor, error) {
	logs.Infow("begin to build protobuf.", "package", b.Package.FullName, "path", b.Path)

	compiler := protobuf.NewCompiler()
	err := compiler.CompilePackages(b.Package.GetAllPackages())
	if err != nil {
		logs.Errorw("failed to compile protobuf", "package", b.Package.FullName, "error", err.Error())
		return nil, err
	}

	if !b.APIEnabled {
		logs.Infow("disable generation, skip to generate protobuf.")
		return compiler.Files, nil
	}

	output := path2.Join(b.GetAbsolutePath(), "protobuf")
	if len(b.Output) > 0 {
		output = b.Output
	}

	generator := protobuf.NewGenerator(compiler.Files)
	return generator.Generate(output)
}
