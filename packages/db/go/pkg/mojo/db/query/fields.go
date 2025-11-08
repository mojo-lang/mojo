package query

type Fields map[string]Field

func (f Fields) Repeated(field string) bool {
	if f != nil {
		if info, ok := f[field]; ok {
			return info.Repeated
		}
	}
	return false
}

func (f Fields) Type(field string) FieldType {
	if f != nil {
		if info, ok := f[field]; ok {
			return info.Type
		}
	}
	return FieldType(0)
}
