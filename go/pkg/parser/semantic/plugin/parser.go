package plugin

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type Parser interface {
	Parse(ctx *Context, pkg *lang.Package, options map[string]interface{}) error
}
