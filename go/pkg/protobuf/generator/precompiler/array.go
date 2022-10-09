package precompiler

import (
	"errors"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/compiler/transformer"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

func CompileArrayToStruct(ctx context.Context, t *lang.NominalType) (*lang.StructDecl, error) {
	if t.GetFullName() != core.ArrayTypeFullName {
		return nil, errors.New("")
	}

	if len(t.GenericArguments) != 1 {
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

	val := t.GenericArguments[0]

	if name, _ := lang.GetStringAttribute(t.Attributes, lang.OriginalTypeAliasName); len(name) > 0 {
		if strings.Contains(name, "<") {
			nominalType, _ := lang.ParseNominalTypeFullName(name)
			s.Name = transformer.GenericTypeNamer{}.Name(nominalType)
		} else {
			s.Name = name
		}
	} else {
		switch val.GetFullName() {
		case core.BoolTypeFullName:
			s.Name = core.BoolValuesTypeName
		case core.Int32TypeFullName:
			s.Name = core.Int32ValuesTypeName
		case core.Int64TypeFullName:
			s.Name = core.Int64ValuesTypeName
		case core.UInt32TypeFullName:
			s.Name = core.UInt32ValuesTypeName
		case core.UInt64TypeFullName:
			s.Name = core.UInt64ValuesTypeName
		case core.StringTypeFullName:
			s.Name = core.StringValuesTypeName
		default:
			s.Name = transformer.Plural(strcase.ToCamel(val.Name))
		}
	}

	return s, nil
}
