package _go

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/go"
	desc "github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	path2 "path"
)

type Builder struct {
	PWD     string
	Path    string
	Output  string
	Package *lang.Package
	Files   []*desc.FileDescriptor
}

func (b Builder) Build() error {
	logs.Infow("begin to build go.", "pwd", b.PWD, "path", b.Path)

	compiler := _go.Compiler{
		Path:    path2.Join(b.PWD, b.Path),
		Package: b.Package,
		Files:   b.Files,
	}
	files, err := compiler.Compile()
	if err != nil {
		logs.Errorw("failed to compile go", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	output := path2.Join(b.Path, "go")
	if len(b.Output) > 0 {
		output = b.Output
	}
	generator := _go.NewGenerator(files)
	err = generator.Generate(output)
	if err != nil {
		logs.Errorw("generate go failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
		return err
	}

	return nil
}
