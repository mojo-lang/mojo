package data

// HTTPBinding represents one of potentially several bindings from a gRPC
// service method to a particular HTTP path/verb.
type HTTPBinding struct {
	// Label is the name of this method, plus the english word for the index of
	// this binding in this methods slice of bindings. So if this binding where
	// the first binding in the slice of bindings for the method "Sum", the
	// label for this binding would be "SumZero". If it where the third
	// binding, it would be named "SumTwo".
	Label string

	Verb string
	Path string

	// BasePath is the longest static portion of the full PathTemplate, and is
	// given to the net/http mux as the path for the route for this binding.
	BasePath string

	// There is one HTTPParameter for each of the fields on parent service
	// methods Request.
	Parameters HttpParameters

	// reference to the parameter which located in body
	Body *HTTPParameter

	Response *HTTPResponse

	// A pointer back to the parent method of this binding. Used within some
	// binding methods
	Parent *Method

	Extensions map[string]interface{}
}

func (b *HTTPBinding) GetResponseBody() *Field {
	if b != nil && b.Response != nil {
		return b.Response.Body
	}
	return nil
}
