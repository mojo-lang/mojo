package commander

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/compiler/mojo/mpm"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

// Bootstrap regenerates the selected standard library languages into the
// checkout's shared go/ and java/ directories, then refreshes embedded snapshots.
func Bootstrap(start string, targets ...string) error {
	target := "go"
	if len(targets) > 0 {
		target = strings.Join(targets, ",")
	}
	generateGo := false
	for _, name := range strings.Split(target, ",") {
		switch strings.TrimSpace(name) {
		case "go", "golang":
			generateGo = true
		case "java":
		default:
			return fmt.Errorf("unsupported bootstrap target %q; use go or java", name)
		}
	}
	root := util.MojoRepositoryRoot(start)
	if root == "" {
		return fmt.Errorf("cannot find Mojo source repository from %q", start)
	}

	if _, err := os.Stat(filepath.Join(root, "package.mojo")); err == nil {
		b := Builder{Pwd: root, Path: root, Targets: target}
		if err := b.Execute(); err != nil {
			return fmt.Errorf("bootstrap: %w", err)
		}
	} else {
		for _, name := range mpm.MojoPackageNames {
			b := Builder{Pwd: root, Path: util.MojoPackagePath(root, name), Targets: target}
			if err := b.Execute(); err != nil {
				return fmt.Errorf("bootstrap mojo.%s: %w", name, err)
			}
		}
	}

	if generateGo {
		if err := generateMojoOptionsGo(root); err != nil {
			return err
		}
	}
	return mpm.GenerateMojoPackages(root)
}

// The handwritten custom options proto has no Mojo AST descriptor, so it needs
// a separate protoc invocation after the standard library has been generated.
func generateMojoOptionsGo(root string) error {
	protoDir := filepath.Join(root, "protobuf")
	if _, err := os.Stat(filepath.Join(root, "package.mojo")); os.IsNotExist(err) {
		protoDir = filepath.Join(util.MojoPackagePath(root, "core"), "protobuf")
	}
	if _, err := os.Stat(filepath.Join(protoDir, "mojo", "mojo.proto")); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	tempDir, err := os.MkdirTemp("", "mojo-options-go-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)
	cmd := exec.Command("protoc", "-I.", "--go_out="+tempDir, "--go_opt=paths=source_relative", "mojo/mojo.proto")
	cmd.Dir = protoDir
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("bootstrap Go custom options mojo/mojo.proto: %w\n%s", err, out)
	}
	contents, err := os.ReadFile(filepath.Join(tempDir, "mojo", "mojo.pb.go"))
	if err != nil {
		return err
	}
	output := filepath.Join(root, "go", "pkg", "mojo", "mojo.pb.go")
	if err := os.MkdirAll(filepath.Dir(output), 0755); err != nil {
		return err
	}
	return os.WriteFile(output, contents, 0644)
}
