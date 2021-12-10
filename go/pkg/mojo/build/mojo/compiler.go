package mojo

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/compiler"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type Compiler struct {
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) Compile(pkg *lang.Package) error {
	if pkg.GetExtraBool("compiled") {
		return nil
	}

	for _, dependency := range pkg.ResolvedDependencies {
		if dependency.GetExtraBool("compiled") {
			continue
		}

		if err := c.Compile(dependency); err != nil {
			return err
		}
	}

	compiler := compiler.New()
	if err := compiler.CompilePackage(context.Empty(), pkg); err != nil {
		return err
	}

	for _, p := range pkg.Children {
		if err := compiler.CompilePackage(context.Empty(), p); err != nil {
			return err
		}
	}

	pkg.SetExtraBool("compiled", true)
	return nil
}
