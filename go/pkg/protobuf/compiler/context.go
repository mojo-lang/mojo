package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
)

type Context struct {
	Options     []map[string]interface{}
	MojoItems   []interface{}
	Descriptors []interface{}
}

func (c *Context) SetOption(key string, value interface{}) {
	if len(c.Options) > 0 {
		c.Options[len(c.Options)-1][key] = value
	}
}

func (c *Context) DeleteOption(key string) {
	if len(c.Options) > 0 {
		delete(c.Options[len(c.Options)-1], key)
	}
}

// find the option from current context option all all parents
func (c *Context) GetOption(key string) interface{} {
	if len(c.Options) > 0 {
		for i := len(c.Options) - 1; i >= 0; i-- {
			if value := c.Options[i][key]; value != nil {
				return value
			}
		}
	}
	return nil
}

func (c *Context) Open(mojo interface{}, descriptor interface{}) {
	c.MojoItems = append(c.MojoItems, mojo)
	c.Descriptors = append(c.Descriptors, descriptor)
	c.Options = append(c.Options, make(map[string]interface{}))
}

func (c *Context) Close() {
	if len(c.MojoItems) > 0 {
		c.MojoItems = c.MojoItems[:len(c.MojoItems)-1]
	}
	if len(c.Descriptors) > 0 {
		c.Descriptors = c.Descriptors[:len(c.Descriptors)-1]
	}
	if len(c.Options) > 0 {
		c.Options = c.Options[:len(c.Options)-1]
	}
}

func (c *Context) GetPackage() *lang.Package {
	for i := len(c.MojoItems) - 1; i >= 0; i-- {
		if pkg, ok := c.MojoItems[i].(*lang.Package); ok {
			return pkg
		}
	}
	return nil
}

func (c *Context) GetSourceFile() *lang.SourceFile {
	for i := len(c.MojoItems) - 1; i >= 0; i-- {
		if file, ok := c.MojoItems[i].(*lang.SourceFile); ok {
			return file
		}
	}
	return nil
}

func (c *Context) GetFileDescriptor() *descriptor.FileDescriptor {
	for i := len(c.Descriptors) - 1; i >= 0; i-- {
		if d, ok := c.Descriptors[i].(*descriptor.FileDescriptor); ok {
			return d
		}
	}
	return nil
}

func (c *Context) GetCurrent() interface{} {
	if len(c.MojoItems) == 0 {
		return nil
	}

	return c.MojoItems[len(c.MojoItems)-1]
}

func (c *Context) GetCurrentDescriptor() interface{} {
	if len(c.Descriptors) == 0 {
		return nil
	}

	return c.Descriptors[len(c.Descriptors)-1]
}

func (c *Context) GetParentMessageDescriptor() *descriptor.MessageDescriptor {
	if len(c.Descriptors) > 1 {
		for i := len(c.Descriptors) - 2; i >= 0; i-- {
			if d, ok := c.Descriptors[i].(*descriptor.MessageDescriptor); ok {
				return d
			}
		}
	}

	return nil
}
