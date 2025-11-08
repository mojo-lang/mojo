package lang

import (
	"errors"
	"strings"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

func NewBoolAttribute(pkg string, name string) *Attribute {
	return newBoolAttribute(pkg, name, false)
}

func NewIntegerAttribute(pkg string, name string, value int64) *Attribute {
	return newIntegerAttribute(pkg, name, value, false)
}

func NewFloatAttribute(pkg string, name string, value float64) *Attribute {
	return newFloatAttribute(pkg, name, value, false)
}

func NewStringAttribute(pkg string, name string, value string) *Attribute {
	return newStringAttribute(pkg, name, value, false)
}

func NewImplicitBoolAttribute(pkg string, name string) *Attribute {
	return newBoolAttribute(pkg, name, true)
}

func NewImplicitIntegerAttribute(pkg string, name string, value int64) *Attribute {
	return newIntegerAttribute(pkg, name, value, true)
}

func NewImplicitFloatAttribute(pkg string, name string, value float64) *Attribute {
	return newFloatAttribute(pkg, name, value, true)
}

func NewImplicitStringAttribute(pkg string, name string, value string) *Attribute {
	return newStringAttribute(pkg, name, value, false)
}

// IsNumber return true if the attribute is number attribute
func (x *Attribute) IsNumber() bool {
	return x != nil && x.Name == core.NumberAttributeName
}

func (x *Attribute) IsRequired() bool {
	return x != nil && x.Name == core.RequiredAttributeName
}

func (x *Attribute) IsOptional() bool {
	return x != nil && x.Name == core.OptionalAttributeName
}

func (x *Attribute) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *Attribute) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *Attribute) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		return MergeCommentToTerms(comment, x.GetTerms())
	}

	return false, errors.New("nil ArrayLiteralExpr")
}

func (x *Attribute) GetTerms() []interface{} {
	if x != nil {
		var terms []interface{}
		terms = append(terms, &Term{StartPosition: x.NamePosition,
			EndPosition: &Position{
				Line:   x.NamePosition.Line,
				Column: x.NamePosition.Column + int64(len(x.Name)),
			},
			Type:  "Name",
			Value: x.Name,
		})

		for _, argument := range x.GenericArguments {
			terms = append(terms, argument)
		}
		for _, argument := range x.Arguments {
			terms = append(terms, argument)
		}
		return terms
	}
	return nil
}

func (x *Attribute) IsSameName(fullName string) bool {
	if x == nil {
		return false
	}

	removeMojoPackage := func(pkg string) string {
		switch pkg {
		case "mojo", "mojo.core":
			return ""
		default:
			return strings.TrimPrefix(pkg, "mojo.")
		}
	}

	pkg, name := ParseIdentifierName(fullName)
	if len(pkg) > 0 {
		if x.PackageName == pkg || removeMojoPackage(x.PackageName) == removeMojoPackage(pkg) {
			return x.Name == name
		}
		return false
	} else {
		return x.Name == name
	}
}

func (x *Attribute) GetFullName() string {
	fullName := ""
	if x != nil {
		if len(x.PackageName) > 0 {
			fullName = x.PackageName + "."
		}
		fullName = fullName + x.Name
	}
	return fullName
}

func (x *Attribute) Repeatable() bool {
	if x != nil && x.Declaration != nil {
		if v, err := GetBoolAttribute(x.Declaration.Attributes, "repeatable"); err == nil {
			return v
		}
	}
	return false
}

func (x *Attribute) GetArrayLiteralArgument() *ArrayLiteralExpr {
	array := &ArrayLiteralExpr{
		Implicit: true,
	}

	if len(x.Arguments) == 1 {
		if arrayLiteral := x.Arguments[0].Value.GetArrayLiteralExpr(); arrayLiteral != nil {
			return arrayLiteral
		}
	}

	for _, argument := range x.Arguments {
		if len(argument.Label) > 0 {
			return nil
		}

		array.Elements = append(array.Elements, argument.Value)
	}
	return array
}

func (x *Attribute) GetObjectLiteralArgument() *ObjectLiteralExpr {
	if x != nil {
		object := &ObjectLiteralExpr{
			Implicit: true,
		}
		if len(x.Arguments) == 1 {
			if objectLiteral := x.Arguments[0].Value.GetObjectLiteralExpr(); objectLiteral != nil {
				return objectLiteral
			}
		}

		object, _ = Arguments(x.Arguments).ToObjectLiteralExpr()
		return object
	}
	return nil
}

func (x *Attribute) GetMapLiteralArgument() *MapLiteralExpr {
	if len(x.Arguments) == 1 {
		if mapLiteral := x.Arguments[0].Value.GetMapLiteralExpr(); mapLiteral != nil {
			return mapLiteral
		}
	}

	object := x.GetObjectLiteralArgument()
	if object != nil {
		mapLiteral := &MapLiteralExpr{
			Implicit: true,
		}
		for _, field := range object.Fields {
			mapLiteral.Entries = append(mapLiteral.Entries, &MapLiteralExpr_Entry{
				Key:   field.Name,
				Value: field.Value,
			})
		}
		return mapLiteral
	}
	return nil
}

func (x *Attribute) GetObjectArgument(object interface{}) error {
	o := x.GetObjectLiteralArgument()
	if o != nil {
		return o.To(object)
	}
	return errors.New("not an object literal argument")
}

func (x *Attribute) SetImplicit(value bool) *Attribute {
	if x != nil {
		x.Implicit = value
	}
	return x
}

// GetBool eval the argument first, then use the default value of the attribute declaration if no argument
func (x *Attribute) GetBool() bool {
	if x != nil {
		if len(x.Arguments) == 0 {
			if x.Declaration != nil && x.Declaration.DefaultValue != nil {
				val, _ := x.Declaration.DefaultValue.EvalBoolLiteral()
				return val
			}
			return true
		}
		val, _ := x.Arguments[0].GetValue().EvalBoolLiteral()
		return val
	}
	return false
}

func (x *Attribute) GetInteger() (int64, error) {
	if x != nil {
		if len(x.Arguments) > 0 {
			return x.Arguments[0].GetInteger()
		} else {
			return x.GetValue().EvalIntegerLiteral()
		}
	}
	return 0, errors.New("")
}

func (x *Attribute) GetString() (string, error) {
	if x != nil {
		if len(x.Arguments) > 0 {
			return x.Arguments[0].GetString()
		} else {
			return x.GetValue().EvalStringLiteral()
		}
	}
	return "", errors.New("")
}

func (x *Attribute) SetString(value string) *Attribute {
	if x != nil {
		if len(x.Arguments) > 0 {
			x.Arguments[0] = NewStringArgument(value)
		} else {
			x.Arguments = append(x.Arguments, NewStringArgument(value))
		}

		x.Value = x.Arguments[0].GetValue()
	}
	return x
}

func newBoolAttribute(pkg string, name string, implicit bool) *Attribute {
	return &Attribute{
		PackageName: pkg,
		Name:        name,
		Arguments: []*Argument{{
			Value: NewBoolLiteralExpression(&BoolLiteralExpr{
				Kind:     0,
				Implicit: implicit,
				Value:    true,
			}),
		},
		},
	}
}

func newIntegerAttribute(pkg string, name string, value int64, implicit bool) *Attribute {
	v := uint64(value)
	if value < 0 {
		v = uint64(-value)
	}

	return &Attribute{
		PackageName: pkg,
		Name:        name,
		Arguments: []*Argument{{
			Value: NewIntegerLiteralExpression(&IntegerLiteralExpr{
				Kind:       0,
				Implicit:   implicit,
				IsNegative: value < 0,
				Value:      v,
			}),
		},
		},
	}
}

func newFloatAttribute(pkg string, name string, value float64, implicit bool) *Attribute {
	return &Attribute{
		PackageName: pkg,
		Name:        name,
		Arguments: []*Argument{{
			Value: NewFloatLiteralExpression(&FloatLiteralExpr{
				Kind:       0,
				Implicit:   implicit,
				IsNegative: value < 0,
				Value:      value,
			}),
		},
		},
	}
}

func newStringAttribute(pkg string, name string, value string, implicit bool) *Attribute {
	return &Attribute{
		PackageName: pkg,
		Name:        name,
		Arguments: []*Argument{{
			Value: NewStringLiteralExpression(&StringLiteralExpr{
				Kind:     0,
				Implicit: implicit,
				Value:    value,
			}),
		},
		},
	}
}
