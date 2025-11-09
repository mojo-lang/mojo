package query

type FieldProjection struct {
	Name      string   `json:"name,omitempty"`
	Uniques   []string `json:"uniques,omitempty"`
	Functions []string `json:"functions,omitempty"` // sum, min, max, avg, count
	Alias     []string `json:"alias,omitempty"`     // alias to sum(field)
}

type ProjectField struct {
	Projection *FieldProjection
	Index      int
}

func (f *ProjectField) GetFunction() string {
	if f != nil && f.Projection != nil && f.Index >= 0 && len(f.Projection.Functions) > f.Index {
		return f.Projection.Functions[f.Index]
	}
	return ""
}

func (f *ProjectField) GetName() string {
	if f != nil && f.Projection != nil {
		if len(f.Projection.Name) > 0 {
			return f.Projection.Name
		} else if f.Index >= 0 && len(f.Projection.Uniques) > f.Index {
			return f.Projection.Uniques[f.Index]
		}
	}
	return ""
}
