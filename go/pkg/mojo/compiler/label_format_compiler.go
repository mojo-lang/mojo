package compiler

import (
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
)

const labelFormatName = "compiler.label-format"

func init() {
    plugin.RegisterPlugin(NewLabelFormatCompiler(nil))
}

// LabelFormatCompiler used for union, tuple
type LabelFormatCompiler struct {
    plugin.BasicPlugin
}

func NewLabelFormatCompiler(options core.Options) *LabelFormatCompiler {
    return &LabelFormatCompiler{
        BasicPlugin: plugin.BasicPlugin{
            Name:          labelFormatName,
            Group:         "compiler",
            GroupPriority: 10,
            Priority:      13,
            Creator: func(options core.Options) plugin.Plugin {
                return NewLabelFormatCompiler(options)
            },
        },
    }
}

func (c *LabelFormatCompiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
    logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompilePackage", "pkg", pkg.FullName)

    thisCtx := context.WithType(ctx, pkg)
    for _, sourceFile := range pkg.SourceFiles {
        fileCtx := context.WithType(thisCtx, sourceFile)
        for _, statement := range sourceFile.Statements {
            if decl := statement.GetDeclaration(); decl != nil {
                switch decl.Declaration.(type) {
                case *lang.Declaration_StructDecl:
                    //structDecl := decl.GetStructDecl()
                case *lang.Declaration_TypeAliasDecl:
                    if err := c.CompileTypeAlias(fileCtx, decl.GetTypeAliasDecl()); err != nil {
                        return err
                    }
                case *lang.Declaration_EnumDecl:
                case *lang.Declaration_InterfaceDecl:
                }
            }
        }
    }
    return nil
}

func (c *LabelFormatCompiler) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
    //for _, s := range decl.StructDecls {
    //
    //}
    //
    //for _, field := range decl.Type.
    return nil
}

func (c *LabelFormatCompiler) CompileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
    if decl.Type.GetFullName() == core.UnionTypeName {
        if lang.HasAttribute(decl.Attributes, core.LabelFormatAttributeName) {
        }
    }
    return nil
}
