package data

import "github.com/mojo-lang/lang/go/pkg/mojo/lang"

// Message represents a protobuf Message, though greatly simplified.
type Message struct {
	Decl        *lang.StructDecl
	PackageName string // full package name
	Name        string
	Entity      bool

	Fields []*Field

	Go         *GoMessage
	Extensions map[string]interface{}
}

type GoMessage struct {
	PackageName string
	ImportPath  string
}

func (m *Message) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}
