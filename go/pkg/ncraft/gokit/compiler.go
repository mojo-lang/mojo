package gokit

import (
	"errors"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/compiler"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/types"
)

type Compiler struct {
	Context *compiler.Context

	Services map[string]*types.Service
}

func NewCompiler() *Compiler {
	return &Compiler{
		Context:  compiler.NewContext(),
		Services: make(map[string]*types.Service),
	}
}

func (c *Compiler) Compile(packages map[string]*lang.Package) error {
	for _, pkg := range packages {
		err := c.compilePackage(c.Context, pkg)
		if err != nil {
			return err
		}
	}

	// fill the `AllInterfaceNames` field in the `Service`
	interfaceGroup := make(map[string][]string)
	for _, s := range c.Services {
		pkg := s.FullPkgName
		interfaceGroup[pkg] = append(interfaceGroup[pkg], s.Interface.Name)
	}
	for _, s := range c.Services {
		pkg := s.FullPkgName
		s.AllInterfaceNames = interfaceGroup[pkg]
	}

	return nil
}

func (c *Compiler) CompilePackage(pkg *lang.Package) error {
	return c.compilePackage(c.Context, pkg)
}

func (c *Compiler) CompileFile(file *lang.SourceFile) error {
	return c.compileFile(c.Context, file)
}

func (c *Compiler) compilePackage(ctx *compiler.Context, pkg *lang.Package) error {
	ctx.Open(pkg)
	defer func() {
		ctx.Close()
	}()

	for _, sourceFile := range pkg.SourceFiles {
		err := c.compileFile(ctx, sourceFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Compiler) compileFile(ctx *compiler.Context, file *lang.SourceFile) error {
	ctx.Open(file)
	defer func() {
		ctx.Close()
	}()

	// compile statements
	for _, statement := range file.Statements {
		switch statement.Statement.(type) {
		case *lang.Statement_Declaration:
			decl := statement.GetDeclaration()
			if decl == nil {
				return errors.New("declaration statement is nil")
			}

			switch decl.Declaration.(type) {
			case *lang.Declaration_InterfaceDecl:
				err := c.compileInterface(ctx, decl.GetInterfaceDecl())
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (c *Compiler) compileInterface(ctx *compiler.Context, decl *lang.InterfaceDecl) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	service, err := compiler.CompileInterface(ctx, decl)
	if err != nil {
		return err
	}

	c.Services[decl.GetFullName()] = service
	return nil
}
