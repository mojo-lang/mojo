package commander

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/mojo-lang/mojo/go/pkg/compiler/mojo/mpm"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestBootstrapLocalSources(t *testing.T) {
	for _, command := range []string{"protoc", "protoc-gen-go", "protoc-gen-go-grpc"} {
		if _, err := exec.LookPath(command); err != nil {
			t.Skipf("integration test requires %s", command)
		}
	}
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
	module := "module " + lang.MojoGoModule + "\n\ngo 1.24.7\n"
	manual := "package core\nconst HandWritten = true\n"
	write("go/go.mod", module)
	write("go/pkg/mojo/core/manual.go", manual)
	for _, name := range mpm.MojoPackageNames {
		dependencies := ""
		source := "type BootstrapReference { probe: mojo.core.BootstrapProbe @1 }"
		if name == "core" {
			source = "type String\ntype BootstrapProbe { value: String @1 }"
		} else {
			dependencies = "dependencies: { 'mojo.core': {repository: 'github.com/mojo-lang/mojo/packages/core'} }"
		}
		write("packages/"+name+"/package.mojo", fmt.Sprintf("package mojo.%s {\nrepository: 'github.com/mojo-lang/mojo/packages/%s'\n%s\n}", name, name, dependencies))
		write("packages/"+name+"/mojo/"+name+"/bootstrap.mojo", source)
	}
	require.NoError(t, Bootstrap(filepath.Join(root, "go")))
	require.Contains(t, read("go/pkg/mojo/document/bootstrap.pb.go"), lang.MojoGoModule+"/pkg/mojo/core")
	require.Contains(t, read("packages/core/protobuf/mojo/core/bootstrap.proto"), lang.MojoGoModule+"/pkg/mojo/core;core")
	require.NoFileExists(t, filepath.Join(root, "packages/core/go/go.mod"))
	require.Equal(t, manual, read("go/pkg/mojo/core/manual.go"))
	require.Equal(t, module, read("go/go.mod"))

	// A second pass must read the updated source, and update both the generated
	// Go file and the snapshots that the next CLI binary embeds.
	write("packages/core/mojo/core/bootstrap.mojo", "type String\ntype BootstrapProbe { value: String @1\n revision: String @2 }")
	require.NoError(t, Bootstrap(root))
	require.Contains(t, read("go/pkg/mojo/core/bootstrap.pb.go"), "GetRevision()")
	binary, err := mpm.DecodeBinaryFile([]byte(read("go/pkg/compiler/mojo/mpm/mojo/core.pb.binary")))
	require.NoError(t, err)
	require.Contains(t, string(binary.Parts["mojo/core/bootstrap.proto"]), "string revision = 2;")
	pkg := &lang.Package{}
	require.NoError(t, proto.Unmarshal([]byte(read("go/pkg/compiler/mojo/mpm/mojo/core.binary")), pkg))
	require.Contains(t, pkg.String(), "revision")
	require.Equal(t, manual, read("go/pkg/mojo/core/manual.go"))
	require.Equal(t, module, read("go/go.mod"))

	// --output changes the Go destination while protoc must consume the newly
	// generated intermediate protobuf, rather than a stale file in that output.
	write("packages/core/mojo/core/bootstrap.mojo", "type String\ntype BootstrapProbe { fresh: String @1 }")
	b := Builder{Pwd: root, Path: filepath.Join(root, "packages/core"), Targets: "go", Output: "custom-go"}
	require.NoError(t, b.Execute())
	require.Contains(t, read("custom-go/pkg/mojo/core/bootstrap.pb.go"), "GetFresh()")
	require.Contains(t, read("packages/core/protobuf/mojo/core/bootstrap.proto"), "string fresh = 1;")
	require.NoFileExists(t, filepath.Join(root, "custom-go/go.mod"))
}

func TestUnsupportedBuildTargets(t *testing.T) {
	for _, target := range []string{"java", "cpp", "invalid"} {
		b := Builder{Targets: target}
		require.ErrorContains(t, b.Execute(), "unsupported build target")
	}
	b := Builder{Targets: "service", Engine: "boot"}
	require.ErrorContains(t, b.Execute(), "unsupported engine")
}
