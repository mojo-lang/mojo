package plugin

import "sort"

type Plugin interface {
	GetName() string
	GetPriority() int
}

type basicPlugin struct {
	Name     string
	Priority int
}

func (p *basicPlugin) GetName() string {
	return p.Name
}

func (p *basicPlugin) GetPriority() int {
	return p.Priority
}

type Plugins []Plugin

func Sort(plugins Plugins) {
	sort.Sort(plugins)
}

func (p Plugins) Len() int           { return len(p) }
func (p Plugins) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Plugins) Less(i, j int) bool { return p[i].GetPriority() < p[j].GetPriority() }
