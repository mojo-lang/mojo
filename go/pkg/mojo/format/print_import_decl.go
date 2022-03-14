package format

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func (p *Printer) PrintImportDecl(ctx context.Context, file *lang.ImportDecl) *Printer {
    if file == nil || p == nil || p.Error != nil {
        return p
    }

    return p
}
