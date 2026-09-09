package generator_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	mojoc "github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/mojo"
	"github.com/mojo-lang/mojo/go/pkg/compiler/document/generator"
	openapigen "github.com/mojo-lang/mojo/go/pkg/compiler/openapi/generator"
	"github.com/mojo-lang/mojo/go/pkg/markdown"
	"github.com/stretchr/testify/require"
	extast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
)

func TestStructDescriptionCodeBlockTable(t *testing.T) {
	root := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(root, "package.mojo"), []byte("package example { repository: 'github.com/example/docs' }"), 0644))
	require.NoError(t, os.MkdirAll(filepath.Join(root, "mojo/example"), 0755))
	source := strings.Join([]string{
		"type Example {",
		"    /// Example payload:",
		"    ///",
		"    /// ```json",
		"    /// {",
		"    ///   \"value\": \"left | right\"",
		"    /// }",
		"    /// ```",
		"    ///",
		"    /// - First item",
		"    ///   continuation",
		"    /// - Second item",
		"    payload: String @1",
		"    /// The following field must remain in the table.",
		"    next: String @2",
		"}",
	}, "\n")
	require.NoError(t, os.WriteFile(filepath.Join(root, "mojo/example/example.mojo"), []byte(source), 0644))
	pkg, err := (mojoc.Builder{Builder: builder.Builder{Path: root}}).Build()
	require.NoError(t, err)
	apis := openapigen.NewCompiler()
	require.NoError(t, apis.CompilePackages(pkg.GetAllPackages()))
	docs, err := generator.NewCompiler(root, pkg, apis.OpenAPIs).Compile()
	require.NoError(t, err)
	output := t.TempDir()
	require.NoError(t, generator.NewGenerator(docs).Generate(output))
	content, err := os.ReadFile(filepath.Join(output, "example/example.md"))
	require.NoError(t, err)
	require.Len(t, strings.Split(strings.TrimSuffix(string(content), "\n"), "\n"), 4, string(content))
	parsed := markdown.NewParser().MdEngine.Parser().Parse(text.NewReader(content))
	require.NotNil(t, parsed.FirstChild())
	require.Nil(t, parsed.FirstChild().NextSibling(), string(content))
	table, ok := parsed.FirstChild().(*extast.Table)
	require.True(t, ok, string(content))
	require.Equal(t, 3, table.ChildCount())
	for row := table.FirstChild(); row != nil; row = row.NextSibling() {
		require.Equal(t, 6, row.ChildCount(), string(content))
	}
	require.Contains(t, string(content), "left \\| right")
	require.Contains(t, string(content), "Second item")
	require.Contains(t, string(content), "The following field must remain in the table.")
}
