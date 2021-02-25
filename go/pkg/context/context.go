package context

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type Context struct {
	Options []map[string]interface{}
	Values  []interface{}
}

func (c *Context) GetNamesForLogs() []interface{} {
	var ns []interface{}

	names := c.GetNames()
	for _, name := range names {
		ns = append(ns, name)
	}
	return ns
}

func (c *Context) GetNames() []string {
	var names []string
	if c == nil {
		return names
	}

	for _, value := range c.Values {
		switch v := value.(type) {
		case *lang.Package:
			names = append(names, "package", v.GetName())
		case *lang.SourceFile:
			names = append(names, "file", v.GetName())
		case *lang.StructDecl:
			names = append(names, "struct", v.GetName())
		case *lang.EnumDecl:
			names = append(names, "enum", v.GetName())
		case *lang.TypeAliasDecl:
			names = append(names, "type_alias", v.GetName())
		case *lang.InterfaceDecl:
			names = append(names, "interface", v.GetName())
		}
	}
	return names
}

func (c *Context) SetOptions(options map[string]interface{}) {
	for key, value := range options {
		c.SetOption(key, value)
	}
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

func (c *Context) Open(value interface{}) {
	c.Values = append(c.Values, value)
	c.Options = append(c.Options, make(map[string]interface{}))
}

func (c *Context) Close() {
	if len(c.Values) > 0 {
		c.Values = c.Values[:len(c.Values)-1]
	}

	if len(c.Options) > 0 {
		c.Options = c.Options[:len(c.Options)-1]
	}
}

func (c *Context) GetPackage() *lang.Package {
	for i := len(c.Values) - 1; i >= 0; i-- {
		if pkg, ok := c.Values[i].(*lang.Package); ok {
			return pkg
		}
	}
	return nil
}

func (c *Context) GetSourceFile() *lang.SourceFile {
	for i := len(c.Values) - 1; i >= 0; i-- {
		if file, ok := c.Values[i].(*lang.SourceFile); ok {
			return file
		}
	}
	return nil
}

func (c *Context) PreviousValue(index int) interface{} {
	if len(c.Values) == 0 {
		return nil
	}

	if index < 0 || index >= len(c.Values) {
		return nil
	}

	return c.Values[len(c.Values)-1-index]
}

func (c *Context) CurrentValue() interface{} {
	if len(c.Values) == 0 {
		return nil
	}

	return c.Values[len(c.Values)-1]
}

func (c *Context) OrderlyValues() []interface{} {
	if len(c.Values) == 0 {
		return []interface{}{}
	}

	values := make([]interface{}, 0, len(c.Values))
	for i := len(c.Values) - 1; i >= 0; i-- {
		values = append(values, c.Values[i])
	}
	return values
}
