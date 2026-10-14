package plugin

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
)

func (p *lifecyclePlugin) ParsePackage(context.Context, *lang.Package) error {
	return p.record("parse", p.compileErr)
}
func (p *lifecyclePlugin) ParseSourceFile(context.Context, *lang.SourceFile) error {
	return p.record("parse", p.compileErr)
}

func TestParserLifecycle(t *testing.T) {
	sentinel := errors.New("parse failure")
	for _, dispatch := range []struct {
		name string
		call func(interface{}) error
	}{
		{"package", func(p interface{}) error { return ParsePackage(p, context.Empty(), &lang.Package{}) }},
		{"source file", func(p interface{}) error { return ParseSourceFile(p, context.Empty(), &lang.SourceFile{}) }},
	} {
		t.Run(dispatch.name, func(t *testing.T) {
			require.NoError(t, dispatch.call(nil))
			require.True(t, core.IsSkipError(dispatch.call(struct{}{})))
			for _, tt := range []struct {
				name                string
				pre, parse, wantErr error
				events              []string
			}{
				{"success", nil, nil, nil, []string{"pre", "parse", "post"}},
				{"skip", core.NewSkipError(), nil, nil, []string{"pre"}},
				{"pre error", sentinel, nil, sentinel, []string{"pre"}},
				{"parse error", nil, sentinel, sentinel, []string{"pre", "parse", "post"}},
			} {
				t.Run(tt.name, func(t *testing.T) {
					p := &lifecyclePlugin{preErr: tt.pre, compileErr: tt.parse}
					require.ErrorIs(t, dispatch.call(p), tt.wantErr)
					require.Equal(t, tt.events, p.events)
				})
			}
		})
	}
}

type stringParserFunc func(context.Context, string) (*lang.SourceFile, error)

func (f stringParserFunc) ParseString(ctx context.Context, s string) (*lang.SourceFile, error) {
	return f(ctx, s)
}

type fileParserFunc func(context.Context, string) (*lang.SourceFile, error)

func (f fileParserFunc) ParseFile(ctx context.Context, s string) (*lang.SourceFile, error) {
	return f(ctx, s)
}

type sourceFileParserFunc func(context.Context, *lang.SourceFile) error

func (f sourceFileParserFunc) ParseSourceFile(ctx context.Context, s *lang.SourceFile) error {
	return f(ctx, s)
}

func TestParseFileLoadsContentAndSetsNames(t *testing.T) {
	for _, virtual := range []bool{false, true} {
		name := "virtual"
		if !virtual {
			name = "disk"
		}
		t.Run(name, func(t *testing.T) {
			filename := "nested/schema.mojo"
			ctx := context.Empty()
			if virtual {
				ctx = WithFs(ctx, fstest.MapFS{filename: &fstest.MapFile{Data: []byte("package sample")}})
			} else {
				filename = filepath.Join(t.TempDir(), "schema.mojo")
				require.NoError(t, os.WriteFile(filename, []byte("package sample"), 0600))
			}
			expected := &lang.SourceFile{}
			parser := stringParserFunc(func(ctx context.Context, content string) (*lang.SourceFile, error) {
				require.Equal(t, filename, ContextFilename(ctx))
				require.Equal(t, "package sample", content)
				return expected, nil
			})
			got, err := ParseFile(parser, ctx, filename)
			require.NoError(t, err)
			require.Same(t, expected, got)
			require.Equal(t, "schema.mojo", got.Name)
			require.Equal(t, filename, got.FullName)
		})
	}
}

func TestParseFileErrorsAndFallback(t *testing.T) {
	sentinel := errors.New("syntax error")
	parser := stringParserFunc(func(context.Context, string) (*lang.SourceFile, error) { return nil, sentinel })
	ctx := WithFs(context.Empty(), fstest.MapFS{"bad.mojo": &fstest.MapFile{Data: []byte("bad")}})
	file, err := ParseFile(parser, ctx, "missing.mojo")
	require.ErrorIs(t, err, os.ErrNotExist)
	require.Nil(t, file)
	file, err = ParseFile(parser, ctx, "bad.mojo")
	require.ErrorIs(t, err, sentinel)
	require.Nil(t, file)
	_, err = ParseString(struct{}{}, ctx, "source")
	require.Error(t, err)
	_, err = ParseFile(struct{}{}, ctx, "source")
	require.Error(t, err)
	expected := &lang.SourceFile{Name: "direct"}
	direct := fileParserFunc(func(ctx context.Context, name string) (*lang.SourceFile, error) {
		require.Equal(t, "source", name)
		return expected, nil
	})
	file, err = ParseFile(direct, ctx, "source")
	require.NoError(t, err)
	require.Same(t, expected, file)
}

func TestParsePackageTraversesFilesAndChildren(t *testing.T) {
	rootFile := &lang.SourceFile{Name: "root.mojo"}
	childFile := &lang.SourceFile{Name: "child.mojo"}
	child := &lang.Package{FullName: "root.child", SourceFiles: []*lang.SourceFile{childFile}}
	root := &lang.Package{FullName: "root", SourceFiles: []*lang.SourceFile{rootFile}, Children: []*lang.Package{child}}
	for _, fail := range []bool{false, true} {
		var visits []string
		sentinel := errors.New("parse failed")
		parser := sourceFileParserFunc(func(ctx context.Context, file *lang.SourceFile) error {
			visits = append(visits, file.Name)
			expected := root
			if file == childFile {
				expected = child
			}
			require.Same(t, expected, context.Package(ctx))
			if fail {
				return sentinel
			}
			return nil
		})
		err := ParsePackage(parser, context.Empty(), root)
		if fail {
			require.ErrorIs(t, err, sentinel)
			require.Equal(t, []string{"root.mojo"}, visits)
		} else {
			require.NoError(t, err)
			require.Equal(t, []string{"root.mojo", "child.mojo"}, visits)
		}
	}
}
