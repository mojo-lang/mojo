package injection

import (
	"context"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/fatih/structtag"
	"github.com/stretchr/testify/require"
)

func TestTagInjectorMergesOverridesAndSkipsInternalFields(t *testing.T) {
	source := []byte("package sample\n" +
		"type Item struct {\n" +
		" Name string `json:\"name,omitempty\" xml:\"name\"`\n" +
		" XXX_state int `json:\"state\"`\n" +
		" hidden int `json:\"hidden\"`\n" +
		" Untagged int\n" +
		"}\n")
	override, err := structtag.Parse("json:\"display_name\" db:\"name\"")
	require.NoError(t, err)
	injector := NewTagInjector([]string{"json", "xml"}, TagInjections{"Item.Name": *override})
	var changed []string
	require.Same(t, injector, injector.RegisterTagChanger(func(ctx context.Context, field *ast.Field, tags *structtag.Tags) {
		changed = append(changed, GetStructName(ctx)+"."+field.Names[0].Name)
	}))
	result, err := injector.Inject("item.go", source)
	require.NoError(t, err)
	require.Equal(t, []string{"Item.Name", "Item.XXX_state", "Item.hidden"}, changed)
	parsed, err := parser.ParseFile(token.NewFileSet(), "item.go", result, 0)
	require.NoError(t, err)
	fields := parsed.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List
	tags, err := structtag.Parse(strings.Trim(fields[0].Tag.Value, "`"))
	require.NoError(t, err)
	for key, want := range map[string]string{"json": "display_name", "xml": "name", "db": "name"} {
		tag, err := tags.Get(key)
		require.NoError(t, err)
		require.Equal(t, want, tag.Name)
		require.Empty(t, tag.Options)
	}
	for _, field := range fields[1:3] {
		tags, err = structtag.Parse(strings.Trim(field.Tag.Value, "`"))
		require.NoError(t, err)
		for _, key := range []string{"json", "xml"} {
			tag, err := tags.Get(key)
			require.NoError(t, err)
			require.Equal(t, "-", tag.Name)
		}
	}
	require.Nil(t, fields[3].Tag)
	again, err := injector.Inject("item.go", result)
	require.NoError(t, err)
	require.Equal(t, result, again)
}

func TestTagInjectorLeavesUnchangedAndMalformedTags(t *testing.T) {
	for _, tag := range []string{"json:\"name\"", "malformed"} {
		source := []byte("package sample\ntype Item struct { Name string `" + tag + "` }")
		got, err := NewTagInjector(nil, nil).Inject("item.go", source)
		require.NoError(t, err)
		require.Equal(t, source, got)
	}
}
