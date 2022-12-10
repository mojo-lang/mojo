package mpm

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type Dependency struct {
	Root     *lang.Package
	Packages map[string]*lang.Package
}

func NewDependency(root *lang.Package) *Dependency {
	return &Dependency{
		Root:     root,
		Packages: make(map[string]*lang.Package),
	}
}

func (d *Dependency) Resolve() (*lang.Package, error) {
	if err := d.resolve(d.Root); err != nil {
		return nil, err
	}
	return d.Root, nil
}

func (d *Dependency) resolve(pkg *lang.Package) error {
	for _, dependency := range pkg.ResolvedDependencies {
		p := d.Packages[dependency.FullName]
		if p == nil {
			p = dependency
			d.Packages[dependency.FullName] = dependency
		} else {
			pkg.ResolvedDependencies[dependency.FullName] = p
		}

		if err := d.resolve(dependency); err != nil {
			return err
		}
	}

	return nil
}
