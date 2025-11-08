package openapi

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/yaml/go/pkg/mojo/yaml"
)

type APIMarshaller func(v interface{}) ([]byte, error)
type APIUnmarshaler func(data []byte, v interface{}) error

func New() *OpenAPI {
	return &OpenAPI{
		Openapi: &core.Version{
			Major: 3,
			Minor: 0,
			Patch: 3,
		},
		Info:       &Info{},
		Paths:      &Paths{Vals: map[string]*PathItem{}},
		Components: &Components{Schemas: map[string]*Schema{}},
	}
}

func Parse(api []byte) (*OpenAPI, error) {
	return parse(api, yaml.Marshal, yaml.Unmarshal)
}

func ParseFile(filename string) (*OpenAPI, error) {
	if yaml.IsYAMLSuffix(filename) || strings.HasSuffix(filename, ".json") {
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to read the file (%s) %w", filename, err)
		}

		var marshaller APIMarshaller
		var unmarshaler APIUnmarshaler

		if yaml.IsYAMLSuffix(filename) {
			marshaller, unmarshaler = yaml.Marshal, yaml.Unmarshal
		} else if strings.HasSuffix(filename, ".json") {
			marshaller, unmarshaler = jsoniter.Marshal, jsoniter.Unmarshal
		}

		return parse(data, marshaller, unmarshaler)
	}
	return nil, fmt.Errorf("invalid OpenAPI file suffix, filename: %s", filename)
}

func parse(data []byte, marshaller APIMarshaller, unmarshaler APIUnmarshaler) (*OpenAPI, error) {
	oa := &OpenAPI{}
	err := unmarshaler(data, oa)
	if err != nil {
		logs.Warnw("failed to parse the openapi", "error", err)
		return nil, fmt.Errorf("failed to parse the openapi due to %w", err)
	}

	if !oa.CheckVersion("3.0") {
		// try to upgrade openapi to latest version
		logs.Info("try to upgrade openapi to latest version due to not supported openapi version!")
		o, err := Upgrade(data, marshaller, unmarshaler)
		if err != nil {
			logs.Warnw("failed to upgrade openAPI to latest version", "error", err)
			return nil, fmt.Errorf("failed to upgrade openAPI to latest version due to %w", err)
		}
		oa = o
	}

	if oa.Info == nil || oa.Openapi == nil {
		logs.Warnw("invalid the openapi document, need info & openapi fields")
		return nil, fmt.Errorf("invalid the openapi document, need info & openapi fields")
	}

	return oa, nil
}

func Upgrade(data []byte, marshaller APIMarshaller, unmarshaler APIUnmarshaler) (*OpenAPI, error) {
	var err error
	start := time.Now()
	data, err = lintAndRefine(data, marshaller, unmarshaler)
	if err != nil {
		return nil, err
	}

	v2 := &openapi2.T{}
	if err = unmarshaler(data, v2); err != nil {
		logs.Warnw("failed to unmarshal openapi2 document", "error", err, "buffer", string(data))
		return nil, err
	}

	if len(v2.Schemes) == 0 {
		var scheme string
		u, err := core.ParseUrl(v2.Host)
		if err == nil {
			scheme = u.Scheme
		}

		if scheme == "" {
			scheme = "http"
		}

		v2.Schemes = []string{scheme}
	}

	if len(v2.BasePath) > 0 && len(strings.TrimPrefix(v2.BasePath, "/")) > 0 {
		newPaths := make(map[string]*openapi2.PathItem, len(v2.Paths))
		for key, item := range v2.Paths {
			if key == "/" {
				key = v2.BasePath
			} else {
				key = path.Join(v2.BasePath, key)
			}
			newPaths[key] = item
		}
		v2.BasePath = ""
		v2.Paths = newPaths
	}

	v3, err := openapi2conv.ToV3(v2)
	if err != nil {
		logs.Warnw("failed to convert swagger2 to openapi3", "error", err)
		return nil, err
	}

	data, err = marshaller(v3)
	if err != nil {
		logs.Warnw("failed to marshal openapi3", "error", err)
		return nil, err
	}

	oa := &OpenAPI{}
	err = unmarshaler(data, oa)
	if err != nil {
		logs.Warnw("failed to unmarshal response body", "error", err)
		return nil, err
	}

	logs.Infow("api upgrade successfully", "upgraded version", oa.GetOpenapi().ToString(), "duration", time.Since(start).String())
	return oa, nil
}

func lintAndRefine(buffer []byte, marshaller APIMarshaller, unmarshaler APIUnmarshaler) ([]byte, error) {
	lintResult := EasyLint(buffer)
	if lintResult.Valid {
		return buffer, nil
	}

	v2 := &openapi2.T{}
	if err := unmarshaler(buffer, v2); err != nil {
		logs.Warnw("failed to unmarshal openapi2 document", "error", err, "buffer", string(buffer))
		return nil, err
	}

	for _, operation := range lintResult.Operations {
		if operation.Valid {
			continue
		}

		operations, found := v2.Paths[operation.Path]
		if !found {
			continue
		}

		// remove invalid operations
		switch strings.ToUpper(operation.Method) {
		case "GET":
			operations.Get = nil
		case "POST":
			operations.Post = nil
		case "PUT":
			operations.Put = nil
		case "DELETE":
			operations.Delete = nil
		case "HEAD":
			operations.Head = nil
		case "OPTIONS":
			operations.Options = nil
		case "PATCH":
			operations.Patch = nil
		default:
			continue
		}

		logs.Debugw("invalid operation has been deleted", "operation", operation)
	}

	buf, err := marshaller(v2)
	if err != nil {
		logs.Warnw("failed to marshal openapi2 document", "error", err, "v2", v2)
		return nil, err
	}

	return buf, nil
}

func (x *OpenAPI) GetPath(name string) *PathItem {
	if x != nil && x.Paths != nil && x.Paths.Vals != nil {
		if pi, ok := x.Paths.Vals[name]; ok {
			return pi
		}
	}
	return nil
}

// CheckVersion check the version is "3.0", or "3.1"
func (x *OpenAPI) CheckVersion(version string) bool {
	if x != nil && x.Openapi != nil {
		v, err := core.ParseVersion(version)
		if err != nil {
			return false
		}

		if v.Level > 0 && v.Major != x.Openapi.Major {
			return false
		}
		if v.Level > 1 && v.Minor != x.Openapi.Minor {
			return false
		}
		if v.Level > 2 && v.Patch != x.Openapi.Patch {
			return false
		}

		return true
	}
	return false
}

func (x *OpenAPI) GetOperation(id string) (path string, method string, operation *Operation, err error) {
	if x != nil && x.Paths != nil {
		for k, v := range x.Paths.Vals {
			for _, method = range []string{"Get", "Put", "Post", "Delete", "Options", "Head", "Patch", "Trace"} {
				ok := false
				op := reflect.ValueOf(v).Elem().FieldByName(method)
				if !op.IsNil() {
					if operation, ok = op.Interface().(*Operation); ok && operation.OperationId == id {
						path = k
						method = strings.ToUpper(method)
						return
					}
				}
			}
		}
	}
	return "", "", nil, core.NewNotFoundError("id: %s", id)
}

func (x *OpenAPI) GenerateExample() {
	if x != nil && x.Paths != nil {
		for _, p := range x.Paths.Vals {
			p.GenerateExample(x.Components)
		}
	}
}

func (x *OpenAPI) SupplementExample() {
	if x != nil && x.Paths != nil {
		for _, p := range x.Paths.Vals {
			p.SupplementExample(x.Components)
		}
	}
}
