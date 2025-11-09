package printer

import (
	"github.com/mojo-lang/mojo/go/pkg/compiler/printer"
	"github.com/mojo-lang/mojo/go/pkg/compiler/sql/printer/ansi"
)

type SqlitePrinter struct {
	ansi.SQLPrinter
}

func New(printer *printer.Printer) *SqlitePrinter {
	return &SqlitePrinter{ansi.SQLPrinter{Printer: printer}}
}
