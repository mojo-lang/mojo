package openapi

func (x *RequestBody) GenerateExample(index map[string]*Schema) {
	if x != nil {
		for key, mt := range x.Content {
			mt.GenerateExample(index, key)
		}
	}
}

func (x *RequestBody) SupplementExample(index map[string]*Schema) {
	if x != nil {
		for key, mt := range x.Content {
			mt.SupplementExample(index, key)
		}
	}
}
