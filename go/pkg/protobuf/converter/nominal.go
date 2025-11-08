package converter

import (
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/context"
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
	scope := lang.GetScope(context.StructDecl(ctx))

	getName := func() string {
		if pkg != nil && pkg.FullName == t.PackageName {
			if decl := t.TypeDeclaration; decl != nil {
				if s := decl.GetStructDecl(); s != nil && s.Enclosing != nil {
					if scope == nil || scope.Identifiers == nil || scope.Identifiers[s.Name] == nil {
						return lang.GetFullName("", t.GetEnclosingNames(), t.Name)
					}
					return s.Name
				} else if e := decl.GetEnumDecl(); e != nil && e.Enclosing != nil {
					if scope == nil || scope.Identifiers == nil || scope.Identifiers[e.Name] == nil {
						return lang.GetFullName("", t.GetEnclosingNames(), t.Name)
					}
					return e.Name
				}
			}
			return t.Name
		} else {
			return t.GetFullName()
		}
	}

	if t.IsScalar() {
		return "Scalar", t.GetFullName(), nil
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
		tp, typeName, err := p.ConvertTo(ctx, t)
		if err == nil && len(typeName) > 0 {
			return tp, typeName, err
		}
	}

	return tp, getName(), nil
}
