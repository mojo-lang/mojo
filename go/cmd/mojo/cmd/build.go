package cmd

import (
	"fmt"
	"github.com/mojo-lang/mojo/go/pkg/cmd/commander"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

type BuildCmd struct {
	BaseCmd
	commander.Builder
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
		Builder: commander.Builder{
			Pwd: getPwd(),
		},
	}
}

func (b *BuildCmd) Build() {
	b.BaseCmd.Command.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Usage:       "the mojo package root path to compile",
			Destination: &b.Path,
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
			Name:        "targets",
			Aliases:     []string{"t"},
			Usage:       "the target to generate",
			Destination: &b.Targets,
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
			Usage:       "the git repository of NCraft related code",
			Destination: &b.Repository,
		},
	}

	b.BaseCmd.Command.Action = b.Execute
}

func (b *BuildCmd) Execute(ctx *cli.Context) error {
	if ctx.NArg() > 0 {
		b.Path = ctx.Args().Get(0)
		if strings.HasPrefix(b.Path, "--") {
			return fmt.Errorf("failed to parse path from commandline, path: %s", b.Path)
		}
	}
	if ctx.NArg() > 1 {
		b.Output = ctx.Args().Get(1)
		if strings.HasPrefix(b.Path, "--") {
			return fmt.Errorf("failed to parse output from commandline, output: %s", b.Output)
		}
	}
	return b.Builder.Execute()
}
