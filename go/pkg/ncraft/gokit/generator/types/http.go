package types

// HTTPBinding represents one of potentially several bindings from a gRPC
// service method to a particuar HTTP path/verb.
type HTTPBinding struct {
    Verb string
    Path string
    // There is one HTTPParamter for each of the fields on parent service
    // methods RequestType.
    Params []*HTTPParameter
}

// HTTPParameter represents the location of one field for a given HTTPBinding.
// Each HTTPParameter corresponds to one Field of the parent
// InterfaceMethod.RequestType.Fields
type HTTPParameter struct {
    // Field points to a Field on the Parent service methods "RequestType".
    Field *Field
    // Location will be either "body", "path", or "query"
    Location string
}

type HTTPResponse struct {
    BodyField *Field
    Headers   map[string]string
}
