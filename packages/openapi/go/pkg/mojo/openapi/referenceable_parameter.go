package openapi

import "strings"

const ParameterReferenceRoot = "/components/parameters/"

func NewReferenceableParameter(parameter *Parameter) *ReferenceableParameter {
	s := &ReferenceableParameter{}
	s.SetParameter(parameter)
	return s
}

func NewReferencedParameter(reference *Reference) *ReferenceableParameter {
	s := &ReferenceableParameter{}
	s.SetReference(reference)
	return s
}

func (x *ReferenceableParameter) SetParameter(parameter *Parameter) {
	if x != nil {
		x.ReferenceableParameter = &ReferenceableParameter_Parameter{
			Parameter: parameter,
		}
	}
}

func (x *ReferenceableParameter) SetReference(reference *Reference) {
	if x != nil {
		x.ReferenceableParameter = &ReferenceableParameter_Reference{
			Reference: reference,
		}
	}
}

func (x *ReferenceableParameter) GetParameterOf(index map[string]*Parameter) *Parameter {
	if x != nil {
		reference := x.GetReference()
		if reference != nil {
			fragment := reference.GetRef().GetFragment()
			key := strings.TrimPrefix(fragment, ParameterReferenceRoot)
			if s, ok := index[key]; ok {
				return s
			}
			return nil
		}
		return x.GetParameter()
	}
	return nil
}
