package openapi

func (x *Response) GenerateExample(index map[string]*Schema) {
	if x != nil {
		for key, mt := range x.Content {
			mt.GenerateExample(index, key)
		}
	}
}

func (x *Response) SupplementExample(index map[string]*Schema) {
	if x != nil {
		for key, mt := range x.Content {
			mt.SupplementExample(index, key)
		}
	}
}
