package lang

import (
	"errors"
)

func (x *ValueDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *ValueDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *ValueDecl) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil StructDecl")
}

func (x *ValueDecl) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		if x.Document != nil && !x.Document.Following {
			terms = append(terms, x.Document)
		}

		for _, attribute := range x.Attributes {
			terms = append(terms, attribute)
		}

		if x.KeywordPosition != nil {
			terms = append(terms,
				NewSymbolTerm(x.KeywordPosition, TermTypeKeyword, KeywordType))
		}

		if x.NamePosition != nil {
			terms = append(terms,
				NewSymbolTerm(x.NamePosition, TermTypeName, x.Name))
		}

		if x.Type != nil {
			terms = append(terms, x.Type)
		}

		if x.Initializer != nil {
			terms = append(terms, x.Initializer)
		}

		return terms
	}
	return nil
}

func (x *ValueDecl) GetInitialValue() *Expression {
	if x != nil {
		return x.GetInitializer().GetValue()
	}
	return nil
}

func (x *ValueDecl) GetAttributeArguments(name string) ([]*Argument, error) {
	if x != nil {
		arguments, err := GetAttributeArguments(x.Attributes, name)
		if err != nil && x.Type != nil {
			return GetAttributeArguments(x.Type.Attributes, name)
		}
		return arguments, err
	}

	return nil, errors.New("ValueDecl is nil")
}

func (x *ValueDecl) GetAttributeArgument(name string) (*Argument, error) {
	if x != nil {
		argument, err := GetAttributeArgument(x.Attributes, name)
		if err != nil && x.Type != nil {
			return GetAttributeArgument(x.Type.Attributes, name)
		}
		return argument, err
	}

	return nil, errors.New("ValueDecl is nil")
}

func (x *ValueDecl) GetAttribute(name string) *Attribute {
	if x != nil {
		attr := GetAttribute(x.Attributes, name)
		if attr == nil && x.Type != nil {
			attr = GetAttribute(x.Type.Attributes, name)
		}
		return attr
	}
	return nil
}

func (x *ValueDecl) HasAttribute(name string) bool {
	if x != nil {
		attr := HasAttribute(x.Attributes, name)
		if !attr && x.Type != nil {
			attr = HasAttribute(x.Type.Attributes, name)
		}
		return attr
	}

	return false
}

func (x *ValueDecl) GetBoolAttribute(name string) (bool, error) {
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

func (x *ValueDecl) SetBoolAttribute(name string, value bool) *Attribute {
	if x != nil && x.Type != nil {
		x.Type.Attributes = SetBoolAttribute(x.Type.Attributes, name, value)
		return x.Type.Attributes[len(x.Type.Attributes)-1]
	}
	return nil
}

func (x *ValueDecl) SetImplicitBoolAttribute(name string, value bool) *Attribute {
	if x != nil && x.Type != nil {
		return x.SetBoolAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *ValueDecl) GetIntegerAttribute(name string) (int64, error) {
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

func (x *ValueDecl) SetIntegerAttribute(name string, value int64) *Attribute {
	if x != nil && x.Type != nil {
		x.Type.Attributes = SetIntegerAttribute(x.Type.Attributes, name, value)
		return x.Type.Attributes[len(x.Type.Attributes)-1]
	}
	return nil
}

func (x *ValueDecl) SetImplicitIntegerAttribute(name string, value int64) *Attribute {
	if x != nil && x.Type != nil {
		return x.SetIntegerAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *ValueDecl) GetStringAttribute(name string) (string, error) {
	argument, err := x.GetAttributeArgument(name)
	if err != nil {
		return "", err
	}

	if argument != nil {
		return argument.GetString()
	} else {
		return "", nil
	}
}

func (x *ValueDecl) SetStringAttribute(name string, value string) *Attribute {
	if x != nil && x.Type != nil {
		x.Type.Attributes = SetStringAttribute(x.Type.Attributes, name, value)
		return x.Type.Attributes[len(x.Type.Attributes)-1]
	}
	return nil
}

func (x *ValueDecl) SetImplicitStringAttribute(name string, value string) *Attribute {
	if x != nil && x.Type != nil {
		return x.SetStringAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *ValueDecl) RemoveAttribute(name string) {
	if x != nil {
		x.Attributes = RemoveAttribute(x.Attributes, name)
		if x.Type != nil {
			x.Type.Attributes = RemoveAttribute(x.Type.Attributes, name)
		}
	}
}

func (x *ValueDecl) IsArrayType() bool {
	return x.GetType().IsArrayType()
}

func (x *ValueDecl) IsMapType() bool {
	return x.GetType().IsMapType()
}

func (x *ValueDecl) IsUnionType() bool {
	return x.GetType().IsUnionType()
}

func (x *ValueDecl) IsScalarLike() bool {
	return x.GetType().IsScalar()
}

func (x *ValueDecl) IsEnum() bool {
	return x.GetType().IsEnum()
}
