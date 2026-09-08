package cmd

import (
	"fmt"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/commander"
	"github.com/urfave/cli/v2"
)

func init() {
	commands = append(commands, &BaseCmd{Command: &cli.Command{
		Name:      "bootstrap",
		Usage:     "regenerate Mojo's standard library and embedded packages from local sources",
		ArgsUsage: "[repository path]",
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() > 1 {
				return fmt.Errorf("bootstrap accepts at most one repository path")
			}
			start := ctx.Args().First()
			if start == "" {
				start = getPwd()
			}
			return commander.Bootstrap(start)
		},
	}})
}
