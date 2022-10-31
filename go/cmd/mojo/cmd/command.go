package cmd

import (
	"github.com/urfave/cli/v2"
)

type Cmd interface {
	GetCmd() *BaseCmd
}

type Cmds []Cmd

func (c Cmds) Len() int {
	return len(c)
}
func (c Cmds) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Cmds) Less(i, j int) bool {
	return c[i].GetCmd().Name < c[j].GetCmd().Name
}

var commands []Cmd

type BaseCmd struct {
	*cli.Command

	source      string
	baseURL     string
	environment string

	buildWatch bool

	// Profile flags (for debugging of performance problems)
	// cpuprofile   string
	// memprofile   string
	// mutexprofile string
	// traceprofile string
	// printm       bool

	// TODO(bep) var vs string
	mode string // verbose, quiet

	cfgFile string
	cfgDir  string
	logFile string
}

func (b *BaseCmd) GetCmd() *BaseCmd {
	return b
}
