package commander

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/builder"
	"github.com/mojo-lang/mojo/go/pkg/mojo/create/scaffolding"
	"github.com/mojo-lang/mojo/go/pkg/mojo/create/scaffolding/types"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/handlers"
	path2 "path"
)

type Creator struct {
	PackageName  string
	Repository   string // the git repository for the generated code
	Version      string
	License      string
	Author       string
	Email        string
	Organization string

	Data   *scaffolding.Data
	Output string

	Pwd string

	HelloWorldMode bool
}

func (c *Creator) Execute() error {
	if err := c.createScaffolding(); err != nil {
		return err
	}

	if c.HelloWorldMode {
		output := builder.GetAbsolutePath(c.Pwd, c.Output)
		helloWorldRoot := path2.Join(output, scaffolding.GetPackageDirName(c.Data))
		if err := c.CompileHelloWorld(helloWorldRoot); err != nil {
			return err
		}
	}

	return nil
}

func (c *Creator) createScaffolding() error {
	c.Data = &scaffolding.Data{Package: &types.MojoPackage{
		Name:         lang.GetPackageName(c.PackageName),
		FullName:     c.PackageName,
		Version:      c.Version,
		License:      c.License,
		Repository:   c.Repository,
		Author:       c.Author,
		Email:        c.Email,
		Organization: c.Organization,
	}}

	g := &scaffolding.Generator{}
	err := g.Generate(c.Data, c.Output)
	if err != nil {
		return err
	}

	if c.HelloWorldMode {
		output := builder.GetAbsolutePath(c.Pwd, c.Output)
		if err = scaffolding.GenerateHelloWorldFiles(output); err != nil {
			return err
		}
	}

	return nil
}

func (c *Creator) CompileHelloWorld(root string) error {
	builder := &Builder{
		PackageName: scaffolding.HelloWorldPackageName,
		Targets:     "api,service",
		Pwd:         root,
	}

	handlers.ResetHandlerMethods(scaffolding.HandlerMethod)
	return builder.Execute()
}
