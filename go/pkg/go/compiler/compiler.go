package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/context"
	"strings"
)

func GetGoPackage(pkg string) string {
	segments := strings.Split(pkg, ".")
	if len(segments) > 0 {
		return segments[len(segments)-1]
	}
	return ""
}

func GetFullName(enclosing string, name string) string {
	if len(enclosing) == 0 {
		return name
	}
	return strings.ReplaceAll(enclosing, ".", "_") + "_" + name
}

type Compiler struct {
	*Data
}

func (c *Compiler) CompileDbJSON(ctx context.Context, decl *lang.StructDecl) error {
	dbJSONs := make(DbJSONs, 0, 0)
	dbJSONs.CompileStruct(ctx, decl)
	if len(dbJSONs) > 0 {
		c.Data.DbJSONs = append(c.Data.DbJSONs, dbJSONs...)
	}
	return nil
}

func (c *Compiler) Compile(ctx context.Context, decl interface{}) error {
	switch declaration := decl.(type) {
	case *lang.StructDecl:
		if union, err := lang.GetBoolAttribute(declaration.Attributes, "union"); union && err != nil {
			boxedUnion := &BoxedUnion{}
			err = boxedUnion.Compile(declaration)
			if err != nil {
				return err
			}
			c.Data.BoxedUnions = append(c.Data.BoxedUnions, boxedUnion)
		} else {
			t := declaration.Type.Fields[0].Type.Name
			switch t {
			case "Array":
				boxedArray := &BoxedArray{}
				err = boxedArray.Compile(declaration)
				if err != nil {
					return err
				}

				if !c.Data.IsExist(boxedArray.FullName) {
					c.Data.NameIndex[boxedArray.FullName] = true
					c.Data.BoxedArrays = append(c.Data.BoxedArrays, boxedArray)
				}
			case "Map":
				//boxedMap := &BoxedMap{}
				//err = boxedMap.Compile(declaration)
				//if err != nil {
				//	return err
				//}
				//c.Data.BoxedDictionaries = append(c.Data.BoxedDictionaries, boxedMap)
			}
		}
	case *lang.EnumDecl:
		enum := &Enum{}
		err := enum.Compile(declaration)
		if err != nil {
			return err
		}
		c.Data.Enums = append(c.Data.Enums, enum)
	case *lang.Package:
		//mod := &GoMod{}
		//err := mod.Compile(declaration)
		//if err != nil {
		//	return err
		//}
		//c.Data.GoMod = mod
	}

	return nil
}
