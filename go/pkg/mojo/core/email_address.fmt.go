package core

import "strings"

func (x *EmailAddress) Format() string {
	if x != nil {
		return x.LocalPart + "@" + x.Domain.Format()
	}
	return ""
}

func (x *EmailAddress) ToString() string {
	return x.Format()
}

func (x *EmailAddress) Parse(value string) error {
	if x != nil && len(value) > 0 {
		segments := strings.Split(value, "@")
		if len(segments) == 2 {
			x.LocalPart = segments[0]
			domain, err := ParseDomain(segments[1])
			if err != nil {
				return err
			}
			x.Domain = domain
		} else {
			return NewMalformedSyntaxError("invalid email address: %s", value)
		}
	}
	return nil
}

func ParseEmailAddress(value string) (*EmailAddress, error) {
	email := &EmailAddress{}
	if err := email.Parse(value); err != nil {
		return nil, err
	}
	return email, nil
}
