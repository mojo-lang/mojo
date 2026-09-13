package generator

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit/generator/handlers"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit/generator/templates"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
	"github.com/stretchr/testify/require"
)

const handlerTarget = "pkg/catalog-service/handlers/handlers.go"
const previousHandler = `package handlers
import ("context";pb "example.com/api")
type catalogServer struct{}
func(s catalogServer)GetItem(c context.Context,r *pb.GetItemRequest)(*pb.Item,error){return nil,nil}
`
const siblingHandler = `package handlers
import ("context";pb "example.com/api")
func(s catalogServer)ListItems(c context.Context,r *pb.ListItemsRequest)(*pb.Item,error){return nil,nil}
`

func sourceService() *data.Service {
	svc := &data.Service{Interface: &data.Interface{Name: "Catalog", BaredName: "Catalog", ServerName: "CatalogServer"}, Go: &data.GoService{PackageName: "catalog", ApiImportPath: "example.com/api"}, FuncMap: template.FuncMap{
		"ToLowerCamel": strcase.ToLowerCamel, "GoName": strcase.ToCamel, "GoPackageName": func(string) string { return "pb" },
	}}
	for _, name := range []string{"get_item", "list_items", "create_item"} {
		svc.Interface.Methods = append(svc.Interface.Methods, &data.Method{Name: name, Request: &data.Message{Name: strcase.ToCamel(name) + "Request"}, Response: &data.Message{Name: "Item"}})
	}
	return svc
}
func TestHandlerPackageFromOutputDirectory(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "pkg/catalog-service/handlers")
	require.NoError(t, os.MkdirAll(dir, 0755))
	require.NoError(t, os.WriteFile(filepath.Join(root, handlerTarget), []byte(previousHandler), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "lists.go"), []byte(siblingHandler), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "ignored_test.go"), []byte("not valid Go"), 0644))
	options := &Options{Output: root}
	files, err := options.generateTemplatedFiles(sourceService(), []string{handlers.ServerHandlerPath}, templates.Service)
	require.NoError(t, err)
	require.Len(t, files, 1)
	generated, err := io.ReadAll(files[0].Reader)
	require.NoError(t, err)
	require.NotContains(t, string(generated), "func (s catalogServer) ListItems")
	require.Contains(t, string(generated), "func (s catalogServer) CreateItem")
	sibling, err := os.ReadFile(filepath.Join(dir, "lists.go"))
	require.NoError(t, err)
	require.Equal(t, siblingHandler, string(sibling))
	require.NoError(t, os.WriteFile(filepath.Join(root, handlerTarget), generated, 0644))
	again, err := options.generateTemplatedFiles(sourceService(), []string{handlers.ServerHandlerPath}, templates.Service)
	require.NoError(t, err)
	next, err := io.ReadAll(again[0].Reader)
	require.NoError(t, err)
	require.Equal(t, generated, next)
}
func TestPreviousFileSnapshotsCanBeReused(t *testing.T) {
	options := &Options{PreviousFiles: map[string]io.Reader{
		handlerTarget:                                 strings.NewReader(previousHandler),
		"pkg/catalog-service/handlers/lists.go":       strings.NewReader(siblingHandler),
		"pkg/catalog-service/handlers/middlewares.go": strings.NewReader("package handlers\n\n// Preserve middleware customization.\n"),
	}}
	var previous [][]byte
	for i := 0; i < 2; i++ {
		files, err := options.generateTemplatedFiles(sourceService(), []string{handlers.ServerHandlerPath, handlers.MiddlewaresPath}, templates.Service)
		require.NoError(t, err)
		var outputs [][]byte
		for _, file := range files {
			content, err := io.ReadAll(file.Reader)
			require.NoError(t, err)
			outputs = append(outputs, content)
		}
		require.Contains(t, string(outputs[1]), "Preserve middleware customization.")
		if i == 1 {
			require.Equal(t, previous, outputs)
		}
		previous = outputs
	}
}
func TestBrokenSiblingAbortsBeforeWriting(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "pkg/catalog-service/handlers")
	require.NoError(t, os.MkdirAll(dir, 0755))
	require.NoError(t, os.WriteFile(filepath.Join(root, handlerTarget), []byte(previousHandler), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "broken.go"), []byte("package handlers\nfunc broken("), 0644))
	options := &Options{Output: root}
	files, err := options.generateTemplatedFiles(sourceService(), []string{handlers.ServerHandlerPath}, templates.Service)
	require.ErrorContains(t, err, "broken.go")
	require.Nil(t, files)
	original, err := os.ReadFile(filepath.Join(root, handlerTarget))
	require.NoError(t, err)
	require.Equal(t, previousHandler, string(original))
}
