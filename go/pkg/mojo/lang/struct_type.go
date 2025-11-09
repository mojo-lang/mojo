package lang

import (
	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

type FieldNamOption int32

const (
	FieldNamOptionDefault FieldNamOption = iota
	FieldNamOptionUseAlias
)

func (x *StructType) SetStartPosition(position *Position) {
	if x != nil {
		x.StartPosition = PatchPosition(x.StartPosition, position)
	}
}

func (x *StructType) SetEndPosition(position *Position) {
	if x != nil {
		x.EndPosition = PatchPosition(x.EndPosition, position)
	}
}

func (x *StructType) GetAllFieldNames(option FieldNamOption) []string {
	if x != nil {
		var fieldNames []string
		for _, inherit := range x.Inherits {
			names := inherit.GetTypeDeclaration().GetStructDecl().GetAllFieldNames(option)
			fieldNames = append(fieldNames, names...)
		}

		for _, field := range x.Fields {
			if option == FieldNamOptionUseAlias {
				if alias, _ := field.GetStringAttribute(core.AliasAttributeName); len(alias) > 0 {
					fieldNames = append(fieldNames, alias)
					continue
				}
			}

			fieldNames = append(fieldNames, field.Name)
		}
		return fieldNames
	}
	return nil
}

func (x *StructType) GetField(name string) *ValueDecl {
	if x != nil {
		for _, inherit := range x.Inherits {
			if field := inherit.GetTypeDeclaration().GetStructDecl().GetField(name); field != nil {
				return field
			}
		}

		for _, field := range x.Fields {
			if field.Name == name {
				return field
			}
		}
	}
	return nil
}

func (x *StructType) GetInheritFields() []*ValueDecl {
	if x != nil {
		var fields []*ValueDecl
		for _, inherit := range x.Inherits {
			fields = append(fields, inherit.GetTypeDeclaration().GetStructDecl().GetAllFields()...)
		}
		return fields
	}
	return nil
}

func (x *StructType) GetAllFields() []*ValueDecl {
	if x != nil {
		var fields []*ValueDecl

		for _, inherit := range x.Inherits {
			fields = append(fields, inherit.GetTypeDeclaration().GetStructDecl().GetAllFields()...)
		}

		fields = append(fields, x.Fields...)

		return fields
	}
	return nil
}

func (x *StructType) EachField(iter func(decl *ValueDecl) error) error {
	if x != nil {
		for _, inherit := range x.Inherits {
			if err := inherit.GetTypeDeclaration().GetStructDecl().EachField(iter); err != nil {
				return err
			}
		}

		for _, field := range x.Fields {
			if err := iter(field); err != nil {
				return err
			}
		}
	}
	return nil
}

func MergeFields(fields []*ValueDecl, merged []*ValueDecl) []*ValueDecl {
	index := make(map[string]*ValueDecl)
	for _, field := range fields {
		index[field.Name] = field
	}

	for _, field := range merged {
		if f, ok := index[field.Name]; !ok {
			fields = append(fields, field)
		} else if f.Type.GetFullName() != field.Type.GetFullName() {
			logs.Warnw("merge fields with same name but without same type, will skip it",
				"name", field.Name,
				"srcType", f.Type.GetFullName(),
				"mergedType", field.Type.GetFullName())
		}
	}
	return fields
}

func (x *StructType) MergeFields(merged []*ValueDecl) *StructType {
	if x != nil {
		x.Fields = MergeFields(x.Fields, merged)
	}
	return x
}
