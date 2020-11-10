package protobuf

// ProtoFile struct for the protobuffer template
type ProtoFile struct {
	Name    string
	
	Package string

	Imports []string

	Options map[string]interface{}

	Services []*ProtoService

	Enums []*ProtoEnum

	Messages []*ProtoMessage `json:"messages"`
}
