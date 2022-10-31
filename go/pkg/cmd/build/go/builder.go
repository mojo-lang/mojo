package _go

import (
	"os/exec"
	path2 "path"

	"github.com/mojo-lang/mojo/go/pkg/go/generator"
	"github.com/mojo-lang/mojo/go/pkg/util"

	"github.com/mojo-lang/core/go/pkg/logs"
	desc "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
)

type Builder struct {
	builder.Builder
	Output string
	Files  []*desc.File
}

func (b Builder) Build() error {
	if !b.APIEnabled {
		logs.Infow("disable generation, skip to build go.")
		return nil
	}

	logs.Infow("begin to build go.", "pwd", b.PWD, "path", b.Path)

	compiler := generator.NewCompiler(b.GetAbsolutePath(), b.Files)
	files, err := compiler.CompilePackage(b.Package)
	if err != nil {
		logs.Errorw("failed to compile go", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	output := path2.Join(b.GetAbsolutePath(), "go")
	if len(b.Output) > 0 {
		output = util.GetAbsolutePath(b.PWD, b.Output)
	}
	generator := generator.NewGenerator(files, compiler.Data)
	err = generator.Generate(output)
	if err != nil {
		logs.Errorw("generate go failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
		return err
	}

	GoModTidy(output)
	return nil
}

func GoModTidy(pwd string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = pwd

	logs.Info("begin to running go mod tidy")
	out, err := cmd.CombinedOutput()
	if err != nil {
		logs.Errorw("failed to run go mod tidy", "error", string(out))
		return err
	}
	logs.Info("finish to run go mod tidy")
	return nil
}
