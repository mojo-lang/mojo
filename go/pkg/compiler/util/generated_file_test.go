package util

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type failingReader struct{ err error }

func (r failingReader) Read([]byte) (int, error) { return 0, r.err }

func TestGeneratedFileWriteTo(t *testing.T) {
	for _, tt := range []struct {
		name, content string
		reader        io.Reader
	}{
		{"content", "package generated\n", nil},
		{"reader", "", strings.NewReader("package generated\n")},
		{"content takes precedence", "package generated\n", failingReader{errors.New("must not read")}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			file := &GeneratedFile{Name: "nested/output.go", Content: tt.content, Reader: tt.reader}
			require.NoError(t, file.WriteTo(dir, &PathGuard{}))
			got, err := os.ReadFile(filepath.Join(dir, file.Name))
			require.NoError(t, err)
			require.Equal(t, "package generated\n", string(got))
			require.Equal(t, string(got), file.Content)
		})
	}
}

func TestGeneratedFileSkipExisting(t *testing.T) {
	for _, tt := range []struct {
		name, existing                    string
		skipExisting, skipMixed, wantSkip bool
	}{
		{"existing", generatedHeader, true, false, true},
		{"manual", "package manual\n", false, true, true},
		{"generated", generatedHeader, false, true, false},
		{"overwrite", "package manual\n", false, false, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			name := writeTestFile(t, dir, "output.go", tt.existing)
			file := &GeneratedFile{Name: "output.go", Content: "replacement", SkipIfExist: tt.skipExisting, SkipIfUserCodeMixed: tt.skipMixed}
			require.NoError(t, file.WriteTo(dir, &PathGuard{DisableClear: true}))
			got, err := os.ReadFile(name)
			require.NoError(t, err)
			want := "replacement"
			if tt.wantSkip {
				want = tt.existing
			}
			require.Equal(t, want, string(got))
		})
	}
}

func TestGeneratedFileWriteErrors(t *testing.T) {
	dir := t.TempDir()
	require.ErrorContains(t, (&GeneratedFile{Content: "data"}).WriteTo(dir, nil), "file name")
	empty := &GeneratedFile{Name: "empty.go"}
	require.NoError(t, empty.WriteTo(dir, nil))
	require.NoFileExists(t, filepath.Join(dir, empty.Name))
	sentinel := errors.New("read failed")
	file := &GeneratedFile{Name: "output.go", Reader: failingReader{sentinel}}
	require.ErrorIs(t, file.WriteTo(dir, nil), sentinel)
	require.NoFileExists(t, filepath.Join(dir, file.Name))
	file = &GeneratedFile{Name: "missing/output.go", Content: "data"}
	require.Error(t, file.WriteTo(dir, nil))
	blocked := writeTestFile(t, dir, "blocked", "data")
	require.Error(t, file.WriteTo(blocked, &PathGuard{}))
}
