package core

import (
	"net/url"
)

func (x *Url_Query) Parse(value string) error {
	if x != nil {
		values, err := url.ParseQuery(value)
		if err != nil {
			return err
		}
		x.FromUrlValues(values)
	}
	return nil
}

func (x *Url_Query) Format() string {
	if x != nil {
		if values := x.ToUrlValues(); values != nil {
			return x.ToUrlValues().Encode()
		}
	}
	return ""
}
