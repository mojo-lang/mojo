package core

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/logs"
	"strings"
)

type Options map[string]interface{}

func NewOptions(kvs ...interface{}) Options {
	options := make(Options)
	return options.SetValues(kvs...)
}

func (o Options) SetValue(key string, value interface{}) Options {
	o[key] = value
	return o
}

func (o Options) HasValue(key string) bool {
	if _, ok := o[key]; ok {
		return true
	}
	return false
}

func (o Options) GetValue(key string) interface{} {
	if val, ok := o[key]; ok {
		return val
	}
	return nil
}

func (o Options) GetValueByPath(key string) interface{} {
	if len(key) == 0 {
		return nil
	}

	segments := strings.Split(key, ".")
	options := o
	for i, segment := range segments {
		if option, ok := options[segment]; ok {
			if i == len(segments)-1 {
				return option
			} else if op, ok := option.(map[string]interface{}); ok {
				options = op
			} else {
				logs.Warnw("not found", "key", key, "options", options)
				return nil
			}
		} else {
			logs.Warnw("not found", "key", key, "options", options)
			return nil
		}
	}
	return nil
}

func (o Options) GetBool(key string) bool {
	if val, ok := o[key]; ok {
		if v, ok := val.(bool); ok {
			return v
		}
	}
	return false
}

func (o Options) GetString(key string) string {
	if val, ok := o[key]; ok {
		if v, ok := val.(string); ok {
			return v
		}
	}
	return ""
}

func (o Options) GetInteger(key string) int64 {
	if val, ok := o[key]; ok {
		switch v := val.(type) {
		case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64:
			return int64(v.(int))
		}
	}
	return 0
}

func (o Options) GetFloat(key string) float64 {
	if val, ok := o[key]; ok {
		switch v := val.(type) {
		case float32:
			return float64(v)
		case float64:
			return v
		case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64:
			return float64(v.(int))
		default:
			return 0
		}
	}
	return 0
}

func (o Options) SetValues(kvs ...interface{}) Options {
	for i := 0; i < len(kvs)-1; i += 2 {
		o[kvs[i].(string)] = kvs[i+1]
	}
	return o
}

func (o Options) Merge(options Options) Options {
	for k, v := range options {
		o[k] = v
	}
	return o
}

func (o Options) KeyValues() []interface{} {
	var kvs []interface{}
	if o != nil {
		for k, v := range o {
			kvs = append(kvs, k)
			kvs = append(kvs, v)
		}
	}
	return kvs
}

func (o Options) To(newOptions interface{}) error {
	if len(o) > 0 {
		bytes, err := jsoniter.ConfigFastest.Marshal(o)
		if err != nil {
			return err
		}
		return jsoniter.ConfigFastest.Unmarshal(bytes, newOptions)
	}
	return nil
}
