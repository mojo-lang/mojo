package openapi

import (
	"errors"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/openapi/compiler"
	"github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

type Compiler struct {
	*OpenAPIs

	Context *compiler.Context
}

func NewCompiler() *Compiler {
	c := &Compiler{
		OpenAPIs: &OpenAPIs{},
	}

	c.Components = &openapi.Components{
		Schemas: make(map[string]*openapi.Schema),
	}
	c.Context = &compiler.Context{
		Context:    &context.Context{},
		Components: c.Components,
	}
	c.APIs = make(map[string]*openapi.OpenAPI)
	return c
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
			case *lang.Declaration_TypeAliasDecl:
				err := c.compileTypeAlias(ctx, decl.GetTypeAliasDecl())
				if err != nil {
					return err
				}
			case *lang.Declaration_StructDecl:
				err := c.compileStruct(ctx, decl.GetStructDecl())
				if err != nil {
					return err
				}
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

func (c *Compiler) compileTypeAlias(ctx *compiler.Context, decl *lang.TypeAliasDecl) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	_, err := compiler.CompileTypeAliasDecl(ctx, decl)
	return err
}

func (c *Compiler) compileStruct(ctx *compiler.Context, decl *lang.StructDecl) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	return compiler.CompileStructDecl(ctx, decl)
}

func (c *Compiler) compileInterface(ctx *compiler.Context, decl *lang.InterfaceDecl) error {
	if len(decl.GenericParameters) > 0 {
		return nil
	}

	ctx.Open(decl)
	defer func() {
		ctx.Close()
	}()

	api, err := compiler.CompileInterface(ctx, decl)
	if err != nil {
		return err
	}

	api.Components = ctx.Components
	key := lang.TypeNameToFileName(decl.GetFullName())
	c.APIs[key] = api
	return nil
}
