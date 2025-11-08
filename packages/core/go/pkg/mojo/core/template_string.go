package core

import "strconv"

const TemplateStringTypeName = "TemplateString"
const TemplateStringTypeFullName = "mojo.core.TemplateString"

func NewTemplateString(str string) *TemplateString {
	ts := &TemplateString{}
	if err := ts.Parse(str); err != nil {
		return nil
	}

	return ts
}

func (x *TemplateString) AppendSegment(segment string) {
	if x != nil {
		x.Segments = append(x.Segments, &TemplateString_Segment{Content: segment})
	}
}

func (x *TemplateString) AppendTemplatedSegment(segment string) {
	if x != nil {
		x.Segments = append(x.Segments,
			&TemplateString_Segment{Content: segment, Templated: true},
		)
	}
}

// Apply the values to the template string
// todo currently only support simplest string value substitution
func (x *TemplateString) Apply(values map[string]interface{}) (string, error) {
	if x != nil {
		for _, segment := range x.Segments {
			if segment.Templated {
				if value, ok := values[segment.Content]; ok {
					switch v := value.(type) {
					case string:
						segment.Templated = false
						segment.Content = v
					case int32:
						segment.Templated = false
						segment.Content = strconv.Itoa(int(v))
					case int64:
						segment.Templated = false
						segment.Content = strconv.Itoa(int(v))
					}
				}
			}
		}
		return x.Format(), nil
	}

	return "", nil
}

func (x *TemplateString) ApplyStrings(values map[string]string) (string, error) {
	if x != nil {
		vs := make(map[string]interface{}, len(values))
		for k, v := range values {
			vs[k] = v
		}
		return x.Apply(vs)
	}

	return "", nil
}
