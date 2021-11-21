package scaffolding

import (
	"github.com/mojo-lang/mojo/go/pkg/util"
	"io/ioutil"
	path2 "path"
)

const HelloWorldPackageName = "hello_world"
const HelloWorldOrg = "mojo-lang.org"
const HelloWorldRepo = "git.company.com/examples/hello-world"

func GenerateFiles(output string) error {
	files := []*util.GeneratedFile{{
		Name:    path2.Join(output, "hello-world/mojo/hello-world/v1/echo.mojo"),
		Content: helloWorldEcho,
	}, {
		Name:    path2.Join(output, "hello-world/mojo/hello-world/v1/hello_world.mojo"),
		Content: helloWorldService,
	}}

	for _, f := range files {
		path := path2.Dir(f.Name)
		if err := util.CreateDir(path); err != nil {
			return err
		}
		if err := ioutil.WriteFile(f.Name, []byte(f.Content), 0666); err != nil {
			return err
		}
	}
	return nil
}

func CompileHelloWorld(output string) error {
	return nil
}
