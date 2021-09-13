package compiler

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

type UnionField struct {
	Type   string // bool, number, string, array, object
	Format string

	Name string

	Boxed       bool
	UnionFields []*UnionField
}

type BoxedUnion struct {
	GoPackageName string
	Name          string
	FullName      string
	EnclosingName string
	Discriminator string

	BoolField    *UnionField
	NumberFields []*UnionField
	StringFields []*UnionField
	ArrayFields  []*UnionField
	ObjectFields []*UnionField
}

func (b *BoxedUnion) Compile(decl *lang.StructDecl) error {
	return nil
}
