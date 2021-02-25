package plugin

import (
	"sort"
)

type Plugin struct {
	Name string
	Priority int
	Parser
}

var Plugins ParserPlugins

func AddPlugin(name string, priority int, parser Parser) {
	Plugins = append(Plugins, &Plugin{
		Name: name,
		Priority: priority,
		Parser: parser,
	})
}

func SortPlugins() {
	sort.Sort(Plugins)
}

type ParserPlugins []*Plugin
func (p ParserPlugins) Len() int           { return len(p) }
func (p ParserPlugins) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ParserPlugins) Less(i, j int) bool { return p[i].Priority < p[j].Priority }
