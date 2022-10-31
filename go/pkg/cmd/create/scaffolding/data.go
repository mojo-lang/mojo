package scaffolding

import (
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/cmd/create/scaffolding/types"
)

type Data struct {
	Package *types.MojoPackage
}

func (d *Data) IsMojoPackage() bool {
	return d != nil && strings.HasPrefix(d.Package.FullName, "mojo.")
}
