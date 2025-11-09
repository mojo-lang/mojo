package cmd

import (
	"fmt"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/commander"
	"github.com/urfave/cli/v2"
)

type FormatCmd struct {
	BaseCmd
	commander.Formatter
}

func init() {
	cmd := NewFormatCmd()
	cmd.Build()
	commands = append(commands, cmd)
}

func NewFormatCmd() *FormatCmd {
	return &FormatCmd{
		BaseCmd: BaseCmd{
			Command: &cli.Command{
				Name:  "format",
				Usage: "format the mojo code source",
			},
		},
		Formatter: commander.Formatter{
			WorkingDir: getPwd(),
		},
	}
}

func (b *FormatCmd) Build() {
	b.BaseCmd.Command.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "package",
			Aliases:     []string{"p"},
			Usage:       "the mojo package name to compile",
			Destination: &b.TargetPackage,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "the Output to generate",
			Destination: &b.Output,
			DefaultText: ".",
		},
		&cli.StringSliceFlag{
			Name:        "files",
			Aliases:     []string{"f"},
			Usage:       "the version of the package to be created",
			Destination: &b.TargetFiles,
		},
		&cli.BoolFlag{
			Name:        "backup",
			Aliases:     []string{"b"},
			Usage:       "backup the original mojo source file with add .back suffix, otherwise will override it",
			Destination: &b.BackupSource,
		},
	}

	b.BaseCmd.Command.Action = b.Execute
}

func (b *FormatCmd) Execute(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() > 0 {
		b.Path = args.Get(0)
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

	// create the mojo scaffolding code
	if err := b.Formatter.Execute(); err != nil {
		return err
	}

	return nil
}
