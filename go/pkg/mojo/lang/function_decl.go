package lang

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

func (x *FunctionDecl) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *FunctionDecl) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *FunctionDecl) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil StructDecl")
}

func (x *FunctionDecl) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}

		if x.Document != nil {
			terms = append(terms, x.Document)
		}
		for _, attribute := range x.Attributes {
			terms = append(terms, attribute)
		}

		if x.Signature != nil {
			terms = append(terms, x.Signature)
		}

		if x.Body != nil {
			terms = append(terms, x.Body)
		}

		return terms
	}
	return nil
}

func (x *FunctionDecl) GetAttributeArguments(name string) ([]*Argument, error) {
	if x != nil {
		return GetAttributeArguments(x.Attributes, name)
	}
	return nil, errors.New("FunctionDecl is nil")
}

func (x *FunctionDecl) GetAttributeArgument(name string) (*Argument, error) {
	if x != nil {
		return GetAttributeArgument(x.Attributes, name)
	}
	return nil, errors.New("FunctionDecl is nil")
}

func (x *FunctionDecl) HasAttribute(name string) bool {
	if x != nil {
		return HasAttribute(x.Attributes, name)
	}
	return false
}

func (x *FunctionDecl) GetAttribute(name string) *Attribute {
	if x != nil {
		return GetAttribute(x.Attributes, name)
	}
	return nil
}

func (x *FunctionDecl) GetBoolAttribute(name string) (bool, error) {
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

func (x *FunctionDecl) SetBoolAttribute(name string, value bool) *Attribute {
	if x != nil {
		x.Attributes = SetBoolAttribute(x.Attributes, name, value)
		return x.GetAttribute(name)
	}
	return nil
}

func (x *FunctionDecl) SetImplicitBoolAttribute(name string, value bool) *Attribute {
	if x != nil {
		return x.SetBoolAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *FunctionDecl) GetIntegerAttribute(name string) (int64, error) {
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

func (x *FunctionDecl) SetIntegerAttribute(name string, value int64) *Attribute {
	if x != nil {
		x.Attributes = SetIntegerAttribute(x.Attributes, name, value)
		return x.GetAttribute(name)
	}
	return nil
}

func (x *FunctionDecl) SetImplicitIntegerAttribute(name string, value int64) *Attribute {
	if x != nil {
		return x.SetIntegerAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *FunctionDecl) GetStringAttribute(name string) (string, error) {
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

func (x *FunctionDecl) SetStringAttribute(name string, value string) *Attribute {
	if x != nil {
		x.Attributes = SetStringAttribute(x.Attributes, name, value)
		return x.GetAttribute(name)
	}
	return nil
}

func (x *FunctionDecl) SetImplicitStringAttribute(name string, value string) *Attribute {
	if x != nil {
		return x.SetStringAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (x *FunctionDecl) RemoveAttribute(name string) {
	if x != nil {
		x.Attributes = RemoveAttribute(x.Attributes, name)
	}
}

func (x *FunctionDecl) GetResource() string {
	if x != nil {
		// if x.GetStringAttribute()
	}
	return ""
}

// Copy shallow copy for the FunctionDecl
// deepFields will using deep copy
func (x *FunctionDecl) Copy(deepFields ...string) *FunctionDecl {
	decl := &FunctionDecl{
		StartPosition:     x.StartPosition,
		EndPosition:       x.EndPosition,
		Implicit:          x.Implicit,
		Document:          x.Document,
		PackageName:       x.PackageName,
		SourceFileName:    x.SourceFileName,
		KeywordPosition:   x.KeywordPosition,
		Name:              x.Name,
		FullName:          x.FullName,
		Attributes:        x.Attributes,
		GenericParameters: x.GenericParameters,
		Enclosing:         x.Enclosing,
		NamePosition:      x.NamePosition,
		Signature:         x.Signature,
		Body:              x.Body,
		Scope:             x.Scope,
	}

	index := make(map[string]bool)
	for _, field := range deepFields {
		index[field] = true
	}

	if index["attributes"] {
		var attributes []*Attribute
		for _, attribute := range x.Attributes {
			attributes = append(attributes, proto.Clone(attribute).(*Attribute))
		}
		decl.Attributes = attributes
	}

	if index["signature"] {
		decl.Signature = proto.Clone(x.Signature).(*FunctionSignature)
	}

	if index["body"] {
		decl.Body = proto.Clone(x.Body).(*BlockStmt)
	}

	return decl
}
