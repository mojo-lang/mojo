package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type Compiler struct {
}

func New() *Compiler {
	return &Compiler{}
}

func (c *Compiler) CompilePackages(packages map[string]*lang.Package) error {
	for _, pkg := range packages {
		err := c.compilePackage(context.Empty(), pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
	return c.compilePackage(ctx, pkg)
}

func (c *Compiler) compilePackage(ctx context.Context, pkg *lang.Package) error {
	typeAlias := &TypeAliasCompiler{}
	if err := typeAlias.CompilePackage(ctx, pkg); err != nil {
		return err
	}

	generic := &GenericCompiler{}
	if err := generic.CompilePackage(ctx, pkg); err != nil {
		return err
	}

	genericNaming := NewGenericRenamingCompiler()
	if err := genericNaming.CompilePackage(ctx, pkg); err != nil {
		return err
	}

	if err := typeAlias.CompilePackage(ctx, pkg); err != nil {
		return err
	}

	circle := &CircleCompiler{}
	if err := circle.Compile(ctx, pkg); err != nil {
		return err
	}

	goPkg := &GoPackageNameCompiler{}
	if err := goPkg.CompilePackage(ctx, pkg); err != nil {
		return err
	}

	return nil
}
