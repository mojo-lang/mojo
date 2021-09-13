package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
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

func (c *Compiler) Compile(decl interface{}) error {
	switch declaration := decl.(type) {
	case *lang.StructDecl:
		if union, err := lang.GetBoolAttribute(declaration.Attributes, "union"); union && err != nil {
			//boxedUnion := &BoxedUnion{}
			//err = boxedUnion.Compile(declaration)
			//if err != nil {
			//	return err
			//}
			//c.Data.BoxedUnions = append(c.Data.BoxedUnions, boxedUnion)
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
			case "Dictionary":
				//boxedDictionary := &BoxedDictionary{}
				//err = boxedDictionary.Compile(declaration)
				//if err != nil {
				//	return err
				//}
				//c.Data.BoxedDictionaries = append(c.Data.BoxedDictionaries, boxedDictionary)
			case "String":
				//boxedString := &BoxedString{}
				//err = boxedString.Compile(declaration)
				//if err != nil {
				//	return err
				//}
				//c.Data.BoxedStrings = append(c.Data.BoxedStrings, boxedString)
			}
		}
	case *lang.EnumDecl:
		//enum := &Enum{}
		//err := enum.Compile(declaration)
		//if err != nil {
		//	return err
		//}
		//c.Data.Enums = append(c.Data.Enums, enum)
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
