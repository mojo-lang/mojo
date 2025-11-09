package core

import (
	"time"
)

const DateTimeTypeName = "DateTime"
const DateTimeTypeFullName = "mojo.core.DateTime"

func FromDateTime(dt time.Time) *DateTime {
	y, m, d := dt.Date()
	name, offset := dt.Zone()

	return &DateTime{
		Year:        int32(y),
		Month:       int32(m),
		Day:         int32(d),
		Hour:        int32(dt.Hour()),
		Minute:      int32(dt.Minute()),
		Seconds:     int32(dt.Second()),
		Nanoseconds: int32(dt.Nanosecond()),
		TimeZone:    NewTimeZone(name, int32(offset)),
	}
}

func FromTimestamp(timestamp *Timestamp) *DateTime {
	if timestamp != nil {
		return FromDateTime(timestamp.ToTime())
	}
	return nil
}
