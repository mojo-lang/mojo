package mojo

import (
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/builder"
)

type Builder struct {
	builder.Builder
}

func (b Builder) Build() (*lang.Package, error) {
	logs.Infow("begin to parse mojo package.", "pwd", b.PWD, "path", b.Path)
	parser := NewParser(b.PWD)
	err := parser.Parse(b.Path)
	if err != nil {
		return nil, err
	}

	b.Package = parser.Root

	// compile all the package
	logs.Infow("begin to compile mojo package.", "pwd", b.PWD, "path", b.Path)
	if err := NewCompiler().Compile(b.Package); err != nil {
		return nil, err
	}

	return b.Package, nil
}
