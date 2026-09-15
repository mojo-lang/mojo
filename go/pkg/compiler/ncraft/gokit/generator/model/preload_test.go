package model

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"text/template"

	"github.com/mojo-lang/mojo/go/pkg/compiler/ncraft/data"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func preloadEntity(name string, fields ...*lang.ValueDecl) *data.Message {
	key := &lang.ValueDecl{Name: "id", Type: &lang.NominalType{PackageName: "mojo.core", Name: "String"}}
	decl := &lang.StructDecl{Name: name, PackageName: "sample", Type: &lang.StructType{Fields: append([]*lang.ValueDecl{key}, fields...)}}
	// Automatic entity detection marks the resolved declaration this way too.
	decl.SetImplicitBoolAttribute(core.EntityAttributeName, true)
	return &data.Message{Decl: decl, KeyField: key, Go: &data.GoMessage{ImportPath: "example.com/preload/entity"}}
}

func preloadField(name string, decl *lang.StructDecl, repeated, json bool) *lang.ValueDecl {
	typ := &lang.NominalType{PackageName: decl.PackageName, Name: decl.Name, TypeDeclaration: lang.NewStructTypeDeclaration(decl)}
	if repeated {
		typ = &lang.NominalType{PackageName: "mojo.core", Name: "Array", GenericArguments: []*lang.NominalType{typ},
			TypeDeclaration: lang.NewStructTypeDeclaration(&lang.StructDecl{PackageName: "mojo.core", Name: "Array"})}
	}
	field := &lang.ValueDecl{Name: name, Type: typ}
	if json {
		field.SetBoolAttribute(db.JSONAttributeFullName, true)
	}
	return field
}

func TestGeneratedModelsPreloadRules(t *testing.T) {
	associated := preloadEntity("Associated").Decl
	associated.PackageName = "dependency.v1" // Dependency Entities need no local model.
	ordinary := &lang.StructDecl{Name: "Ordinary", PackageName: "sample", Type: &lang.StructType{}}
	for _, tc := range []struct {
		name   string
		fields []*lang.ValueDecl
		want   bool
	}{
		{name: "no associations"},
		{name: "ordinary struct", fields: []*lang.ValueDecl{preloadField("value", ordinary, false, false)}},
		{name: "ordinary struct array", fields: []*lang.ValueDecl{preloadField("values", ordinary, true, false)}},
		{name: "one to one", fields: []*lang.ValueDecl{preloadField("value", associated, false, false)}, want: true},
		{name: "one to many", fields: []*lang.ValueDecl{preloadField("values", associated, true, false)}, want: true},
		{name: "JSON entity", fields: []*lang.ValueDecl{preloadField("value", associated, false, true)}},
		{name: "JSON entity array", fields: []*lang.ValueDecl{preloadField("values", associated, true, true)}},
		{name: "mixed", fields: []*lang.ValueDecl{preloadField("snapshot", associated, false, true), preloadField("values", associated, true, false)}, want: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			files, err := (Model{}).GenerateEntities([]*data.Message{preloadEntity("Parent", tc.fields...)})
			require.NoError(t, err)
			var source string
			for _, file := range files {
				if file.Name == "pkg/model/parent_model.go" {
					source = file.Content
				}
			}
			require.NotEmpty(t, source)
			count := 0
			if tc.want {
				count = 3
			}
			require.Equal(t, count, strings.Count(source, ".Preload(clause.Associations)"))
		})
	}
}

func TestGeneratedModelsPreload(t *testing.T) {
	root := t.TempDir()
	write := func(name, content string) {
		name = filepath.Join(root, name)
		require.NoError(t, os.MkdirAll(filepath.Dir(name), 0755))
		require.NoError(t, os.WriteFile(name, []byte(content), 0644))
	}
	// Exercise generation from resolved entity metadata and execute the result
	// against GORM's actual schema with multiple associations and JSON fields.
	parent := preloadEntity("Parent",
		preloadField("profile", preloadEntity("Profile").Decl, false, false),
		preloadField("settings", preloadEntity("Settings").Decl, false, false),
		preloadField("children", preloadEntity("Child").Decl, true, false),
		preloadField("notes", preloadEntity("Note").Decl, true, false),
		preloadField("json_entity", preloadEntity("Snapshot").Decl, false, true),
		preloadField("json_entities", preloadEntity("Snapshot").Decl, true, true),
	)
	files, err := (Model{}).GenerateEntities([]*data.Message{parent, preloadEntity("Plain")})
	require.NoError(t, err)
	for _, file := range files {
		write(file.Name, file.Content)
	}
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
