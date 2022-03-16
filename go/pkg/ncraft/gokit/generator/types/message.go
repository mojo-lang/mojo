package types

// Message represents a protobuf Message, though greatly simplified.
type Message struct {
    Name   string
    Fields []*Field
}

type Enum struct {
    Name string
}

type Map struct {
    // KeyType will always be a base type, e.g. string, int64, etc.
    KeyType   *FieldType
    ValueType *FieldType
}

// Field represents a field on a protobuf message.
type Field struct {
    Name     string
    FullName string
    Type     *FieldType

    Unwrap bool // raw data encode/decode in query parameters

    Enclosing *Field
}

// FieldType contains information about the type of one Field on a message,
// such as if that Field is a slice or if it's a pointer, as well as a
// reference to the definition of the type of this Field.
type FieldType struct {
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
    Map *Map
    // StarExpr is True if this FieldType represents a pointer to a type.
    StarExpr bool
    // ArrayType is True if this FieldType represents a slice of a type.
    ArrayType bool
}

func (f *Field) IsUnwrappedMap() bool {
    if f != nil && f.Type != nil && f.Type.Map != nil {
        return f.Unwrap
    }
    return false
}

func (f *Field) GetType() *FieldType {
    if f != nil {
        return f.Type
    }
    return nil
}
