package http

import (
	"bytes"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"text/template"
)

const (
	RouterTemplateValuesAttributeName     = "router_template_values"
	RouterTemplateValuesAttributeFullName = "http.router_template_values"
)

func GetRouterTemplateValues(attributes ...*lang.Attribute) map[string]string {
	if attribute := lang.Attributes(attributes).GetAttribute(RouterTemplateValuesAttributeFullName); attribute != nil {
		values := make(map[string]string)
		for _, argument := range attribute.Arguments {
			if len(argument.Label) > 0 {
				if value, err := argument.GetString(); err != nil {
					logs.Warnw("failed to parse the string argument",
						"attribute", RouterTemplateValuesAttributeFullName, "label", argument.Label)
				} else {
					values[argument.Label] = value
				}
			}
		}
		return values
	}

	return nil
}

func ApplyRouterTemplateValues(attributes lang.Attributes, values map[string]string) lang.Attributes {
	if len(values) > 0 {
		for _, attribute := range attributes {
			if attribute.PackageName != "http" {
				continue
			}

			switch attribute.Name {
			case GetAttributeName, PostAttributeName, PutAttributeName, PatchAttributeName,
				DeleteAttributeName, OptionsAttributeName, HeadAttributeName, TraceAttributeName:
				if path, _ := attribute.GetString(); len(path) > 0 {
					if tmpl, err := template.New("http-router-template").Parse(path); err == nil {
						buffer := bytes.NewBuffer(nil)
						if err = tmpl.Execute(buffer, values); err == nil {
							attribute.SetString(buffer.String())
						}
					}
				}
			}
		}
	}
	return attributes
}
