package gokit_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	mojoc "github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/mojo"
	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/compiler"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit"
	"github.com/stretchr/testify/require"
)

func TestLegacyModelServiceRegeneration(t *testing.T) {
	root := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(root, "package.mojo"), []byte("package sample {repository: 'example.com/sample'}"), 0644))
	require.NoError(t, os.MkdirAll(filepath.Join(root, "mojo/sample"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(root, "mojo/sample/item.mojo"), []byte(`type Item { value: String @1 }
 interface Items {
 @http.get('/items')
 get_item() -> Item
 }`), 0644))
	pkg, err := (mojoc.Builder{Builder: builder.Builder{Path: root}}).Build()
	require.NoError(t, err)
	services, err := compiler.CompilePackage(context.Empty(), pkg)
	require.NoError(t, err)
	require.Len(t, services, 1)
	output := filepath.Join(root, "service-go")
	options := gokit.Options{Output: output, Repository: "example.com/sample/service-go", ApiRepository: "example.com/sample/go"}
	require.NoError(t, gokit.GenerateService(services[0], options))
	require.NoDirExists(t, filepath.Join(output, "internal/model"))
	require.DirExists(t, filepath.Join(output, "pkg/items-service"))
	old := filepath.Join(output, "internal/model")
	require.NoError(t, os.MkdirAll(old, 0755))
	for _, name := range []string{"ENTITY_model.go", "ENTITY_query.go"} {
		require.NoError(t, os.WriteFile(filepath.Join(old, name), nil, 0644))
	}
	require.NoError(t, gokit.GenerateService(services[0], options))
	require.NoDirExists(t, old)
}
