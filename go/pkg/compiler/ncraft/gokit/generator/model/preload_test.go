package model

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"text/template"

	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/gokit/generator/model/templates"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
	"github.com/stretchr/testify/require"
)

func TestGeneratedModelsPreload(t *testing.T) {
	root := t.TempDir()
	write := func(name, content string) {
		name = filepath.Join(root, name)
		require.NoError(t, os.MkdirAll(filepath.Dir(name), 0755))
		require.NoError(t, os.WriteFile(name, []byte(content), 0644))
	}
	// Relations are supplied by GORM's runtime schema, so exercise the actual
	// generated models with multiple mapped associations and a JSON field.
	for _, name := range []string{"Parent", "Plain"} {
		file, err := render("pkg/model/"+name+".go", templates.EntityModel, entityModel{
			Name: name, ImportPath: "example.com/preload/entity", KeyName: "Id", KeyType: "string",
		})
		require.NoError(t, err)
		write(file.Name, file.Content)
	}
	database, err := render("pkg/model/db.go", templates.DB, nil)
	require.NoError(t, err)
	write(database.Name, database.Content)
	for source, target := range map[string]string{
		"testdata/preload_entities.go.txt": "entity/entity.go",
		"testdata/preload_test.go.txt":     "pkg/model/preload_test.go",
	} {
		content, err := os.ReadFile(source)
		require.NoError(t, err)
		write(target, string(content))
	}
	_, source, _, _ := runtime.Caller(0)
	goRoot := filepath.Clean(filepath.Join(filepath.Dir(source), "../../../../../.."))
	// Use the production SQL template emitted for @db.json fields, including
	// Entity values and arrays of Entities; they must remain column values.
	jsonTemplate, err := os.ReadFile(filepath.Join(goRoot, "pkg/compiler/go/generator/generator/template/go/DB_JSON.sql.go.tmpl"))
	require.NoError(t, err)
	for _, name := range []string{"Snapshot", "ParentSnapshots"} {
		reader, err := util.ApplyTemplate("db-json", string(jsonTemplate), map[string]interface{}{
			"PackageName": "sample", "GoPackageName": "entity", "FullName": name,
			"InDbPkg": false, "UnderlyingTypeName": "",
		}, template.FuncMap{"IsMojoPackage": func(string) bool { return false }})
		require.NoError(t, err)
		content, err := io.ReadAll(reader)
		require.NoError(t, err)
		write("entity/"+name+".sql.go", string(content))
	}
	module := fmt.Sprintf("module example.com/preload\n\ngo 1.24.7\nrequire github.com/mojo-lang/mojo/go v0.0.0\nreplace github.com/mojo-lang/mojo/go => %q\n", filepath.ToSlash(goRoot))
	if ncraftRoot := os.Getenv("NCRAFT_GO_ROOT"); ncraftRoot != "" {
		ncraftRoot, err = filepath.Abs(ncraftRoot)
		require.NoError(t, err)
		require.FileExists(t, filepath.Join(ncraftRoot, "go.mod"))
		module += fmt.Sprintf("replace github.com/ncraft-io/ncraft/go => %q\n", filepath.ToSlash(ncraftRoot))
	}
	write("go.mod", module)
	sum, err := os.ReadFile(filepath.Join(goRoot, "go.sum"))
	require.NoError(t, err)
	write("go.sum", string(sum))
	cmd := exec.Command("go", "test", "-mod=mod", "./pkg/model")
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOWORK=off")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, string(output))
	t.Log(string(output))
}
