package compiler

import (
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/compiler/context"
	data2 "github.com/mojo-lang/mojo/go/pkg/compiler/go/generator/data"
)

type BoxedArray struct {
	*data2.Data
}

func (b *BoxedArray) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
	_ = ctx
	ba := &data2.BoxedArray{}

	ba.PackageName = decl.GetPackageName()
	ba.GoPackageName = GetGoPackage(decl.GetPackageName())

	if pkg, _ := lang.GetStringAttribute(decl.Attributes, "go_package_name"); len(pkg) > 0 {
		ba.GoPackageName = pkg
	}

	ba.Name = decl.Name
	ba.EnclosingName = strings.Join(lang.GetEnclosingNames(decl.Enclosing), ".")
	ba.FullName = GetFullName(ba.EnclosingName, ba.Name)
	ba.FieldName = strcase.ToCamel(decl.Type.Fields[0].Name)

	b.BoxedArrays = append(b.BoxedArrays, ba)
	return nil
}
