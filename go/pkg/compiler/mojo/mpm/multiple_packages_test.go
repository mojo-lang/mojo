package mpm

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/plugin"
	"github.com/stretchr/testify/require"
)

func TestMultiplePackageDeclarations(t *testing.T) {
	root := t.TempDir()
	write := func(name, content string) {
		filename := filepath.Join(root, name)
		require.NoError(t, os.MkdirAll(filepath.Dir(filename), 0755))
		require.NoError(t, os.WriteFile(filename, []byte(content), 0644))
	}
	// Deliberately declare the dependent package first.
	write("package.mojo", `package mojo.sample {
 repository: 'github.com/mojo-lang/mojo'
 version: '0.2.0'
 dependencies: { 'mojo.core': '^0.1' }
 }
 package mojo.core {
 repository: 'github.com/mojo-lang/mojo'
 version: '0.1.0'
 }`)
	write("mojo/core/types.mojo", "type String\ntype Shared { value: String @1 }")
	write("mojo/sample/package.mojo", "type Consumer { shared: mojo.core.Shared @1 }")
	for _, groups := range [][]string{{"mpm", "syntax"}, {"mpm", "syntax", "semantic", "compiler"}} {
		pkg, err := plugin.NewPlugins(groups...).ParsePath(context.Empty(), root)
		require.NoError(t, err)
		require.True(t, pkg.GetExtraBool("package-set"))
		require.Len(t, pkg.Children, 2)
		sample, core := pkg.Children[0], pkg.Children[1]
		require.Equal(t, "mojo.sample", sample.FullName)
		require.Equal(t, uint64(2), sample.Version.Minor)
		require.Same(t, core, sample.ResolvedDependencies["mojo.core"])
		require.Len(t, core.SourceFiles, 1)
		require.Len(t, sample.SourceFiles, 1)
		selected, err := plugin.NewPlugins(groups...).ParsePath(context.Empty(), filepath.Join(root, "mojo/sample"))
		require.NoError(t, err)
		require.Equal(t, "mojo.sample", selected.FullName)
		require.Equal(t, root, selected.GetExtraString("path"))
	}
}

func TestInvalidPackageManifests(t *testing.T) {
	for _, tt := range []struct{ source, message string }{
		{"", "no package declarations"},
		{"type Invalid {}", "expected a package declaration"},
		{"package mojo.a {}\npackage mojo.a {}", "duplicate package mojo.a"},
		{"package mojo.a { dependencies: { 'mojo.b': '*' } }\npackage mojo.b { dependencies: { 'mojo.a': '*' } }", "cyclic package dependency"},
	} {
		t.Run(tt.message, func(t *testing.T) {
			root := t.TempDir()
			require.NoError(t, os.WriteFile(filepath.Join(root, "package.mojo"), []byte(tt.source), 0644))
			_, err := NewDependencyParser(nil).ParsePath(context.Empty(), root)
			require.ErrorContains(t, err, tt.message)
		})
	}
}
