package syntax

import (
	"context"
	"regexp"

	"github.com/antlr4-go/antlr/v4"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/plugin"
	"github.com/mojo-lang/mojo/go/pkg/util"
)

type Parser struct {
}

func New(options core.Options) *Parser {
	_ = options
	return &Parser{}
}

func (p *Parser) ParseString(ctx context.Context, content string) (*lang.SourceFile, error) {
	input := antlr.NewInputStream(content)
	return p.ParseStream(plugin.ContextFilename(ctx), input)
}

var protoMismatched = regexp.MustCompile(`.*mismatched.+proto2.+proto3.+`)

func (p *Parser) ParseStream(fileName string, input antlr.CharStream) (*lang.SourceFile, error) {
	lexer := NewProtobuf3Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	errorListener := util.NewErrorListener(fileName, false)

	parser := NewProtobuf3Parser(stream)
	parser.AddErrorListener(errorListener)
	parser.BuildParseTrees = true

	tree := parser.Proto()
	if sourceFile, ok := tree.Accept(NewProtoVisitor()).(*lang.SourceFile); ok {
		if len(errorListener.Errors) == 0 {
			return sourceFile, nil
		}
	}

	if len(errorListener.Errors) > 0 {
		for _, e := range errorListener.Errors {
			msg := e.Error()
			if protoMismatched.MatchString(msg) {
				return nil, &ProtoError{}
			}
		}
		return nil, logs.NewErrorw("failed to parse proto3 file", "file", fileName, "error", errorListener.Errors.Error())
	}

	return nil, logs.NewErrorw("failed to parse proto3 file", "file", fileName)
}

func (p *Parser) ParseFile(ctx context.Context, fileName string) (*lang.SourceFile, error) {
	return plugin.ParseFile(p, ctx, fileName)
}
