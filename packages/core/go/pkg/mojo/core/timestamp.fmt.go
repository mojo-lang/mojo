package core

import (
	"time"

	"github.com/araddon/dateparse"
)

var normalFormatLen = len("2006-01-02 15:04:05")

func (x *Timestamp) Format() string {
	if x != nil {
		return x.ToTime().Format("2006-01-02T15:04:05.000Z07:00")
	}
	return ""
}

func (x *Timestamp) ToString() string {
	return x.Format()
}

func ParseTimestamp(value string) (*Timestamp, error) {
	ts := &Timestamp{}
	if err := ts.Parse(value); err != nil {
		return nil, err
	}
	return ts, nil
}

func (x *Timestamp) Parse(value string) error {
	if x != nil {
		// ts has timezone info, like "2006-01-02 15:04:05+0800"
		// since '+' will be replaced by space in url, we restore it to '+' if possible
		if len(value) > normalFormatLen && value[normalFormatLen] == ' ' {
			value = value[:normalFormatLen] + "+" + value[normalFormatLen+1:]
		}

		t, err := dateparse.ParseIn(value, time.UTC)
		if err != nil {
			return err
		}

		x.FromTime(t)
	}
	return nil
}
