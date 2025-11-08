package protobuf

import (
	"path"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/converter"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/generator"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/protobuf/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
)

type Builder struct {
	builder.Builder
	Output string
}

func (b Builder) Build() ([]*descriptor.File, error) {
	logs.Infow("begin to build protobuf.", "package", b.Package.FullName, "path", b.Path)

	compiler := converter.New()
	err := compiler.CompilePackage(context.Empty(), b.Package)
	if err != nil {
		logs.Errorw("failed to compile protobuf", "package", b.Package.FullName, "error", err.Error())
		return nil, err
	}

	files := compiler.Descriptors.Filter(b.Package.FullName, false)
	if !b.APIEnabled {
		logs.Infow("disable generation, skip to generate protobuf.")
		return files, nil
	}

	output := path.Join(b.GetAbsolutePath(), "protobuf")
	if len(b.Output) > 0 {
		output = b.Output
	}

	return files, generator.New().GenerateDescriptorsTo(files, output)
}
