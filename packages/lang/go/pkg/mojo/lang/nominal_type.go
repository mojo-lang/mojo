package lang

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

var (
	StringNominalType = &NominalType{
		PackageName: "mojo.core",
		Name:        "String",
	}
)

// ParseNominalTypeFullName
// support 'mojo.alias.Foo<mojo.alias.Bar>'
// support 'mojo.alias.Foo'
func ParseNominalTypeFullName(name string) (*NominalType, error) {
	nominalType := &NominalType{}
	return nominalType, nominalType.ParseFullName(name)
}

func NewArrayNominalType(typ *NominalType) *NominalType {
	return &NominalType{
		Implicit:         false,
		PackageName:      "mojo.core",
		Name:             "Array",
		GenericArguments: []*NominalType{typ},
	}
}

func NewUnionNominalType(types ...*NominalType) *NominalType {
	return &NominalType{
		Implicit:         false,
		PackageName:      "mojo.core",
		Name:             "Union",
		GenericArguments: types,
	}
}

func NewMapNominalType(key *NominalType, value *NominalType) *NominalType {
	return &NominalType{
		Implicit:         false,
		PackageName:      "mojo.core",
		Name:             "Map",
		GenericArguments: []*NominalType{key, value},
	}
}

func NewStringMapNominalType(value *NominalType) *NominalType {
	return &NominalType{
		Implicit:         false,
		PackageName:      "mojo.core",
		Name:             "Map",
		GenericArguments: []*NominalType{StringNominalType, value},
	}
}

func (x *NominalType) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *NominalType) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *NominalType) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil NominalType")
}

func (x *NominalType) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		// terms = append(terms, NewSymbolTerm(x.StartPosition, TermTypeName, x.Name))

		for _, attribute := range x.Attributes {
			terms = append(terms, attribute)
		}

		if x.Document != nil && x.Document.Following {
			terms = append(terms, x.Document)
		}

		return terms
	}
	return nil
}

func (x *NominalType) ToIdentifier() *Identifier {
	return x.NewIdentifier()
}

func (x *NominalType) NewIdentifier() *Identifier {
	if x != nil {
		identifier := &Identifier{
			StartPosition:  x.StartPosition,
			EndPosition:    x.EndPosition,
			Kind:           Identifier_KIND_UNSPECIFIED,
			PackageName:    x.PackageName,
			SourceFileName: x.TypeDeclaration.GetSourceFileName(),
			Enclosing:      GetEnclosingIdentifier(x.Enclosing),
			Name:           x.Name,
			FullName:       x.GetFullName(),
			Declaration:    NewDeclarationFromTypeDeclaration(x.TypeDeclaration),
			Implicit:       x.TypeDeclaration.Implicit(),
		}

		if decl := x.TypeDeclaration.GetTypeDeclaration(); decl != nil {
			switch decl.(type) {
			case *TypeDeclaration_EnumDecl:
				identifier.Kind = Identifier_KIND_ENUM
			case *TypeDeclaration_StructDecl:
				identifier.Kind = Identifier_KIND_STRUCT
			case *TypeDeclaration_TypeAliasDecl:
				identifier.Kind = Identifier_KIND_TYPE_ALIAS
			case *TypeDeclaration_InterfaceDecl:
				identifier.Kind = Identifier_KIND_INTERFACE
			}
		}

		return identifier
	}
	return nil
}

func (x *NominalType) GetFullName() string {
	if x != nil {
		if strings.Contains(x.Name, ".") {
			return GetFullName(x.PackageName, nil, x.Name)
		} else {
			return GetFullName(x.PackageName, x.GetEnclosingNames(), x.Name)
		}
	}
	return ""
}

func (x *NominalType) GetGenericFullName() string {
	if x != nil {
		if strings.Contains(x.Name, ".") {
			return GetFullName(x.PackageName, nil, x.GetGenericName())
		} else {
			return GetFullName(x.PackageName, x.GetEnclosingGenericNames(), x.GetGenericName())
		}
	}
	return ""
}

func (x *NominalType) GetGenericName() string {
	if x != nil {
		name := x.GetName()
		buf := bytes.NewBufferString(name)
		if len(x.GenericArguments) > 0 {
			buf.WriteByte('<')
			for i, argument := range x.GenericArguments {
				if i > 0 {
					buf.WriteByte(',')
				}
				buf.WriteString(argument.GetGenericFullName())
			}
			buf.WriteByte('>')
		}
		return buf.String()
	}
	return ""
}

func (x *NominalType) GetEnclosingNames() []string {
	if x == nil {
		return []string{}
	}
	return GetEnclosingNames(x.Enclosing)
}

func (x *NominalType) GetEnclosingGenericNames() []string {
	if x == nil {
		return []string{}
	}
	return GetEnclosingGenericNames(x.Enclosing)
}

func (x *NominalType) IsTypeAlias() bool {
	if x == nil {
		return false
	}
	return x.TypeDeclaration != nil && x.TypeDeclaration.GetTypeAliasDecl() != nil
}

func (x *NominalType) IsInstantiatedGeneric() bool {
	if x.IsGeneric() {
		for _, argument := range x.GetGenericArguments() {
			if argument.IsGenericParameter() {
				return false
			}
		}
		return true
	}
	return false
}

func (x *NominalType) IsGeneric() bool {
	return len(x.GetGenericArguments()) > 0
}

func (x *NominalType) IsGenericParameter() bool {
	if x == nil {
		return false
	}
	return x.TypeDeclaration != nil && x.TypeDeclaration.GetGenericParameter() != nil
}

func (x *NominalType) IsStruct() bool {
	if x == nil {
		return false
	}
	return x.TypeDeclaration != nil && x.TypeDeclaration.GetStructDecl() != nil
}

func (x *NominalType) IsInterface() bool {
	if x == nil {
		return false
	}
	return x.TypeDeclaration != nil && x.TypeDeclaration.GetInterfaceDecl() != nil
}

func (x *NominalType) IsEnum() bool {
	if x == nil {
		return false
	}
	return x.TypeDeclaration != nil && x.TypeDeclaration.GetEnumDecl() != nil
}

func (x *NominalType) IsScalar() bool {
	if x == nil {
		return false
	}
	if x.TypeDeclaration != nil {
		if decl := x.TypeDeclaration.GetStructDecl(); decl != nil {
			switch decl.GetFullName() {
			case core.BoolTypeFullName,
				core.Int8TypeFullName, core.Int16TypeFullName, core.Int32TypeFullName, core.Int64TypeFullName,
				core.UInt8TypeFullName, core.UInt16TypeFullName, core.UInt32TypeFullName, core.UInt64TypeFullName,
				core.PositiveTypeFullName, core.NegativeTypeFullName,
				core.Float32TypeFullName, core.Float64TypeFullName,
				core.StringTypeFullName, core.BytesTypeFullName:
				return true
			default:
				return false
			}
		} else if alias := x.TypeDeclaration.GetTypeAliasDecl(); alias != nil {
			switch alias.GetFullName() {
			case core.ByteTypeFullName, core.IntTypeFullName, core.UIntTypeFullName, core.SizeTypeFullName,
				core.FloatTypeFullName, core.DoubleTypeFullName:
				return true
			default:
				return false
			}
		}
	}
	return false
}

func (x *NominalType) IsArrayType() bool {
	return x.isType(core.ArrayTypeFullName)
}

func (x *NominalType) IsMapType() bool {
	return x.isType(core.MapTypeFullName)
}

func (x *NominalType) IsTupleType() bool {
	return x.isType(core.TupleTypeFullName)
}

func (x *NominalType) IsIntersectionType() bool {
	return x.isType(core.IntersectionTypeFullName)
}

func (x *NominalType) IsUnionType() bool {
	return x.isType(core.UnionTypeFullName)
}

func (x *NominalType) isType(typeName string) bool {
	if x == nil {
		return false
	}
	if x.TypeDeclaration != nil {
		if decl := x.TypeDeclaration.GetStructDecl(); decl != nil {
			return decl.GetFullName() == typeName
		}
	}
	return false
}

func (x *NominalType) ParseFullName(fullName string) error {
	_, err := x.parseFullGenericName(fullName)
	return err
}

func (x *NominalType) parseFullGenericName(fullName string) (string, error) {
	if x != nil {
		var err error
		originalFullName := fullName
		fullName = strings.TrimSpace(fullName)
		fullName, err = x.parseFullName(fullName)
		if err != nil {
			return "", err
		}

		fullName = strings.TrimSpace(fullName)
		if len(fullName) > 0 {
			if strings.HasPrefix(fullName, "<") {
				fullName = fullName[1:]
				fullName = strings.TrimSpace(fullName)
				for !strings.HasPrefix(fullName, ">") {
					argument := &NominalType{}
					fullName, err = argument.parseFullGenericName(fullName)
					if err != nil {
						return "", err
					}
					x.GenericArguments = append(x.GenericArguments, argument)
					fullName = strings.TrimSpace(fullName)

					if len(fullName) == 0 {
						return "", fmt.Errorf("malformed generic full name (%s)", originalFullName)
					}
				}
				fullName = fullName[1:]
				fullName = strings.TrimSpace(fullName)
				if strings.HasPrefix(fullName, ",") {
					fullName = fullName[1:]
					fullName = strings.TrimSpace(fullName)
				}
				return fullName, nil
			} else if strings.HasPrefix(fullName, ",") {
				fullName = strings.TrimPrefix(fullName, ",")
				fullName = strings.TrimSpace(fullName)
				return fullName, nil
			} else if strings.HasPrefix(fullName, ">") {
				return fullName, nil
			} else {
				return "", fmt.Errorf("malformed generic full name (%s)", originalFullName)
			}
		}
		return "", nil
	}
	return "", errors.New("nil NominalType")
}

func (x *NominalType) parseFullName(fullName string) (string, error) {
	if x != nil {
		if pkgReg, err := regexp.Compile(`^([a-z][a-z_0-9]*\.)+`); err != nil {
			return "", err
		} else {
			pkg := pkgReg.FindString(fullName)
			if len(pkg) > 0 {
				x.PackageName = strings.TrimSuffix(pkg, ".")
				fullName = fullName[len(pkg):]
			}
		}

		if enclosingReg, err := regexp.Compile(`^([A-Z][a-zA-Z0-9_]*\.)+`); err != nil {
			return "", err
		} else {
			enclosing := enclosingReg.FindString(fullName)
			fullName = fullName[len(enclosing):]
			enclosing = strings.TrimSuffix(enclosing, ".")
			if len(enclosing) > 0 {
				segments := strings.Split(enclosing, ".")
				var enclosingType *NominalType
				for _, segment := range segments {
					if enclosingType == nil {
						enclosingType = &NominalType{
							PackageName: x.PackageName,
							Name:        segment,
						}
					} else {
						enclosingType = &NominalType{
							PackageName: x.PackageName,
							Name:        segment,
							Enclosing:   enclosingType,
						}
					}
				}
				x.Enclosing = enclosingType
			}
		}

		if nameReg, err := regexp.Compile(`^[A-Z][a-zA-Z0-9_]*`); err != nil {
			return "", err
		} else {
			name := nameReg.FindString(fullName)
			if len(name) > 0 {
				x.Name = name
				fullName = fullName[len(name):]
			}
		}
	}

	return fullName, nil
}

func (x *NominalType) GetAttributeArguments(name string) ([]*Argument, error) {
	if x != nil {
		return GetAttributeArguments(x.Attributes, name)
	}
	return nil, errors.New("NominalType is nil")
}

func (x *NominalType) GetAttributeArgument(name string) (*Argument, error) {
	if x != nil {
		return GetAttributeArgument(x.Attributes, name)
	}
	return nil, errors.New("NominalType is nil")
}

func (x *NominalType) HasAttribute(name string) bool {
	if x != nil {
		return HasAttribute(x.Attributes, name)
	}
	return false
}

func (x *NominalType) GetAttribute(name string) *Attribute {
	if x != nil {
		return GetAttribute(x.Attributes, name)
	}
	return nil
}

func (x *NominalType) GetBoolAttribute(name string) (bool, error) {
	argument, err := x.GetAttributeArgument(name)
	if err != nil {
		return false, err
	}

	if argument != nil {
		return argument.GetBool()
	} else {
		// TODO using the default value of the attribute declaration
		return true, nil
	}
}

func (x *NominalType) SetBoolAttribute(name string, value bool) *Attribute {
	if x != nil {
		x.Attributes = SetBoolAttribute(x.Attributes, name, value)
		return x.Attributes[len(x.Attributes)-1]
	}
	return nil
}

func (x *NominalType) SetImplicitBoolAttribute(name string, value bool) *Attribute {
	if x != nil {
		return x.SetBoolAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *NominalType) GetIntegerAttribute(name string) (int64, error) {
	if argument, err := x.GetAttributeArgument(name); err != nil {
		return 0, err
	} else {
		if argument != nil {
			return argument.GetInteger()
		} else {
			// TODO using the default value of the attribute declaration
			return 0, nil
		}
	}
}

func (x *NominalType) SetIntegerAttribute(name string, value int64) *Attribute {
	if x != nil {
		x.Attributes = SetIntegerAttribute(x.Attributes, name, value)
		return x.Attributes[len(x.Attributes)-1]
	}
	return nil
}

func (x *NominalType) SetImplicitIntegerAttribute(name string, value int64) *Attribute {
	if x != nil {
		return x.SetIntegerAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *NominalType) GetStringAttribute(name string) (string, error) {
	if argument, err := x.GetAttributeArgument(name); err != nil {
		return "", err
	} else {
		if argument != nil {
			return argument.GetString()
		} else {
			return "", nil
		}
	}
}

func (x *NominalType) SetStringAttribute(name string, value string) *Attribute {
	if x != nil {
		x.Attributes = SetStringAttribute(x.Attributes, name, value)
		return x.Attributes[len(x.Attributes)-1]
	}
	return nil
}

func (x *NominalType) SetImplicitStringAttribute(name string, value string) *Attribute {
	if x != nil {
		return x.SetStringAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *NominalType) RemoveAttribute(name string) {
	if x != nil {
		x.Attributes = RemoveAttribute(x.Attributes, name)
	}
}
