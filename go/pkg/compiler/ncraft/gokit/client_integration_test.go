package gokit_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	buildbase "github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	clientbuild "github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/ncraft/gokit"
	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/commander"
	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/compiler"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/require"
)

func TestGeneratedHTTPAndGRPCClients(t *testing.T) {
	for _, name := range []string{"protoc", "protoc-gen-go", "protoc-gen-go-grpc"} {
		if _, err := exec.LookPath(name); err != nil {
			t.Skipf("requires %s", name)
		}
	}
	root := t.TempDir()
	write := func(name, content string) {
		name = filepath.Join(root, name)
		require.NoError(t, os.MkdirAll(filepath.Dir(name), 0755))
		require.NoError(t, os.WriteFile(name, []byte(content), 0644))
	}
	write("package.mojo", `package sample { repository:'example.com/acme/library'
 authors:[{organization:'example.com'}] }`)
	write("mojo/sample/v1/books.mojo", `type Book { name: String @1 }
 interface BookService {
 @http.get('/books/{isbn}')
 @http.get('/alternative/{isbn}')
 get_book(isbn: String @1, count: Int32 @2, labels: [String] @3) -> Book
 @http.post('/books')
 create_book(book: Book @1 @http.body) -> Book
 @http.delete('/books/{isbn}')
 delete_book(isbn: String @1)
 @http.post('/raw')
 raw_body(body: String @1 @http.body @http.style('raw')) -> Book
 @http.post('/lookup/{book.name}')
 lookup(book: Book @1 @http.body) -> Book
 rpc_only() -> Book
 @http.get('/fail')
 fail() -> Book
 }
 interface HealthService {
 @http.get('/ping')
 ping() -> String
 }`)

	// Exercise the actual build target. Dependency resolution runs below in the
	// generated module, with the current checkout replacing the Mojo runtime.
	bin := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(bin, "go"), []byte("#!/bin/sh\n[ \"$1 $2\" = 'mod tidy' ] || exit 1\npwd > tidy-directory.txt\n"), 0755))
	originalPath := os.Getenv("PATH")
	t.Setenv("PATH", bin+string(os.PathListSeparator)+originalPath)
	build := commander.Builder{Path: root, Targets: "api,client"}
	require.NoError(t, build.Execute())
	t.Setenv("PATH", originalPath)
	pkg := build.Package
	imports := make(core.Options)
	for _, p := range pkg.GetAllPackages() {
		imports[p.FullName] = p.GetGoPackageImport()
	}
	for _, p := range pkg.GetAllDependentPackages() {
		imports[p.FullName] = p.GetGoPackageImport()
	}
	compiled, err := compiler.CompilePackage(compiler.WithGoPackageImports(context.Empty(), imports), pkg)
	require.NoError(t, err)
	require.Len(t, compiled, 2)
	output := filepath.Join(root, "client-go")
	require.FileExists(t, filepath.Join(output, "tidy-directory.txt"))
	options := gokit.Options{Output: output, Repository: "example.com/acme/library/client-go", ApiRepository: "example.com/acme/library/go", MixedInAPI: true}
	for _, name := range []string{"book", "health"} {
		for _, file := range []string{"endpoints.go", "grpc.go", "grpc_client.go", "http_client.go", "options.go"} {
			require.FileExists(t, filepath.Join(output, "pkg", name+"-client", file))
		}
		require.NoDirExists(t, filepath.Join(output, name+"-client"))
	}
	require.NoDirExists(t, filepath.Join(root, "service-go"))
	// Client-only builds keep the same layout, including a custom module name
	// and a relocated output directory pointing back to the shared API module.
	t.Setenv("PATH", bin+string(os.PathListSeparator)+originalPath)
	require.NoError(t, (clientbuild.Builder{
		Builder: buildbase.Builder{Path: root, Package: pkg}, Type: "client",
	}).Build())
	require.NoError(t, (clientbuild.Builder{
		Builder: buildbase.Builder{Path: root, Package: pkg}, Type: "client",
		Output: filepath.Join(root, "generated"), Repository: "example.com/acme/sdk",
	}).Build())
	t.Setenv("PATH", originalPath)
	relocated, err := os.ReadFile(filepath.Join(root, "generated/client-go/go.mod"))
	require.NoError(t, err)
	require.Contains(t, string(relocated), "module example.com/acme/sdk")
	require.Contains(t, string(relocated), "example.com/acme/library/go => ../../go")
	require.NoDirExists(t, filepath.Join(root, "generated/sdk"))
	write("client-go/pkg/book-client/custom.go", "package book_client\nconst HandWritten=true\n")
	for _, service := range compiled {
		require.NoError(t, gokit.GenerateClient(service, options))
	}
	require.FileExists(t, filepath.Join(output, "pkg/book-client/custom.go"))
	// Compile and exercise the generated module against the actual API output.
	_, source, _, _ := runtime.Caller(0)
	goRoot := filepath.Clean(filepath.Join(filepath.Dir(source), "../../../.."))
	write("go/go.mod", fmt.Sprintf("module example.com/acme/library/go\n\ngo 1.24.7\nrequire github.com/mojo-lang/mojo/go v0.0.0\nreplace github.com/mojo-lang/mojo/go => %s\n", filepath.ToSlash(goRoot)))
	module, err := os.ReadFile(filepath.Join(output, "go.mod"))
	require.NoError(t, err)
	require.Contains(t, string(module), "example.com/acme/library/go => ../go")
	write("client-go/go.mod", string(module)+"\nreplace github.com/mojo-lang/mojo/go => "+filepath.ToSlash(goRoot)+"\n")
	sum, err := os.ReadFile(filepath.Join(goRoot, "go.sum"))
	require.NoError(t, err)
	write("client-go/go.sum", string(sum))
	tests, err := os.ReadFile("testdata/client/transports_test.go.txt")
	require.NoError(t, err)
	write("client-go/pkg/book-client/transports_test.go", string(tests))
	cmd := exec.Command("go", "test", "-mod=mod", "./...")
	cmd.Dir = output
	cmd.Env = append(os.Environ(), "GOWORK=off")
	result, err := cmd.CombinedOutput()
	require.NoError(t, err, string(result))
	t.Log(string(result))
	// All transport sources must be independent from the service module.
	require.NoError(t, filepath.WalkDir(output, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		content, err := os.ReadFile(path)
		require.NoError(t, err)
		require.NotContains(t, string(content), "service-go/")
		return nil
	}))
}
