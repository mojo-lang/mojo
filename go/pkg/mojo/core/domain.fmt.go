package core

import (
	"strings"
)

func (x *Domain) Format() string {
	if x != nil {
		first := true
		sb := strings.Builder{}
		for i := len(x.Labels) - 1; i >= 0; i-- {
			if first {
				first = false
			} else {
				sb.WriteByte('.')
			}
			sb.WriteString(x.Labels[i])
		}
		return sb.String()
	}
	return ""
}

func (x *Domain) ToString() string {
	return x.Format()
}

func (x *Domain) Parse(value string) error {
	if x != nil && len(value) > 0 {
		segments := strings.Split(value, ".")
		for i := len(segments) - 1; i >= 0; i-- {
			x.Labels = append(x.Labels, segments[i])
		}
	}
	return nil
}

func ParseDomain(value string) (*Domain, error) {
	domain := &Domain{}
	if err := domain.Parse(value); err != nil {
		return nil, err
	}
	return domain, nil
}
