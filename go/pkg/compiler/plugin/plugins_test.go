package plugin

import (
	"errors"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

type pipelinePlugin struct {
	BasicPlugin
	events               *[]string
	parseErr, compileErr error
}

func (p *pipelinePlugin) ParsePackage(context.Context, *lang.Package) error {
	*p.events = append(*p.events, p.Name+":parse")
	return p.parseErr
}
func (p *pipelinePlugin) CompilePackage(context.Context, *lang.Package) error {
	*p.events = append(*p.events, p.Name+":compile")
	return p.compileErr
}

func TestPluginsIterationAndCopy(t *testing.T) {
	empty := &Plugins{}
	require.Nil(t, empty.Next())
	require.Nil(t, empty.Next())
	first, second := &BasicPlugin{Name: "first"}, &BasicPlugin{Name: "second"}
	p := &Plugins{plugins: plugins{first, second}, parsedPackages: map[string]bool{"done": true}}
	require.Same(t, first, p.plugin())
	require.Same(t, p, p.Next())
	require.Same(t, second, p.plugin())
	copied := p.Copy()
	require.Same(t, first, copied.plugin())
	require.True(t, copied.parsedPackages["done"])
	copied.parsedPackages["shared"] = true
	require.True(t, p.parsedPackages["shared"])
	require.Nil(t, p.Next())
	require.Nil(t, p.plugin())
	require.Nil(t, p.Next())
	require.Same(t, first, copied.plugin())
}

func TestPluginsSortByGroupThenPriority(t *testing.T) {
	a := &BasicPlugin{Name: "a", GroupPriority: 2, Priority: 1}
	b := &BasicPlugin{Name: "b", GroupPriority: 1, Priority: 20}
	c := &BasicPlugin{Name: "c", GroupPriority: 1, Priority: 10}
	p := (&Plugins{plugins: plugins{a, b, c}}).sort()
	require.Equal(t, plugins{c, b, a}, p.plugins)
}

func TestPluginsParsePackageOrderAndCache(t *testing.T) {
	var events []string
	p := &Plugins{plugins: plugins{
		&BasicPlugin{Name: "unsupported"},
		&pipelinePlugin{BasicPlugin: BasicPlugin{Name: "first"}, events: &events},
		&pipelinePlugin{BasicPlugin: BasicPlugin{Name: "second"}, events: &events},
	}, parsedPackages: make(map[string]bool)}
	pkg := &lang.Package{FullName: "sample"}
	require.NoError(t, p.ParsePackage(context.Empty(), pkg))
	require.Equal(t, []string{"first:parse", "first:compile", "second:parse", "second:compile"}, events)
	require.True(t, p.parsedPackages["sample"])
	require.NoError(t, p.Copy().ParsePackage(context.Empty(), pkg))
	require.Len(t, events, 4)
}

func TestPluginsPipelineErrors(t *testing.T) {
	sentinel := errors.New("pipeline failed")
	for _, tt := range []struct {
		name           string
		parse, compile error
		want           []string
		wantErr        error
	}{
		{"parse error", sentinel, nil, []string{"first:parse"}, sentinel},
		{"compile error", nil, sentinel, []string{"first:parse", "first:compile"}, sentinel},
		{"skip", core.NewSkipError(), core.NewSkipError(), []string{"first:parse", "first:compile", "second:parse", "second:compile"}, nil},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var events []string
			p := &Plugins{plugins: plugins{
				&pipelinePlugin{BasicPlugin: BasicPlugin{Name: "first"}, events: &events, parseErr: tt.parse, compileErr: tt.compile},
				&pipelinePlugin{BasicPlugin: BasicPlugin{Name: "second"}, events: &events},
			}, parsedPackages: make(map[string]bool)}
			err := p.ParsePackage(context.Empty(), &lang.Package{FullName: "sample"})
			require.ErrorIs(t, err, tt.wantErr)
			require.Equal(t, tt.want, events)
			require.Equal(t, tt.wantErr == nil, p.parsedPackages["sample"])
		})
	}
}

func TestPluginsCompilePackage(t *testing.T) {
	sentinel := errors.New("compile failed")
	for _, err := range []error{nil, core.NewSkipError(), sentinel} {
		var events []string
		p := &Plugins{plugins: plugins{
			&pipelinePlugin{BasicPlugin: BasicPlugin{Name: "first"}, events: &events, compileErr: err},
			&pipelinePlugin{BasicPlugin: BasicPlugin{Name: "second"}, events: &events},
		}}
		got := p.CompilePackage(context.Empty(), &lang.Package{})
		if err == sentinel {
			require.ErrorIs(t, got, sentinel)
			require.Equal(t, []string{"first:compile"}, events)
		} else {
			require.NoError(t, got)
			require.Equal(t, []string{"first:compile", "second:compile"}, events)
		}
	}
}

func TestPluginsParsePathWithoutParser(t *testing.T) {
	p := &Plugins{plugins: plugins{&BasicPlugin{Name: "unsupported"}}}
	pkg, err := p.ParsePath(context.Empty(), "missing")
	require.ErrorContains(t, err, "missing")
	require.Nil(t, pkg)
}
