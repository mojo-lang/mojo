package precompiler

import (
	"errors"
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
	"github.com/mojo-lang/mojo/go/pkg/mojo/compiler/transformer"
)

func CompileMapToStruct(ctx context.Context, t *lang.NominalType) (*lang.StructDecl, error) {
	if t.GetFullName() != core.MapTypeFullName {
		return nil, errors.New("")
	}

	if len(t.GenericArguments) != 2 {
		return nil, errors.New("")
	}

	t.Attributes = lang.SetIntegerAttribute(t.Attributes, core.NumberAttributeName, 1)

	s := &lang.StructDecl{}
	s.Type = &lang.StructType{
		Fields: []*lang.ValueDecl{{
			Name: "vals",
			Type: t,
		}},
	}

	if name, _ := lang.GetStringAttribute(t.Attributes, lang.OriginalTypeAliasName); len(name) > 0 {
		if strings.Contains(name, "<") {
			nominalType, _ := lang.ParseNominalTypeFullName(name)
			s.Name = transformer.GenericTypeNamer{}.Name(nominalType)
		} else {
			s.Name = name
		}
	} else {
		val := t.GenericArguments[1]
		// TODO need precompile firstly
		if transformer.IsSingular(val.Name) {
			s.Name = transformer.Plural(strcase.ToCamel(val.Name))
		} else {
			s.Name = val.Name + "Map"
		}
	}

	return s, nil
}
