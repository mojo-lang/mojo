package document

import (
	path2 "path"

	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/openapi"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	"github.com/mojo-lang/mojo/go/pkg/compiler/document/generator"
)

type Builder struct {
	builder.Builder
	Output   string
	OpenAPIs *openapi.OpenAPIs
}

func (b Builder) Build() error {
	if !b.APIEnabled {
		logs.Infow("disable generation, skip to build document.")
		return nil
	}

	logs.Infow("begin to build document.", "pwd", b.PWD, "path", b.Path)

	compiler := generator.NewCompiler(b.GetAbsolutePath(), b.Package, b.OpenAPIs)
	documents, err := compiler.Compile()
	if err != nil {
		logs.Errorw("failed to compile document", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	output := path2.Join(b.GetAbsolutePath(), "document")
	if len(b.Output) > 0 {
		output = b.Output
	}
	gen := generator.NewGenerator(documents)
	err = gen.Generate(output)
	if err != nil {
		logs.Errorw("generate document failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
		return err
	}

	return nil
}
