package gokit

import (
    "errors"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/compiler"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/types"
)

type Compiler struct {
    Services map[string]*types.Service
}

func NewCompiler() *Compiler {
    return &Compiler{
        Services: make(map[string]*types.Service),
    }
}

func (c *Compiler) Compile(ctx context.Context, packages map[string]*lang.Package) error {
    for _, pkg := range packages {
        err := c.compilePackage(ctx, pkg)
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

func (c *Compiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
    return c.compilePackage(ctx, pkg)
}

func (c *Compiler) CompileFile(ctx context.Context, file *lang.SourceFile) error {
    return c.compileFile(ctx, file)
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

func (c *Compiler) compileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
    if len(decl.GenericParameters) > 0 {
        return nil
    }

    thisCtx := context.WithType(ctx, decl)
    service, err := compiler.CompileInterface(thisCtx, decl)
    if err != nil {
        return err
    }

    c.Services[decl.GetFullName()] = service
    return nil
}
