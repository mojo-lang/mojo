package compiler

import (
	"strings"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/packages/db/go/pkg/mojo/db"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	context2 "github.com/mojo-lang/mojo/go/pkg/context"
	data2 "github.com/mojo-lang/mojo/go/pkg/go/generator/data"
)

type DbJson struct {
	*data2.Data
}

func (j *DbJson) CompileStruct(ctx context2.Context, decl *lang.StructDecl) error {
	if decl.Type == nil || len(decl.Type.Fields) == 0 {
		return nil
	}

	pkg := context2.Package(ctx)
	for _, field := range decl.Type.Fields {
		if !field.HasAttribute(db.JSONAttributeName) {
			continue
		}

		newName := decl.Name + strcase.ToCamel(field.Name)
		dbJSON := &data2.DbJSON{
			PackageName:        pkg.FullName,
			GoPackageName:      lang.GetGoPackageName(pkg.FullName),
			Name:               newName,
			UnderlyingTypeName: field.Type.Name,
			StructType:         true,
			FullName:           newName,
		}

		dbJSON.InDbPkg = dbJSON.GoPackageName == "db"

		fieldTypeFullName := field.Type.GetFullName()
		if fieldTypeFullName == core.ArrayTypeFullName || fieldTypeFullName == core.MapTypeFullName {
			dbJSON.UnderlyingTypeName = ""
			dbJSON.StructType = false
		}

		if field.Type.PackageName == pkg.FullName {
			dbJSON.Name = field.Type.Name
			dbJSON.FullName = GetFullName(strings.Join(field.Type.GetEnclosingNames(), "."), field.Type.Name)
			dbJSON.UnderlyingTypeName = ""
		}

		j.DbJSONs = append(j.DbJSONs, dbJSON)
	}

	return nil
}
