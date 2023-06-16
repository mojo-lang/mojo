package syntax

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"

	sqlite "github.com/mojo-lang/mojo/go/pkg/sqlite/parser/syntax"
)

type StreamParser interface {
	ParseStream(fileName string, input antlr.CharStream) (*sql.SourceFile, error)
}

type Parser struct {
	StreamParser
}

func New(options core.Options) *Parser {
	p := &Parser{}

	if dialect, ok := options["dialect"]; ok {
		if d, ok := dialect.(sql.Dialect); ok {
			switch d {
			case sql.Dialect_DIALECT_SQLITE:
				p.StreamParser = sqlite.New(options)
			}
		}
	}
	return p
}

func (p Parser) ParseString(sql string) (*sql.SourceFile, error) {
	return p.ParseStream("", antlr.NewInputStream(sql))
}
