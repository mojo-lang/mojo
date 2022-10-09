package syntax

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

func CombineValueDecl(decl *lang.ValueDecl, types ...*lang.NominalType) *lang.ValueDecl {
	var name string
	var typeDecl *lang.TypeDeclaration
	var attributes []*lang.Attribute
	for _, typ := range types {
		if len(name) == 0 && len(typ.Name) > 0 {
			name = typ.Name
		}
		if typeDecl == nil && typ.TypeDeclaration != nil {
			typeDecl = typ.TypeDeclaration
		}
		if len(typ.Attributes) > 0 {
			attributes = append(attributes, typ.Attributes...)
		}
	}
	decl.Type = &lang.NominalType{
		Name:            name,
		TypeDeclaration: typeDecl,
		Attributes:      attributes,
	}

	return decl
}
