package commander

import (
	"fmt"
	"path/filepath"

	"github.com/mojo-lang/mojo/go/pkg/compiler/mojo/mpm"
	"github.com/mojo-lang/mojo/go/pkg/compiler/util"
)

// Bootstrap regenerates the standard library into the checkout's shared Go
// module, then refreshes the source and protobuf snapshots embedded in the CLI.
func Bootstrap(start string) error {
	root := util.MojoRepositoryRoot(start)
	if root == "" {
		return fmt.Errorf("cannot find Mojo source repository from %q", start)
	}
	for _, name := range mpm.MojoPackageNames {
		b := Builder{Pwd: root, Path: filepath.Join(root, "packages", name), Targets: "go"}
		if err := b.Execute(); err != nil {
			return fmt.Errorf("bootstrap mojo.%s: %w", name, err)
		}
	}
	return mpm.GenerateMojoPackages(root)
}
