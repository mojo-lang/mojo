package printer

import (
	"context"
	"io"

	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"

	"github.com/mojo-lang/mojo/go/pkg/printer"
	"github.com/mojo-lang/mojo/go/pkg/sql/printer/ansi"

	sqlite "github.com/mojo-lang/mojo/go/pkg/sqlite/printer"
)

type Config struct {
	*printer.Config

	Dialect sql.Dialect
}

// Printer dialect
type Printer struct {
	P   *printer.Printer
	SQL SQLPrinter
}

func New(config *Config, writer io.Writer) *Printer {
	p := &Printer{
		P: printer.New(config.Config, writer),
	}
	switch config.Dialect {
	case sql.Dialect_DIALECT_SQLITE:
		p.SQL = sqlite.New(p.P)
	default:
		p.SQL = ansi.New(p.P)
	}

	return p
}

func (p *Printer) PrintExpression(ctx context.Context, expr *sql.Expression) *Printer {
	if p != nil && expr != nil {
		p.SQL.PrintExpression(ctx, expr)
	}
	return p
}
