package lang

import (
	"errors"
)

func (x *StructDecl) GetFullName() string {
	if x != nil {
		return GetFullName(x.PackageName, GetEnclosingNames(x.Enclosing), x.Name)
	}
	return ""
}

func (x *StructDecl) SetScope(scope *Scope) {
	if x != nil {
		x.Scope = scope
	}
}

func (x *StructDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *StructDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *StructDecl) IsBoxed() bool {
	if x != nil && x.Type != nil && len(x.Type.Inherits) == 1 && len(x.Type.Fields) == 0 {
		inherit := x.Type.Inherits[0]
		structDecl := inherit.TypeDeclaration.GetStructDecl()
		return structDecl != nil && structDecl.Type == nil
	}
	return false
}

func (x *StructDecl) IsOpacity() bool {
	return x != nil && x.Type == nil
}

func (x *StructDecl) IsGeneric() bool {
	return x != nil && len(x.GenericParameters) > 0
}

func (x *StructDecl) IsBodyEmpty() bool {
	if x != nil {
		return len(x.TypeAliasDecls) == 0 && len(x.EnumDecls) == 0 && len(x.StructDecls) == 0 && x.GetFieldCount() == 0
	}
	return true
}

func (x *StructDecl) GetFieldCount() int {
	return len(x.GetType().GetFields())
}

func (x *StructDecl) GetAllFieldNames(option FieldNamOption) []string {
	if x != nil {
		return x.Type.GetAllFieldNames(option)
	}
	return nil
}

func (x *StructDecl) GetInheritFields() []*ValueDecl {
	if x != nil {
		return x.Type.GetInheritFields()
	}
	return nil
}

func (x *StructDecl) GetAllFields() []*ValueDecl {
	if x != nil {
		return x.Type.GetAllFields()
	}
	return nil
}

func (x *StructDecl) GetField(name string) *ValueDecl {
	if x != nil {
		return x.Type.GetField(name)
	}
	return nil
}

func (x *StructDecl) EachField(iter func(decl *ValueDecl) error) error {
	if x != nil {
		return x.Type.EachField(iter)
	}
	return nil
}

func (x *StructDecl) EnclosingTypeDecl() interface{} {
	return x.GetEnclosing().GetTypeDeclaration().GetDecl()
}

func (x *StructDecl) GetAttributeArguments(name string) ([]*Argument, error) {
	if x != nil {
		return GetAttributeArguments(x.Attributes, name)
	}
	return nil, errors.New("StructDecl is nil")
}

func (x *StructDecl) GetAttributeArgument(name string) (*Argument, error) {
	if x != nil {
		return GetAttributeArgument(x.Attributes, name)
	}
	return nil, errors.New("StructDecl is nil")
}

func (x *StructDecl) HasAttribute(name string) bool {
	if x != nil {
		return HasAttribute(x.Attributes, name)
	}
	return false
}

func (x *StructDecl) GetAttribute(name string) *Attribute {
	if x != nil {
		return GetAttribute(x.Attributes, name)
	}
	return nil
}

func (x *StructDecl) GetBoolAttribute(name string) (bool, error) {
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

func (x *StructDecl) SetBoolAttribute(name string, value bool) *Attribute {
	if x != nil {
		x.Attributes = SetBoolAttribute(x.Attributes, name, value)
		return x.Attributes[len(x.Attributes)-1]
	}
	return nil
}

func (x *StructDecl) SetImplicitBoolAttribute(name string, value bool) *Attribute {
	if x != nil {
		return x.SetBoolAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *StructDecl) GetIntegerAttribute(name string) (int64, error) {
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

func (x *StructDecl) SetIntegerAttribute(name string, value int64) *Attribute {
	if x != nil {
		x.Attributes = SetIntegerAttribute(x.Attributes, name, value)
		return x.Attributes[len(x.Attributes)-1]
	}
	return nil
}

func (x *StructDecl) SetImplicitIntegerAttribute(name string, value int64) *Attribute {
	if x != nil {
		return x.SetIntegerAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *StructDecl) GetStringAttribute(name string) (string, error) {
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

func (x *StructDecl) SetStringAttribute(name string, value string) *Attribute {
	if x != nil {
		x.Attributes = SetStringAttribute(x.Attributes, name, value)
		return x.Attributes[len(x.Attributes)-1]
	}
	return nil
}

func (x *StructDecl) SetImplicitStringAttribute(name string, value string) *Attribute {
	if x != nil {
		return x.SetStringAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *StructDecl) RemoveAttribute(name string) {
	if x != nil {
		x.Attributes = RemoveAttribute(x.Attributes, name)
	}
}

func (x *StructDecl) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms()...)
	}

	return false, errors.New("nil StructDecl")
}

func (x *StructDecl) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		if x.Document != nil {
			terms = append(terms, x.Document)
		}

		for _, attribute := range x.Attributes {
			terms = append(terms, attribute)
		}

		terms = append(terms,
			NewSymbolTerm(x.KeywordPosition, TermTypeKeyword, KeywordType),
			NewSymbolTerm(x.NamePosition, TermTypeName, x.Name))

		for _, parameter := range x.GenericParameters {
			terms = append(terms, parameter)
		}

		if x.Type != nil {
			if len(x.Type.Inherits) > 0 {
				terms = append(terms, NewSymbolTerm(x.Type.InheritPosition, TermTypeSymbol, ":"))

				for _, inherit := range x.Type.Inherits {
					terms = append(terms, inherit)
				}
			}

			terms = append(terms, NewSymbolTerm(x.Type.StartPosition, TermTypeStart, ""))

			for _, alias := range x.TypeAliasDecls {
				terms = append(terms, alias)
			}

			for _, enum := range x.EnumDecls {
				terms = append(terms, enum)
			}

			for _, structDecl := range x.StructDecls {
				terms = append(terms, structDecl)
			}

			for _, field := range x.Type.Fields {
				terms = append(terms, field)
			}

			terms = append(terms, NewSymbolTerm(x.Type.EndPosition, TermTypeEnd, ""))
		}

		return terms
	}
	return nil
}
