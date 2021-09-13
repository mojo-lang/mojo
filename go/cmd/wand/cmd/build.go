package cmd

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	"github.com/mojo-lang/mojo/go/pkg/wand/builder"
	"github.com/mojo-lang/mojo/go/pkg/wand/document"
	_go "github.com/mojo-lang/mojo/go/pkg/wand/go"
	"github.com/mojo-lang/mojo/go/pkg/wand/java"
	"github.com/mojo-lang/mojo/go/pkg/wand/mojo"
	"github.com/mojo-lang/mojo/go/pkg/wand/ncraft/gokit"
	"github.com/mojo-lang/mojo/go/pkg/wand/openapi"
	"github.com/mojo-lang/mojo/go/pkg/wand/protobuf"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

type BuildCmd struct {
	BaseCmd

	Package  *lang.Package
	Files    []*descriptor.FileDescriptor
	OpenAPIs *openapi.OpenAPIs

	PackageName string

	Target string
	Engine string

	Output string

	pwd  string
	path string

	disableMojoGeneration bool

	// the git repository for the generated code
	Repository string
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
		&cli.StringFlag{
			Name:        "engine",
			Aliases:     []string{"e"},
			Usage:       "the ncraft engine to generate",
			Destination: &b.Engine,
		},
		&cli.StringFlag{
			Name:        "target",
			Aliases:     []string{"t"},
			Usage:       "the target to generate",
			Destination: &b.Target,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "the Output to generate",
			Destination: &b.Output,
		},
		&cli.StringFlag{
			Name:        "repo",
			Aliases:     []string{"sr"},
			Usage:       "the git repository of NCraft releated code",
			Destination: &b.Repository,
		},
	}

	b.BaseCmd.Command.Action = b.Execute
}

func (b *BuildCmd) Execute(ctx *cli.Context) error {
	ncraft := false
	if strings.HasPrefix(b.Target, "ncraft") {
		ncraft = true
		b.disableMojoGeneration = true
	}

	// build the mojo package
	if err := b.buildMojo(); err != nil {
		return err
	}

	// compile the target package to openapi
	if err := b.buildOpenapi(); err != nil {
		return err
	}

	// compile the target package to document
	if err := b.buildDocument(); err != nil {
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

	if ncraft {
		switch b.Engine {
		case "gokit":
			if err := b.buildGokit(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *BuildCmd) buildMojo() (err error) {
	b.Package, err = mojo.Builder{
		Builder: builder.Builder{
			PWD:  b.pwd,
			Path: b.path,
		},
	}.Build()
	return err
}

func (b *BuildCmd) buildProtobuf() (err error) {
	b.Files, err = protobuf.Builder{
		Builder: builder.Builder{
			PWD:               b.pwd,
			Path:              b.path,
			Package:           b.Package,
			DisableGeneration: b.disableMojoGeneration,
		},
		Output: b.Output,
	}.Build()
	return err
}

func (b *BuildCmd) buildGo() error {
	return _go.Builder{
		Builder: builder.Builder{
			PWD:               b.pwd,
			Path:              b.path,
			Package:           b.Package,
			DisableGeneration: b.disableMojoGeneration,
		},
		Output: b.Output,
		Files:  b.Files,
	}.Build()
}

func (b *BuildCmd) buildJava() error {
	return java.Builder{
		Builder: builder.Builder{
			PWD:               b.pwd,
			Path:              b.path,
			Package:           b.Package,
			DisableGeneration: b.disableMojoGeneration,
		},
		Output: b.Output,
		Files:  b.Files,
	}.Build()
}

func (b *BuildCmd) buildOpenapi() (err error) {
	b.OpenAPIs, err = openapi.Builder{
		Builder: builder.Builder{
			PWD:               b.pwd,
			Path:              b.path,
			Package:           b.Package,
			DisableGeneration: b.disableMojoGeneration,
		},
		Output: b.Output,
	}.Build()
	return err
}

func (b *BuildCmd) buildDocument() error {
	return document.Builder{
		Builder: builder.Builder{
			PWD:               b.pwd,
			Path:              b.path,
			Package:           b.Package,
			DisableGeneration: b.disableMojoGeneration,
		},
		Output:   b.Output,
		OpenAPIs: b.OpenAPIs,
	}.Build()
}

func (b *BuildCmd) buildGokit() error {
	return gokit.Builder{
		Builder: builder.Builder{
			PWD:     b.pwd,
			Path:    b.path,
			Package: b.Package,
		},
		Type:       strings.TrimPrefix(b.Target, "ncraft."),
		Output:     b.Output,
		Repository: b.Repository,
	}.Build()
}
