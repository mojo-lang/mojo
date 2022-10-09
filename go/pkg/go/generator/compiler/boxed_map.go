package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	data2 "github.com/mojo-lang/mojo/go/pkg/go/generator/data"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"strings"
)

type BoxedMap struct {
	*data2.Data
}

func (b *BoxedMap) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
	bm := &data2.BoxedMap{}
	bm.PackageName = decl.GetPackageName()
	bm.GoPackageName = GetGoPackage(decl.GetPackageName())
	bm.Name = decl.Name
	bm.EnclosingName = strings.Join(lang.GetEnclosingNames(decl.Enclosing), ".")
	bm.FullName = GetFullName(bm.EnclosingName, bm.Name)
	bm.FieldName = strcase.ToCamel(decl.Type.Fields[0].Name)

	b.BoxedMaps = append(b.BoxedMaps, bm)
	return nil
}
