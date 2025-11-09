package format

import (
	"path"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/compiler/mojo/printer"
	"github.com/mojo-lang/mojo/go/pkg/compiler/plugin"
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
	pkg, err := plugins.ParsePath(context.Empty(), path.Join(f.WorkingDir, f.Path))
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
		p := printer.New(&printer.Config{}).PrintSourceFile(ctx, file)
		content := p.Buffer.String()
		logs.Info(file.FullName, "\n", content)
	}
	return nil
}
