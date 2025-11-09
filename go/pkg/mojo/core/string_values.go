package core

import "regexp"

const StringValuesTypeName = "StringValues"
const StringValuesTypeFullName = "mojo.core.StringValues"

func NewStringValues(values ...string) *StringValues {
	return &StringValues{Vals: values}
}

func (x *StringValues) Unique() *StringValues {
	if x != nil {
		keys := make(map[string]bool)
		var list []string
		for _, entry := range x.Vals {
			if _, ok := keys[entry]; !ok {
				keys[entry] = true
				list = append(list, entry)
			}
		}
		x.Vals = list
	}
	return x
}

func (x *StringValues) Contains(element string) bool {
	if x != nil {
		for _, e := range x.Vals {
			if e == element {
				return true
			}
		}
	}
	return false
}

func (x *StringValues) Matched(expr string) bool {
	if rgx, err := regexp.Compile(expr); err == nil {
		for _, e := range x.Vals {
			if rgx.MatchString(e) {
				return true
			}
		}
	}
	return false
}

func (x *StringValues) Matches(expr string) *StringValues {
	values := &StringValues{}

	if rgx, err := regexp.Compile(expr); err == nil {
		for _, e := range x.Vals {
			if rgx.MatchString(e) {
				values.Vals = append(values.Vals, e)
			}
		}
	}
	return values
}

func (x *StringValues) ToArray() interface{} {
	if x != nil {
		return x.Vals
	}
	return []string{}
}

func (x *StringValues) Append(vs ...string) *StringValues {
	if x != nil {
		x.Vals = append(x.Vals, vs...)
	}
	return x
}
