package openapi

import "strings"

const LinkReferenceRoot = "/components/links/"

func NewReferenceableLink(link *Link) *ReferenceableLink {
	s := &ReferenceableLink{}
	s.SetLink(link)
	return s
}

func NewReferencedLink(reference *Reference) *ReferenceableLink {
	s := &ReferenceableLink{}
	s.SetReference(reference)
	return s
}

func (x *ReferenceableLink) SetLink(link *Link) {
	if x != nil {
		x.ReferenceableLink = &ReferenceableLink_Link{
			Link: link,
		}
	}
}

func (x *ReferenceableLink) SetReference(reference *Reference) {
	if x != nil {
		x.ReferenceableLink = &ReferenceableLink_Reference{
			Reference: reference,
		}
	}
}

func (x *ReferenceableLink) GetLinkOf(index map[string]*Link) *Link {
	if x != nil {
		reference := x.GetReference()
		if reference != nil {
			fragment := reference.GetRef().GetFragment()
			key := strings.TrimPrefix(fragment, LinkReferenceRoot)
			if s, ok := index[key]; ok {
				return s
			}
			return nil
		}
		return x.GetLink()
	}
	return nil
}
