package lang

import "strings"

func NewStructTypeDeclaration(decl *StructDecl) *TypeDeclaration {
	return &TypeDeclaration{
		TypeDeclaration: &TypeDeclaration_StructDecl{
			StructDecl: decl,
		},
	}
}

func NewTypeDeclarationFromDeclaration(decl *Declaration) *TypeDeclaration {
	switch decl.Declaration.(type) {
	case *Declaration_StructDecl:
		return &TypeDeclaration{
			TypeDeclaration: &TypeDeclaration_StructDecl{
				StructDecl: decl.GetStructDecl(),
			},
		}
	case *Declaration_InterfaceDecl:
		return &TypeDeclaration{
			TypeDeclaration: &TypeDeclaration_InterfaceDecl{
				InterfaceDecl: decl.GetInterfaceDecl(),
			},
		}
	case *Declaration_EnumDecl:
		return &TypeDeclaration{
			TypeDeclaration: &TypeDeclaration_EnumDecl{
				EnumDecl: decl.GetEnumDecl(),
			},
		}
	case *Declaration_TypeAliasDecl:
		return &TypeDeclaration{
			TypeDeclaration: &TypeDeclaration_TypeAliasDecl{
				TypeAliasDecl: decl.GetTypeAliasDecl(),
			},
		}
	case *Declaration_GenericParameter:
		return &TypeDeclaration{
			TypeDeclaration: &TypeDeclaration_GenericParameter{
				GenericParameter: decl.GetGenericParameter(),
			},
		}
	}

	return nil
}

func (x *TypeDeclaration) IsUnion() {
}

func (x *TypeDeclaration) GetSourceFileName() string {
	if x == nil {
		return ""
	}

	switch x.TypeDeclaration.(type) {
	case *TypeDeclaration_StructDecl:
		return x.GetStructDecl().SourceFileName
	case *TypeDeclaration_EnumDecl:
		return x.GetEnumDecl().SourceFileName
	case *TypeDeclaration_InterfaceDecl:
		return x.GetInterfaceDecl().SourceFileName
	case *TypeDeclaration_TypeAliasDecl:
		return x.GetTypeAliasDecl().SourceFileName
	default:
		return ""
	}
}

func (x *TypeDeclaration) Implicit() bool {
	if x != nil {
		switch x.TypeDeclaration.(type) {
		case *TypeDeclaration_StructDecl:
			return x.GetStructDecl().Implicit
		case *TypeDeclaration_EnumDecl:
			return x.GetEnumDecl().Implicit
		case *TypeDeclaration_InterfaceDecl:
			return x.GetInterfaceDecl().Implicit
		case *TypeDeclaration_TypeAliasDecl:
			return x.GetTypeAliasDecl().Implicit
		default:
			return false
		}
	}
	return false
}

func (x *TypeDeclaration) GetScope() *Scope {
	if x != nil {
		switch x.TypeDeclaration.(type) {
		case *TypeDeclaration_StructDecl:
			return x.GetStructDecl().GetScope()
		case *TypeDeclaration_EnumDecl:
			return x.GetEnumDecl().GetScope()
		case *TypeDeclaration_InterfaceDecl:
			return x.GetInterfaceDecl().GetScope()
		case *TypeDeclaration_TypeAliasDecl:
			return x.GetTypeAliasDecl().GetScope()
		default:
			return nil
		}
	}
	return nil
}

func (x *TypeDeclaration) GetEnclosingType() *NominalType {
	if x == nil {
		return nil
	}

	switch x.TypeDeclaration.(type) {
	case *TypeDeclaration_StructDecl:
		return x.GetStructDecl().Enclosing
	case *TypeDeclaration_InterfaceDecl:
		return x.GetInterfaceDecl().Enclosing
	case *TypeDeclaration_EnumDecl:
		return x.GetEnumDecl().Enclosing
	case *TypeDeclaration_TypeAliasDecl:
		return x.GetTypeAliasDecl().Enclosing
	case *TypeDeclaration_GenericParameter:
		return x.GetGenericParameter().Enclosing
	default:
		return nil
	}
}

func (x *TypeDeclaration) GetDecl() interface{} {
	if x != nil {
		switch x.TypeDeclaration.(type) {
		case *TypeDeclaration_StructDecl:
			return x.GetStructDecl()
		case *TypeDeclaration_InterfaceDecl:
			return x.GetInterfaceDecl()
		case *TypeDeclaration_EnumDecl:
			return x.GetEnumDecl()
		case *TypeDeclaration_TypeAliasDecl:
			return x.GetTypeAliasDecl()
		case *TypeDeclaration_GenericParameter:
			return x.GetGenericParameter()
		default:
			return nil
		}
	}
	return nil
}

func (x *TypeDeclaration) IsGeneric() bool {
	switch decl := x.TypeDeclaration.(type) {
	case *TypeDeclaration_TypeAliasDecl:
		return decl.TypeAliasDecl.IsGeneric()
	case *TypeDeclaration_StructDecl:
		return decl.StructDecl.IsGeneric()
	}
	return false
}

func (x *TypeDeclaration) Merge(dependencies ...[]*Identifier) {
	var merged [][]*Identifier
	for _, dep := range dependencies {
		merged = append(merged, dep)
	}

	switch x.TypeDeclaration.(type) {
	case *TypeDeclaration_StructDecl:
		decl := x.GetStructDecl()
		merged = append(merged, decl.ResolvedIdentifiers)
		decl.ResolvedIdentifiers = MergeDependencies(merged...)
	case *TypeDeclaration_EnumDecl:
		decl := x.GetEnumDecl()
		merged = append(merged, decl.ResolvedIdentifiers)
		decl.ResolvedIdentifiers = MergeDependencies(merged...)
	case *TypeDeclaration_InterfaceDecl:
		decl := x.GetInterfaceDecl()
		merged = append(merged, decl.ResolvedIdentifiers)
		decl.ResolvedIdentifiers = MergeDependencies(merged...)
	case *TypeDeclaration_TypeAliasDecl:
		decl := x.GetTypeAliasDecl()
		merged = append(merged, decl.ResolvedIdentifiers)
		decl.ResolvedIdentifiers = MergeDependencies(merged...)
	}
}

func IsGenericTypeName(typeName string) bool {
	return strings.ContainsAny(typeName, "<>")
}
