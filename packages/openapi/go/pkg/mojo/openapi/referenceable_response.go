package openapi

import "strings"

const ResponseReferenceRoot = "/components/responses/"

func NewReferenceableResponse(response *Response) *ReferenceableResponse {
	s := &ReferenceableResponse{}
	s.SetResponse(response)
	return s
}

func NewReferencedResponse(reference *Reference) *ReferenceableResponse {
	s := &ReferenceableResponse{}
	s.SetReference(reference)
	return s
}

func (x *ReferenceableResponse) SetResponse(response *Response) {
	if x != nil {
		x.ReferenceableResponse = &ReferenceableResponse_Response{
			Response: response,
		}
	}
}

func (x *ReferenceableResponse) SetReference(reference *Reference) {
	if x != nil {
		x.ReferenceableResponse = &ReferenceableResponse_Reference{
			Reference: reference,
		}
	}
}

func (x *ReferenceableResponse) GetResponseOf(index map[string]*Response) *Response {
	if x != nil {
		reference := x.GetReference()
		if reference != nil {
			fragment := reference.GetRef().GetFragment()
			key := strings.TrimPrefix(fragment, ResponseReferenceRoot)
			if s, ok := index[key]; ok {
				return s
			}
			return nil
		}
		return x.GetResponse()
	}
	return nil
}
