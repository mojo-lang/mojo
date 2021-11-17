package compiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
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
	WrapName      string

	Items []EnumItem
}

func (e *Enum) Compile(decl *lang.EnumDecl) error {
	e.PackageName = decl.PackageName
	e.GoPackageName = GetGoPackage(decl.PackageName)
	e.EnclosingName = strings.Join(lang.GetEnclosingNames(decl.EnclosingType), "_")
	e.Name = decl.Name
	e.FullName = e.Name
	if len(e.EnclosingName) > 0 {
		e.WrapName = e.EnclosingName
		e.FullName = strings.Join([]string{e.EnclosingName, e.Name}, "_")
	} else {
		e.WrapName = e.Name
	}

	caseStyle, _ := lang.GetStringAttribute(decl.Attributes, core.CaseStyleAttributeName)
	//if len(caseStyle) == 0 {
	//	caseStyle = "kebab"
	//}

	for i, value := range decl.Type.Enumerators {
		if value.Name == "unspecified" {
			continue
		}

		item := EnumItem{
			RawValue: value.Name,
			Value:    value.Name,
		}

		if number, err := lang.GetIntegerAttribute(value.Attributes, "number"); err == nil {
			item.Number = int(number)
		} else {
			item.Number = i
		}

		if alias, err := lang.GetStringAttribute(value.Attributes, "alias"); err == nil && len(alias) > 0 {
			item.Value = alias
		} else {
			item.Value = core.CaseStyler(caseStyle)(item.Value)
		}

		e.Items = append(e.Items, item)
	}

	return nil
}

func (b *Enum) GetPackageName() string {
	return b.PackageName
}

func (b *Enum) GetFullName() string  {
	return b.FullName
}