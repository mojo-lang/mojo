package printer

import (
	"github.com/mojo-lang/mojo/go/pkg/printer"
	"github.com/mojo-lang/mojo/go/pkg/sql/printer/ansi"
)

type SqlitePrinter struct {
	ansi.SQLPrinter
}

func New(printer *printer.Printer) *SqlitePrinter {
	return &SqlitePrinter{ansi.SQLPrinter{P: printer}}
}
