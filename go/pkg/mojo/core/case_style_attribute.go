package core

import "github.com/mojo-lang/mojo/go/pkg/mojo/core/strcase"

const CaseStyleAttributeName = "case_style"
const CaseStyleAttributeFullName = "mojo.core.case_style"

func CaseStyler(style string) func(string) string {
	switch style {
	case "snake":
		return strcase.ToSnake
	case "screaming_snake":
		return strcase.ToScreamingSnake
	case "kebab":
		return strcase.ToKebab
	case "screaming_kebab":
		return strcase.ToScreamingSnake
	case "camel":
		return strcase.ToCamel
	case "lower_camel":
		return strcase.ToLowerCamel
	default:
		return func(s string) string {
			return s
		}
	}
}
