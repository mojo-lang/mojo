package syntax

import (
	"context"
	"io/fs"
	"path"

	"github.com/antlr4-go/antlr/v4"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"

	"github.com/mojo-lang/mojo/go/pkg/util"
)

type Parser struct {
}

func New(options core.Options) *Parser {
	_ = options
	return &Parser{}
}

func (p Parser) ParseString(mojo string) (*sql.SourceFile, error) {
	input := antlr.NewInputStream(mojo)
	return p.ParseStream("", input)
}

func (p Parser) ParseStream(fileName string, input antlr.CharStream) (*sql.SourceFile, error) {
	lexer := NewSQLiteLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	errorListener := util.NewErrorListener(fileName, false)

	parser := NewSQLiteParser(stream)
	parser.AddErrorListener(errorListener)
	parser.BuildParseTrees = true

	tree := parser.Parse()
	if sourceFile, ok := tree.Accept(NewSqlSmtVisitor()).(*sql.SourceFile); ok {
		if len(errorListener.Errors) == 0 {
			return sourceFile, nil
		}
	}

	if len(errorListener.Errors) > 0 {
		return nil, logs.NewErrorw("failed to parse proto3 file", "file", fileName, "error", errorListener.Errors.Error())
	}

	return nil, logs.NewErrorw("failed to parse proto3 file", "file", fileName)
}

func (p Parser) ParseFile(ctx context.Context, fileName string, fileSys fs.FS) (*sql.SourceFile, error) {
	_ = ctx
	if bytes, err := fs.ReadFile(fileSys, fileName); err != nil {
		return nil, err
	} else {
		if sourceFile, err := p.ParseStream(fileName, antlr.NewInputStream(string(bytes))); err != nil {
			return nil, err
		} else {
			sourceFile.Name = path.Base(fileName)
			sourceFile.FullName = fileName

			return sourceFile, nil
		}
	}
}
