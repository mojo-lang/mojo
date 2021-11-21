package commander

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/create/scaffolding"
	"github.com/mojo-lang/mojo/go/pkg/mojo/create/scaffolding/types"
)

type Creator struct {
	PackageName  string
	Repository   string // the git repository for the generated code
	Version      string
	License      string
	Author       string
	Email        string
	Organization string

	Output string

	Pwd string

	HelloWorldMode bool
}

func (b *Creator) Execute() error {
	if err := b.createScaffolding(); err != nil {
		return err
	}

	if b.HelloWorldMode {
		if err := scaffolding.CompileHelloWorld(b.Output); err != nil {
			return err
		}
	}

	return nil
}

func (b *Creator) createScaffolding() error {
	data := &scaffolding.Data{Package: &types.MojoPackage{
		Name:         lang.GetPackageName(b.PackageName),
		FullName:     b.PackageName,
		Version:      b.Version,
		License:      b.License,
		Repository:   b.Repository,
		Author:       b.Author,
		Email:        b.Email,
		Organization: b.Organization,
	}}

	g := &scaffolding.Generator{}
	err := g.Generate(data, b.Output)
	if err != nil {
		return err
	}

	if b.HelloWorldMode {
		if err = scaffolding.GenerateFiles(b.Output); err != nil {
			return err
		}
	}

	return nil
}
