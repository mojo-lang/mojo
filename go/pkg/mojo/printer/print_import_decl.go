package printer

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintImportDecl(ctx context.Context, file *lang.ImportDecl) *Printer {
	if file == nil || p.GetError() != nil {
		return p
	}

	return p
}
