package syntax

import (
	"regexp"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/plugin"
	proto2 "github.com/mojo-lang/mojo/go/pkg/protobuf2/parser/syntax"
	proto3 "github.com/mojo-lang/mojo/go/pkg/protobuf3/parser/syntax"
)

type Parser struct {
	Proto2 *proto2.Parser
	Proto3 *proto3.Parser
}

func New(options core.Options) *Parser {
	return &Parser{
		Proto2: proto2.New(options),
		Proto3: proto3.New(options),
	}
}

var proto3Regex = regexp.MustCompile(`syntax[ \t\r\n]*=[ \t\r\n]*['"]proto3['"]`)

func (p *Parser) ParseString(ctx context.Context, content string) (*lang.SourceFile, error) {
	if proto3Regex.MatchString(content) {
		return p.Proto3.ParseString(ctx, content)
	} else {
		return p.Proto2.ParseString(ctx, content)
	}
}

func (p *Parser) ParseFile(ctx context.Context, filename string) (*lang.SourceFile, error) {
	return plugin.ParseFile(p, ctx, filename)
}
