package compiler

import (
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
)

var _ = Array{}
var _ = Map{}
var _ = Tuple{}
var _ = Union{}

type Nominal struct {
}

// Compile
// type alias transform
// type: Scalar, Enum, Struct
func (n Nominal) Compile(ctx context.Context, t *lang.NominalType) (string, string, error) { // type, typeName, error
    pkg := context.Package(ctx)
    getName := func() string {
        if pkg != nil && pkg.FullName == t.PackageName {
            return t.Name
        } else {
            return t.GetFullName()
        }
    }

    if t.IsScalar() {
        return "Scalar", t.Name, nil
    }

    tp := "Struct"

    if decl := t.TypeDeclaration; decl != nil {
        if aliasDecl := decl.GetTypeAliasDecl(); aliasDecl != nil {
            if aliasDecl.Type != nil {
                //	// type_alias_original_name
                //	_, err := lang.GetStringAttribute(aliasDecl.Type.Attributes, lang.OriginalTypeAliasName)
                //	if err != nil { // not found
                //		aliasDecl.Type.Attributes = lang.SetStringAttribute(aliasDecl.Type.Attributes, lang.OriginalTypeAliasName, t.Name)
                //	}
                //
                //	return CompileNominalType(ctx, aliasDecl.Type)
            }
        }

        if decl.GetEnumDecl() != nil {
            tp = "Enum"
        }
    }

    ps := NominalPlugins()[t.GetFullName()]
    for _, p := range ps {
        tp, typeName, err := p.Compile(ctx, t)
        if err == nil && len(typeName) > 0 {
            return tp, typeName, err
        }
    }

    return tp, getName(), nil
}
