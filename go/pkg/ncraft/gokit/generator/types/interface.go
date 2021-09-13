package types

type Interface struct {
	Name    string
	Methods []*InterfaceMethod
}

type InterfaceMethod struct {
	Name         string
	CamelName 	 string
	RequestType  *Message
	ResponseType *Message

	// Bindings contains information for mapping http paths and parameters onto
	// the fields of this InterfaceMethods RequestType.
	Bindings []*HTTPBinding

	HTTPResponse *HTTPResponse
}