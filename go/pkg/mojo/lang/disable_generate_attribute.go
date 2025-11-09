package lang

import (
	"errors"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

type GenerateOption map[string][]string

func (g *GenerateOption) Parse(attribute *Attribute) error {
	if g != nil && attribute != nil {
		if mapLiteral := attribute.GetMapLiteralArgument(); mapLiteral != nil {
			for _, entry := range mapLiteral.Entries {
				if null := entry.Value.GetNullLiteralExpr(); null != nil {
					(*g)[entry.Key] = []string{}
				}
				if array := entry.Value.GetArrayLiteralExpr(); array != nil {
					for _, elem := range array.Elements {
						if val := elem.GetStringLiteralExpr().GetValue(); len(val) > 0 {
							(*g)[entry.Key] = append((*g)[entry.Key], val)
						}
					}
				}
			}
		}
	}
	return nil
}

func (g GenerateOption) Including(module string, category string) bool {
	if len(g) == 1 {
		if _, ok := g["all"]; ok {
			return true
		}
	}

	if categorys, ok := g[module]; ok {
		if len(categorys) == 1 && categorys[0] == "all" {
			return true
		}
		if len(category) == 0 {
			return true
		}
		for _, c := range categorys {
			if c == category {
				return true
			}
		}
	}
	return false
}

func GetDisableGenerateAttribute(attributes []*Attribute) (GenerateOption, error) {
	attribute := GetAttribute(attributes, core.DisableGenerateAttributeName)
	if attribute == nil {
		return nil, errors.New("NotFound")
	}

	if len(attribute.Arguments) == 0 {
		return GenerateOption{"all": []string{}}, nil
	}

	generateOption := make(GenerateOption)
	if err := generateOption.Parse(attribute); err != nil {
		return nil, err
	}

	return generateOption, nil
}
