package data

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

// Field represents a field on a protobuf message.
type Field struct {
	Decl     *lang.ValueDecl
	Name     string
	FullName string // exploded message's field will include the parent field name

	Exploded      bool // raw data encode/decode in query parameters
	ExplodePrefix string

	Type *FieldType

	Enclosing *Field

	Go         *GoField
	Java       *JavaField
	Extensions map[string]interface{}
}

type GoField struct {
}

type JavaField struct {
	RequestBody       bool
	ParamBindingName  string
	ParamBindingValue string
}

func (f *Field) JavaRequestBody() bool {
	return f != nil && f.Java != nil && f.Java.RequestBody
}

func (f *Field) HasJavaParamBinding() bool {
	return f != nil && f.Java != nil && len(f.Java.ParamBindingName) > 0
}

func (f *Field) JavaParamBinding() string {
	if f != nil && f.Java != nil && f.HasJavaParamBinding() {
		return f.Java.ParamBindingName + "(\"" + f.Java.ParamBindingValue + "\")"
	}
	return ""
}

func (f *Field) GetType() *FieldType {
	if f != nil {
		return f.Type
	}
	return nil
}

func (f *Field) GetKeyType() *FieldType {
	if f != nil && f.Type != nil {
		return f.Type.KeyType
	}
	return nil
}

func (f *Field) GetElementType() *FieldType {
	if f != nil && f.Type != nil {
		return f.Type.ElementType
	}
	return nil
}

func (f *Field) GetGoType() *GoFieldType {
	return f.GetType().GetGoType()
}

func (f *Field) GetElementGoType() *GoFieldType {
	return f.GetType().GetElementGoType()
}

func (f *Field) GetGoTypeName() string {
	if g := f.GetType().GetGoType(); g != nil {
		return g.Name
	}
	return ""
}

func (f *Field) GetElementGoTypeName() string {
	if g := f.GetType().GetElementGoType(); g != nil {
		return g.Name
	}
	return ""
}

func (f *Field) GetEnclosingField() *Field {
	if f != nil {
		return f.Enclosing
	}
	return nil
}
