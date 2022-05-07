package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/data"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "strings"
)

type BoxedMap struct {
    *data.Data
}

func (b *BoxedMap) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
    bm := &data.BoxedMap{}
    bm.PackageName = decl.GetPackageName()
    bm.GoPackageName = GetGoPackage(decl.GetPackageName())
    bm.Name = decl.Name
    bm.EnclosingName = strings.Join(lang.GetEnclosingNames(decl.EnclosingType), ".")
    bm.FullName = GetFullName(bm.EnclosingName, bm.Name)
    bm.FieldName = strcase.ToCamel(decl.Type.Fields[0].Name)

    b.BoxedMaps = append(b.BoxedMaps, bm)
    return nil
}
