package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type Compiler struct {
	Context *context.Context
}

func New() *Compiler {
	return &Compiler{
		Context: &context.Context{},
	}
}

func (c *Compiler) CompilePackages(packages map[string]*lang.Package) error {
	for _, pkg := range packages {
		err := c.compilePackage(c.Context, pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compiler) CompilePackage(pkg *lang.Package) error {
	return c.compilePackage(c.Context, pkg)
}

func (c *Compiler) compilePackage(ctx *context.Context, pkg *lang.Package) error {
	ctx.Open(pkg)
	defer func() {
		ctx.Close()
	}()

	typeAlias := &TypeAliasCompiler{}
	if err := typeAlias.Compile(ctx, pkg); err != nil {
		return err
	}

	generic := &GenericCompiler{}
	if err := generic.Compile(ctx, pkg); err != nil {
		return err
	}

	if err := typeAlias.Compile(ctx, pkg); err != nil {
		return err
	}

	circle := &CircleCompiler{}
	if err := circle.Compile(ctx, pkg); err != nil {
		return err
	}

	goPkg := &GoPackageNameCompiler{}
	if err := goPkg.Compile(ctx, pkg); err != nil {
		return err
	}

	return nil
}
