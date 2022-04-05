package openapi

import (
    "errors"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/openapi/compiler"
    "github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
)

type Compiler struct {
    *openapi.OpenAPIs
}

func NewCompiler() *Compiler {
    c := &Compiler{
        OpenAPIs: &openapi.OpenAPIs{
            APIs:       make(map[string]*openapi.OpenAPI),
            Components: openapi.NewComponents(),
        },
    }
    return c
}

func (c *Compiler) NewContext() context.Context {
    return context.WithComponents(context.Empty(), c.Components)
}

func (c *Compiler) CompilePackages(packages map[string]*lang.Package) error {
    for _, pkg := range packages {
        err := c.compilePackage(c.NewContext(), pkg)
        if err != nil {
            return err
        }
    }
    return nil
}

func (c *Compiler) CompilePackage(pkg *lang.Package) error {
    return c.compilePackage(c.NewContext(), pkg)
}

func (c *Compiler) CompileFile(file *lang.SourceFile) error {
    return c.compileFile(c.NewContext(), file)
}

func (c *Compiler) compilePackage(ctx context.Context, pkg *lang.Package) error {
    thisCtx := context.WithType(ctx, pkg)
    for _, sourceFile := range pkg.SourceFiles {
        err := c.compileFile(thisCtx, sourceFile)
        if err != nil {
            return err
        }
    }
    return nil
}

func (c *Compiler) compileFile(ctx context.Context, file *lang.SourceFile) error {
    if file.IsGenericInstantiated() {
        return nil
    }

    thisCtx := context.WithType(ctx, file)

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
                err := c.compileTypeAlias(thisCtx, decl.GetTypeAliasDecl())
                if err != nil {
                    return err
                }
            case *lang.Declaration_StructDecl:
                err := c.compileStruct(thisCtx, decl.GetStructDecl())
                if err != nil {
                    return err
                }
            case *lang.Declaration_InterfaceDecl:
                err := c.compileInterface(thisCtx, decl.GetInterfaceDecl())
                if err != nil {
                    return err
                }
            }
        }
    }

    return nil
}

func (c *Compiler) compileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
    if len(decl.GenericParameters) > 0 {
        return nil
    }
    _, err := compiler.CompileTypeAliasDecl(context.WithType(ctx, decl), decl)
    return err
}

func (c *Compiler) compileStruct(ctx context.Context, decl *lang.StructDecl) error {
    if len(decl.GenericParameters) > 0 {
        return nil
    }
    return compiler.CompileStructDecl(context.WithType(ctx, decl), decl)
}

func (c *Compiler) compileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
    if len(decl.GenericParameters) > 0 {
        return nil
    }

    thisCtx := context.WithType(ctx, decl)
    api, err := compiler.CompileInterface(thisCtx, decl)
    if err != nil {
        return err
    }

    api.Components = context.Components(thisCtx)
    //key := lang.TypeNameToFileName(decl.GetFullName())
    c.APIs[decl.GetFullName()] = api
    return nil
}
