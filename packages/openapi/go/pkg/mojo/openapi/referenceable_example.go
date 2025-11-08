package openapi

import "strings"

const ExampleReferenceRoot = "/components/examples/"

func NewReferenceableExample(example *Example) *ReferenceableExample {
	s := &ReferenceableExample{}
	s.SetExample(example)
	return s
}

func NewReferencedExample(reference *Reference) *ReferenceableExample {
	s := &ReferenceableExample{}
	s.SetReference(reference)
	return s
}

func (x *ReferenceableExample) SetExample(example *Example) {
	if x != nil {
		x.ReferenceableExample = &ReferenceableExample_Example{
			Example: example,
		}
	}
}

func (x *ReferenceableExample) SetReference(reference *Reference) {
	if x != nil {
		x.ReferenceableExample = &ReferenceableExample_Reference{
			Reference: reference,
		}
	}
}

func (x *ReferenceableExample) GetExampleOf(index map[string]*Example) *Example {
	if x != nil {
		reference := x.GetReference()
		if reference != nil {
			fragment := reference.GetRef().GetFragment()
			key := strings.TrimPrefix(fragment, ExampleReferenceRoot)
			if s, ok := index[key]; ok {
				return s
			}
			return nil
		}
		return x.GetExample()
	}
	return nil
}
