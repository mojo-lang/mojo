package compiler

import (
    "fmt"
    "sort"
    "strings"

    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/db/go/pkg/mojo/db"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
)

func init() {
    plugin.RegisterPlugin(NewEntityCompiler(nil))
}

const (
    entityName = "compiler.entity"

    entityNodeTarget = "node"
    entityEdgeTarget = "edge"
    interfaceTarget  = "interface"

    attributeTarget = "attribute"

    conflictPostfix = "<conflict>"
)

type EntityCompiler struct {
    plugin.BasicPlugin
}

func NewEntityCompiler(options core.Options) *EntityCompiler {
    return &EntityCompiler{
        BasicPlugin: plugin.BasicPlugin{
            Name:          entityName,
            Group:         "compiler",
            GroupPriority: 10,
            Priority:      20,
            Creator: func(options core.Options) plugin.Plugin {
                return NewEntityCompiler(options)
            },
        },
    }
}

func (c *EntityCompiler) CompilePackage(ctx context.Context, pkg *lang.Package) error {
    logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompilePackage", "pkg", pkg.FullName)

    for _, target := range []string{entityNodeTarget, entityEdgeTarget} {
        if err := c.compilePackage(ctx, pkg, target); err != nil {
            return err
        }
    }

    if err := c.refineEntityEdge(pkg); err != nil {
        return err
    }
    if err := c.compilePackage(ctx, pkg, attributeTarget); err != nil {
        return err
    }

    return c.compilePackage(ctx, pkg, interfaceTarget)
}

func (c *EntityCompiler) compilePackage(ctx context.Context, pkg *lang.Package, target string) error {
    thisCtx := context.WithType(ctx, pkg)
    for _, sourceFile := range pkg.SourceFiles {
        fileCtx := context.WithType(thisCtx, sourceFile)
        for _, statement := range sourceFile.Statements {
            if decl := statement.GetDeclaration(); decl != nil {
                switch decl.Declaration.(type) {
                case *lang.Declaration_StructDecl:
                    if target == entityNodeTarget {
                        if err := c.compileEntityNode(fileCtx, decl.GetStructDecl()); err != nil {
                            return err
                        }
                    } else if target == entityEdgeTarget {
                        if err := c.compileEntityEdge(fileCtx, decl.GetStructDecl()); err != nil {
                            return err
                        }
                    } else if target == attributeTarget {
                        if err := c.compileAttribute(fileCtx, decl.GetStructDecl()); err != nil {
                            return err
                        }
                    }
                case *lang.Declaration_InterfaceDecl:
                    if target == interfaceTarget {
                        if err := c.CompileInterface(fileCtx, decl.GetInterfaceDecl()); err != nil {
                            return err
                        }
                    }
                }
            }
        }
    }
    return nil
}

func (c *EntityCompiler) compileEntityNode(ctx context.Context, decl *lang.StructDecl) error {
    pkg := context.Package(ctx)

    makeEntity := func(field *lang.ValueDecl) {
        pkg.SetEntityNode(decl.GetFullName(), &lang.EntityNode{
            Name:            decl.GetFullName(),
            TypeDeclaration: lang.NewStructTypeDeclaration(decl),
            KeyField:        field,
        })
        decl.SetImplicitBoolAttribute(core.EntityAttributeName, true)
    }

    if field := decl.GetField("id"); field != nil {
        makeEntity(field)
        return nil
    }

    decl.EachField(func(decl *lang.ValueDecl) error {
        if decl.HasAttribute(core.KeyAttributeName) || decl.HasAttribute(db.PrimaryKeyAttributeName) {
            makeEntity(decl)
            return core.NewBreakError()
        }
        return nil
    })
    return nil
}

func refineLabel(label string, node *lang.EntityNode) string {
    if decl := node.GetTypeDeclaration().GetStructDecl(); decl != nil {
        for _, name := range decl.FieldNames(lang.FieldNamOptionDefault) {
            suffix := "_" + name
            if strings.HasSuffix(label, suffix) {
                return strings.TrimSuffix(label, suffix)
            }
        }
    }

    return label
}

func (c *EntityCompiler) compileEntityEdge(ctx context.Context, decl *lang.StructDecl) error {
    if decl == nil || decl.Type == nil || !decl.HasAttribute(core.EntityAttributeName) {
        return nil
    }

    pkg := context.Package(ctx)
    makeEdge := func(field *lang.ValueDecl, toName string, inverse bool, implicit bool) {
        id := ""
        if inverse {
            id = fmt.Sprintf("%s~>%s", decl.GetFullName(), toName)
        } else {
            id = fmt.Sprintf("%s->%s", decl.GetFullName(), toName)
        }

        edge := &lang.EntityEdge{
            Id:             id,
            Name:           field.Name,
            From:           pkg.GetEntityNode(decl.GetFullName()),
            To:             pkg.GetEntityNode(toName),
            Inverse:        inverse,
            Multiple:       field.IsArrayType(),
            Implicit:       implicit,
            ReferenceField: field,
        }
        edge.Name = refineLabel(edge.Name, edge.To)

        if pkg.GetEntityEdge(edge.Name) != nil {
            edge.Id = edge.Id + conflictPostfix
            pkg.SetEntityEdge(edge.Id, edge)
        } else {
            pkg.SetEntityEdge(edge.Id, edge)
        }
    }

    for _, field := range decl.Type.Fields {
        typeDecl := field.GetType().GetTypeDeclaration().GetStructDecl()
        if field.IsArrayType() {
            typeDecl = field.GetType().GenericArguments[0].GetTypeDeclaration().GetStructDecl()
        } else if field.IsMapType() {
            //TODO not support yet
            continue
        }

        if field.HasAttribute(core.ReferenceAttributeName) {
            if reference, err := field.GetStringAttribute(core.ReferenceAttributeName); err == nil && len(reference) > 0 {
                if identifier := pkg.GetIdentifier(reference); identifier != nil {
                    makeEdge(field, identifier.FullName, false, false)
                } else {
                    return fmt.Errorf("not found the identifier (%s) with reference in the field '%s' of %s", reference, field.Name, decl.GetFullName())
                }
            } else if typeDecl.HasAttribute(core.EntityAttributeName) {
                makeEdge(field, typeDecl.GetFullName(), false, false)
            } else {
                return fmt.Errorf("not set the reference type nor an entity type in the field '%s' of %s", field.Name, decl.GetFullName())
            }
        } else if field.HasAttribute(core.BackReferenceAttributeName) {
            if backReference, err := field.GetStringAttribute(core.BackReferenceAttributeName); err == nil && len(backReference) > 0 {
                if identifier := pkg.GetIdentifier(backReference); identifier != nil {
                    makeEdge(field, identifier.FullName, true, false)
                } else {
                    return fmt.Errorf("not found the identifier (%s) with back reference in the field '%s' of %s", backReference, field.Name, decl.GetFullName())
                }
            } else if typeDecl.HasAttribute(core.EntityAttributeName) {
                makeEdge(field, typeDecl.GetFullName(), true, false)
            } else {
                return fmt.Errorf("not set the back reference type nor an entity type in the field '%s' of %s", field.Name, decl.GetFullName())
            }
        } else if typeDecl.HasAttribute(core.EntityAttributeName) {
            makeEdge(field, typeDecl.GetFullName(), false, true)
        }
    }
    return nil
}

func (c *EntityCompiler) refineEntityEdge(pkg *lang.Package) error {
    group := make(map[string][]*lang.EntityEdge)
    for _, edge := range pkg.EntityRelationSet.GetEdges() {
        names := []string{edge.From.Name, edge.To.Name}
        sort.Strings(names)
        name := strings.Join(names, "-")

        group[name] = append(group[name], edge)
    }

    resetEdge := func(edge *lang.EntityEdge, inverse bool) {
        pkg.DeleteEntityEdge(edge.Id)

        if inverse {
            edge.Id = fmt.Sprintf("%s~>%s", edge.From.Name, edge.To.Name)
        } else {
            edge.Id = fmt.Sprintf("%s->%s", edge.From.Name, edge.To.Name)
        }
        edge.Inverse = inverse
        pkg.SetEntityEdge(edge.Id, edge)
    }

    for _, edges := range group {
        if len(edges) == 1 {
            continue
        }

        if edges[0].Multiple && edges[1].Multiple { // M2M
            if edges[0].Inverse == edges[1].Inverse {
                if edges[0].Implicit {
                    resetEdge(edges[0], !edges[0].Inverse)
                } else if edges[1].Implicit {
                    resetEdge(edges[1], !edges[1].Inverse)
                } else {
                    return fmt.Errorf("the edges has same direction %s (%s), %s (%s)", edges[0].Id, edges[0].Name, edges[1].Id, edges[1].Name)
                }
            }
        } else if !edges[0].Multiple && !edges[1].Multiple { // O2O
            if edges[0].Inverse == edges[1].Inverse {
                if edges[0].Implicit {
                    resetEdge(edges[0], !edges[0].Inverse)
                } else if edges[1].Implicit {
                    resetEdge(edges[1], !edges[1].Inverse)
                } else {
                    return fmt.Errorf("the edges has same direction %s (%s), %s (%s)", edges[0].Id, edges[0].Name, edges[1].Id, edges[1].Name)
                }
            }
        } else {
            if edges[0].Inverse == edges[1].Inverse {
                if edges[0].Implicit {
                    if edges[1].Implicit {
                        resetEdge(edges[0], !edges[0].Multiple)
                        resetEdge(edges[1], !edges[1].Multiple)
                    } else {
                        resetEdge(edges[0], !edges[0].Inverse)
                    }
                } else {
                    if edges[1].Implicit {
                        resetEdge(edges[1], !edges[1].Inverse)
                    } else {
                        return fmt.Errorf("the edges has same direction %s (%s), %s (%s)", edges[0].Id, edges[0].Name, edges[1].Id, edges[1].Name)
                    }
                }
            }
        }
    }
    return nil
}

func (c *EntityCompiler) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
    for key, methods := range decl.GetMethodGroups() {
        for _, method := range methods {
            //TODO: check the entity type with input or output type
            method.SetStringAttribute(core.EntityAttributeName, key)
        }
    }
    return nil
}

func (c *EntityCompiler) compileAttribute(ctx context.Context, decl *lang.StructDecl) error {
    pkg := context.Package(ctx)
    if pkg.EntityRelationSet == nil {
        return nil
    }

    for _, edge := range pkg.EntityRelationSet.Edges {
        if edge.GetFrom().Name == decl.GetFullName() {
            if edge.Inverse {
                if !edge.ReferenceField.HasAttribute(core.BackReferenceAttributeName) {
                    edge.ReferenceField.SetImplicitBoolAttribute(core.BackReferenceAttributeName, true)
                }
            } else {
                if !edge.ReferenceField.HasAttribute(core.ReferenceAttributeName) {
                    edge.ReferenceField.SetImplicitBoolAttribute(core.ReferenceAttributeName, true)
                }
            }
        }
    }

    return nil
}
