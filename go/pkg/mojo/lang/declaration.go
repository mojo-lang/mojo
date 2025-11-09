package lang

import (
	"errors"
	"reflect"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

const (
	PackageDeclName        = "PackageDecl"
	ImportDeclName         = "ImportDecl"
	EnumDeclName           = "EnumDecl"
	StructDeclName         = "StructDecl"
	InterfaceDeclName      = "InterfaceDecl"
	TypeAliasDeclName      = "TypeAliasDecl"
	ConstantDeclName       = "ConstantDecl"
	VariableDeclName       = "VariableDecl"
	AttributeDeclName      = "AttributeDecl"
	AttributeAliasDeclName = "AttributeAliasDecl"
	FunctionDeclName       = "FunctionDecl"
	ConstructorDeclName    = "ConstructorDecl"
	GenericParameterName   = "GenericParameter"
)

var declInfos map[reflect.Type]StructJsonInfo

func init() {
	declInfos = GetStructJsonInfos(jsoniter.Config{},
		PackageDecl{},
		ImportDecl{},
		EnumDecl{},
		StructDecl{},
		InterfaceDecl{},
		TypeAliasDecl{},
		ConstantDecl{},
		VariableDecl{},
		AttributeDecl{},
		AttributeAliasDecl{},
		FunctionDecl{},
		ConstructorDecl{},
		GenericParameter{},
	)
}

func NewEnumDeclaration(decl *EnumDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_EnumDecl{
			EnumDecl: decl,
		},
	}
}

func NewStructDeclaration(decl *StructDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_StructDecl{
			StructDecl: decl,
		},
	}
}

func NewInterfaceDeclaration(decl *InterfaceDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_InterfaceDecl{
			InterfaceDecl: decl,
		},
	}
}

func NewTypeAliasDeclaration(decl *TypeAliasDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_TypeAliasDecl{
			TypeAliasDecl: decl,
		},
	}
}

func NewFunctionDeclaration(decl *FunctionDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_FunctionDecl{
			FunctionDecl: decl,
		},
	}
}

func NewVariableDeclaration(decl *VariableDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_VariableDecl{
			VariableDecl: decl,
		},
	}
}

func NewVariableDeclarationFromValueDecl(decl *ValueDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_VariableDecl{
			VariableDecl: NewVariableDeclFromValueDecl(decl),
		},
	}
}

func NewConstantDeclaration(decl *ConstantDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_ConstantDecl{
			ConstantDecl: decl,
		},
	}
}

func NewAttributeDeclaration(decl *AttributeDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_AttributeDecl{
			AttributeDecl: decl,
		},
	}
}

func NewAttributeAliasDeclaration(decl *AttributeAliasDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_AttributeAliasDecl{
			AttributeAliasDecl: decl,
		},
	}
}

func NewPackageDeclaration(decl *PackageDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_PackageDecl{
			PackageDecl: decl,
		},
	}
}

func NewImportDeclaration(decl *ImportDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_ImportDecl{
			ImportDecl: decl,
		},
	}
}

func NewConstructorDeclaration(decl *ConstructorDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_ConstructorDecl{
			ConstructorDecl: decl,
		},
	}
}

func NewGenericParameterDeclaration(parameter *GenericParameter) *Declaration {
	return &Declaration{
		Declaration: &Declaration_GenericParameter{
			GenericParameter: parameter,
		},
	}
}

func NewDeclarationFromTypeDeclaration(decl *TypeDeclaration) *Declaration {
	if decl != nil {
		switch decl.TypeDeclaration.(type) {
		case *TypeDeclaration_StructDecl:
			return &Declaration{
				Declaration: &Declaration_StructDecl{
					StructDecl: decl.GetStructDecl(),
				},
			}
		case *TypeDeclaration_InterfaceDecl:
			return &Declaration{
				Declaration: &Declaration_InterfaceDecl{
					InterfaceDecl: decl.GetInterfaceDecl(),
				},
			}
		case *TypeDeclaration_EnumDecl:
			return &Declaration{
				Declaration: &Declaration_EnumDecl{
					EnumDecl: decl.GetEnumDecl(),
				},
			}
		case *TypeDeclaration_TypeAliasDecl:
			return &Declaration{
				Declaration: &Declaration_TypeAliasDecl{
					TypeAliasDecl: decl.GetTypeAliasDecl(),
				},
			}
		case *TypeDeclaration_GenericParameter:
			return &Declaration{
				Declaration: &Declaration_GenericParameter{
					GenericParameter: decl.GetGenericParameter(),
				},
			}
		default:
			return nil
		}
	}
	return nil
}

func (x *Declaration) IsUnion() {
}

func (x *Declaration) GetStartPosition() *Position {
	return GetUnionPosition(x, StartPositionFieldName)
}

func (x *Declaration) GetEndPosition() *Position {
	return GetUnionPosition(x, EndPositionFieldName)
}

func (x *Declaration) SetStartPosition(position *Position) {
	SetUnionPosition(x, StartPositionFieldName, position)
}

func (x *Declaration) SetEndPosition(position *Position) {
	SetUnionPosition(x, EndPositionFieldName, position)
}

func (x *Declaration) GetName() string {
	if x != nil {
		return core.GetUnionField(x, "Name").String()
	}
	return ""
}

func (x *Declaration) GetEnclosingType() *NominalType {
	if x != nil {
		switch x.Declaration.(type) {
		case *Declaration_StructDecl:
			return x.GetStructDecl().Enclosing
		case *Declaration_InterfaceDecl:
			return x.GetInterfaceDecl().Enclosing
		case *Declaration_EnumDecl:
			return x.GetEnumDecl().Enclosing
		case *Declaration_TypeAliasDecl:
			return x.GetTypeAliasDecl().Enclosing
		case *Declaration_GenericParameter:
			return x.GetGenericParameter().Enclosing
		default:
			return nil
		}
	}
	return nil
}

func (x *Declaration) GetPackageName() string {
	if x != nil {
		return core.GetUnionField(x, "PackageName").String()
	}
	return ""
}

func (x *Declaration) SetPackageName(name string) *Declaration {
	if x != nil {
		core.SetUnionField(x, "PackageName", reflect.ValueOf(name))
	}
	return x
}

func (x *Declaration) IsGeneric() bool {
	switch decl := x.Declaration.(type) {
	case *Declaration_TypeAliasDecl:
		return decl.TypeAliasDecl.IsGeneric()
	case *Declaration_StructDecl:
		return decl.StructDecl.IsGeneric()
	}
	return false
}

func (x *Declaration) IsInstantiatedGeneric() bool {
	return strings.ContainsAny(x.GetStructDecl().GetName(), "<>")
}

func (x *Declaration) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		value := reflect.ValueOf(x.Declaration)
		value = reflect.Indirect(value).Field(0)
		if merger, ok := value.Interface().(CommentMerger); ok {
			return merger.MergeComment(comment)
		} else {
			// error
		}
	}

	return false, errors.New("nil Declaration")
}
