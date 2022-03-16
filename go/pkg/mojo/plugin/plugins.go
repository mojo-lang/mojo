package plugin

import (
    "errors"
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin/parser"
    "io/fs"
    "sort"
)

type Plugins struct {
    plugins
    cursor         int
    parsedPackages map[string]bool
}

func NewPlugins(plugins ...string) *Plugins {
    ps := &Plugins{
        parsedPackages: make(map[string]bool),
    }

    for _, name := range plugins {
        if p := GetPlugin(name); p != nil {
            ps.plugins = append(ps.plugins, p.Create(nil))
        } else if plugs := GetPluginGroup(name); len(plugs) > 0 {
            for _, p = range plugs {
                ps.plugins = append(ps.plugins, p.Create(nil))
            }
        } else {
            logs.Warnw("the plugin has not been register", "plugin", name)
        }
    }
    return ps.sort()
}

func (p *Plugins) Next() *Plugins {
    p.cursor++
    if p.cursor == len(p.plugins) {
        return nil
    }
    return p
}

func (p *Plugins) plugin() Plugin {
    if p.cursor >= len(p.plugins) {
        return nil
    }
    return p.plugins[p.cursor]
}

func (p *Plugins) sort() *Plugins {
    sort.Sort(p.plugins)
    return p
}

func (p *Plugins) copy() *Plugins {
    return &Plugins{
        plugins:        p.plugins,
        parsedPackages: p.parsedPackages,
        cursor:         0,
    }
}

// ParsePackagePath if there is no mpm plugin, need set the package name to the ctx, using `parser.WithPackageName(ctx, pkgName)`
func (p *Plugins) ParsePackagePath(ctx context.Context, pkgPath string, fileSys fs.FS) (pkg *lang.Package, err error) {
    thisCtx := WithPlugins(parser.WithFsCache(ctx, make(parser.FsCache)), p)
    for p.plugin() != nil {
        if psr, ok := p.plugin().(parser.PackagePathParser); ok {
            if pkg, err = psr.ParsePackagePath(thisCtx, pkgPath, fileSys); err != nil {
                return nil, err
            }
            break
        }

        p.Next()
    }

    for _, dep := range pkg.ResolvedDependencies {
        logs.Infow("begin to parse dependency", "dependency", dep.FullName)

        cloned := p.copy()
        if err = cloned.ParsePackage(WithPlugins(thisCtx, cloned), dep); err != nil && !errors.Is(err, SkipError{}) {
            return nil, err
        }
    }

    if err = p.ParsePackage(thisCtx, pkg); err != nil {
        return nil, err
    }

    return pkg, nil
}

func (p *Plugins) ParseFile(ctx context.Context, fileName string, fileSys fs.FS) (file *lang.SourceFile, err error) {
    thisCtx := WithPlugins(parser.WithFsCache(ctx, make(parser.FsCache)), p)
    for p.plugin() != nil {
        if psr, ok := p.plugin().(parser.FileParser); ok {
            if file, err = psr.ParseFile(thisCtx, fileName, fileSys); err != nil {
                return nil, err
            }
        }

        if err = ParseSourceFile(p.plugin(), thisCtx, file); err != nil && !errors.Is(err, SkipError{}) {
            return nil, err
        }
        p.Next()
    }

    return file, nil
}

func (p *Plugins) ParseString(content string) (*lang.SourceFile, error) {
    return p.ParseFile(context.Empty(), "", core.StringFS(content))
}

func (p *Plugins) ParsePackage(ctx context.Context, pkg *lang.Package) error {
    if p.parsedPackages[pkg.FullName] {
        logs.Infow("skip when the package has been parsed", "pkg", pkg.FullName)
        return nil
    }

    for p.plugin() != nil {
        if err := ParsePackage(p.plugin(), ctx, pkg); err != nil && !errors.Is(err, SkipError{}) {
            return err
        }

        if err := CompilePackage(p.plugin(), ctx, pkg); err != nil && !errors.Is(err, SkipError{}) {
            return err
        }

        p.Next()
    }

    if !pkg.IsPadding() {
        p.parsedPackages[pkg.FullName] = true
    }
    return nil
}

func (p *Plugins) CompilePackage(ctx context.Context, pkg *lang.Package) error {
    for p.plugin() != nil {
        if err := CompilePackage(p.plugin(), ctx, pkg); err != nil && !errors.Is(err, SkipError{}) {
            return err
        }

        p.Next()
    }
    return nil
}

type plugins []Plugin

func (p plugins) Len() int      { return len(p) }
func (p plugins) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p plugins) Less(i, j int) bool {
    if p[i].GetGroupPriority() == p[j].GetGroupPriority() {
        return p[i].GetPriority() < p[j].GetPriority()
    } else {
        return p[i].GetGroupPriority() < p[j].GetGroupPriority()
    }
}
