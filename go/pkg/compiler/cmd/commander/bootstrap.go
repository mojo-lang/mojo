package commander

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/compiler/mojo/mpm"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

// DefaultBootstrapTargets regenerates all standard library outputs. The api
// target includes Go, Protobuf, Markdown documentation, and OpenAPI definitions.
const DefaultBootstrapTargets = "api,java"

// Bootstrap regenerates the selected standard library outputs, then refreshes
// embedded snapshots. With no targets, it regenerates every supported output.
func Bootstrap(start string, targets ...string) error {
	target := DefaultBootstrapTargets
	if len(targets) > 0 {
		target = strings.Join(targets, ",")
	}
	generateGo := false
	generateAPI := false
	for _, name := range strings.Split(target, ",") {
		switch strings.TrimSpace(name) {
		case "api", "go", "golang":
			generateGo = true
			generateAPI = generateAPI || strings.TrimSpace(name) == "api"
		case "java":
		default:
			return fmt.Errorf("unsupported bootstrap target %q; use api, go, or java", name)
		}
	}
	root := util.MojoRepositoryRoot(start)
	if root == "" {
		return fmt.Errorf("cannot find Mojo source repository from %q", start)
	}

	var staging string
	if _, err := os.Stat(filepath.Join(root, "package.mojo")); err == nil {
		b := Builder{Pwd: root, Path: root, Targets: target}
		if generateAPI {
			staging, err = os.MkdirTemp("", "mojo-bootstrap-docs-")
			if err != nil {
				return err
			}
			defer os.RemoveAll(staging)
			b.documentOutput = filepath.Join(staging, "document")
			b.openapiOutput = filepath.Join(staging, "openapi")
		}
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
	if staging != "" {
		if err := publishBootstrapOutput(filepath.Join(staging, "document"), filepath.Join(root, "document"), ".md"); err != nil {
			return err
		}
		if err := publishBootstrapOutput(filepath.Join(staging, "openapi"), filepath.Join(root, "openapi"), ".schema.json", ".yaml"); err != nil {
			return err
		}
	}
	return mpm.GenerateMojoPackages(root)
}

// Publish only after every package has built successfully. Comparing complete
// trees also removes obsolete types in directories no generator visits anymore.
func publishBootstrapOutput(staging, output string, suffixes ...string) error {
	current := make(map[string]bool)
	if _, err := os.Stat(staging); err == nil {
		err = filepath.WalkDir(staging, func(name string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() {
				return nil
			}
			rel, err := filepath.Rel(staging, name)
			if err != nil {
				return err
			}
			content, err := os.ReadFile(name)
			if err != nil {
				return err
			}
			target := filepath.Join(output, rel)
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}
			if err := os.WriteFile(target, content, 0644); err != nil {
				return err
			}
			current[rel] = true
			return nil
		})
		if err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	if _, err := os.Stat(output); os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	return filepath.WalkDir(output, func(name string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(output, name)
		if err != nil {
			return err
		}
		if !current[rel] {
			for _, suffix := range suffixes {
				if strings.HasSuffix(name, suffix) {
					return os.Remove(name)
				}
			}
		}
		return nil
	})
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
