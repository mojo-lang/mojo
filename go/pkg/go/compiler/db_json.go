package compiler

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/db/go/pkg/mojo/db"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/data"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

type DbJson struct {
    *data.Data
}

func (j *DbJson) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
    if decl.Type == nil || len(decl.Type.Fields) == 0 {
        return nil
    }

    pkg := context.Package(ctx)
    for _, field := range decl.Type.Fields {
        if !field.HasAttribute(db.JSONAttributeName) {
            continue
        }

        newName := decl.Name + strcase.ToCamel(field.Name)
        dbJSON := &data.DbJSON{
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
            dbJSON.FullName = lang.GetFullName("", field.Type.GetEnclosingNames(), field.Type.Name)
            dbJSON.UnderlyingTypeName = ""
        }

        j.DbJSONs = append(j.DbJSONs, dbJSON)
    }

    return nil
}
