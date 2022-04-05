package compiler

import (
    "strings"

    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/go/data"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

type Enum struct {
    *data.Data
}

func (e *Enum) CompileEnum(ctx context.Context, decl *lang.EnumDecl) error {
    file := context.SourceFile(ctx)
    if file != nil && file.PackageName != decl.PackageName {
        return nil
    }

    enum := &data.Enum{}

    enum.PackageName = decl.PackageName
    enum.GoPackageName = GetGoPackage(decl.PackageName)
    enum.EnclosingName = strings.Join(lang.GetEnclosingNames(decl.EnclosingType), "_")
    enum.Name = decl.Name
    enum.FullName = enum.Name
    if len(enum.EnclosingName) > 0 {
        enum.WrapName = enum.EnclosingName
        enum.FullName = strings.Join([]string{enum.EnclosingName, enum.Name}, "_")
    } else {
        enum.WrapName = enum.Name
    }

    caseStyle, _ := lang.GetStringAttribute(decl.Attributes, core.CaseStyleAttributeName)
    //if len(caseStyle) == 0 {
    //	caseStyle = "kebab"
    //}

    for i, value := range decl.Type.Enumerators {
        if value.Name == "unspecified" {
            continue
        }

        item := data.EnumItem{
            RawValue: value.Name,
            Value:    value.Name,
        }

        if number, err := lang.GetIntegerAttribute(value.Attributes, core.NumberAttributeName); err == nil {
            item.Number = int(number)
        } else {
            item.Number = i
        }

        if alias, err := lang.GetStringAttribute(value.Attributes, "alias"); err == nil && len(alias) > 0 {
            item.Value = alias
        } else {
            item.Value = core.CaseStyler(caseStyle)(item.Value)
        }

        enum.Items = append(enum.Items, item)
    }

    e.Enums = append(e.Enums, enum)
    return nil
}
