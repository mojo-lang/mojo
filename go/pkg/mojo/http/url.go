package http

import "github.com/mojo-lang/mojo/go/pkg/mojo/core"

func UnmarshalPathParam(params map[string]string, value interface{}, keys ...string) error {
	val := ""
	key := ""
	for _, key = range keys {
		if val, _ = params[key]; len(val) > 0 {
			break
		}
	}

	if len(val) == 0 {
		return core.NewNotFoundError("can't found the parameter, %s", key)
	}

	return core.UnmarshalParam(val, value)
}

func UnmarshalQueryParam(params *core.Url_Query, value interface{}, keys ...string) error {
	key := ""
	for _, k := range keys {
		if params.Has(k) {
			key = k
		}
	}
	if len(key) == 0 {
		return core.NewNotFoundError("can't found the parameter, %s", key)
	}
	return params.Unmarshal(key, value)
}
