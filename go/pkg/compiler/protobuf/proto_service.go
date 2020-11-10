package protobuf

// ProtoMethod is
type ProtoMethod struct {
	Document []string
	Options  map[string]interface{}

	Name string

	InputType  string
	OutputType string

	ClientStreaming bool
	ServerStreaming bool
}

// ProtoService is
type ProtoService struct {
	Document []string
	Options  map[string]interface{}

	Name string

	Methods []*ProtoMethod
}
