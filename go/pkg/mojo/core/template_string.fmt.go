package core

import (
	"bytes"
	"regexp"
)

// A pattern may contain nested braces, e.g. {id:[0-9]{2,4}}. Match only the
// opening name here and scan balanced braces for the complete segment.
var segmentStartRegex = regexp.MustCompile(`\{[a-zA-Z\d_.]+[}:]`)

func (x *TemplateString) Parse(str string) error {
	if x != nil && len(str) > 0 {
		cur := 0
		for scan := 0; scan < len(str); {
			index := segmentStartRegex.FindStringIndex(str[scan:])
			if index == nil {
				break
			}
			left, right := scan+index[0], scan+index[1]
			if str[right-1] == ':' {
				depth := 1
				for right < len(str) && depth > 0 {
					switch str[right] {
					case '{':
						depth++
					case '}':
						depth--
					}
					right++
				}
				// Preserve incomplete templates as literal text, as for {name.
				if depth != 0 {
					break
				}
			}
			scan = right

			// {{something}} will be skipped
			if left > 0 && right < len(str) && str[left-1] == '{' && str[right] == '}' {
				continue
			}

			if left > cur {
				x.Segments = append(x.Segments, &TemplateString_Segment{Content: str[cur:left]})
			}

			content := str[left+1 : right-1]
			x.Segments = append(x.Segments, &TemplateString_Segment{Content: content, Templated: true})
			cur = right
		}

		if cur < len(str) {
			x.Segments = append(x.Segments, &TemplateString_Segment{Content: str[cur:]})
		}
	}
	return nil
}

func (x *TemplateString) Format() string {
	if x != nil {
		buffer := bytes.Buffer{}
		for _, segment := range x.Segments {
			if segment.Templated {
				buffer.WriteByte('{')
				buffer.WriteString(segment.Content)
				buffer.WriteByte('}')
			} else {
				buffer.WriteString(segment.Content)
			}
		}
		return buffer.String()
	}
	return ""
}

func (x *TemplateString) ToString() string {
	return x.Format()
}
