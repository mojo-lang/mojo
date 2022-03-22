package templates

const Middlewares = `
package handlers

import (
	"github.com/ncraft-io/ncraft-gokit/pkg/middleware"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"

	{{range $i := .ExternalMessageImports}}
	"{{$i}}"
	{{- end}}

	"{{.RepositoryPath -}}/pkg/{{ToKebab .Interface.Name}}-service/svc"
	pb "{{.ApiImportPath -}}"
)

var ({{range $name := .ExternalStructs}}
	_ = {{PackageName $name}}.{{$name}}{}
{{- end}}{{range $name := .ExternalEnums}}
	_ = {{PackageName $name}}.{{$name}}(0)
{{- end}})

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints, options map[string]interface{}) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/frankee/ncraft/_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)
	
	var tracer stdopentracing.Tracer
	if value, ok := options["tracer"]; ok && value != nil{
		tracer = value.(stdopentracing.Tracer)
	}
	var count *kitprometheus.Counter
	if value, ok := options["count"]; ok && value != nil {
		count = value.(*kitprometheus.Counter)
	}
	var latency *kitprometheus.Histogram
	if value, ok := options["latency"]; ok && value != nil {
		latency = value.(*kitprometheus.Histogram)
	}
	//var validator *middleware.Validator
	//if value, ok := options["validator"]; ok && value != nil {
	//	validator = value.(*middleware.Validator)
	//}

	{{range $i := .Interface.Methods}}
	{ // {{$i.Name}}
		if tracer != nil {
			in.{{ToCamel $i.Name}}Endpoint = opentracing.TraceServer(tracer, "{{$i.Name}}")(in.{{ToCamel $i.Name}}Endpoint)
 		}
		if count != nil && latency != nil {
			in.{{ToCamel $i.Name}}Endpoint = middleware.Instrumenting(latency.With("method", "{{$i.Name}}"), count.With("method", "{{$i.Name}}"))(in.{{ToCamel $i.Name}}Endpoint)
		}
		//if validator != nil {
		//	in.{{ToCamel $i.Name}}Endpoint = validator.Validate()(in.{{ToCamel $i.Name}}Endpoint)
		//}
	}
	{{- end}}

	return in
}

func WrapService(in pb.{{.Interface.Name}}Server, options map[string]interface{}) pb.{{.Interface.Name}}Server {
	return in
}
`
