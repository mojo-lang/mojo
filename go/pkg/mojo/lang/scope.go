package lang

import (
	"strings"
)

func NewScope() *Scope {
	return &Scope{
		Identifiers: make(map[string]*Identifier),
	}
}

func (x *Scope) Declare(declaration *Declaration) *Identifier {
	if declaration == nil {
		return nil
	}

	switch declaration.Declaration.(type) {
	case *Declaration_StructDecl:
		decl := declaration.GetStructDecl()

		packageName := decl.PackageName
		enclosingNames := GetEnclosingNames(decl.Enclosing)
		fullName := GetFullName(packageName, enclosingNames, decl.Name)

		identifier := &Identifier{
			PackageName: packageName,
			Enclosing:   decl.Enclosing.ToIdentifier(),
			Kind:        Identifier_KIND_STRUCT,
			Name:        decl.Name,
			FullName:    fullName,
			Declaration: declaration,
		}
		x.Identifiers[identifier.Name] = identifier
		return identifier
	case *Declaration_EnumDecl:
		decl := declaration.GetEnumDecl()

		packageName := decl.PackageName
		enclosingNames := GetEnclosingNames(decl.Enclosing)
		fullName := GetFullName(packageName, enclosingNames, decl.Name)

		identifier := &Identifier{
			PackageName: packageName,
			Enclosing:   decl.Enclosing.ToIdentifier(),
			Kind:        Identifier_KIND_ENUM,
			Name:        decl.Name,
			FullName:    fullName,
			Declaration: declaration,
		}
		x.Identifiers[identifier.Name] = identifier
		return identifier
	case *Declaration_InterfaceDecl:
		decl := declaration.GetInterfaceDecl()

		packageName := decl.PackageName
		fullName := strings.Join([]string{packageName, decl.Name}, ".")

		identifier := &Identifier{
			PackageName: packageName,
			Enclosing:   nil,
			Kind:        Identifier_KIND_INTERFACE,
			Name:        decl.Name,
			FullName:    fullName,
			Declaration: declaration,
		}
		x.Identifiers[identifier.Name] = identifier
		return identifier
	case *Declaration_TypeAliasDecl:
		decl := declaration.GetTypeAliasDecl()

		packageName := decl.PackageName
		fullName := strings.Join([]string{packageName, decl.Name}, ".")

		identifier := &Identifier{
			PackageName: packageName,
			Enclosing:   nil,
			Kind:        Identifier_KIND_TYPE_ALIAS,
			Name:        decl.Name,
			FullName:    fullName,
			Declaration: declaration,
		}
		x.Identifiers[identifier.Name] = identifier
		return identifier
	case *Declaration_ImportDecl:
		// err := c.compileImport(c.Context, decl.GetImportDecl())
		// if err != nil {
		//	return err
		// }
		return nil
	case *Declaration_GenericParameter:
		parameter := declaration.GetGenericParameter()

		identifier := &Identifier{
			PackageName: parameter.PackageName,
			Kind:        Identifier_KIND_GENERIC_PARAMETER,
			Name:        parameter.Name,
			Declaration: declaration,
		}
		x.Identifiers[identifier.Name] = identifier
		return identifier
	}
	return nil
}

func (x *Scope) GetIdentifier(name string) *Identifier {
	if x != nil && x.Identifiers != nil {
		return x.Identifiers[GetTypeTypeName(name)]
	}
	return nil
}
