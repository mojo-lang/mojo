package data

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

// Message represents a protobuf Message, though greatly simplified.
type Message struct {
    Decl        *lang.StructDecl
    PackageName string // full package name
    Name        string

    Fields []*Field

    Entity     bool
    Extensions map[string]interface{}

    Go GoMessage
}

type GoMessage struct {
    PackageName string
    ImportPath  string
}

// Field represents a field on a protobuf message.
type Field struct {
    Name     string
    FullName string // exploded message's field will include the parent field name

    Exploded      bool // raw data encode/decode in query parameters
    ExplodePrefix string

    Type *FieldType

    Enclosing *Field

    Decl *lang.ValueDecl
}

func (f *Field) GetType() *FieldType {
    if f != nil {
        return f.Type
    }
    return nil
}

func (f *Field) GetGoType() *GoFieldType {
    return f.GetType().GetGo()
}

func (f *Field) GetElementGoType() *GoFieldType {
    return f.GetType().GetElementGo()
}

func (f *Field) GetGoTypeName() string {
    if g := f.GetType().GetGo(); g != nil {
        return g.Name
    }
    return ""
}

func (f *Field) GetElementGoTypeName() string {
    if g := f.GetType().GetElementGo(); g != nil {
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

// FieldType contains information about the type of one Field on a message,
// such as if that Field is a slice or if it's a pointer, as well as a
// reference to the definition of the type of this Field.
type FieldType struct {
    PackageName string

    // Name will contain the name of the type, for example "string" or "bool"
    Name string

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

    // Protobuf Enums need to be handled uniquely when parsing queryparameters
    IsEnum bool

    IsMap bool

    // Repeated is true if this arg corresponds to a protobuf field which is
    // given an identifier of "repeated", meaning it will represented in Go as
    // a slice of it's type.
    IsArray bool

    Go GoFieldType

    ElementGo GoFieldType
}

func (f *FieldType) GetGo() *GoFieldType {
    if f != nil {
        return &f.Go
    }
    return nil
}

func (f *FieldType) GetElementGo() *GoFieldType {
    if f != nil {
        return &f.ElementGo
    }
    return nil
}

type GoFieldType struct {
    Name string

    // Indicates whether this field represents a basic protobuf type such as
    // one of the ints, floats, strings, bools, etc. Since we can only create
    // automatic marshaling of base types, if this is false a warning is given
    // to the user.
    IsBaseType bool

    // IsPointer is True if this FieldType represents a pointer to a type.
    IsPointer bool
}
