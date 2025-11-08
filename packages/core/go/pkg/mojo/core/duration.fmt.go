package core

import "time"

func (x *Duration) Format() string {
	if x != nil {
		return x.ToDuration().String()
	}
	return ""
}

func (x *Duration) ToString() string {
	return x.Format()
}

func ParseDuration(value string) (*Duration, error) {
	duration := &Duration{}
	if err := duration.Parse(value); err != nil {
		return nil, err
	}
	return duration, nil
}

func (x *Duration) Parse(value string) error {
	if x != nil {
		d, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		x.FromDuration(d)
	}
	return nil
}
