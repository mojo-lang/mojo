package lang

import "errors"

const OriginalTypeAliasName = "original_type_alias_name"

func (x *TypeAliasDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *TypeAliasDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *TypeAliasDecl) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil NominalType")
}

func (x *TypeAliasDecl) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		for _, attribute := range x.Attributes {
			terms = append(terms, attribute)
		}

		terms = append(terms,
			NewSymbolTerm(x.KeywordPosition, TermTypeKeyword, KeywordType),
			NewSymbolTerm(x.NamePosition, TermTypeName, x.Name))

		for _, parameter := range x.GenericParameters {
			terms = append(terms, parameter)
		}

		terms = append(terms, x.Type)

		return terms
	}
	return nil
}

func (x *TypeAliasDecl) GetFullName() string {
	if x != nil {
		return GetFullName(x.PackageName, GetEnclosingNames(x.Enclosing), x.Name)
	}
	return ""
}

func (x *TypeAliasDecl) SetScope(scope *Scope) {
	if x != nil {
		x.Scope = scope
	}
}

func (x *TypeAliasDecl) IsGeneric() bool {
	return x != nil && len(x.GenericParameters) > 0
}

func (x *TypeAliasDecl) EnclosingTypeDecl() interface{} {
	return x.GetEnclosing().GetTypeDeclaration().GetDecl()
}

func (x *TypeAliasDecl) GetAttributeArguments(name string) ([]*Argument, error) {
	if x != nil {
		return GetAttributeArguments(x.Attributes, name)
	}
	return nil, errors.New("StructDecl is nil")
}

func (x *TypeAliasDecl) GetAttributeArgument(name string) (*Argument, error) {
	if x != nil {
		return GetAttributeArgument(x.Attributes, name)
	}
	return nil, errors.New("StructDecl is nil")
}

func (x *TypeAliasDecl) HasAttribute(name string) bool {
	if x != nil {
		return HasAttribute(x.Attributes, name)
	}
	return false
}

func (x *TypeAliasDecl) GetAttribute(name string) *Attribute {
	if x != nil {
		return GetAttribute(x.Attributes, name)
	}
	return nil
}

func (x *TypeAliasDecl) GetBoolAttribute(name string) (bool, error) {
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

func (x *TypeAliasDecl) SetBoolAttribute(name string, value bool) *Attribute {
	if x != nil {
		x.Attributes = SetBoolAttribute(x.Attributes, name, value)
		return x.Attributes[len(x.Attributes)-1]
	}
	return nil
}

func (x *TypeAliasDecl) SetImplicitBoolAttribute(name string, value bool) *Attribute {
	if x != nil {
		return x.SetBoolAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *TypeAliasDecl) GetIntegerAttribute(name string) (int64, error) {
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

func (x *TypeAliasDecl) SetIntegerAttribute(name string, value int64) *Attribute {
	if x != nil {
		x.Attributes = SetIntegerAttribute(x.Attributes, name, value)
		return x.Attributes[len(x.Attributes)-1]
	}
	return nil
}

func (x *TypeAliasDecl) SetImplicitIntegerAttribute(name string, value int64) *Attribute {
	if x != nil {
		return x.SetIntegerAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *TypeAliasDecl) GetStringAttribute(name string) (string, error) {
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

func (x *TypeAliasDecl) SetStringAttribute(name string, value string) *Attribute {
	if x != nil {
		x.Attributes = SetStringAttribute(x.Attributes, name, value)
		return x.Attributes[len(x.Attributes)-1]
	}
	return nil
}

func (x *TypeAliasDecl) SetImplicitStringAttribute(name string, value string) *Attribute {
	if x != nil {
		return x.SetStringAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *TypeAliasDecl) RemoveAttribute(name string) {
	if x != nil {
		x.Attributes = RemoveAttribute(x.Attributes, name)
	}
}
