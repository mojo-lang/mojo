package handlers

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
	"github.com/stretchr/testify/require"
)

func testService(names ...string) *data.Service {
	svc := &data.Service{Interface: &data.Interface{Name: "Catalog", ServerName: "CatalogServer"}, Go: &data.GoService{ApiImportPath: "example.com/api"}, FuncMap: template.FuncMap{
		"ToLowerCamel": strcase.ToLowerCamel, "GoName": strcase.ToCamel, "GoPackageName": func(string) string { return "pb" },
	}}
	for _, name := range names {
		svc.Interface.Methods = append(svc.Interface.Methods, &data.Method{Name: name, Request: &data.Message{Name: strcase.ToCamel(name) + "Request"}, Response: &data.Message{Name: "Item"}})
	}
	return svc
}

const entry = `package handlers
import (
 "context"
 pb "example.com/api"
)
type catalogServer struct { pb.UnimplementedCatalogServer }
func NewService() pb.CatalogServer { return catalogServer{} }
// Keep the body and comment exactly as written.
func (s catalogServer) GetItem(ctx context.Context, in *pb.GetItemRequest) (*pb.Item, error) { return &pb.Item{}, nil }
`

func renderPackage(t *testing.T, svc *data.Service, files map[string][]byte) (string, error) {
	t.Helper()
	h, err := NewPackage(svc.Interface, "handlers.go", files)
	if err != nil {
		return "", err
	}
	reader, err := h.Render(ServerHandlerPath, svc)
	if err != nil {
		return "", err
	}
	code, err := io.ReadAll(reader)
	require.NoError(t, err)
	formatted, err := format.Source(code)
	require.NoError(t, err, "%s", code)
	return string(formatted), nil
}
func methodNames(t *testing.T, source string) []string {
	t.Helper()
	file, err := parser.ParseFile(token.NewFileSet(), "handlers.go", source, 0)
	require.NoError(t, err)
	var names []string
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && receiverName(fn) == "catalogServer" {
			names = append(names, fn.Name.Name)
		}
	}
	return names
}

func TestPackageRegenerationPreservesSplitMethods(t *testing.T) {
	svc := testService("get_item", "list_items", "create_item", "delete_item")
	sibling := `package handlers
import (ctx "context"; api "example.com/api")
func (s *catalogServer) ListItems(c ctx.Context, r *api.ListItemsRequest) (*api.Item,error) { panic("user implementation") }
`
	files := map[string][]byte{"handlers.go": []byte(entry), "columns.go": []byte(sibling), "columns_test.go": []byte("not valid Go"), "nested/extra.go": []byte("not valid Go")}
	generated, err := renderPackage(t, svc, files)
	require.NoError(t, err)
	require.Equal(t, []string{"GetItem", "CreateItem", "DeleteItem"}, methodNames(t, generated))
	require.Contains(t, generated, "Keep the body and comment exactly as written.")
	require.Equal(t, sibling, string(files["columns.go"]))
	files["handlers.go"] = []byte(generated)
	again, err := renderPackage(t, svc, files)
	require.NoError(t, err)
	require.Equal(t, generated, again)
}

func TestPreserveHelpersOtherReceiversAndRemovedRPCs(t *testing.T) {
	source := entry + `
func (s catalogServer) Helper() int { return 42 }
func (s catalogServer) Removed(ctx context.Context, in *pb.GetItemRequest) (*pb.Item,error) { return nil,nil }
type different struct{}
func (s different) ListItems() {}
`
	generated, err := renderPackage(t, testService("get_item", "list_items"), map[string][]byte{"handlers.go": []byte(source)})
	require.NoError(t, err)
	require.Contains(t, generated, "return 42")
	require.Contains(t, generated, "func (s different) ListItems()")
	require.Equal(t, []string{"GetItem", "Helper", "Removed", "ListItems"}, methodNames(t, generated))
}

func TestSignatureAndDuplicateDiagnostics(t *testing.T) {
	for _, test := range []struct{ name, source, want string }{
		{"request", "func (s catalogServer) ListItems(ctx context.Context, in *pb.OldRequest) (*pb.Item,error) {return nil,nil}", "signature mismatch"},
		{"return", "func (s catalogServer) ListItems(ctx context.Context, in *pb.ListItemsRequest) (*pb.Other,error) {return nil,nil}", "signature mismatch"},
		{"pointer", "func (s catalogServer) ListItems(ctx context.Context, in pb.ListItemsRequest) (*pb.Item,error) {return nil,nil}", "signature mismatch"},
		{"arity", "func (s catalogServer) ListItems(ctx context.Context) {}", "signature mismatch"},
		{"duplicate", "func (s *catalogServer) GetItem(ctx context.Context, in *pb.GetItemRequest) (*pb.Item,error) {return nil,nil}", "duplicate RPC"},
	} {
		t.Run(test.name, func(t *testing.T) {
			sibling := "package handlers\nimport (\"context\"; pb \"example.com/api\")\n" + test.source
			files := map[string][]byte{"handlers.go": []byte(entry), "columns.go": []byte(sibling)}
			_, err := renderPackage(t, testService("get_item", "list_items"), files)
			require.ErrorContains(t, err, test.want)
			require.ErrorContains(t, err, "columns.go:")
			require.Equal(t, entry, string(files["handlers.go"]))
		})
	}
}

func TestConditionalRPCsAreDiagnosed(t *testing.T) {
	for _, name := range []string{"columns_linux.go", "columns_amd64.go", "columns.go"} {
		t.Run(name, func(t *testing.T) {
			source := `package handlers
import ("context"; pb "example.com/api")
func (s catalogServer) ListItems(ctx context.Context, in *pb.ListItemsRequest) (*pb.Item,error) {return nil,nil}
`
			if name == "columns.go" {
				source = "//go:build feature\n\n" + source
			}
			_, err := renderPackage(t, testService("get_item", "list_items"), map[string][]byte{"handlers.go": []byte(entry), name: []byte(source)})
			require.ErrorContains(t, err, "build constraints")
			require.ErrorContains(t, err, name)
		})
	}
}

func TestImportsRespectAliasesAndCollisions(t *testing.T) {
	source := strings.ReplaceAll(entry, "pb.", "api.")
	source = strings.Replace(source, `pb "example.com/api"`, `api "example.com/api"`, 1)
	source = strings.Replace(source, "\"context\"", "ctx \"context\"", 1)
	source = strings.ReplaceAll(source, "context.Context", "ctx.Context")
	generated, err := renderPackage(t, testService("get_item", "list_items"), map[string][]byte{"handlers.go": []byte(source)})
	require.NoError(t, err)
	require.Contains(t, generated, "in *api.ListItemsRequest")
	require.NotContains(t, generated, "pb.")
	// The output file does not already use the API; a top-level pb name is legal.
	source = "package handlers\nvar pb = 42\n"
	generated, err = renderPackage(t, testService("list_items"), map[string][]byte{"handlers.go": []byte(source), "instance.go": []byte("package handlers\ntype catalogServer struct{}")})
	require.NoError(t, err)
	require.Contains(t, generated, `pb2 "example.com/api"`)
	require.Contains(t, generated, "in *pb2.ListItemsRequest")
}

func TestMissingEntryFileDoesNotDuplicateMovedDeclarations(t *testing.T) {
	generated, err := renderPackage(t, testService("get_item", "list_items"), map[string][]byte{"instance.go": []byte(entry)})
	require.NoError(t, err)
	require.NotContains(t, generated, "type catalogServer")
	require.NotContains(t, generated, "func NewService")
	require.Equal(t, []string{"ListItems"}, methodNames(t, generated))
}

func TestForeignPackageAndUnimplementedEmbeddingAreNotImplementations(t *testing.T) {
	source := `package handlers
import pb "example.com/api"
type catalogServer struct { pb.UnimplementedCatalogServer }
`
	generated, err := renderPackage(t, testService("get_item"), map[string][]byte{"handlers.go": []byte(source), "foreign.go": []byte("package other\nfunc(s catalogServer) GetItem() {}")})
	require.NoError(t, err)
	require.Equal(t, []string{"GetItem"}, methodNames(t, generated))
}

func TestRegeneratedPackageCompiles(t *testing.T) {
	root := t.TempDir()
	api := `package api
import "context"
type Item struct{}
type GetItemRequest struct{}
type ListItemsRequest struct{}
type CreateItemRequest struct{}
type UnimplementedCatalogServer struct{}
type CatalogServer interface {
 GetItem(context.Context,*GetItemRequest)(*Item,error)
 ListItems(context.Context,*ListItemsRequest)(*Item,error)
 CreateItem(context.Context,*CreateItemRequest)(*Item,error)
}
`
	files := map[string][]byte{"handlers.go": []byte(strings.Replace(entry, "return catalogServer{}", "return &catalogServer{}", 1)), "lists.go": []byte(`package handlers
import ("context"; api "example.com/api")
func(s *catalogServer) ListItems(c context.Context,in *api.ListItemsRequest)(*api.Item,error){return nil,nil}
`)}
	svc := testService("get_item", "list_items", "create_item")
	generated, err := renderPackage(t, svc, files)
	require.NoError(t, err)
	files["handlers.go"] = []byte(generated)
	// Simulate moving all entry-point declarations out of handlers.go as well.
	files["moved.go"] = files["handlers.go"]
	delete(files, "handlers.go")
	generated, err = renderPackage(t, svc, files)
	require.NoError(t, err)
	files["handlers.go"] = []byte(generated)
	require.NoError(t, os.WriteFile(filepath.Join(root, "go.mod"), []byte("module example.com\ngo 1.24.0\n"), 0644))
	require.NoError(t, os.MkdirAll(filepath.Join(root, "api"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(root, "api", "api.go"), []byte(api), 0644))
	require.NoError(t, os.MkdirAll(filepath.Join(root, "handlers"), 0755))
	for name, source := range files {
		require.NoError(t, os.WriteFile(filepath.Join(root, "handlers", name), source, 0644))
	}
	command := exec.Command("go", "test", "./...")
	command.Dir = root
	command.Env = append(os.Environ(), "GOWORK=off")
	output, err := command.CombinedOutput()
	require.NoError(t, err, "%s", output)
}
