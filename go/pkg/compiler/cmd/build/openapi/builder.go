package openapi

import (
	path2 "path"

	"github.com/mojo-lang/mojo/go/pkg/compiler/openapi/generator"

	"github.com/mojo-lang/mojo/go/pkg/logs"
	api "github.com/mojo-lang/mojo/go/pkg/mojo/openapi"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
)

type Builder struct {
	builder.Builder
	Output string
}

func (b Builder) Build() (*api.OpenAPIs, error) {
	logs.Infow("begin to build openapi.", "pwd", b.PWD, "path", b.Path)

	compiler := generator.NewCompiler()
	err := compiler.CompilePackages(b.Package.GetAllPackages())
	if err != nil {
		logs.Errorw("failed to compile openapi", "package", b.Package.FullName, "error", err.Error())
		return nil, err
	}

	if !b.APIEnabled {
		logs.Infow("disable generation, skip to generate openapi.")
		return compiler.OpenAPIs, nil
	}

	output := path2.Join(b.GetAbsolutePath(), "openapi")
	if len(b.Output) > 0 {
		output = b.Output
	}
	gen := generator.NewGenerator(b.Package, compiler.APIs, compiler.Components)
	err = gen.Generate(output)
	if err != nil {
		logs.Errorw("generate openapi failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
		return nil, err
	}

	return compiler.OpenAPIs, nil
}
