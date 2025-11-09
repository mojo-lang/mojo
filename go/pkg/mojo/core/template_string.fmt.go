package core

import (
	"bytes"
	"regexp"
)

var segmentRegex *regexp.Regexp

func init() {
	segmentRegex = regexp.MustCompile(`\{[a-zA-Z\d_.]+\}`)
}

func (x *TemplateString) Parse(str string) error {
	if x != nil && len(str) > 0 {
		index := segmentRegex.FindAllStringIndex(str, -1)
		cur := 0
		for _, i := range index {
			left := i[0]
			right := i[1]

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
