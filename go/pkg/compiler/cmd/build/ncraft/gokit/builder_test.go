package gokit

import (
	"fmt"
	"github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/builder"
	mojoc "github.com/mojo-lang/mojo/go/pkg/compiler/cmd/build/mojo"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnsupportedBuildType(t *testing.T) {
	for _, buildType := range []string{"sidecar", "ncraft.sidecar", "unknown"} {
		t.Run(buildType, func(t *testing.T) {
			// Reject removed or unknown types before reading the package or
			// falling through to service generation.
			b := Builder{Type: buildType}
			require.ErrorContains(t, b.Build(), "unsupported ncraft build type")
		})
	}
}

func TestEntityOnlyServiceOutput(t *testing.T) {
	root := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(root, "package.mojo"), []byte("package sample { repository: 'example.com/acme/store' }"), 0644))
	require.NoError(t, os.MkdirAll(filepath.Join(root, "mojo/sample"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(root, "mojo/sample/record.mojo"), []byte("@entity\ntype Record { name: String @1 }"), 0644))
	pkg, err := (mojoc.Builder{Builder: builder.Builder{Path: root}}).Build()
	require.NoError(t, err)
	for _, withAPI := range []bool{false, true} {
		t.Run(fmt.Sprint(withAPI), func(t *testing.T) {
			// Module resolution is outside this routing test; real generated code is
			// compiled and exercised against SQLite by the model integration test.
			bin := t.TempDir()
			require.NoError(t, os.WriteFile(filepath.Join(bin, "go"), []byte("#!/bin/sh\n[ \"$1 $2\" = 'mod tidy' ] || exit 1\npwd > tidy-directory.txt\n"), 0755))
			t.Setenv("PATH", bin+string(os.PathListSeparator)+os.Getenv("PATH"))
			b := Builder{Builder: builder.Builder{Path: root, Package: pkg, APIEnabled: withAPI}, Type: "service"}
			require.NoError(t, b.Build())
			output := filepath.Join(root, "service-go")
			require.FileExists(t, filepath.Join(output, "pkg/model/record_model.go"))
			require.FileExists(t, filepath.Join(output, "pkg/model/db.go"))
			require.FileExists(t, filepath.Join(output, "go.mod"))
			cwd, err := os.ReadFile(filepath.Join(output, "tidy-directory.txt"))
			require.NoError(t, err)
			require.Equal(t, output, strings.TrimSpace(string(cwd)))
			require.NoDirExists(t, filepath.Join(output, "internal/model"))
		})
	}
}
