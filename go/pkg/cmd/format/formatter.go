package format

import (
	"os"
	"strings"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin/parser"
	"github.com/mojo-lang/mojo/go/pkg/mojo/printer"
)

type Formatter struct {
	WorkingDir string
	Path       string
}

func (f *Formatter) Format() error {
	logs.Infow("begin to parse mojo package.", "pwd", f.WorkingDir, "path", f.Path)

	plugins := plugin.NewPlugins("mpm", "syntax", "semantic")
	if strings.HasPrefix(f.Path, f.WorkingDir) {
		f.Path = strings.TrimPrefix(f.Path, f.WorkingDir)
	}
	pkg, err := plugins.ParsePackagePath(parser.WithWorkingDir(context.Empty(), f.WorkingDir), f.Path, os.DirFS(f.WorkingDir))
	if err != nil {
		return err
	}
	return f.FormatPackage(pkg)
}

func (f *Formatter) FormatPackage(pkg *lang.Package) error {
	if err := f.formatPackage(pkg); err != nil {
		return err
	}

	for _, child := range pkg.Children {
		return f.FormatPackage(child)
	}

	return nil
}

func (f *Formatter) formatPackage(pkg *lang.Package) error {
	ctx := context.WithType(context.Empty(), pkg)
	for _, file := range pkg.SourceFiles {
		sb := strings.Builder{}
		printer := printer.New(printer.Config{}, &sb)
		printer.PrintSourceFile(ctx, file)

		content := sb.String()
		logs.Info(file.FullName, "\n", content)
	}
	return nil
}
