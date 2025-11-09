package openapi

import "strings"

const RequestBodyReferenceRoot = "/components/requestBodies/"

func NewReferenceableRequestBody(requestBody *RequestBody) *ReferenceableRequestBody {
	s := &ReferenceableRequestBody{}
	s.SetRequestBody(requestBody)
	return s
}

func NewReferencedRequestBody(reference *Reference) *ReferenceableRequestBody {
	s := &ReferenceableRequestBody{}
	s.SetReference(reference)
	return s
}

func (x *ReferenceableRequestBody) SetRequestBody(requestBody *RequestBody) {
	if x != nil {
		x.ReferenceableRequestBody = &ReferenceableRequestBody_RequestBody{
			RequestBody: requestBody,
		}
	}
}

func (x *ReferenceableRequestBody) SetReference(reference *Reference) {
	if x != nil {
		x.ReferenceableRequestBody = &ReferenceableRequestBody_Reference{
			Reference: reference,
		}
	}
}

func (x *ReferenceableRequestBody) GetRequestBodyOf(index map[string]*RequestBody) *RequestBody {
	if x != nil {
		reference := x.GetReference()
		if reference != nil {
			fragment := reference.GetRef().GetFragment()
			key := strings.TrimPrefix(fragment, RequestBodyReferenceRoot)
			if s, ok := index[key]; ok {
				return s
			}
			return nil
		}
		return x.GetRequestBody()
	}
	return nil
}
