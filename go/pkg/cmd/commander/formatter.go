package commander

import (
    "github.com/mojo-lang/mojo/go/pkg/cmd/format"
    "github.com/urfave/cli/v2"
)

type Formatter struct {
    WorkingDir string
    Path       string
    Output     string

    BackupSource  bool
    TargetFiles   cli.StringSlice
    TargetPackage string
}

func (f *Formatter) Execute() error {
    formatter := &format.Formatter{
        WorkingDir: f.WorkingDir,
        Path:       f.Path,
    }
    return formatter.Format()
}
