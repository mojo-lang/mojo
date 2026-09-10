package model_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	mojoc "github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/mojo"
	pbc "github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/protobuf"
	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	gogen "github.com/mojo-lang/mojo/go/pkg/compiler/go/generator"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/compiler"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit"
	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit/generator/model"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/stretchr/testify/require"
)

func TestGeneratedModelsCRUD(t *testing.T) {
	for _, command := range []string{"protoc", "protoc-gen-go", "protoc-gen-go-grpc"} {
		if _, err := exec.LookPath(command); err != nil {
			t.Skipf("requires %s", command)
		}
	}
	root := t.TempDir()
	write := func(name, content string) {
		name = filepath.Join(root, name)
		require.NoError(t, os.MkdirAll(filepath.Dir(name), 0755))
		require.NoError(t, os.WriteFile(name, []byte(content), 0644))
	}
	write("package.mojo", `package sample {
 repository: 'example.com/acme/store'
 authors: [{organization:'example.com'}]
 }`)
	write("mojo/sample/entity.mojo", `type Record {
 id: String @1
 name: String @2
 count: Int32 @3
 active: Bool @4
 }
 type Keyed {
 code: String @1 @key
 score: Int32 @2
 }
 @entity
 type Manual { name: String @1 }
	 type Ordinary { value: String @1 }
	 type Numbered {
	 id: Int64 @1
	 name: String @2
	 }
 interface Store {
 @http.get('/records/{id}')
 get_record(id: String @1) -> Record
 }
 interface Admin {
 @http.get('/admin/records/{id}')
 get_record(id: String @1) -> Record
 }`)
	pkg, err := (mojoc.Builder{Builder: builder.Builder{Path: root}}).Build()
	require.NoError(t, err)
	manual := pkg.GetEntityNode("sample.Manual")
	require.NotNil(t, manual)
	require.True(t, manual.KeyField.Implicit)
	require.Equal(t, "mojo.core.String", manual.KeyField.Type.GetFullName())
	require.Len(t, manual.GetTypeDeclaration().GetStructDecl().GetAllFields(), 2)
	number, err := manual.KeyField.GetIntegerAttribute(core.NumberAttributeName)
	require.NoError(t, err)
	require.EqualValues(t, 536870911, number)
	require.Nil(t, pkg.GetEntityNode("sample.Ordinary"))
	require.Equal(t, "code", pkg.GetEntityNode("sample.Keyed").KeyField.Name)
	base := builder.Builder{Path: root, Package: pkg, APIEnabled: true}
	descriptors, err := (pbc.Builder{Builder: base}).Build()
	require.NoError(t, err)
	goCompiler := gogen.NewCompiler(root, descriptors)
	files, err := goCompiler.CompilePackage(pkg)
	require.NoError(t, err)
	require.NoError(t, gogen.NewGenerator(files, goCompiler.Data).Generate(filepath.Join(root, "go")))
	ctx := compiler.WithGoPackageImports(context.Empty(), core.Options{"sample": "example.com/acme/store/go/pkg/sample"})
	compiled, err := compiler.Compile(ctx, pkg)
	require.NoError(t, err)
	require.Len(t, compiled.Data, 2)
	require.Len(t, compiled.Entities, 4)
	output := filepath.Join(root, "service-go")
	options := gokit.Options{Repository: "example.com/acme/store/service-go", ApiRepository: "example.com/acme/store/go", MixedInAPI: true, Output: output}
	for _, svc := range compiled.Data {
		require.Len(t, svc.Entities, 4)
		require.NoError(t, gokit.GenerateService(svc, options))
	}
	for _, name := range []string{"record_model.go", "keyed_model.go", "manual_model.go", "numbered_model.go", "db.go"} {
		require.FileExists(t, filepath.Join(output, "pkg/model", name))
	}
	require.NoFileExists(t, filepath.Join(output, "pkg/model/ordinary_model.go"))
	require.NoDirExists(t, filepath.Join(output, "internal/model"))
	// Rebuilding multiple services preserves handwritten files and regenerates all models.
	write("service-go/pkg/model/custom.go", "package model\nconst Custom=true\n")
	for _, svc := range compiled.Data {
		require.NoError(t, gokit.GenerateService(svc, options))
	}
	require.FileExists(t, filepath.Join(output, "pkg/model/custom.go"))
	// Use the actual generated protobuf types and the checkout's DB runtime.
	// This module only needs model dependencies, independent of server transport versions.
	_, source, _, _ := runtime.Caller(0)
	goRoot := filepath.Clean(filepath.Join(filepath.Dir(source), "../../../../../.."))
	mod := fmt.Sprintf("module example.com/acme/store/go\n\ngo 1.24.7\nrequire github.com/mojo-lang/mojo/go v0.0.0\nreplace github.com/mojo-lang/mojo/go => %s\n", filepath.ToSlash(goRoot))
	write("go/go.mod", mod)
	mod = fmt.Sprintf("module example.com/acme/store/service-go\n\ngo 1.24.7\nrequire (\nexample.com/acme/store/go v0.0.0\ngithub.com/mojo-lang/mojo/go v0.0.0\n)\nreplace example.com/acme/store/go => ../go\nreplace github.com/mojo-lang/mojo/go => %s\n", filepath.ToSlash(goRoot))
	write("service-go/go.mod", mod)
	test, err := os.ReadFile("testdata/crud_test.go.txt")
	require.NoError(t, err)
	write("service-go/pkg/model/crud_test.go", string(test))
	sum, err := os.ReadFile(filepath.Join(goRoot, "go.sum"))
	require.NoError(t, err)
	write("service-go/go.sum", string(sum))
	cmd := exec.Command("go", "test", "-mod=mod", "./pkg/model")
	cmd.Dir = output
	cmd.Env = append(os.Environ(), "GOWORK=off")
	result, err := cmd.CombinedOutput()
	require.NoError(t, err, string(result))
	t.Log(string(result))
	// Missing metadata must fail generation, not silently omit CRUD files.
	compiled.Entities[0].KeyField = nil
	_, err = (model.Model{}).GenerateEntities(compiled.Entities)
	require.ErrorContains(t, err, "has no primary key")
}
