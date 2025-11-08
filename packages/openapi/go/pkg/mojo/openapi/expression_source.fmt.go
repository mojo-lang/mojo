package openapi

import (
	"strings"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

const (
	sourceHeaderPrefix = "header."
	sourceQueryPrefix  = "query."
	sourcePathPrefix   = "path."
	sourceBodyPrefix   = "body"
)

func ParseExpressionSource(value string) (*Expression_Source, error) {
	v := &Expression_Source{}
	if err := v.Parse(value); err != nil {
		return v, err
	}
	return v, nil
}

func (x *Expression_Source) Format() string {
	if x != nil {
		builder := new(strings.Builder)

		builder.WriteString(x.Location.Format())
		if len(x.Name) > 0 {
			builder.WriteByte('.')
			builder.WriteString(x.Name)
		} else if len(x.Path) > 0 {
			builder.WriteByte('#')
			builder.WriteString(x.Path)
		}
		return builder.String()
	}
	return ""
}

func (x *Expression_Source) Parse(value string) error {
	if x != nil {
		if len(value) >= 4 {
			switch value[0:4] {
			case "head":
				if strings.HasPrefix(value, sourceHeaderPrefix) {
					x.Location = Expression_LOCATION_HEADER
					x.Name = value[len(sourceHeaderPrefix):]
				} else {
					return core.NewMalformedSyntaxError("invalid expression source in header: %s", value)
				}
			case "quer":
				if strings.HasPrefix(value, sourceQueryPrefix) {
					x.Location = Expression_LOCATION_QUERY
					x.Name = value[len(sourceQueryPrefix):]
				} else {
					return core.NewMalformedSyntaxError("invalid expression source in query: %s", value)
				}
			case "path":
				if strings.HasPrefix(value, sourcePathPrefix) {
					x.Location = Expression_LOCATION_PATH
					x.Name = value[len(sourcePathPrefix):]
				} else {
					return core.NewMalformedSyntaxError("invalid expression source in path: %s", value)
				}
			case "body":
				if strings.HasPrefix(value, sourceBodyPrefix) {
					x.Location = Expression_LOCATION_BODY
					if len(value) > 4 {
						x.Path = value[len(sourcePathPrefix):]
					}
				} else {
					return core.NewMalformedSyntaxError("invalid expression source in path: %s", value)
				}
			default:
				return core.NewMalformedSyntaxError("invalid expression source: %s", value)
			}
			return nil
		}
		return core.NewMalformedSyntaxError("invalid expression source: %s", value)
	}
	return nil
}
