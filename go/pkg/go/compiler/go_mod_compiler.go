package compiler

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

const GoModTemplate = `
`
type GoMod struct {
}

func (g *GoMod) Generate(pkg *lang.Package, path string) error {
	return nil
}
