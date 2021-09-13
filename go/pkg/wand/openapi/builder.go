package openapi

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/openapi"
	"github.com/mojo-lang/mojo/go/pkg/wand/builder"
	path2 "path"
)

type OpenAPIs = openapi.OpenAPIs

type Builder struct {
	builder.Builder
	Output string
}

func (b Builder) Build() (*OpenAPIs, error) {
	logs.Infow("begin to build openapi.", "pwd", b.PWD, "path", b.Path)

	compiler := openapi.NewCompiler()
	err := compiler.CompilePackages(b.Package.GetAllPackage())
	if err != nil {
		logs.Errorw("failed to compile openapi", "package", b.Package.FullName, "error", err.Error())
		return nil, err
	}

	if b.DisableGeneration {
		logs.Infow("disable generation, skip to generate openapi.")
		return compiler.OpenAPIs, nil
	}

	output := path2.Join(b.Path, "openapi")
	if len(b.Output) > 0 {
		output = b.Output
	}
	generator := openapi.NewGenerator(compiler.APIs, compiler.Components)
	err = generator.Generate(output)
	if err != nil {
		logs.Errorw("generate openapi failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
		return nil, err
	}

	return compiler.OpenAPIs, nil
}
