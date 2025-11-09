package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/commander"
	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/create/scaffolding"
)

type CreateCmd struct {
	BaseCmd
	commander.Creator
}

func init() {
	cmd := NewCreateCmd()
	cmd.Build()
	commands = append(commands, cmd)
}

func NewCreateCmd() *CreateCmd {
	return &CreateCmd{
		BaseCmd: BaseCmd{
			Command: &cli.Command{
				Name:  "create",
				Usage: "create the scaffolding code for mojo module",
			},
		},
		Creator: commander.Creator{
			Pwd: getPwd(),
		},
	}
}

func (b *CreateCmd) Build() {
	b.BaseCmd.Command.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "package",
			Aliases:     []string{"p"},
			Usage:       "the mojo package name to compile",
			Destination: &b.PackageName,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "the Output to generate",
			Destination: &b.Output,
			DefaultText: ".",
		},
		&cli.StringFlag{
			Name:        "repo",
			Aliases:     []string{"r"},
			Usage:       "the git repository of NCraft releated code",
			Destination: &b.Repository,
		},
		&cli.StringFlag{
			Name:        "version",
			Aliases:     []string{"v"},
			Usage:       "the version of the package to be created",
			Destination: &b.Version,
			DefaultText: "0.1.0",
		},
		&cli.StringFlag{
			Name:        "license",
			Aliases:     []string{"l"},
			Usage:       "the license of the package to be created",
			Destination: &b.License,
		},
		&cli.StringFlag{
			Name:        "Author",
			Aliases:     []string{"a"},
			Usage:       "the license of the package to be created",
			Destination: &b.Author,
			DefaultText: "YOUR NAME",
		},
		&cli.StringFlag{
			Name:        "email",
			Aliases:     []string{"e"},
			Usage:       "the email of the author of the package to be created",
			Destination: &b.Email,
			DefaultText: "YOUR EMAIL",
		},
		&cli.StringFlag{
			Name:        "organization",
			Aliases:     []string{"org"},
			Usage:       "the organization of the package to be created, for generating java code",
			Destination: &b.Organization,
			DefaultText: "YOUR ORGANIZATION (like: mojolang.org)",
		},
		&cli.BoolFlag{
			Name:        "run",
			Aliases:     []string{},
			Usage:       "run the hello world server immediately",
			Destination: &b.RunImmediately,
		},
	}

	b.BaseCmd.Command.Action = b.Execute
}

func (b *CreateCmd) Execute(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() > 0 {
		b.PackageName = args.Get(0)
	}
	if args.Len() > 1 {
		b.Output = args.Get(1)
	}

	if len(b.PackageName) == 0 {
		b.PackageName = scaffolding.HelloWorldPackageName
		b.Organization = scaffolding.HelloWorldOrg
		b.Repository = scaffolding.HelloWorldRepo
		b.HelloWorldMode = true
	}

	if len(b.Output) == 0 {
		b.Output = "./"
	}

	// create the mojo scaffolding code
	if err := b.Creator.Execute(); err != nil {
		return err
	}

	return nil
}
