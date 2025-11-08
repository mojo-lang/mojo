package lang

import (
	"errors"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

type AttributesGetter interface {
	GetAttributes() []*Attribute
}

func GetAttributes(decl interface{}) Attributes {
	if _, ok := decl.(core.IsUnion); ok {
		decl = core.GetUnionPrimeType(decl)
	}

	if getter, ok := decl.(AttributesGetter); ok {
		return getter.GetAttributes()
	}
	return nil
}

type Attributes []*Attribute

func (a Attributes) GetNumberAttribute() *Attribute {
	return a.GetAttribute(core.NumberAttributeName)
}

func (a Attributes) GetRequiredAttribute() *Attribute {
	return a.GetAttribute(core.RequiredAttributeName)
}

func (a Attributes) GetOptionalAttribute() *Attribute {
	return a.GetAttribute(core.OptionalAttributeName)
}

func (a Attributes) GetAttribute(name string) *Attribute {
	return GetAttribute(a, name)
}

func (a Attributes) HasAttribute(name string) bool {
	return HasAttribute(a, name)
}

func (a Attributes) GetStartPosition() *Position {
	if len(a) > 0 {
		return a[0].GetStartPosition()
	}
	return nil
}

func (a Attributes) GetEndPosition() *Position {
	if len(a) > 0 {
		return a[len(a)-1].GetEndPosition()
	}
	return nil
}

func GetAttribute(attributes []*Attribute, name string) *Attribute {
	for _, attribute := range attributes {
		if attribute.IsSameName(name) {
			return attribute
		}
	}
	return nil
}

// GetAttributeArguments make sure returned Argument list which length greater than 0
func GetAttributeArguments(attributes []*Attribute, name string) ([]*Argument, error) {
	for _, attribute := range attributes {
		if attribute.IsSameName(name) {
			if len(attribute.Arguments) > 0 {
				return attribute.Arguments, nil
			} else {
				break
			}
		}
	}
	return nil, errors.New("NotFound")
}

func GetAttributeArgument(attributes []*Attribute, name string) (*Argument, error) {
	for _, attribute := range attributes {
		if attribute.IsSameName(name) {
			if len(attribute.Arguments) > 0 {
				return attribute.Arguments[0], nil
			} else {
				return nil, nil
			}
		}
	}
	return nil, errors.New("NotFound")
}

func HasAttribute(attributes []*Attribute, name string) bool {
	for _, attribute := range attributes {
		if attribute.IsSameName(name) {
			return true
		}
	}
	return false
}

func GetBoolAttribute(attributes []*Attribute, name string) (bool, error) {
	for _, attribute := range attributes {
		if attribute.IsSameName(name) {
			if len(attribute.Arguments) > 0 {
				if value := attribute.Arguments[0].Value.GetBoolLiteralExpr(); value != nil {
					return value.Value, nil
				}
			} else { // default value is true
				return true, nil
			}
		}
	}
	return false, errors.New("NotFound")
}

func SetBoolAttribute(attributes []*Attribute, name string, value bool) []*Attribute {
	for _, attribute := range attributes {
		if attribute.IsSameName(name) {
			if len(attribute.Arguments) > 0 {
				attribute.Arguments[0].Value = NewBoolLiteralExpression(&BoolLiteralExpr{
					Kind:     0,
					Implicit: false,
					Value:    value,
				})
				return attributes
			}
		}
	}

	pkg, n := ParseIdentifierName(name)
	attributes = append(attributes, &Attribute{
		PackageName: pkg,
		Name:        n,
		Arguments: []*Argument{
			{
				Value: NewBoolLiteralExpression(&BoolLiteralExpr{
					Kind:     0,
					Implicit: false,
					Value:    value,
				}),
			},
		},
	})

	return attributes
}

func GetIntegerAttribute(attributes []*Attribute, name string) (int64, error) {
	argument, err := GetAttributeArgument(attributes, name)
	if err != nil {
		return 0, err
	}
	if argument != nil {
		return argument.GetInteger()
	} else {
		// TODO using the default value of the attribute declaration
		return 0, nil
	}
}

func SetIntegerAttribute(attributes []*Attribute, name string, value int64) []*Attribute {
	for _, attribute := range attributes {
		if attribute.IsSameName(name) {
			if len(attribute.Arguments) > 0 {
				attribute.Arguments[0] = NewIntegerLiteralArgument(NewIntegerLiteralExpr(value))
				return attributes
			}
		}
	}

	pkg, n := ParseIdentifierName(name)
	attributes = append(attributes, &Attribute{
		PackageName: pkg,
		Name:        n,
		Arguments:   []*Argument{NewIntegerLiteralArgument(NewIntegerLiteralExpr(value))},
	})

	return attributes
}

func GetStringAttribute(attributes []*Attribute, name string) (string, error) {
	argument, err := GetAttributeArgument(attributes, name)
	if err != nil {
		return "", err
	}
	if argument != nil {
		return argument.GetString()
	} else {
		return "", nil
	}
}

func GetStringValuesAttribute(attributes []*Attribute, name string) ([]string, error) {
	var values []string
	arguments, err := GetAttributeArguments(attributes, name)
	if err != nil {
		return nil, err
	}

	if len(arguments) == 1 {
		if strArray, err := arguments[0].GetStringArray(); err == nil {
			return strArray, nil
		}
	}

	for _, argument := range arguments {
		if val, err := argument.GetString(); err != nil {
			return nil, err
		} else {
			values = append(values, val)
		}
	}
	return values, nil
}

func SetStringAttribute(attributes []*Attribute, name string, value string) []*Attribute {
	for _, attribute := range attributes {
		if attribute.IsSameName(name) {
			if len(attribute.Arguments) > 0 {
				attribute.Arguments[0] = NewStringLiteralArgument(&StringLiteralExpr{
					Kind:     0,
					Implicit: false,
					Value:    value,
				})
				return attributes
			}
		}
	}

	pkg, n := ParseIdentifierName(name)
	attributes = append(attributes, &Attribute{
		PackageName: pkg,
		Name:        n,
		Arguments: []*Argument{NewStringLiteralArgument(&StringLiteralExpr{
			Kind:     0,
			Implicit: false,
			Value:    value,
		})},
	})

	return attributes
}

func RemoveAttribute(attributes []*Attribute, name string) []*Attribute {
	var newAttributes []*Attribute
	for _, attribute := range attributes {
		if attribute.IsSameName(name) {
			continue
		}
		newAttributes = append(newAttributes, attribute)
	}

	return newAttributes
}

func MergeAttributes(dst []*Attribute, src []*Attribute) []*Attribute {
	index := make(map[string][]*Attribute)
	setAttributes := func(attributes []*Attribute) {
		for _, attr := range attributes {
			fullName := attr.GetFullName()
			if attr.Repeatable() || len(index[fullName]) == 0 {
				index[fullName] = append(index[fullName], attr)
			}
		}
	}
	setAttributes(dst)
	setAttributes(src)

	var attributes []*Attribute
	for _, attrs := range index {
		attributes = append(attributes, attrs...)
	}
	return attributes
}
