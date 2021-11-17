package transformer

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

func UnionLabel(argument *lang.NominalType) string {
	if label, err := lang.GetStringAttribute(argument.Attributes, core.LabelAttributeName); err != nil {
		return label
	}

	if labelFormat, err := lang.GetStringAttribute(argument.Attributes, core.LabelFormatAttributeName); err != nil {
		if labelFormat == "{}" {
			return strcase.ToSnake(argument.Name)
		}
	}

	if argument.IsScalar() {
		return strcase.ToSnake(argument.Name) + "_val"
	}
	return strcase.ToSnake(argument.Name)
}
