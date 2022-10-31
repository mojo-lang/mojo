package transformer

import (
	"strings"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type GenericTypeNamer struct {
}

func (g GenericTypeNamer) NameFrom(name string) string {
	if t, err := lang.ParseNominalTypeFullName(name); err != nil {
		return name
	} else {
		return g.Name(t)
	}
}

// Name
// typename: typename templates
// "Cached": "Cached{T}"
func (g GenericTypeNamer) Name(t *lang.NominalType) string {
	name := t.Name
	if strings.HasSuffix(name, "ed") || strings.HasSuffix(name, "able") {
		for _, argument := range t.GenericArguments {
			name += argument.Name
		}
	} else {
		argumentName := ""
		for _, argument := range t.GenericArguments {
			argumentName += argument.Name
		}
		name = argumentName + name
	}
	return name
}
