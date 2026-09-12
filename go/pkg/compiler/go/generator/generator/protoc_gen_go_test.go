package generator

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/mojo/mpm"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/protobuf/descriptor"
	"github.com/stretchr/testify/require"
)

func TestProtocGenGoUnifiedProtoDirectory(t *testing.T) {
	for _, sharedWithInput := range []bool{true, false} {
		name := "dependency repository"
		if sharedWithInput {
			name = "input repository"
		}
		t.Run(name, func(t *testing.T) {
			root := t.TempDir()
			bin := filepath.Join(root, "bin")
			require.NoError(t, os.MkdirAll(bin, 0755))
			require.NoError(t, os.MkdirAll(filepath.Join(root, "protobuf"), 0755))
			// Capture the actual invocation without installing or running plugins.
			for _, command := range []string{"protoc", "protoc-gen-go", "protoc-gen-go-grpc"} {
				script := "#!/bin/sh\nexit 0\n"
				if command == "protoc" {
					script = "#!/bin/sh\nprintf '%s\\n' \"$@\" > \"$PROTOC_TEST_ARGS\"\n"
				}
				require.NoError(t, os.WriteFile(filepath.Join(bin, command), []byte(script), 0755))
			}
			argsPath := filepath.Join(root, "args")
			t.Setenv("PATH", bin+string(os.PathListSeparator)+os.Getenv("PATH"))
			t.Setenv("PROTOC_TEST_ARGS", argsPath)

			depRoot := root
			if !sharedWithInput {
				depRoot = t.TempDir()
			}
			pkg := &lang.Package{ResolvedDependencies: make(map[string]*lang.Package)}
			for _, name := range []string{"core", "document", "lang"} {
				dep := &lang.Package{Name: name, FullName: "mojo." + name}
				dep.SetExtraString("path", depRoot)
				pkg.ResolvedDependencies[dep.FullName] = dep
			}
			file := descriptor.NewFileWithName("example/probe.proto", "example")
			file.Messages = []*descriptor.Message{{}}
			_, err := ProtocGenGo(root, pkg, []*descriptor.File{file})
			require.NoError(t, err)
			contents, err := os.ReadFile(argsPath)
			require.NoError(t, err)
			args := strings.Split(strings.TrimSpace(string(contents)), "\n")
			var includes []string
			for _, arg := range args {
				if strings.HasPrefix(arg, "-I") || strings.HasPrefix(arg, "--proto_path=") {
					includes = append(includes, arg)
				}
			}
			want := []string{"-I."}
			if !sharedWithInput {
				want = append(want, "--proto_path="+filepath.Join(depRoot, "protobuf"))
			}
			require.Equal(t, want, includes)
			require.Contains(t, args, file.GetName())
			require.NoDirExists(t, filepath.Join(root, "go.out"))
		})
	}
}

func TestProtocGenGoEmbeddedProtoDirectory(t *testing.T) {
	for _, mode := range []string{"success", "failure", "real protoc"} {
		t.Run(mode, func(t *testing.T) {
			if mode == "real protoc" {
				for _, command := range []string{"protoc", "protoc-gen-go", "protoc-gen-go-grpc"} {
					filename, err := exec.LookPath(command)
					if err != nil {
						t.Skipf("integration test requires %s", command)
					}
					if command == "protoc" {
						t.Setenv("PROTOC_TEST_REAL", filename)
					}
				}
			}
			root := t.TempDir()
			bin := filepath.Join(root, "bin")
			require.NoError(t, os.MkdirAll(bin, 0755))
			require.NoError(t, os.MkdirAll(filepath.Join(root, "protobuf"), 0755))
			capture := filepath.Join(root, "capture")
			argsPath := filepath.Join(root, "args")
			t.Setenv("PROTOC_TEST_CAPTURE", capture)
			t.Setenv("PROTOC_TEST_ARGS", argsPath)
			script := `#!/bin/sh
set -eu
printf '%s\n' "$@" > "$PROTOC_TEST_ARGS"
for arg in "$@"; do
    case "$arg" in
        --proto_path=*) cp -R "${arg#--proto_path=}" "$PROTOC_TEST_CAPTURE" ;;
    esac
done
`
			switch mode {
			case "failure":
				script += "exit 1\n"
			case "real protoc":
				script += "exec \"$PROTOC_TEST_REAL\" \"$@\"\n"
			}
			require.NoError(t, os.WriteFile(filepath.Join(bin, "protoc"), []byte(script), 0755))
			if mode != "real protoc" {
				for _, command := range []string{"protoc-gen-go", "protoc-gen-go-grpc"} {
					require.NoError(t, os.WriteFile(filepath.Join(bin, command), []byte("#!/bin/sh\nexit 0\n"), 0755))
				}
			}
			t.Setenv("PATH", bin+string(os.PathListSeparator)+os.Getenv("PATH"))
			pkg := &lang.Package{ResolvedDependencies: make(map[string]*lang.Package)}
			for _, name := range mpm.MojoPackageNames {
				pkg.ResolvedDependencies["mojo."+name] = &lang.Package{Name: name, FullName: "mojo." + name}
			}
			// HTTP Request imports both HTTP and core protos, exercising cross-module imports.
			const source = `syntax = "proto3";
package example;
import "mojo/http/request.proto";
option go_package = "example.com/probe;example";
enum Probe { PROBE_UNSPECIFIED = 0; }
`
			require.NoError(t, os.WriteFile(filepath.Join(root, "protobuf/probe.proto"), []byte(source), 0644))
			file := descriptor.NewFileWithName("probe.proto", "example")
			file.Enums = []*descriptor.Enum{{}}
			generated, err := ProtocGenGo(root, pkg, []*descriptor.File{file})
			if mode == "failure" {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if mode == "real protoc" {
					require.Len(t, generated, 1)
					require.Contains(t, generated[0].Content, "type Probe int32")
				}
			}
			contents, err := os.ReadFile(argsPath)
			require.NoError(t, err)
			var includes []string
			for _, arg := range strings.Split(string(contents), "\n") {
				if strings.HasPrefix(arg, "--proto_path=") {
					includes = append(includes, strings.TrimPrefix(arg, "--proto_path="))
				}
			}
			require.Len(t, includes, 1, "all embedded modules must share one include root")
			require.NoDirExists(t, includes[0], "temporary protos must be cleaned up on success and failure")
			for _, name := range mpm.MojoPackageNames {
				files := mpm.GetMojoPbFile(name)
				require.NotNil(t, files)
				for filename, expected := range files.Parts {
					actual, err := os.ReadFile(filepath.Join(capture, filename))
					require.NoError(t, err)
					require.Equal(t, expected, actual, filename)
				}
			}
			require.NoDirExists(t, filepath.Join(root, "go.out"))
		})
	}
}
