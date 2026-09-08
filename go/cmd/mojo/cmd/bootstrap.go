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
		Flags:     []cli.Flag{&cli.StringFlag{Name: "targets", Aliases: []string{"t"}, Value: "go", Usage: "languages to regenerate: go, java, or go,java"}},
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() > 1 {
				return fmt.Errorf("bootstrap accepts at most one repository path")
			}
			start := ctx.Args().First()
			if start == "" {
				start = getPwd()
			}
			return commander.Bootstrap(start, ctx.String("targets"))
		},
	}})
}
