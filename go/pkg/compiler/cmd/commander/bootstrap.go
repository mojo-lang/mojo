package commander

import (
	"fmt"
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
	for _, name := range strings.Split(target, ",") {
		switch strings.TrimSpace(name) {
		case "go", "golang", "java":
		default:
			return fmt.Errorf("unsupported bootstrap target %q; use go or java", name)
		}
	}
	root := util.MojoRepositoryRoot(start)
	if root == "" {
		return fmt.Errorf("cannot find Mojo source repository from %q", start)
	}
	for _, name := range mpm.MojoPackageNames {
		b := Builder{Pwd: root, Path: filepath.Join(root, "packages", name), Targets: target}
		if err := b.Execute(); err != nil {
			return fmt.Errorf("bootstrap mojo.%s: %w", name, err)
		}
	}
	return mpm.GenerateMojoPackages(root)
}
