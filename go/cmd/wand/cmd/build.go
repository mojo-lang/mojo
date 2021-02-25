package cmd

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	"github.com/mojo-lang/mojo/go/pkg/wand/document"
	_go "github.com/mojo-lang/mojo/go/pkg/wand/go"
	"github.com/mojo-lang/mojo/go/pkg/wand/java"
	"github.com/mojo-lang/mojo/go/pkg/wand/mojo"
	"github.com/mojo-lang/mojo/go/pkg/wand/openapi"
	"github.com/mojo-lang/mojo/go/pkg/wand/protobuf"
	"github.com/urfave/cli/v2"
	"os"
)

type BuildCmd struct {
	BaseCmd

	Package  *lang.Package
	Files    []*descriptor.FileDescriptor
	OpenAPIs *openapi.OpenAPIs

	PackageName string

	pwd    string
	path   string
	output string
}

func init() {
	cmd := NewBuildCmd()
	cmd.Build()
	commands = append(commands, cmd)
}

func getPwd() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path

	//ex, err := os.Executable()
	//if err != nil {
	//	return ""
	//}
	//return filepath.Dir(ex)
}

func NewBuildCmd() *BuildCmd {
	return &BuildCmd{
		BaseCmd: BaseCmd{
			Command: &cli.Command{
				Name:  "build",
				Usage: "build the mojo source files",
			},
		},
		pwd: getPwd(),
	}
}

func (b *BuildCmd) Build() {
	b.BaseCmd.Command.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Usage:       "the mojo package root path to compile",
			Destination: &b.path,
		},
		&cli.StringFlag{
			Name:        "package",
			Aliases:     []string{"p"},
			Usage:       "the mojo package name to compile",
			Destination: &b.PackageName,
		},
	}

	b.BaseCmd.Command.Action = b.Execute
}

func (b *BuildCmd) Execute(ctx *cli.Context) error {
	// build the mojo package
	if err := b.buildMojo(); err != nil {
		return err
	}

	// compile the target package to protobuf & generate the protobuf files
	if err := b.buildProtobuf(); err != nil {
		return err
	}

	// generate the target package to golang
	if err := b.buildGo(); err != nil {
		return err
	}

	// generate the target package to java
	if err := b.buildJava(); err != nil {
		return err
	}

	// compile the resource to sql orm file (including sql script, create table)

	// compile the target package to openapi
	if err := b.buildOpenapi(); err != nil {
		return err
	}

	// compile the target package to document
	if err := b.buildDocument(); err != nil {
		return err
	}

	return nil
}

func (b *BuildCmd) buildMojo() (err error) {
	b.Package, err = mojo.Builder{
		PWD:  b.pwd,
		Path: b.path,
	}.Build()
	return err
}

func (b *BuildCmd) buildProtobuf() (err error) {
	b.Files, err = protobuf.Builder{
		PWD:     b.pwd,
		Path:    b.path,
		Output:  b.output,
		Package: b.Package,
	}.Build()
	return err
}

func (b *BuildCmd) buildGo() error {
	return _go.Builder{
		PWD:     b.pwd,
		Path:    b.path,
		Output:  b.output,
		Package: b.Package,
		Files:   b.Files,
	}.Build()
}

func (b *BuildCmd) buildJava() error {
	return java.Builder{
		PWD:     b.pwd,
		Path:    b.path,
		Output:  b.output,
		Package: b.Package,
		Files:   b.Files,
	}.Build()
}

func (b *BuildCmd) buildOpenapi() (err error) {
	b.OpenAPIs, err = openapi.Builder{
		PWD:     b.pwd,
		Path:    b.path,
		Output:  b.output,
		Package: b.Package,
	}.Build()
	return err
}

func (b *BuildCmd) buildDocument() error {
	return document.Builder{
		PWD:      b.pwd,
		Path:     b.path,
		Output:   b.output,
		Package:  b.Package,
		OpenAPIs: b.OpenAPIs,
	}.Build()
}
