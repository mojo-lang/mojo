package lang

import (
	"sort"
	"strings"
)

// NewIdentifier
// sample: NewIdentifier("mojo", "core", "Int32")
func NewIdentifier(names ...string) *Identifier {
	if len(names) == 0 {
		return nil
	}

	if len(names) == 1 {
		names = strings.Split(names[0], ".")
	}

	identifier := &Identifier{
		Name: names[len(names)-1],
	}
	cur := identifier
	for i := len(names) - 2; i >= 0; i-- {
		cur.Enclosing = &Identifier{Name: names[i]}
		cur = cur.Enclosing
	}
	return identifier
}

func ParseIdentifierName(fullName string) (string, string) {
	index := strings.LastIndex(fullName, ".")
	if index > 0 {
		return fullName[0:index], fullName[index+1:]
	} else {
		return "", fullName
	}
}

func FindIdentifier(identifiers []*Identifier, fullName string) *Identifier {
	for _, id := range identifiers {
		if id.FullName == fullName {
			return id
		}
	}
	return nil
}

func (x *Identifier) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *Identifier) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *Identifier) IsGenericInstantiated() bool {
	return strings.Contains(x.GetName(), "<")
}

func (x *Identifier) ToNominalType() *NominalType {
	if x != nil {
		switch x.Kind {
		case Identifier_KIND_ENUM, Identifier_KIND_STRUCT, Identifier_KIND_TYPE_ALIAS, Identifier_KIND_INTERFACE:
			return &NominalType{
				StartPosition:    x.StartPosition,
				EndPosition:      x.EndPosition,
				PackageName:      x.PackageName,
				Name:             x.Name,
				TypeDeclaration:  NewTypeDeclarationFromDeclaration(x.Declaration),
				GenericArguments: nil,
				Attributes:       nil,
				Enclosing:        nil,
			}
		}
	}
	return nil
}

func (x *Identifier) FullNames() []string {
	if x != nil {
		var names []string
		if len(x.PackageName) > 0 {
			names = strings.Split(x.PackageName, ".")
		}

		enclosingNames := x.EnclosingNames()
		if len(enclosingNames) > 0 {
			names = append(names, enclosingNames...)
		}

		names = append(names, x.Name)
		return names
	}
	return nil
}

func (x *Identifier) EnclosingNames() []string {
	if x != nil && x.Enclosing != nil {
		enclosing := x.Enclosing
		var names []string
		for enclosing != nil {
			names = append(names, enclosing.Name)
			enclosing = enclosing.Enclosing
		}
		if len(names) > 0 {
			sort.Sort(sort.Reverse(sort.StringSlice(names)))
		}
		return names
	}
	return nil
}
