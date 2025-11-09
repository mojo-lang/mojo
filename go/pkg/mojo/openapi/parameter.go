package openapi

func (x *Parameter) GenerateExample(index map[string]*Schema) {
	if x != nil {
		if example := x.Schema.GetSchemaOf(index).GenerateExample(index); example != nil {
			x.Example = example
		}
	}
}

func (x *Parameter) SupplementExample(index map[string]*Schema) {
	if x != nil && x.Example == nil && len(x.Examples) == 0 {
		if x.Schema.GetSchema() == nil || x.Schema.GetSchema().Example == nil {
			x.Example = x.Schema.GetSchemaOf(index).SupplementExample(index)
		}
	}
}
