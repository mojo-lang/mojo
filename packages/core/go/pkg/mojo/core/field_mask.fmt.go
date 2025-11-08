package core

import "strings"

func ParseFieldMask(fieldMask string) (*FieldMask, error) {
	mask := NewFieldMask(fieldMask)
	return mask, nil
}

func (x *FieldMask) Format() string {
	if x != nil && len(x.Paths) > 0 {
		return strings.Join(x.Paths, ",")
	} else {
		return ""
	}
}

func (x *FieldMask) ToString() string {
	return x.Format()
}

func (x *FieldMask) Parse(field string) error {
	var paths []string
	paths = parseFields(field, paths)

	x.Paths = paths
	return nil
}

func parseFields(fields string, segments []string) []string {
	if len(fields) == 0 {
		return segments
	}

	comma := strings.Index(fields, ",")
	leftBrace := strings.Index(fields, "{")

	if comma < 0 {
		if leftBrace < 0 {
			segments = append(segments, strings.TrimSpace(fields))
		} else {
			rightBrace := strings.Index(fields, "}")
			if rightBrace > 0 {
				segments = append(segments, strings.TrimSpace(fields[:leftBrace])+"."+strings.TrimSpace(fields[leftBrace+1:rightBrace]))
			} else {
				segments = nil
			}
		}
	} else if leftBrace < 0 {
		list := strings.Split(fields, ",")
		for i := 0; i < len(list); i++ {
			if len(list[i]) > 0 {
				segments = append(segments, strings.TrimSpace(list[i]))
			}
		}
	} else {
		if leftBrace < comma {
			rightBrace := strings.Index(fields, "}")
			if rightBrace <= 0 {
				segments = nil
			} else {
				parent := strings.TrimSpace(fields[:leftBrace])
				brace := fields[leftBrace+1 : rightBrace]

				var subs []string
				subs = parseFields(brace, subs)
				if len(subs) == 0 {
					segments = append(segments, parent)
				} else {
					for i := 0; i < len(subs); i++ {
						segments = append(segments, parent+"."+strings.TrimSpace(subs[i]))
					}
				}
				segments = parseFields(fields[rightBrace+1:], segments)
			}
		} else {
			segments = append(segments, strings.TrimSpace(fields[:comma]))
			segments = parseFields(fields[comma+1:], segments)
		}
	}

	return segments
}
