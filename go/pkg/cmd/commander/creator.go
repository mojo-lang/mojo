package commander

import (
	"fmt"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/cmd/create/scaffolding"
	"github.com/mojo-lang/mojo/go/pkg/cmd/create/scaffolding/types"
	"github.com/mojo-lang/mojo/go/pkg/mojo/util"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/handlers"
	"os/exec"
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
	RunImmediately bool
}

func (c *Creator) Execute() error {
	if err := c.createScaffolding(); err != nil {
		return err
	}

	if c.HelloWorldMode {
		output := util.GetAbsolutePath(c.Pwd, c.Output)
		helloWorldRoot := path2.Join(output, scaffolding.GetPackageDirName(c.Data))
		if err := c.CompileHelloWorld(helloWorldRoot); err != nil {
			return err
		}

		if c.RunImmediately {
			c.RunHelloWorld(path2.Join(helloWorldRoot, "service-go"))
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
		output := util.GetAbsolutePath(c.Pwd, c.Output)
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

func (c *Creator) RunHelloWorld(root string) error {
	cmd := exec.Command("go", "run", "cmd/hello-world-server/main.go")
	cmd.Dir = root

	logs.Debug("begin to running go run")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout
	if err = cmd.Start(); err != nil {
		return err
	}

	for {
		output := make([]byte, 1024)
		_, err = stdout.Read(output)
		fmt.Print(string(output))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
