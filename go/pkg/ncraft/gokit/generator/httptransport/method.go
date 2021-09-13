package httptransport

import "github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/types"

// Method contains the distillation of information within an
// svcdef.InterfaceMethod that's useful for templating http transport.
type Method struct {
	Name string

	// RequestType is the name of type of the Request, e.g. *EchoRequest
	RequestType  string
	ResponseType string

	Bindings []*Binding

	ResponseBodyField *Field
	ResponseHeaders   map[string]string
}

// NewMethod builds a Method struct from a svcdef.InterfaceMethod.
func NewMethod(meth *types.InterfaceMethod) *Method {
	nMeth := Method{
		Name:         meth.Name,
		RequestType:  meth.RequestType.Name,
		ResponseType: meth.ResponseType.Name,
	}
	//for i := range meth.HttpBindings {
	for i := range meth.Bindings {
		nBinding := NewBinding(i, meth)
		nBinding.Parent = &nMeth
		nMeth.Bindings = append(nMeth.Bindings, nBinding)
	}
	if meth.HTTPResponse != nil {
		nMeth.ResponseHeaders = meth.HTTPResponse.Headers
		nMeth.ResponseBodyField = &Field{
			Name: meth.HTTPResponse.BodyField.Name,
		}
	}

	return &nMeth
}
