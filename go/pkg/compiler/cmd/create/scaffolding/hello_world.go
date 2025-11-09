package scaffolding

import (
	"io/ioutil"
	path2 "path"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"

	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

const HelloWorldPackageName = "hello_world"
const HelloWorldOrg = "mojo-lang.org"
const HelloWorldRepo = "git.company.com/examples/hello-world"

const HandlerMethod = `
func (s helloWorld) GetEcho(ctx context.Context, in *pb.GetEchoRequest) (*pb.Echo, error) {
	resp := pb.Echo{
		Name:    in.Name,
		Message: "Hello, " + in.Name + "!",
	}
	return &resp, nil
}
`

func GenerateHelloWorldFiles(output string) error {
	files := []*util.GeneratedFile{{
		Name:    path2.Join(output, "hello-world/mojo/hello-world/v1/echo.mojo"),
		Content: helloWorldEcho,
	}, {
		Name:    path2.Join(output, "hello-world/mojo/hello-world/v1/hello_world.mojo"),
		Content: helloWorldService,
	}}

	for _, f := range files {
		path := path2.Dir(f.Name)
		if err := core.CreateDir(path); err != nil {
			return err
		}
		if err := ioutil.WriteFile(f.Name, []byte(f.Content), 0666); err != nil {
			return err
		}
	}
	return nil
}
