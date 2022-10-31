package data

// FieldType contains information about the type of one Field on a message,
// such as if that Field is a slice or if it's a pointer, as well as a
// reference to the definition of the type of this Field.
type FieldType struct {
	// Name will contain the name of the type, for example "String" or "Bool"
	// will be the element type name if it is array or map
	Name string

	// PackageName the full package name for the field type
	PackageName string

	// Enum contains a pointer to the Enum type this fieldtype represents, if
	// this FieldType represents an Enum. If not, Enum is nil.
	Enum *Enum

	// Message contains a pointer to the Message type this FieldType
	// represents, if this FieldType represents a Message. If not, Message is
	// nil.
	Message *Message

	// Map contains a pointer to the Map type this FieldType represents, if
	// this FieldType represents a Map. If not, Map is nil.
	// KeyType will always be a base type, e.g. string, int64, etc.
	KeyType *FieldType

	ElementType *FieldType

	IsEnum bool

	//
	IsMap bool

	// IsArray is true if this type corresponds to a protobuf field which is
	// given an identifier of "repeated", meaning it will be represented in Go as
	// a slice of the type.
	IsArray bool

	// Indicates whether this field represents a basic protobuf type such as
	// one of the integers, floats, strings, booleans, etc.
	IsScalar bool

	Go         *GoFieldType
	Extensions map[string]interface{}
}

type GoFieldType struct {
	Name string

	// IsPointer is True if this FieldType represents a pointer to a type.
	IsPointer bool
}

func (f *FieldType) GetName() string {
	if f != nil {
		return f.Name
	}
	return ""
}

func (f *FieldType) IsString() bool {
	if f != nil {
		return f.Name == "String"
	}
	return false
}

func (f *FieldType) GetPackageName() string {
	if f != nil {
		return f.PackageName
	}
	return ""
}

func (f *FieldType) GetKeyType() *FieldType {
	if f != nil {
		return f.KeyType
	}
	return nil
}

func (f *FieldType) GetElementType() *FieldType {
	if f != nil {
		return f.ElementType
	}
	return nil
}

func (f *FieldType) GetGoType() *GoFieldType {
	if f != nil {
		return f.Go
	}
	return nil
}

func (f *FieldType) GetKeyGoType() *GoFieldType {
	return f.GetKeyType().GetGoType()
}

func (f *FieldType) GetElementGoType() *GoFieldType {
	return f.GetElementType().GetGoType()
}
