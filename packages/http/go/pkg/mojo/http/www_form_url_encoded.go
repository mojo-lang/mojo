package http

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core/strcase"
	"net/url"
	"strconv"
)

func WwwFormUrlencodedToJSON(content []byte) ([]byte, error) {
	values, err := url.ParseQuery(string(content))
	if err != nil {
		return nil, err
	}

	json := make(map[string]interface{})
	for k, v := range values {
		key := strcase.ToLowerCamel(k)
		if len(v) == 0 {
			json[key] = nil
		} else if len(v) == 1 {
			json[key] = parseValue(v[0])
		} else {
			var vs []interface{}
			for _, value := range v {
				vs = append(vs, parseValue(value))
			}
			json[key] = vs
		}
	}

	return jsoniter.Marshal(json)
}

func parseValue(value string) interface{} {
	if len(value) == 0 {
		return nil
	}
	if value == "null" {
		return nil
	}
	if v, err := strconv.ParseBool(value); err == nil {
		return v
	}
	if v, err := strconv.ParseInt(value, 10, 64); err == nil {
		return v
	}
	if v, err := strconv.ParseFloat(value, 64); err == nil {
		return v
	}

	return value
}
