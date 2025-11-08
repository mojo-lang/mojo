package data

import "github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"

// Message represents a protobuf Message, though greatly simplified.
type Message struct {
	Decl        *lang.StructDecl
	PackageName string // full package name
	Name        string
	Entity      bool
	IsNull      bool

	Fields []*Field

	Go         *GoMessage
	Java       *JavaMessage
	Extensions map[string]interface{}
}

type GoMessage struct {
	PackageName string
	ImportPath  string
}

type JavaMessage struct {
	Name         string // may be in wrap style, like Result<T>, Pagination<T>
	RawName      string // Result
	ParamName    string // <T>
	GRpcName     string
	IsPagination bool

	NeedConvert      bool
	GRpc2HttpConvert string
	Http2GrpcConvert string
}

func (m *Message) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Message) GetFirstFields() *Field {
	if m != nil && len(m.Fields) > 0 {
		return m.Fields[0]
	}
	return nil
}
