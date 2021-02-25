package openapi

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/openapi"
	path2 "path"
)

type OpenAPIs = openapi.OpenAPIs

type Builder struct {
	PWD     string
	Path    string
	Output  string
	Package *lang.Package
}

func (b Builder) Build() (*OpenAPIs, error) {
	logs.Infow("begin to build openapi.", "pwd", b.PWD, "path", b.Path)

	compiler := openapi.NewCompiler()
	err := compiler.CompilePackages(b.Package.GetAllPackage())
	if err != nil {
		logs.Errorw("failed to compile openapi", "package", b.Package.FullName, "error", err.Error())
		return nil, err
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
