package commander

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/compiler/util"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	api "github.com/mojo-lang/mojo/go/pkg/mojo/openapi"
	"github.com/mojo-lang/mojo/go/pkg/mojo/protobuf/descriptor"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/document"
	_go "github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/go"
	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/java"
	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/mojo"
	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/ncraft/gokit"
	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/openapi"
	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/protobuf"
)

type Builder struct {
	Package  *lang.Package
	Files    []*descriptor.File
	OpenAPIs *api.OpenAPIs

	PackageName string

	Targets string
	Engine  string

	Output string

	// Bootstrap stages complete documentation trees before publishing them.
	documentOutput string
	openapiOutput  string

	Pwd  string
	Path string

	APIEnabled      bool
	GoEnabled       bool
	JavaEnabled     bool
	ProtobufEnabled bool

	NcraftAllEnabled     bool
	NcraftServiceEnabled bool
	NcraftClientEnabled  bool

	// the git repository for the generated code
	Repository string
}

func (b *Builder) Execute() error {
	if len(b.Path) == 0 {
		b.Path = "./"
	}
	if len(b.Targets) == 0 {
		b.Targets = "api"

		if core.IsExist(path.Join(util.GetAbsolutePath(b.Pwd, b.Path), "service-go")) {
			b.Targets = "api,service"
		}
	}

	targets := strings.Split(b.Targets, ",")
	for _, target := range targets {
		switch strings.TrimSpace(target) {
		case "api":
			b.APIEnabled = true
		case "go", "golang":
			b.GoEnabled = true
			b.ProtobufEnabled = true
		case "java":
			b.JavaEnabled = true
			b.ProtobufEnabled = true
		case "protobuf":
			b.ProtobufEnabled = true
		case "ncraft":
			b.NcraftAllEnabled = true
		case "ncraft.service", "service":
			b.NcraftServiceEnabled = true
		case "ncraft.client", "client":
			b.NcraftClientEnabled = true
		default:
			return fmt.Errorf("unsupported build target %q", target)
		}
	}
	if b.Engine != "" && b.Engine != "gokit" {
		return fmt.Errorf("unsupported engine %q; only gokit is supported", b.Engine)
	}

	// build the mojo package
	if err := b.buildMojo(); err != nil {
		return err
	}

	if b.Package.GetExtraBool("package-set") {
		members := make(map[string]*lang.Package)
		for _, pkg := range b.Package.Children {
			members[pkg.FullName] = pkg
		}
		built := make(map[string]bool)
		var build func(*lang.Package) error
		build = func(pkg *lang.Package) error {
			if built[pkg.FullName] {
				return nil
			}
			built[pkg.FullName] = true
			var names []string
			for name := range pkg.ResolvedDependencies {
				names = append(names, name)
			}
			sort.Strings(names)
			for _, name := range names {
				if dep := members[name]; dep != nil && dep == pkg.ResolvedDependencies[name] {
					if err := build(dep); err != nil {
						return err
					}
				}
			}
			child := *b
			// Generators mutate the AST (for example protobuf lowers generic types).
			// Give each package its own compilation pipeline and dependency ASTs.
			child.Package = nil
			child.PackageName = pkg.FullName
			if err := child.Execute(); err != nil {
				return fmt.Errorf("build %s: %w", pkg.FullName, err)
			}
			return nil
		}
		for _, pkg := range b.Package.Children {
			if err := build(pkg); err != nil {
				return err
			}
		}
		return nil
	}

	return b.buildOutputs()
}

func (b *Builder) buildOutputs() error {

	if b.APIEnabled || b.NcraftAllEnabled || b.NcraftServiceEnabled || b.NcraftClientEnabled {
		if err := b.buildOpenapi(); err != nil {
			return err
		}
		if err := b.buildDocument(); err != nil {
			return err
		}
	}

	// compile the target package to protobuf & generate the protobuf files
	if err := b.buildProtobuf(); err != nil {
		return err
	}

	// generate the target package to golang
	if err := b.buildGo(); err != nil {
		return err
	}
	if err := b.buildJava(); err != nil {
		return err
	}

	// compile the resource to sql orm file (including sql script, create table)
	if b.NcraftAllEnabled || b.NcraftServiceEnabled {
		if len(b.Engine) == 0 {
			b.Engine = "gokit"
		}

		switch b.Engine {
		case "gokit":
			if err := b.buildGokit("service"); err != nil {
				return err
			}
		}
	}
	if b.NcraftAllEnabled || b.NcraftClientEnabled {
		if err := b.buildGokit("client"); err != nil {
			return err
		}
	}

	return nil
}

func (b *Builder) buildMojo() (err error) {
	b.Package, err = mojo.Builder{
		PackageName: b.PackageName,
		Builder: builder.Builder{
			PWD:  b.Pwd,
			Path: b.Path,
		},
	}.Build()
	return err
}

func (b *Builder) buildProtobuf() (err error) {
	output := ""
	if !b.APIEnabled && !b.GoEnabled && !b.JavaEnabled && b.Output != "" {
		output = util.GetAbsolutePath(b.Pwd, b.Output)
	}
	b.Files, err = protobuf.Builder{
		Builder: builder.Builder{
			PWD:        b.Pwd,
			Path:       b.Path,
			Package:    b.Package,
			APIEnabled: b.APIEnabled || b.ProtobufEnabled,
		},
		// Protoc reads the package's protobuf directory. --output selects the
		// final Go destination when building Go code.
		Output: output,
	}.Build()
	return err
}

func (b *Builder) buildJava() error {
	return java.Builder{
		Builder: builder.Builder{PWD: b.Pwd, Path: b.Path, Package: b.Package, APIEnabled: b.JavaEnabled},
		Output:  b.Output,
		Files:   b.Files,
	}.Build()
}

func (b *Builder) buildGo() error {
	return _go.Builder{
		Builder: builder.Builder{
			PWD:        b.Pwd,
			Path:       b.Path,
			Package:    b.Package,
			APIEnabled: b.APIEnabled || b.GoEnabled,
		},
		Output: b.Output,
		Files:  b.Files,
	}.Build()
}

func (b *Builder) buildOpenapi() (err error) {
	output := b.Output
	if b.openapiOutput != "" {
		output = b.openapiOutput
	}
	b.OpenAPIs, err = openapi.Builder{
		Builder: builder.Builder{
			PWD:        b.Pwd,
			Path:       b.Path,
			Package:    b.Package,
			APIEnabled: b.APIEnabled,
		},
		Output: output,
	}.Build()
	return err
}

func (b *Builder) buildDocument() error {
	output := b.Output
	if b.documentOutput != "" {
		output = b.documentOutput
	}
	return document.Builder{
		Builder: builder.Builder{
			PWD:        b.Pwd,
			Path:       b.Path,
			Package:    b.Package,
			APIEnabled: b.APIEnabled,
		},
		Output:   output,
		OpenAPIs: b.OpenAPIs,
	}.Build()
}

func (b *Builder) buildGokit(ncraftType string) error {
	return gokit.Builder{
		Builder: builder.Builder{
			PWD:        b.Pwd,
			Path:       b.Path,
			Package:    b.Package,
			APIEnabled: b.APIEnabled,
		},
		Type:       ncraftType,
		Output:     b.Output,
		Repository: b.Repository,
	}.Build()
}
