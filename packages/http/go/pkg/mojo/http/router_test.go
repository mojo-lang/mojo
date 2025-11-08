package http

import (
	"testing"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"
)

var attributes = []*lang.Attribute{
	{PackageName: "http", Name: "get", Arguments: []*lang.Argument{lang.NewStringArgument("/tile/v1/{x}/{y}/{z}")}},
	{PackageName: "http", Name: "header", Arguments: []*lang.Argument{lang.NewStringArgument("Connection"), lang.NewStringArgument("Keep-Alive")}},
}

func TestRouter_FromAttributes(t *testing.T) {
	router := &Router{}
	err := router.FromAttributes(attributes)
	assert.NoError(t, err)

	assert.Equal(t, "GET", router.Method.Format())
	assert.Equal(t, "/tile/v1/{x}/{y}/{z}", router.Path.Format())

	assert.Equal(t, 1, len(router.Headers))
	assert.Equal(t, "Connection", router.Headers[0].Name)
	assert.Equal(t, "Keep-Alive", router.Headers[0].Value.Format())
}

func TestRouter_ToAttributes(t *testing.T) {
	router := &Router{
		Path:    core.NewTemplateString("/tile/v1/{x}/{y}/{z}"),
		Headers: []*TemplateHeader{{Name: "Connection", Value: core.NewTemplateString("Keep-Alive")}},
		Method:  Method_METHOD_GET,
	}

	attrs := router.ToAttributes()
	assert.Equal(t, 2, len(attrs))
	assert.Equal(t, attributes[0].Name, attrs[0].Name)
	assert.Equal(t,
		attributes[0].Arguments[0].Value.GetStringLiteralExpr().EvalValue(),
		attrs[0].Arguments[0].Value.GetStringLiteralExpr().EvalValue())

	assert.Equal(t, attributes[1].Name, attrs[1].Name)
}

func TestRouter_ApplyTemplateValues(t *testing.T) {
	router := &Router{
		Path:    core.NewTemplateString("/tile/v1/{{service}}/{x}/{y}/{z}"),
		Headers: []*TemplateHeader{{Name: "Connection", Value: core.NewTemplateString("Keep-Alive-{{service}}")}},
		Method:  Method_METHOD_GET,
	}

	router.ApplyTemplateValues(map[string]string{"service": "3d"})

	assert.Equal(t, "/tile/v1/3d/{x}/{y}/{z}", router.Path.Format())
	assert.Equal(t, "Keep-Alive-3d", router.Headers[0].Value.Format())
}

func TestRouter_ApplyTemplateValues2(t *testing.T) {
	router := &Router{
		Path:    core.NewTemplateString("/tile/v1/{{service}}/{x}/{y}/{z}"),
		Headers: []*TemplateHeader{{Name: "Connection", Value: core.NewTemplateString("Keep-Alive-{{service}}")}},
		Method:  Method_METHOD_GET,
	}

	router.ApplyTemplateValues(map[string]string{"service": ""})

	assert.Equal(t, "/tile/v1/{x}/{y}/{z}", router.Path.Format())
	assert.Equal(t, "Keep-Alive-", router.Headers[0].Value.Format())
}
