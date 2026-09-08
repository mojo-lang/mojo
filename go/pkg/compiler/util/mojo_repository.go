package util

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"
)

// MojoRepositoryRoot locates the source checkout, from either its root or a
// descendant such as go/ or mojo/core/. It does not depend on the cwd.
func MojoRepositoryRoot(start string) string {
	dir, err := filepath.Abs(start)
	if err != nil {
		return ""
	}
	for {
		contents, err := os.ReadFile(filepath.Join(dir, "go", "go.mod"))
		if err == nil {
			fields := strings.Fields(string(contents))
			if len(fields) >= 2 && fields[0] == "module" && fields[1] == lang.MojoGoModule {
				for _, manifest := range []string{filepath.Join(dir, "package.mojo"), filepath.Join(dir, "packages", "core", "package.mojo")} {
					if info, err := os.Stat(manifest); err == nil && !info.IsDir() {
						return dir
					}
				}
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

// MojoPackagePath supports the unified source tree and legacy component projects.
func MojoPackagePath(root, name string) string {
	if _, err := os.Stat(filepath.Join(root, "package.mojo")); err == nil {
		return filepath.Join(root, "mojo", name)
	}
	return filepath.Join(root, "packages", name)
}
