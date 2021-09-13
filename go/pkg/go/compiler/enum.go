package compiler

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"strings"
)

type EnumItem struct {
	RawValue string
	Value    string
	Number   int
}

type Enum struct {
	PackageName   string
	GoPackageName string
	Name          string
	FullName      string
	EnclosingName string

	Items []EnumItem
}

func (e *Enum) Compile(decl *lang.EnumDecl) error {
	e.GoPackageName = decl.PackageName
	e.EnclosingName = strings.Join(lang.GetEnclosingNames(decl.EnclosingType), ".")
	e.Name = decl.Name

	for i, value := range decl.Type.Enumerators {
		item := EnumItem{
			RawValue: value.Name,
			Value:    value.Name,
		}

		if number, err := lang.GetIntegerAttribute(value.Attributes, "number"); err != nil {
			item.Number = int(number)
		} else {
			item.Number = i
		}

		if alias, err := lang.GetStringAttribute(value.Attributes, "alias"); err != nil {
			item.Value = alias
		}

		e.Items = append(e.Items, item)
	}

	return nil
}
