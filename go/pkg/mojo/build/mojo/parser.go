package mojo

import (
	"errors"
	"fmt"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/builder"
	pkg2 "github.com/mojo-lang/mojo/go/pkg/mojo/pkg"
	"github.com/mojo-lang/mojo/go/pkg/parser"
	path2 "path"
	"sort"
	"strings"
)

type Parser struct {
	Root       *lang.Package
	WorkingDir string
}

func NewParser(pwd string) *Parser {
	return &Parser{
		WorkingDir: pwd,
	}
}

func (p *Parser) ParseDependency(path string) error {
	// parse the mojo package
	err := p.parseDependency(path, nil)
	if err != nil {
		return err
	}

	tree := NewDependency(p.Root)
	return tree.Resolve()
}

func (p *Parser) parseDependency(path string, dependent *lang.Package) error {
	// parse the mojo package
	pkg, err := pkg2.ParsePackageFile(path)
	if err != nil {
		return err
	}

	path = builder.GetAbsolutePath(p.WorkingDir, path)
	pkg.SetExtraString("path", path)

	if dependent == nil {
		p.Root = pkg
	} else {
		if dependent.ResolvedDependencies == nil {
			dependent.ResolvedDependencies = make(map[string]*lang.Package)
		}
		/// TODO: may add version info to the key
		dependent.ResolvedDependencies[pkg.FullName] = pkg
	}

	// parse the dependency
	for _, d := range pkg.Dependencies {
		///TODO: resolve the requirement and download the package if needed
		depPath := builder.GetAbsolutePath(path, d.Path)
		err = p.parseDependency(depPath, pkg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) Parse(path string) error {
	err := p.ParseDependency(path)
	if err != nil {
		return err
	}
	return p.parse(p.Root)
}

func (p *Parser) parse(pkg *lang.Package) error {
	for _, dependency := range pkg.ResolvedDependencies {
		if dependency.GetExtraBool("parsed") {
			continue
		}

		if err := p.parse(dependency); err != nil {
			return err
		}
	}

	parser := parser.New()
	parser.Dependencies = pkg.ResolvedDependencies

	path := pkg.GetExtraString("path")

	if !strings.HasPrefix(pkg.FullName, "mojo.") {
		path = path2.Join(path, "mojo")
	}
	packages, err := parser.ParsePackage(path, pkg.FullName)
	if err != nil {
		return err
	}

	for _, pg := range packages {
		pg.Authors = pkg.Authors
		pg.Repository = pkg.Repository
		pg.Version = pkg.Version
	}

	root, err := p.TreePackages(packages)
	if err != nil {
		logs.Errorw("failed to TreePackages into one root", "error", err.Error(), "package", pkg.FullName)
		return errors.New("failed to TreePackages into one root")
	}

	pkg.Children = root.Children
	pkg.Scope = root.Scope
	pkg.SourceFiles = root.SourceFiles
	pkg.SetExtraBool("parsed", true)

	return nil
}

func (p *Parser) TreePackages(packages map[string]*lang.Package) (*lang.Package, error) {
	packagesCount := len(packages)
	if packagesCount == 0 {
		return nil, errors.New("failed to TreePackages because of empty packages")
	}

	if packagesCount == 1 {
		for _, v := range packages {
			return v, nil
		}
	}

	var names []string
	for key, _ := range packages {
		names = append(names, key)
	}
	sort.Strings(names)

	roots := []string{names[0]}
	for i := 1; i < len(names); i++ {
		root := roots[len(roots)-1]
		if names[i] == root || names[i] == names[i-1] {
			return nil, errors.New(fmt.Sprint("found duplicated two packages: ", root))
		}
		if !strings.HasPrefix(names[i], roots[0]) {
			return nil, errors.New(fmt.Sprint("found duplicated two root packages: ", roots[0], " and ", names[i]))
		}

		if strings.HasPrefix(names[i], names[i-1]) {
			if names[i-1] != root {
				root = names[i-1]
				roots = append(roots, root)
			}
			parent := packages[root]
			child := packages[names[i]]
			if !parent.HasChild(child.FullName) {
				parent.Children = append(parent.Children, child)
			}
		} else {
			for j := len(roots) - 1; j >= 0; j-- {
				if strings.HasPrefix(names[i], roots[j]) {
					parent := packages[roots[j]]
					child := packages[names[i]]
					if !parent.HasChild(child.FullName) {
						parent.Children = append(parent.Children, child)
					}
					roots = roots[0 : j+1]
					break
				}
			}
		}
	}

	return packages[names[0]], nil
}


