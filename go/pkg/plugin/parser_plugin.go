package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/plugin/parser"
)

type parserPlugin struct {
	basicPlugin
	Parser parser.PackageParser
}

var ParserPlugins Plugins

func AddParserPlugin(name string, priority int, parser parser.PackageParser) {
	ParserPlugins = append(ParserPlugins, &parserPlugin{
		basicPlugin: basicPlugin{
			Name:     name,
			Priority: priority,
		},
		Parser: parser,
	})
}

func ParsePackage(p Plugin, ctx context.Context, pkg *lang.Package) error {
	if plugin, ok := p.(*parserPlugin); ok {
		return plugin.Parser.Parse(ctx, pkg)
	}
	return nil
}
