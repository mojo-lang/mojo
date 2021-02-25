package document

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/document"
	"github.com/mojo-lang/mojo/go/pkg/wand/openapi"
	path2 "path"
)

type Builder struct {
	PWD      string
	Path     string
	Output   string
	Package  *lang.Package
	OpenAPIs *openapi.OpenAPIs
}

func (b Builder) Build() error {
	logs.Infow("begin to build document.", "pwd", b.PWD, "path", b.Path)

	compiler := document.NewCompiler(path2.Join(b.PWD, b.Path), b.Package, b.OpenAPIs)
	documents, err := compiler.Compile()
	if err != nil {
		logs.Errorw("failed to compile document", "package", b.Package.FullName, "error", err.Error())
		return err
	}

	output := path2.Join(b.Path, "document")
	if len(b.Output) > 0 {
		output = b.Output
	}
	generator := document.NewGenerator(documents)
	err = generator.Generate(output)
	if err != nil {
		logs.Errorw("generate document failed", "pwd", b.PWD, "path", b.Path, "package", b.Package.FullName, "error", err.Error())
		return err
	}

	return nil
}
