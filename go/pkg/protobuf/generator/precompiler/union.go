package precompiler

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/mojo/compiler/transformer"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"strings"
)

func CompileUnionToStruct(ctx context.Context, t *lang.NominalType) (*lang.StructDecl, error) {
	originName, err := lang.GetStringAttribute(t.Attributes, lang.OriginalTypeAliasName)
	if err != nil { // directly union declaration
		return nil, err
	} else {
		if strings.Contains(originName, "<") {
			nominalType, _ := lang.ParseNominalTypeFullName(originName)
			originName = transformer.GenericTypeNamer{}.Name(nominalType)
		}

		labelFormat, _ := lang.GetStringAttribute(t.Attributes, "label_format")

		// generate the message
		s := &lang.StructDecl{}
		alias, err := lang.GetStringAttribute(t.Attributes, "alias")
		if err == nil {
			s.Name = alias
		} else {
			s.Name = originName
		}
		s.Type = &lang.StructType{}
		for _, argument := range t.GenericArguments {
			label, err := lang.GetStringAttribute(argument.Attributes, "label")
			if err != nil {
				format, _ := lang.GetStringAttribute(argument.Attributes, "label_format")
				if len(format) > 0 {
					labelFormat = format
				}

				// auto generate
				label = generateLabel(argument, labelFormat)
			}

			s.Type.Fields = append(s.Type.Fields, &lang.ValueDecl{
				Name: label,
				Type: argument,
			})
		}

		return s, nil
	}
}

func generateLabel(argument *lang.NominalType, labelFormat string) string {
	if labelFormat == "{}" {
		return strcase.ToSnake(argument.Name)
	} else {
		fullName := argument.GetFullName()
		if argument.IsScalar() {
			return strcase.ToSnake(argument.Name) + "_val"
		} else if fullName == core.ArrayTypeFullName {
			name := argument.GenericArguments[0].Name
			name = strcase.ToSnake(name + "_val")
			return transformer.Plural(name)
		} else if fullName == core.MapTypeFullName {
			name := argument.GenericArguments[1].Name
			return strcase.ToSnake(name + "_map_vals")
		}
		return strcase.ToSnake(argument.Name) + "_val"
	}
}
