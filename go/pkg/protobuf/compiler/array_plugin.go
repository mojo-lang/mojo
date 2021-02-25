package compiler

import (
	"errors"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/descriptor"
)

type ArrayPlugin struct {
	pluralize *pluralize.Client
}

func init() {
	p := plugins["mojo.core.Array"]
	if p == nil {
		p = make([]Plugin, 0)
	}
	p = append(p, &ArrayPlugin{
		pluralize: pluralize.NewClient(),
	})

	plugins["mojo.core.Array"] = p
}

func (p *ArrayPlugin) Compile(ctx *Context, t *lang.NominalType) (string, string, error) {
	if t.Name != "Array" {
		return "", "", errors.New("")
	}

	if len(t.GenericArguments) != 1 {
		return "", "", errors.New("")
	}

	t.Attributes = lang.SetIntegerAttribute(t.Attributes, "number", 1)

	s := &lang.StructDecl{}
	s.Type = &lang.StructType{
		Fields: []*lang.ValueDecl{{
			Name: "values",
			Type: t,
		}},
	}

	val := t.GenericArguments[0]
	s.Name = p.pluralize.Plural(strcase.ToCamel(val.Name))

	err := CompileStruct(ctx, s, descriptor.NewMessageDescriptor(ctx.GetFileDescriptor()))
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("failed to compile the arry field in %s.%s: %s",
			"", //ctx.Message.Name,
			"", //ctx.FieldName,
			err.Error()))
	}

	//ctx.File.Message.AddInnerMessage(context.Message)
	return "struct", s.Name, nil
}
