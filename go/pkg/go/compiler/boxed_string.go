package compiler

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type BoxedString struct {
	GoPackageName string
	Name          string
	FullName      string
	EnclosingName string
}

func (b *BoxedString) Compile(decl *lang.StructDecl) error {
	return nil
}
