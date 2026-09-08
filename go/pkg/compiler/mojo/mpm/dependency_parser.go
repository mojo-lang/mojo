package mpm

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/plugin"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

const pluginName = "mpm.dependency-parser"

var errNotPackageDeclaration = errors.New("expected a package declaration")

func init() {
	plugin.RegisterPlugin(NewDependencyParser(nil))
}

type DependencyParser struct {
	plugin.BasicPlugin

	parsedPackages map[string]*lang.Package
	localMojoRoot  string
	manifests      map[string][]*lang.Package
	resolving      map[string]bool
}

func NewDependencyParser(options core.Options) *DependencyParser {
	return &DependencyParser{
		BasicPlugin: plugin.BasicPlugin{
			Name:          pluginName,
			Group:         "mpm",
			GroupPriority: 0,
			Priority:      1,
			Creator: func(options core.Options) plugin.Plugin {
				return NewDependencyParser(options)
			},
		},
		parsedPackages: make(map[string]*lang.Package),
		manifests:      make(map[string][]*lang.Package),
		resolving:      make(map[string]bool),
	}
}

// ParseFile
// TODO implement the imports
func (p *DependencyParser) ParseFile(ctx context.Context, fileName string) (*lang.SourceFile, error) {
	_ = ctx
	return nil, nil
}

func (p *DependencyParser) ParsePackage(ctx context.Context, pkg *lang.Package) error {
	if util.IsPackageProcessed(pkg, pluginName) {
		return nil
	}
	logs.Infow("enter the plugin", "plugin", p.Name, "method", "ParsePackage", "pkg", pkg.FullName)

	if plugins := plugin.ContextPlugins(ctx); plugins != nil {
		for _, dep := range pkg.ResolvedDependencies {
			logs.Infow("begin to parse mojo dependency", "dependency", dep.FullName)

			cloned := plugins.Copy()
			if err := cloned.ParsePackage(plugin.WithPackageName(plugin.WithPlugins(ctx, cloned), dep.FullName), dep); err != nil && !core.IsSkipError(err) {
				return err
			}
		}
		if plugins.Next() != nil {
			if err := plugins.ParsePackage(ctx, pkg); err != nil {
				return err
			}
		}
	}

	util.SetPackageProcessed(pkg, pluginName)
	return nil
}

func (p *DependencyParser) ParsePath(ctx context.Context, pkgPath string) (*lang.Package, error) {
	workingDir := plugin.ContextWorkingDir(ctx)
	logs.Infow("enter the plugin", "plugin", p.Name, "method", "ParsePackagePath", "workingDir", workingDir, "path", pkgPath)

	fullPath, err := filepath.Abs(util.GetAbsolutePath(workingDir, pkgPath))
	if err != nil {
		return nil, err
	}
	return p.loadPath(ctx, fullPath, plugin.ContextPackageName(ctx))
}

// ReadPackageDeclarations reads every declaration, preserving its metadata and order.
func ReadPackageDeclarations(ctx context.Context, dir string) ([]*lang.Package, error) {
	file, err := plugin.NewPlugins("syntax").ParseFile(ctx, path.Join(dir, "package.mojo"))
	if err != nil {
		return nil, err
	}
	var packages []*lang.Package
	seen := make(map[string]bool)
	for _, statement := range file.Statements {
		decl := statement.GetDeclaration().GetPackageDecl()
		if decl == nil || decl.Package == nil {
			return nil, fmt.Errorf("%s: %w", file.FullName, errNotPackageDeclaration)
		}
		pkg := decl.Package
		if seen[pkg.FullName] {
			return nil, fmt.Errorf("%s: duplicate package %s", file.FullName, pkg.FullName)
		}
		seen[pkg.FullName] = true
		pkg.SetExtraString("path", dir)
		pkg.ResolvedDependencies = make(map[string]*lang.Package)
		packages = append(packages, pkg)
	}
	if len(packages) == 0 {
		return nil, fmt.Errorf("%s: no package declarations", file.FullName)
	}
	return packages, nil
}

func (p *DependencyParser) loadPath(ctx context.Context, requested, name string) (*lang.Package, error) {

	dir := requested
	var packages []*lang.Package
	var sourceFileError error
	for {
		if cached, ok := p.manifests[dir]; ok {
			packages = cached
			break
		}
		if info, err := os.Stat(filepath.Join(dir, "package.mojo")); err == nil && !info.IsDir() {
			packages, err = ReadPackageDeclarations(ctx, dir)
			if err == nil {
				p.manifests[dir] = packages
				break
			}
			// An ordinary source file can also be named package.mojo (mojo.lang.Package).
			// It must not hide the project manifest in an ancestor directory.
			if !errors.Is(err, errNotPackageDeclaration) {
				return nil, err
			}
			sourceFileError = err
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			if sourceFileError != nil {
				return nil, sourceFileError
			}
			return nil, fmt.Errorf("no package.mojo found for %s", requested)
		}
		dir = parent
	}

	if name == "" && requested != dir {
		rel, _ := filepath.Rel(dir, requested)
		name = strings.ReplaceAll(filepath.ToSlash(rel), "/", ".")
		for _, pkg := range packages {
			if "mojo."+pkg.FullName == name {
				name = pkg.FullName
				break
			}
		}
	}
	if name != "" {
		for _, pkg := range packages {
			if pkg.FullName == name {
				return p.resolvePackage(ctx, dir, pkg)
			}
		}
		return nil, fmt.Errorf("package %s is not declared in %s", name, filepath.Join(dir, "package.mojo"))
	}
	if len(packages) == 1 {
		return p.resolvePackage(ctx, dir, packages[0])
	}
	set := &lang.Package{Children: packages, Implicit: true}
	set.SetExtraBool("package-set", true)
	set.SetExtraString("path", dir)
	for _, pkg := range packages {
		if _, err := p.resolvePackage(ctx, dir, pkg); err != nil {
			return nil, err
		}
	}
	return set, nil
}

func (p *DependencyParser) resolvePackage(ctx context.Context, fullPath string, pkg *lang.Package) (*lang.Package, error) {
	key := fullPath + "#" + pkg.FullName
	if cached := p.parsedPackages[key]; cached != nil {
		return cached, nil
	}
	if p.resolving[key] {
		return nil, fmt.Errorf("cyclic package dependency involving %s", pkg.FullName)
	}
	p.resolving[key] = true
	defer delete(p.resolving, key)
	if p.localMojoRoot == "" && pkg.GoModName() == lang.MojoGoModule {
		p.localMojoRoot = util.MojoRepositoryRoot(fullPath)
	}
	var err error

	// parse the dependency
	includedMojoPkg := false
	names := make([]string, 0, len(pkg.Dependencies))
	for name := range pkg.Dependencies {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		d := pkg.Dependencies[name]
		if strings.HasPrefix(name, "mojo.") {
			includedMojoPkg = true
		}
		if d.Path == "" {
			var sibling *lang.Package
			for _, candidate := range p.manifests[fullPath] {
				if candidate.FullName == name {
					sibling = candidate
					break
				}
			}
			if sibling != nil {
				dep, err := p.resolvePackage(ctx, fullPath, sibling)
				if err != nil {
					return nil, err
				}
				pkg.ResolvedDependencies[name] = dep
				continue
			}
		}
		depPath := d.Path
		if depPath == "" && p.localMojoRoot != "" && strings.HasPrefix(name, "mojo.") {
			depPath = util.MojoPackagePath(p.localMojoRoot, strings.TrimPrefix(name, "mojo."))
		}
		if depPath == "" && strings.HasPrefix(name, "mojo.") {
			depPkg := GetMojoPackage(name)
			if depPkg == nil {
				return nil, fmt.Errorf("failed to found the required package %s", name)
			}
			pkg.ResolvedDependencies[depPkg.FullName] = depPkg
			continue
		}

		if len(depPath) == 0 {
			depPath, err = GetPackageCenter().Get(name, d)
			if err != nil {
				return nil, err
			}
		}

		depPath = util.GetAbsolutePath(fullPath, depPath)
		depPkg, err := p.loadPath(ctx, depPath, name)
		if err != nil {
			return nil, err
		}

		pkg.ResolvedDependencies[depPkg.FullName] = depPkg
	}

	// add mojo default packages
	if !strings.HasPrefix(pkg.FullName, "mojo.") {
		if pkg.ResolvedDependencies == nil {
			pkg.ResolvedDependencies = make(map[string]*lang.Package)
		}

		if includedMojoPkg {
			if _, ok := pkg.ResolvedDependencies["mojo.core"]; !ok {
				corePkg := GetMojoPackage("mojo.core")
				pkg.ResolvedDependencies[corePkg.FullName] = corePkg
			}
		} else {
			mojoPkgs := GetMojoPackages()
			for _, mp := range mojoPkgs {
				pkg.ResolvedDependencies[mp.FullName] = mp
			}
		}
	}

	p.parsedPackages[key] = pkg
	return pkg, nil
}
