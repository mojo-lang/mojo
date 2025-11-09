package core

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var urlSchemaPattern = regexp.MustCompile(`^([a-z0-9+\-.]+://)|mailto:|news:`)

func ParseUrl(raw string) (*Url, error) {
	u := &Url{}
	err := u.Parse(raw)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (x *Url) Parse(raw string) error {
	if x == nil {
		return errors.New("")
	}

	u, err := url.Parse(raw)
	if err != nil {
		return err
	}

	x.Scheme = u.Scheme
	x.Authority = &Url_Authority{
		Host: u.Hostname(),
	}

	if len(x.Authority.Host) == 0 && (len(u.Scheme) > 0 || len(u.Opaque) > 0) && !urlSchemaPattern.MatchString(raw) {
		u, err = url.Parse("https://" + raw)
		if err != nil {
			return err
		}
		x.Scheme = ""
		x.Authority.Host = u.Hostname()
		if len(x.Authority.Host) == 0 {
			return errors.New("failed to parse the url")
		}
	}

	port := u.Port()
	if len(port) > 0 {
		v, err := strconv.Atoi(port)
		if err != nil {
			return err
		}
		x.Authority.Port = int64(v)
	}

	x.Path = u.Path
	x.Fragment = u.Fragment

	x.Query = &Url_Query{
		Vals: make(map[string]*StringValues),
	}

	if query, err := url.ParseQuery(u.RawQuery); err != nil {
		return err
	} else {
		for k, v := range query {
			x.Query.Vals[k] = &StringValues{
				Vals: v,
			}
		}
	}

	return nil
}

func (x *Url) Format() string {
	if x == nil {
		return ""
	}

	host := ""
	if x.Authority != nil {
		host = x.Authority.Host
		if x.Authority.Port > 0 {
			host = host + ":" + strconv.FormatInt(x.Authority.Port, 10)
		}
	}

	u := url.URL{
		Scheme:   x.Scheme,
		Host:     host,
		Path:     x.Path,
		Fragment: x.Fragment,
	}

	if x.Query != nil {
		query := url.Values{}
		for k, v := range x.Query.Vals {
			if v != nil {
				query[k] = v.Vals
			}
		}
		u.RawQuery = query.Encode()
	}

	if len(u.Scheme) > 0 {
		return u.String()
	} else {
		return strings.TrimPrefix(u.String(), "//")
	}
}

func (x *Url) ToString() string {
	return x.Format()
}

// FormatWithoutSchema get url like: "apis.company.com/path/to/resource"
func (x *Url) FormatWithoutSchema() string {
	if x == nil {
		return ""
	}

	u := &Url{
		Authority: &Url_Authority{
			Host: x.GetAuthority().GetHost(),
			Port: x.GetAuthority().GetPort(),
		},
		Path:     x.Path,
		Query:    x.Query,
		Fragment: x.Fragment,
	}
	return strings.TrimPrefix(u.Format(), "//")
}
