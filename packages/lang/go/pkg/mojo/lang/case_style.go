package lang

import (
	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
)

type CaseStyle struct {
	*Attribute

	Style   string
	FuncMap map[string]func(string) string
}

func NewCaseStyle(attribute *Attribute) *CaseStyle {
	if attribute == nil || attribute.Name != "case_style" {
		return nil
	}

	style := &CaseStyle{
		Attribute: attribute,
		FuncMap: map[string]func(string) string{
			"snake":           strcase.ToSnake,
			"screaming_snake": strcase.ToScreamingSnake,
			"kebab":           strcase.ToKebab,
			"screaming_kebab": strcase.ToScreamingKebab,
			"camel":           strcase.ToCamel,
			"lower_camel":     strcase.ToLowerCamel,
		},
	}
	if len(style.Arguments) > 0 {
		style.Style, _ = style.Arguments[0].GetString()
	}
	if len(style.Style) == 0 {
		return nil
	}
	return style
}

func (c *CaseStyle) ToCase(str string) string {
	if c != nil {
		if fun, ok := c.FuncMap[c.Style]; ok {
			return fun(str)
		} else {
			logs.Warnw("failed to change case style, because of style is invalid", "style", c.Style)
		}
	}
	return str
}
