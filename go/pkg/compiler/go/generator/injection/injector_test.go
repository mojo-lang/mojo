package injection

import (
	"context"
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInjectorVisitsGroupedDeclarations(t *testing.T) {
	source := []byte("package sample\nimport (\n \"fmt\"\n \"io\"\n)\ntype (\n First struct { Embedded; Name string }\n Alias string\n Second struct { ID int }\n)\nfunc Run() {}\n")
	var events []string
	injector := Injector{
		OnImport: func(ctx context.Context, spec *ast.ImportSpec) {
			require.NotNil(t, GetFile(ctx))
			require.Equal(t, source, GetFileContent(ctx))
			require.Nil(t, GetStructType(ctx))
			events = append(events, "import:"+spec.Path.Value)
		},
		PreStruct: func(ctx context.Context, typ *ast.StructType, name string) {
			require.Equal(t, name, GetStructName(ctx))
			require.Same(t, typ, GetStructType(ctx))
			events = append(events, "pre:"+name)
		},
		OnStructField: func(ctx context.Context, field *ast.Field, appendArea func(Area)) {
			events = append(events, GetStructName(ctx)+"."+field.Names[0].Name)
			appendArea(TextArea{Start: field.Names[0].Pos(), End: field.Names[0].End(), Content: field.Names[0].Name + "Updated"})
		},
		PostStruct: func(ctx context.Context, appendArea func(Area)) { events = append(events, "post:"+GetStructName(ctx)) },
		OnFunction: func(ctx context.Context, f *ast.FuncDecl, appendArea func(Area)) {
			events = append(events, "func:"+f.Name.Name)
		},
	}
	got, err := injector.Inject("sample.go", source)
	require.NoError(t, err)
	require.Equal(t, []string{"import:\"fmt\"", "import:\"io\"", "pre:First", "First.Name", "post:First", "pre:Second", "Second.ID", "post:Second", "func:Run"}, events)
	require.Contains(t, string(got), "NameUpdated string")
	require.Contains(t, string(got), "IDUpdated int")
	require.Contains(t, string(got), "Embedded;")
	require.Contains(t, string(source), "Name string")
}

func TestInjectorInvalidSourceAndNoCallbacks(t *testing.T) {
	injector := Injector{}
	got, err := injector.Inject("broken.go", []byte("package {"))
	require.Error(t, err)
	require.Nil(t, got)
	source := []byte("package sample\ntype Item struct { Name string }\n")
	got, err = injector.Inject("sample.go", source)
	require.NoError(t, err)
	require.Equal(t, source, got)
	require.Nil(t, GetFile(context.Background()))
	require.Nil(t, GetFileContent(context.Background()))
	require.Nil(t, GetStructType(context.Background()))
	require.Empty(t, GetStructName(context.Background()))
}

func TestInjectFileAppliesOriginalPositions(t *testing.T) {
	source := []byte("abcdef")
	got := injectFile(source, []Area{
		TextArea{Start: 5, End: 7, Content: "Z"},
		TextArea{Start: 2, End: 4, Content: "LONG"},
		TextArea{Start: 1, End: 1, Content: "<"},
		TextArea{Start: 7, End: 7, Content: ">"},
	})
	require.Equal(t, "<aLONGdZ>", string(got))
	require.Equal(t, "abcdef", string(source))
	require.Equal(t, source, injectFile(source, nil))
}
