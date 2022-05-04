package circle

import (
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
    "github.com/mojo-lang/mojo/go/pkg/mojo/util"
    path2 "path"
)

const DependencyCircle = "dependency_circle"
const pluginName = "semantic.circle-resolver"

type Resolver struct {
    plugin.BasicPlugin
    Force bool
}

func init() {
    plugin.RegisterPlugin(NewResolver(nil))
}

func NewResolver(options core.Options) *Resolver {
    return &Resolver{
        BasicPlugin: plugin.BasicPlugin{
            Name:          pluginName,
            Group:         "semantic",
            GroupPriority: 3,
            Priority:      5,
            Creator: func(options core.Options) plugin.Plugin {
                return NewResolver(options)
            },
        },
    }
}

func (r *Resolver) ParsePackage(ctx context.Context, pkg *lang.Package) error {
    if !r.Force && util.IsPackageProcessed(pkg, pluginName) {
        logs.Infow("already processed, skip the plugin", "plugin", r.Name, "method", "ParsePackage", "pkg", pkg.FullName)
        return nil
    } else {
        logs.Infow("enter the plugin", "plugin", r.Name, "method", "ParsePackage", "pkg", pkg.FullName)
    }

    thisCtx := context.WithType(ctx, pkg)

    for _, child := range pkg.Children {
        if err := r.ParsePackage(thisCtx, child); err != nil {
            return err
        }
    }

    if err := r.parsePackage(thisCtx, pkg); err != nil {
        return err
    }

    if !pkg.IsPadding() && !r.Force {
        util.SetPackageProcessed(pkg, pluginName)
    }
    return nil
}

func removeDuplicated(values []string) []string {
    index := make(map[string]bool)
    for _, value := range values {
        index[value] = true
    }

    values = []string{}
    for value := range index {
        values = append(values, value)
    }
    return values
}

func (r *Resolver) parsePackage(ctx context.Context, pkg *lang.Package) error {
    // generated all the file node
    _ = ctx
    files := make(map[string]*fileNode)
    for _, file := range pkg.SourceFiles {
        node := &fileNode{
            name: file.FullName,
        }

        for _, dependency := range file.ResolvedIdentifiers {
            if dependency.PackageName == pkg.FullName /*&& !dependency.IsGenericInstantiated()*/ {
                node.dependencies = append(node.dependencies, dependency.SourceFileName)
            }
        }
        node.dependencies = removeDuplicated(node.dependencies)

        files[node.name] = node
    }

    circles := NewSearcher().Search(files)
    if len(circles) == 0 {
        return nil
    }

    circleIndex := make(map[string]bool)
    for _, circle := range circles {
        for _, name := range circle {
            circleIndex[name] = true
        }
    }

    for _, file := range pkg.SourceFiles {
        if circleIndex[file.FullName] {
            circleName := r.generateCircleName(ctx, pkg)
            file.Attributes = lang.SetStringAttribute(file.Attributes, DependencyCircle, circleName)
        }
    }

    return nil
}

func (r *Resolver) generateCircleName(ctx context.Context, pkg *lang.Package) string {
    _ = ctx
    return path2.Join(lang.PackageNameToPath(pkg.FullName), pkg.Name+".mojo")
}
