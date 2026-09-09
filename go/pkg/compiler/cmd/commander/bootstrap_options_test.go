package commander

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBootstrapMojoOptionsGo(t *testing.T) {
	for _, command := range []string{"protoc", "protoc-gen-go", "protoc-gen-go-grpc"} {
		if _, err := exec.LookPath(command); err != nil {
			t.Skipf("integration test requires %s", command)
		}
	}
	for _, target := range []string{"", "api", "golang", "go,java", "java"} {
		t.Run(target, func(t *testing.T) {
			root := t.TempDir()
			write := func(name, contents string) {
				filename := filepath.Join(root, name)
				require.NoError(t, os.MkdirAll(filepath.Dir(filename), 0755))
				require.NoError(t, os.WriteFile(filename, []byte(contents), 0644))
			}
			read := func(name string) string {
				contents, err := os.ReadFile(filepath.Join(root, name))
				require.NoError(t, err)
				return string(contents)
			}
			write("go/go.mod", "module github.com/mojo-lang/mojo/go\n\ngo 1.24.7\n")
			write("package.mojo", `package mojo.core {
 repository: 'github.com/mojo-lang/mojo'
 authors: [{ organization: 'mojolang.org' }]
 }`)
			write("mojo/core/probe.mojo", "type String\ntype Probe { value: String @1 }")
			options := `syntax = "proto2";
package mojo;
option go_package = "github.com/mojo-lang/mojo/go/pkg/mojo";
option java_package = "org.mojolang.mojo";
import "google/protobuf/descriptor.proto";
extend google.protobuf.FieldOptions { optional string bootstrap_note = 51001; }
`
			const source = "protobuf/mojo/mojo.proto"
			const output = "go/pkg/mojo/mojo.pb.go"
			write(source, options)
			write(output, "// stale generated options\npackage mojo\n")
			run := func() {
				var targets []string
				if target != "" {
					targets = []string{target}
				}
				require.NoError(t, Bootstrap(filepath.Join(root, "go"), targets...))
			}
			run()
			if target == "java" {
				require.Equal(t, "// stale generated options\npackage mojo\n", read(output))
				return
			}
			require.Contains(t, read(output), "E_BootstrapNote")
			require.NotContains(t, read(output), "stale generated options")
			write(source, strings.ReplaceAll(options, "bootstrap_note", "bootstrap_revision"))
			run()
			require.Contains(t, read(output), "E_BootstrapRevision")
			require.NotContains(t, read(output), "E_BootstrapNote")
			// A failed protoc run must keep the last successfully generated file.
			previous := read(output)
			write(source, "invalid proto")
			require.ErrorContains(t, generateMojoOptionsGo(root), "mojo/mojo.proto")
			require.Equal(t, previous, read(output))
		})
	}
}
