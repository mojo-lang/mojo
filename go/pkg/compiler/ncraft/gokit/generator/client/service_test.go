package client

import (
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func TestPrepareServicePackageAliases(t *testing.T) {
	message := func(importPath, packageName string) *data.Message {
		return &data.Message{
			Decl: &lang.StructDecl{Name: "Result"},
			Go:   &data.GoMessage{ImportPath: importPath, PackageName: packageName},
		}
	}
	request := message("example.com/api/sample/v1", "sample")
	service := &data.Service{
		Go:        &data.GoService{ApiImportPath: request.Go.ImportPath, PackageName: "sample"},
		Interface: &data.Interface{Name: "BookService", BaredName: "Book"},
	}
	for _, response := range []*data.Message{
		message("example.com/api/core", "core"),
		message("example.com/api/first/model", "model"),
		message("example.com/api/second/model", "model"),
		message("example.com/api/third/model2", "model2"),
		message("example.com/api/http", "http"),
		message("example.com/api/context", "context"),
		message("example.com/api/req", "req"),
		message("example.com/api/fmt", "fmt"),
		message("example.com/api/error", "error"),
		message("example.com/api/jsoniter", "jsoniter"),
		message("example.com/api/client", "client"),
		message("example.com/api/pb", "pb"),
	} {
		service.Interface.Methods = append(service.Interface.Methods, &data.Method{Request: request, Response: response})
	}
	require.NoError(t, PrepareService(service))
	require.Equal(t, "pb", service.Extensions["ClientAPIAlias"])
	want := []string{"core", "model", "model3", "model2", "http2", "context2", "req2", "fmt2", "error2", "jsoniter2", "client2", "pb2"}
	for i, method := range service.Interface.Methods {
		require.Equal(t, "pb.Result", method.Extensions["ClientRequestType"])
		require.Equal(t, want[i]+".Result", method.Extensions["ClientResponseType"])
	}
	imports := service.Extensions["ClientImports"]
	// Generation order and repeated generation must not change the aliases.
	methods := service.Interface.Methods
	for i, j := 0, len(methods)-1; i < j; i, j = i+1, j-1 {
		methods[i], methods[j] = methods[j], methods[i]
	}
	require.NoError(t, PrepareService(service))
	require.Equal(t, imports, service.Extensions["ClientImports"])
}
