package cmd

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/create/scaffolding"
	"github.com/mojo-lang/mojo/go/pkg/mojo/create/scaffolding/types"
	"github.com/urfave/cli/v2"
	"strings"
)

type CreateCmd struct {
	BaseCmd

	PackageName  string
	Repository   string // the git repository for the generated code
	Version      string
	License      string
	Author       string
	Email        string
	Organization string

	Output string

	pwd string
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
		pwd: getPwd(),
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

	// create the mojo scaffolding code
	if err := b.createScaffolding(); err != nil {
		return err
	}

	return nil
}

func getName(fullName string) string {
	segments := strings.Split(fullName, ".")
	if len(segments) > 0 {
		return segments[len(segments)-1]
	}
	return ""
}

func (b *CreateCmd) createScaffolding() (err error) {
	data := &scaffolding.Data{Package: &types.MojoPackage{
		Name:         getName(b.PackageName),
		FullName:     b.PackageName,
		Version:      b.Version,
		License:      b.License,
		Repository:   b.Repository,
		Author:       b.Author,
		Email:        b.Email,
		Organization: b.Organization,
	}}

	g := &scaffolding.Generator{}
	return g.Generate(data, b.Output)
}
