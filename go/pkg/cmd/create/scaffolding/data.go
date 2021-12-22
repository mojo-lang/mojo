package scaffolding

import (
	"github.com/mojo-lang/mojo/go/pkg/cmd/create/scaffolding/types"
	"strings"
)

type Data struct {
	Package *types.MojoPackage
}

func (d *Data) IsMojoPackage() bool {
	return d != nil && strings.HasPrefix(d.Package.FullName, "mojo.")
}
