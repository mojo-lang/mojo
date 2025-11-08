package http

import (
	"errors"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"regexp"
	"strings"
)

func (x *Router) FromAttributes(attributes []*lang.Attribute) error {
	for _, attribute := range attributes {
		if attribute.PackageName != "http" {
			continue
		}

		switch attribute.Name {
		case "get", "post", "put", "patch", "delete", "options", "head", "trace", "connect":
			if err := x.Method.Parse(strings.ToUpper(attribute.Name)); err != nil {
				return err
			}
			if len(attribute.Arguments) > 0 {
				if value := attribute.Arguments[0].GetStringLiteralExpr().EvalValue(); len(value) > 0 {
					x.Path = core.NewTemplateString(value)
					if x.Path == nil {
						return logs.NewErrorw("failed to parse the http router path", "path", value, "method", x.Method.Format())
					}
				}
			}
		case HeaderAttributeName:
			if len(attribute.Arguments) == 2 {
				name, value := "", ""
				if name = attribute.Arguments[0].GetStringLiteralExpr().EvalValue(); len(name) == 0 {
					return errors.New("failed to parse the http header attribute")
				}
				if value = attribute.Arguments[1].GetStringLiteralExpr().EvalValue(); len(value) > 0 {
					x.Headers = append(x.Headers, &TemplateHeader{
						Name:  name,
						Value: core.NewTemplateString(value),
					})
				}
			}
		}
	}
	return nil
}

func (x *Router) ToAttributes() []*lang.Attribute {
	var attributes []*lang.Attribute
	if x.Path != nil {
		attributes = append(attributes, &lang.Attribute{
			PackageName: "http",
			Name:        strings.ToLower(x.Method.Format()),
			Arguments:   []*lang.Argument{lang.NewStringArgument(x.Path.Format())},
		})
	}
	for _, header := range x.Headers {
		attributes = append(attributes, &lang.Attribute{
			PackageName: "http",
			Name:        HeaderAttributeName,
			Arguments: []*lang.Argument{lang.NewStringArgument(header.Name),
				lang.NewStringArgument(header.Value.Format()),
			},
		})
	}

	return attributes
}

var templateRegex *regexp.Regexp

func init() {
	templateRegex = regexp.MustCompile(`\{\{[a-z0-9_.]+\}\}`)
}

func applyTemplateValues(str string, values map[string]string) string {
	return templateRegex.ReplaceAllStringFunc(str, func(s string) string {
		bare := s[2 : len(s)-2]
		if m, ok := values[bare]; ok {
			return m
		}
		return ""
	})
}

func (x *Router) ApplyTemplateValues(values map[string]string) {
	if x != nil {
		if x.Path != nil {
			for _, segment := range x.Path.Segments {
				if !segment.Templated {
					segment.Content = applyTemplateValues(segment.Content, values)
					segment.Content = strings.ReplaceAll(segment.Content, "//", "/")
				}
			}
		}
		for _, header := range x.Headers {
			for _, segment := range header.GetValue().GetSegments() {
				if !segment.Templated {
					segment.Content = applyTemplateValues(segment.Content, values)
				}
			}
		}
	}
}
