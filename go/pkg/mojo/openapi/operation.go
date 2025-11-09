package openapi

func (x *Operation) GetParameter(index int, parameters map[string]*Parameter) *Parameter {
	if x != nil && index >= 0 && index < len(x.Parameters) {
		return x.Parameters[index].GetParameterOf(parameters)
	}
	return nil
}

func (x *Operation) GetLocationParameters() map[Parameter_Location][]*Parameter {
	ps := make(map[Parameter_Location][]*Parameter)
	parameters := x.GetParameters()
	for _, parameter := range parameters {
		p := parameter.GetParameter()
		if p != nil {
			ps[p.In] = append(ps[p.In], p)
		}
	}
	return ps
}

func (x *Operation) GenerateExample(components *Components) {
	if x != nil {
		for _, param := range x.Parameters {
			p := param.GetParameterOf(components.GetParameters())
			p.GenerateExample(components.GetSchemas())
		}

		x.RequestBody.GetRequestBodyOf(components.GetRequestBodies()).GenerateExample(components.GetSchemas())

		for _, resp := range x.Responses.GetVals() {
			r := resp.GetResponseOf(components.GetResponses())
			r.GenerateExample(components.GetSchemas())
		}
	}
}

func (x *Operation) SupplementExample(components *Components) {
	if x != nil {
		for _, param := range x.Parameters {
			p := param.GetParameterOf(components.GetParameters())
			p.SupplementExample(components.GetSchemas())
		}

		x.RequestBody.GetRequestBodyOf(components.GetRequestBodies()).SupplementExample(components.GetSchemas())

		for _, resp := range x.Responses.GetVals() {
			r := resp.GetResponseOf(components.GetResponses())
			r.SupplementExample(components.GetSchemas())
		}
	}
}
