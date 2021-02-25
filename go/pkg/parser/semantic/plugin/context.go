package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type Context struct {
	context.Context
	lang.ScopeContext
}

func NewContext() *Context {
	return &Context{
		Context: context.Context{},
	}
}

func (c *Context) Open(value interface{}) {
	scope := lang.GetScope(value)
	if scope == nil {
		scope = lang.NewScope()
		lang.SetScope(value, scope)
	}

	c.ScopeContext.Open(scope)
	c.Context.Open(value)
}

func (c *Context) Close() {
	c.ScopeContext.Close()
	c.Context.Close()
}
