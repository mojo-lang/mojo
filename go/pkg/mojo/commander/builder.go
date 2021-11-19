package commander

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/builder"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/document"
	_go "github.com/mojo-lang/mojo/go/pkg/mojo/build/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/java"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/mojo"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/ncraft/gokit"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/openapi"
	"github.com/mojo-lang/mojo/go/pkg/mojo/build/protobuf"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
	"strings"
)

type Builder struct {
	Package  *lang.Package
	Files    []*descriptor.FileDescriptor
	OpenAPIs *openapi.OpenAPIs

	PackageName string

	Targets string
	Engine  string

	Output string

	Pwd  string
	Path string

	APIEnabled bool

	NcraftAllEnabled     bool
	NcraftServiceEnabled bool
	NcraftClientEnabled  bool
	NcraftSidecarEnabled bool

	// the git repository for the generated code
	Repository string
}

func (b *Builder) Execute() error {
	if len(b.Path) == 0 {
		b.Path = "./"
	}
	if len(b.Targets) == 0 {
		b.Targets = "api"
	}

	targets := strings.Split(b.Targets, ",")
	for _, target := range targets {
		switch target {
		case "api":
			b.APIEnabled = true
		case "ncraft":
			b.NcraftAllEnabled = true
		case "ncraft.service", "service":
			b.NcraftServiceEnabled = true
		case "ncraft.client", "client":
			b.NcraftClientEnabled = true
		case "ncraft.sidecar", "sidecar":
			b.NcraftSidecarEnabled = true
		}
	}

	// build the mojo package
	if err := b.buildMojo(); err != nil {
		return err
	}

	// compile the target package to openapi
	if err := b.buildOpenapi(); err != nil {
		return err
	}

	// compile the target package to document
	if err := b.buildDocument(); err != nil {
		return err
	}

	// compile the target package to protobuf & generate the protobuf files
	if err := b.buildProtobuf(); err != nil {
		return err
	}

	// generate the target package to golang
	if err := b.buildGo(); err != nil {
		return err
	}

	// generate the target package to java
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
	}
	if b.NcraftAllEnabled || b.NcraftSidecarEnabled {
	}

	return nil
}

func (b *Builder) buildMojo() (err error) {
	b.Package, err = mojo.Builder{
		Builder: builder.Builder{
			PWD:  b.Pwd,
			Path: b.Path,
		},
	}.Build()
	return err
}

func (b *Builder) buildProtobuf() (err error) {
	b.Files, err = protobuf.Builder{
		Builder: builder.Builder{
			PWD:        b.Pwd,
			Path:       b.Path,
			Package:    b.Package,
			APIEnabled: b.APIEnabled,
		},
		Output: b.Output,
	}.Build()
	return err
}

func (b *Builder) buildGo() error {
	return _go.Builder{
		Builder: builder.Builder{
			PWD:        b.Pwd,
			Path:       b.Path,
			Package:    b.Package,
			APIEnabled: b.APIEnabled,
		},
		Output: b.Output,
		Files:  b.Files,
	}.Build()
}

func (b *Builder) buildJava() error {
	return java.Builder{
		Builder: builder.Builder{
			PWD:        b.Pwd,
			Path:       b.Path,
			Package:    b.Package,
			APIEnabled: b.APIEnabled,
		},
		Output: b.Output,
		Files:  b.Files,
	}.Build()
}

func (b *Builder) buildOpenapi() (err error) {
	b.OpenAPIs, err = openapi.Builder{
		Builder: builder.Builder{
			PWD:        b.Pwd,
			Path:       b.Path,
			Package:    b.Package,
			APIEnabled: b.APIEnabled,
		},
		Output: b.Output,
	}.Build()
	return err
}

func (b *Builder) buildDocument() error {
	return document.Builder{
		Builder: builder.Builder{
			PWD:        b.Pwd,
			Path:       b.Path,
			Package:    b.Package,
			APIEnabled: b.APIEnabled,
		},
		Output:   b.Output,
		OpenAPIs: b.OpenAPIs,
	}.Build()
}

func (b *Builder) buildGokit(ncraftType string) error {
	return gokit.Builder{
		Builder: builder.Builder{
			PWD:     b.Pwd,
			Path:    b.Path,
			Package: b.Package,
		},
		Type:       ncraftType,
		Output:     b.Output,
		Repository: b.Repository,
	}.Build()
}
