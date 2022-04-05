package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/data"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "strings"
)

type BoxedArray struct {
    *data.Data
}

func (b *BoxedArray) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {

    ba := &data.BoxedArray{}

    ba.PackageName = decl.GetPackageName()
    ba.GoPackageName = GetGoPackage(decl.GetPackageName())

    if pkg, _ := lang.GetStringAttribute(decl.Attributes, "go_package_name"); len(pkg) > 0 {
        ba.GoPackageName = pkg
    }

    ba.Name = decl.Name
    ba.EnclosingName = strings.Join(lang.GetEnclosingNames(decl.EnclosingType), ".")
    ba.FullName = GetFullName(ba.EnclosingName, ba.Name)
    ba.FieldName = strcase.ToCamel(decl.Type.Fields[0].Name)

    b.BoxedArrays = append(b.BoxedArrays, ba)
    return nil
}
