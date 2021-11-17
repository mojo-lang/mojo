package compiler

import (
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"strings"
)

type BoxedArray struct {
	PackageName   string
	GoPackageName string
	Name          string
	FullName      string
	EnclosingName string

	FieldName string
}

func (b *BoxedArray) Compile(decl *lang.StructDecl) error {
	b.PackageName = decl.GetPackageName()
	b.GoPackageName = GetGoPackage(decl.GetPackageName())

	if pkg, _ := lang.GetStringAttribute(decl.Attributes, "go_package_name"); len(pkg) > 0 {
		b.GoPackageName = pkg
	}

	b.Name = decl.Name
	b.EnclosingName = strings.Join(lang.GetEnclosingNames(decl.EnclosingType), ".")
	b.FullName = GetFullName(b.EnclosingName, b.Name)
	b.FieldName = strcase.ToCamel(decl.Type.Fields[0].Name)
	return nil
}

func (b *BoxedArray) GetPackageName() string {
	return b.PackageName
}

func (b *BoxedArray) GetFullName() string  {
	return b.FullName
}