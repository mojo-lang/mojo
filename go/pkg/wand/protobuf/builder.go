package protobuf

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/protobuf"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	path2 "path"
)

type Builder struct {
	PWD     string
	Path    string
	Output  string
	Package *lang.Package
}

func (b Builder) Build() ([]*descriptor.FileDescriptor, error) {
	logs.Infow("begin to build protobuf.", "package", b.Package.FullName, "path", b.Path)

	compiler := protobuf.NewCompiler()
	err := compiler.CompilePackages(b.Package.GetAllPackage())
	if err != nil {
		logs.Errorw("failed to compile protobuf", "package", b.Package.FullName, "error", err.Error())
		return nil, err
	}

	output := path2.Join(b.Path, "protobuf")
	if len(b.Output) > 0 {
		output = b.Output
	}

	generator := protobuf.NewGenerator(compiler.Files)
	return generator.Generate(output)
}
