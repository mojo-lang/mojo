package httptransport_test

import (
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/compiler"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/mpm"
	_ "github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/compiler"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit/generator/httptransport"
	"github.com/mojo-lang/mojo/go/pkg/compiler/plugin"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/require"
)

func TestHTTPResponseBindingGeneration(t *testing.T) {
	root := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(root, "mojo/sample/v1"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(root, "package.mojo"), []byte(`package sample { repository: 'example.com/sample' }`), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(root, "mojo/sample/v1/files.mojo"), []byte(`
type File { name: String @1 }
interface Files {
 @http.get('/files/{name}')
 @http.head('/files/{name}')
 @http.options('/files/{name}')
 get_file(name: String @1) -> File
}`), 0644))
	pkg, err := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler").ParsePath(context.Empty(), root)
	require.NoError(t, err)
	imports := make(core.Options)
	for _, p := range pkg.GetAllPackages() {
		imports[p.FullName] = p.GetGoPackageImport()
	}
	for _, p := range pkg.GetAllDependentPackages() {
		imports[p.FullName] = p.GetGoPackageImport()
	}
	services, err := compiler.CompilePackage(compiler.WithGoPackageImports(context.Empty(), imports), pkg)
	require.NoError(t, err)
	require.Len(t, services, 1)
	service := services[0]
	require.Len(t, service.Interface.Methods[0].Bindings, 3)
	var first string
	for i := 0; i < 2; i++ {
		generator, err := httptransport.NewServerHttpTransport(service)
		require.NoError(t, err)
		reader, err := generator.Render("", service)
		require.NoError(t, err)
		content, err := io.ReadAll(reader)
		require.NoError(t, err)
		_, err = parser.ParseFile(token.NewFileSet(), "transport_http.go", content, parser.AllErrors)
		require.NoError(t, err)
		source := string(content)
		for _, method := range []string{"GET", "HEAD", "OPTIONS"} {
			require.Contains(t, source, `router.Methods("`+method+`").Path("/files/{name}")`)
		}
		require.Contains(t, source, "httptransport.ServerBefore(nhttp.RequestToContext, headersToContext, queryToContext)")
		require.Contains(t, source, "nhttp.BoundResponseWriter(ctx, response)")
		require.Contains(t, source, `mjhttp.UnmarshalPathParam(pathParams, &req.Name, "name")`)
		if i == 0 {
			first = source
		} else {
			require.Equal(t, first, source)
		}
	}
}
