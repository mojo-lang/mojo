package plugin

import "github.com/mojo-lang/core/go/pkg/mojo/core"

type BasicPlugin struct {
	Name          string
	Group         string
	GroupPriority int
	Priority      int

	Creator func(options core.Options) Plugin

	MarkedPackages map[string]bool
}

func (p *BasicPlugin) GetName() string {
	return p.Name
}

func (p *BasicPlugin) GetGroup() string {
	return p.Group
}

func (p *BasicPlugin) GetPriority() int {
	return p.Priority
}

func (p *BasicPlugin) GetGroupPriority() int {
	return p.GroupPriority
}

func (p *BasicPlugin) Create(options core.Options) Plugin {
	if p.Creator != nil {
		return p.Creator(options)
	}
	return nil
}
