package generator

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit/generator/handlers"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit/generator/templates"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
	"github.com/stretchr/testify/require"
)

// Regenerating the service must retain application-owned startup and transport
// methods while keeping the server entry point fully generated.
func TestServerExtensionsSurviveRegeneration(t *testing.T) {
	const extension = `package handlers
import (
 "context"
 "github.com/gorilla/mux"
 nserver "github.com/ncraft-io/ncraft/go/pkg/gokit/server"
)
func (s catalogServer) RegisterHTTPHandlers(r *mux.Router) { r.Path("/custom") }
func (s catalogServer) Start(cfg nserver.Config) error { return nil }
func (s catalogServer) Shutdown(ctx context.Context) error { return nil }
func (s catalogServer) Errors() <-chan error { return nil }
`
	root := t.TempDir()
	dir := filepath.Join(root, "pkg/catalog-service/handlers")
	require.NoError(t, os.MkdirAll(dir, 0755))
	require.NoError(t, os.WriteFile(filepath.Join(root, handlerTarget), []byte(previousHandler), 0644))
	extensionPath := filepath.Join(dir, "server.go")
	require.NoError(t, os.WriteFile(extensionPath, []byte(extension), 0644))
	service := sourceService()
	service.FuncMap["ToCamel"] = strcase.ToCamel
	service.FuncMap["ToKebab"] = strcase.ToKebab
	options := &Options{Output: root}
	var first []byte
	for i := 0; i < 2; i++ {
		files, err := options.generateTemplatedFiles(service, []string{"internal/NAME-server/run.go.tmpl", handlers.ServerHandlerPath}, templates.Service)
		require.NoError(t, err)
		for _, file := range files {
			content, err := io.ReadAll(file.Reader)
			require.NoError(t, err)
			if strings.HasSuffix(file.Name, "/run.go") {
				require.Contains(t, string(content), "registrar.RegisterHTTPHandlers(r)")
				require.Contains(t, string(content), "starter.Start(cfg)")
				require.Contains(t, string(content), "RegisterService(cfg, r, s, service)")
				require.Contains(t, string(content), "NewEndpoints(options, service)")
				require.Contains(t, string(content), "service.(nserver.Starter)")
				require.Contains(t, string(content), "started[i].(nserver.Shutdowner)")
				require.Contains(t, string(content), "nserver.WaitForError(ctx, errc, started...)")
				require.Contains(t, string(content), "nserver.ShutdownServers(shutdownCtx, s, httpServer, debugServer)")
				require.NotContains(t, string(content), "handlers.InterruptHandler")
				require.Less(t, strings.Index(string(content), "registrar.RegisterHTTPHandlers(r)"), strings.Index(string(content), "svc.RegisterHttpHandler(r,"))
				if i == 0 {
					first = content
				} else {
					require.Equal(t, first, content)
				}
			}
			target := filepath.Join(root, file.Name)
			require.NoError(t, os.MkdirAll(filepath.Dir(target), 0755))
			require.NoError(t, os.WriteFile(target, content, 0644))
		}
		preserved, err := os.ReadFile(extensionPath)
		require.NoError(t, err)
		require.Equal(t, extension, string(preserved))
	}
}

func TestUnifiedServerLifecycle(t *testing.T) {
	service := sourceService()
	service.FuncMap["ToCamel"] = strcase.ToCamel
	service.FuncMap["ToKebab"] = strcase.ToKebab
	service.FuncMap["ToSnake"] = strcase.ToSnake
	service.AllInterfaces = []*data.Interface{service.Interface, {Name: "Audit", BaredName: "Audit", ServerName: "AuditServer"}}
	files, err := (&Options{}).generateTemplatedFiles(service, []string{"internal/server/run.go.tmpl"}, templates.Service)
	require.NoError(t, err)
	content, err := io.ReadAll(files[0].Reader)
	require.NoError(t, err)
	source := string(content)
	// Each endpoint registration must reuse the service whose lifecycle is managed.
	for _, name := range []string{"catalog", "audit"} {
		require.Contains(t, source, name+"Service := "+name+"_handlers.NewService()")
		require.Contains(t, source, name+".RegisterService(cfg, r, s, "+name+"Service)")
	}
	require.Contains(t, source, "service.(nserver.Starter)")
	require.Contains(t, source, "starter.Start(cfg)")
	require.Contains(t, source, "started = append(started, service)")
	require.Contains(t, source, "for i := len(started) - 1; i >= 0; i--")
	require.Contains(t, source, "shutdowner.Shutdown(shutdownCtx)")
	require.Contains(t, source, "nserver.WaitForError(ctx, errc, started...)")
	require.NotContains(t, source, "go InterruptHandler")
	require.Less(t, strings.Index(source, "nserver.ShutdownServers("), strings.Index(source, "shutdowner.Shutdown("))
}
