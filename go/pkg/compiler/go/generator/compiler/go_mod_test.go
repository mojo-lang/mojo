package compiler

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/go/generator/data"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func TestGoModSharedDependencies(t *testing.T) {
	pkg := func(name, repository string) *lang.Package {
		url, err := core.NewUrl(repository)
		require.NoError(t, err)
		return &lang.Package{FullName: name, Repository: url}
	}
	document := pkg("mojo.document", "github.com/mojo-lang/mojo/packages/document")
	document.Dependencies = map[string]*lang.Package_Requirement{"mojo.core": {}}
	document.ResolvedDependencies = map[string]*lang.Package{"mojo.core": pkg("mojo.core", "github.com/mojo-lang/mojo/packages/core")}
	compiler := &GoMod{Data: data.NewData()}
	require.NoError(t, compiler.CompilePackage(context.Empty(), document))
	require.Empty(t, compiler.Data.GoMod.Dependencies)
	app := pkg("example.app", "github.com/example/app")
	app.Dependencies = map[string]*lang.Package_Requirement{"mojo.core": {}, "mojo.document": {}}
	app.ResolvedDependencies = map[string]*lang.Package{"mojo.core": document.ResolvedDependencies["mojo.core"], "mojo.document": document}
	compiler = &GoMod{Data: data.NewData()}
	require.NoError(t, compiler.CompilePackage(context.Empty(), app))
	require.Len(t, compiler.Data.GoMod.Dependencies, 1)
	require.Equal(t, lang.MojoGoModule, compiler.Data.GoMod.Dependencies[0].Name)
}
