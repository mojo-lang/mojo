package boot

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/compiler"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/data"
)

type Compiler struct {
	Services []*data.Service
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	services, err := compiler.CompilePackage(ctx, pkg)
	if err != nil {
		return err
	}

	// add boot compiler here

	c.Services = append(c.Services, services...)
	return nil
}
