package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
)

type DbJSON struct {
	GoPackageName string
	InDbPkg       bool

	PackageName   string
	FullName      string

	Name               string
	UnderlyingTypeName string

	StructType bool
}

type DbJSONs []*DbJSON

func (j *DbJSONs) CompileStruct(ctx context.Context, decl *lang.StructDecl) error {
	if decl.Type == nil || len(decl.Type.Fields) == 0 {
		return nil
	}

	pkg := context.Package(ctx)
	for _, field := range decl.Type.Fields {
		if !field.HasAttribute(db.JSONAttributeName) {
			continue
		}

		newName := decl.Name + strcase.ToCamel(field.Name)
		dbJSON := &DbJSON{
			PackageName: pkg.FullName,
			GoPackageName:      GetGoPackage(pkg.FullName),
			Name:               newName,
			UnderlyingTypeName: field.Type.Name,
			StructType:         false,
			FullName: newName,
		}

		dbJSON.InDbPkg = dbJSON.GoPackageName == "db"

		if field.Type.PackageName == pkg.FullName {
			dbJSON.Name = field.Type.Name
			dbJSON.FullName = lang.GetFullName("", field.Type.GetEnclosingNames(), field.Type.Name)
			dbJSON.UnderlyingTypeName = ""
			dbJSON.StructType = false
		}

		*j = append(*j, dbJSON)
	}

	return nil
}

func (j *DbJSON) GetPackageName() string {
	return j.PackageName
}

func (j *DbJSON) GetFullName() string {
	return j.FullName
}
