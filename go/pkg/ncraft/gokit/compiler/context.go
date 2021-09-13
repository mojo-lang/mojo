package compiler

import (
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type Context struct {
	*context.Context

	// go package repos index for all the package
	GoPackageImports map[string]string
}

func NewContext() *Context {
	return &Context{
		Context: &context.Context{
		},
		GoPackageImports: make(map[string]string),
	}
}
