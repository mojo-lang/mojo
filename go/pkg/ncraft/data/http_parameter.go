package data

// HTTPParameter represents the location of one field for a given HTTPBinding.
// Each HTTPParameter corresponds to one Field of the parent
// InterfaceMethod.Request.Fields
type HTTPParameter struct {
	Name     string
	FullName string

	// Field points to a Field on the Parent service methods "RequestType".
	Field *Field

	// Location will be either "body", "path", or "query"
	Location string

	// In order to support common ways of serializing simple parameters,
	// a set of style values are defined.
	Style string

	Go         *GoHTTPParameter
	Extensions map[string]interface{}
}

type GoHTTPParameter struct {
	//
	ParamUnmarshaler string
}

func (p *HTTPParameter) GetField() *Field {
	if p != nil {
		return p.Field
	}
	return nil
}

func (p *HTTPParameter) GetFieldType() *FieldType {
	return p.GetField().GetType()
}

func (p *HTTPParameter) GetElementType() *FieldType {
	return p.GetField().GetElementType()
}

func (p *HTTPParameter) IsEnum() bool {
	if typ := p.GetFieldType(); typ != nil {
		return typ.IsEnum
	}
	return false
}

func (p *HTTPParameter) IsMap() bool {
	if typ := p.GetFieldType(); typ != nil {
		return typ.IsMap
	}
	return false
}

func (p *HTTPParameter) IsArray() bool {
	if typ := p.GetFieldType(); typ != nil {
		return typ.IsArray
	}
	return false
}

func (p *HTTPParameter) IsScalar() bool {
	if typ := p.GetFieldType(); typ != nil {
		return typ.IsScalar
	}
	return false
}

func (p *HTTPParameter) IsElementScalar() bool {
	if typ := p.GetElementType(); typ != nil {
		return typ.IsScalar
	}
	return false
}

func (p *HTTPParameter) IsRawStyle() bool {
	return p != nil && p.Style == "raw"
}

func (p *HTTPParameter) GetEnclosingField() *Field {
	return p.GetField().GetEnclosingField()
}

func (p *HTTPParameter) GetGoType() *GoFieldType {
	return p.GetFieldType().GetGoType()
}

func (p *HTTPParameter) GetElementGoType() *GoFieldType {
	return p.GetFieldType().GetElementGoType()
}

func (p *HTTPParameter) GetGoTypeName() string {
	if g := p.GetGoType(); g != nil {
		return g.Name
	}
	return ""
}

func (p *HTTPParameter) GetElementGoTypeName() string {
	if g := p.GetElementGoType(); g != nil {
		return g.Name
	}
	return ""
}
