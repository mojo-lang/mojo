package mpm

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/plugin"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func TestLocalMojoDependencies(t *testing.T) {
	root := t.TempDir()
	write := func(name, contents string) {
		filename := filepath.Join(root, name)
		require.NoError(t, os.MkdirAll(filepath.Dir(filename), 0755))
		require.NoError(t, os.WriteFile(filename, []byte(contents), 0644))
	}
	write("go/go.mod", "module "+lang.MojoGoModule+"\n\ngo 1.24.7\n")
	write("packages/core/package.mojo", "package mojo.core { repository: 'github.com/mojo-lang/mojo/packages/core' }")
	write("packages/core/mojo/core/new_type.mojo", "type FreshLocalType { value: String }")
	write("packages/document/package.mojo", `package mojo.document {
 repository: 'github.com/mojo-lang/mojo/packages/document'
 dependencies: { 'mojo.core': {repository: 'github.com/mojo-lang/mojo/packages/core'} }
 }`)
	write("packages/document/mojo/document/new_type.mojo", "type LocalDocument { value: String }")
	for _, start := range []string{root, filepath.Join(root, "go"), filepath.Join(root, "packages/document")} {
		require.Equal(t, root, util.MojoRepositoryRoot(start))
	}
	pkg, err := plugin.NewPlugins("mpm", "syntax").ParsePath(context.Empty(), filepath.Join(root, "packages/document"))
	require.NoError(t, err)
	dep := pkg.ResolvedDependencies["mojo.core"]
	require.NotNil(t, dep)
	require.NotSame(t, GetMojoPackage("mojo.core"), dep)
	require.Equal(t, filepath.Join(root, "packages/core"), dep.GetExtraString("path"))
	require.Len(t, dep.SourceFiles, 1)
	require.Equal(t, "FreshLocalType", dep.SourceFiles[0].Statements[0].GetDeclaration().GetName())
	// Outside a source checkout the normal embedded fallback still works.
	require.NoError(t, os.Remove(filepath.Join(root, "packages/core/package.mojo")))
	_, err = plugin.NewPlugins("mpm", "syntax").ParsePath(context.Empty(), filepath.Join(root, "packages/document"))
	require.NoError(t, err)
}
