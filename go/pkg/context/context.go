package context

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)
//
//const TypeKey = "@type"
//const PackageKey = "@type/Package"
//const SourceFileKey = "@type/SourceFile"
//const InterfaceKey = "@type/Interface"
//const StructKey = "@type/Struct"
//const EnumKey = "@type/Enum"
//const TypeAliasKey = "@type/TypeAlias"
//
//type valuesContext struct {
//	context.Context
//	Options
//}
//
//func (*valuesContext) Deadline() (deadline time.Time, ok bool) {
//	return
//}
//
//func (*valuesContext) Done() <-chan struct{} {
//	return nil
//}
//
//func (*valuesContext) Err() error {
//	return nil
//}
//
//func (c *valuesContext) Value(key interface{}) interface{} {
//	if v, ok := c.Options[key.(string)]; ok {
//		return v
//	}
//	return c.Context.Value(key)
//}
//
//func WithOptions(ctx context.Context, options Options) context.Context {
//	return &valuesContext{
//		Context: ctx,
//		Options: options,
//	}
//}
//
//func WithValues(ctx context.Context, kvs ...interface{}) context.Context {
//	options := make(Options)
//	for i := 0; i < len(kvs)-1; i += 2 {
//		options[kvs[i].(string)] = kvs[i+1]
//	}
//	return WithOptions(ctx, options)
//}
//
//func WithTypeValue(ctx context.Context, value interface{}) context.Context {
//	key := ""
//	switch value.(type) {
//	case *lang.Package:
//		key = PackageKey
//	case *lang.SourceFile:
//		key = SourceFileKey
//	case *lang.StructDecl:
//		key = StructKey
//	case *lang.EnumDecl:
//		key = EnumKey
//	case *lang.TypeAliasDecl:
//		key = TypeAliasKey
//	case *lang.InterfaceDecl:
//		key = InterfaceKey
//	}
//	return WithValues(ctx, TypeKey, value, key, value)
//}
//
//func Package(ctx context.Context) *lang.Package {
//	return ctx.Value(PackageKey).(*lang.Package)
//}
//
//func Values(ctx context.Context, key string) []interface{} {
//	var values []interface{}
//	if c, ok := ctx.(*valuesContext); ok {
//		if v, o := c.Options[key]; o {
//			values = append(values, v)
//		}
//
//		values = append(values, Values(c.Context, key)...)
//	}
//	return values
//}
//
//func PreviousValue(ctx context.Context, key string, index int) interface{} {
//	values := Values(ctx, key)
//	if index < 0 || index >= len(values) {
//		return nil
//	}
//	return values[index]
//}
//
//func TypeValues(ctx context.Context) []interface{} {
//	return Values(ctx, TypeKey)
//}
//
//func TypeValue(ctx context.Context) interface{} {
//	return ctx.Value(TypeKey)
//}
//
//func TypePreviousValue(ctx context.Context, index int) interface{} {
//	return PreviousValue(ctx, TypeKey, index)
//}
//
//func IsTypeEnclosed(ctx context.Context) bool {
//	_, ok := TypePreviousValue(ctx, 1).(EnclosingTypeDecl)
//	return ok
//}
//
//type EnclosingTypeDecl interface {
//	EnclosingTypeDecl()
//}

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

func (c *Context) IsEnclosed() bool {
	previous := c.PreviousValue(1)
	switch previous.(type) {
	case *lang.StructDecl, *lang.InterfaceDecl:
		return true
	default:
		return false
	}
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
