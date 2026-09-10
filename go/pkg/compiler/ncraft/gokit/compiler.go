package gokit

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/compiler"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
)

type Compiler struct {
	Services []*data.Service
	Entities []*data.Message
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	compiled, err := compiler.Compile(ctx, pkg)
	if err != nil {
		return err
	}

	// add gokit compiler here

	c.Services = append(c.Services, compiled.Data...)
	c.Entities = append(c.Entities, compiled.Entities...)
	return nil
}
