package transformer

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

func GenericArgumentLabel(argument *lang.NominalType) string {
	if label, _ := argument.GetStringAttribute(core.LabelAttributeName); len(label) > 0 {
		return label
	}

	if labelFormat, _ := argument.GetStringAttribute(core.LabelFormatAttributeName); len(labelFormat) > 0 {
		if labelFormat == "{}" {
			return strcase.ToSnake(argument.Name)
		}
	}

	if argument.IsScalar() {
		return strcase.ToSnake(argument.Name) + "_val"
	}
	return strcase.ToSnake(argument.Name)
}
