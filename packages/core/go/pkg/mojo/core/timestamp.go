package core

import (
	"time"
)

const TimestampTypeName = "Timestamp"
const TimestampTypeFullName = "mojo.core.Timestamp"

func Now() *Timestamp {
	return FromTime(time.Now())
}

// FromTime convert from time.Time to Timestamp
func FromTime(t time.Time) *Timestamp {
	seconds := t.Unix()
	return &Timestamp{
		Seconds:     seconds,
		Nanoseconds: int32(t.UnixNano() - seconds*1000000000),
	}
}

// Since returns the time elapsed since t.
// It is shorthand for Now().Sub(t).
func Since(t *Timestamp) *Duration {
	if t != nil {
		return FromDuration(time.Since(t.ToTime()))
	}
	return nil
}

// Until returns the duration until t.
// It is shorthand for t.Sub(Now()).
func Until(t *Timestamp) *Duration {
	if t != nil {
		return FromDuration(time.Until(t.ToTime()))
	}
	return nil
}

// FromTime convert from time.Time to Timestamp
func (x *Timestamp) FromTime(t time.Time) *Timestamp {
	if x != nil {
		ts := FromTime(t)
		x.Seconds = ts.Seconds
		x.Nanoseconds = ts.Nanoseconds
	}
	return x
}

// ToTime convert the Timestamp to time.Time
func (x *Timestamp) ToTime() time.Time {
	if x != nil {
		return time.Unix(x.Seconds, int64(x.Nanoseconds))
	} else {
		return time.Time{}
	}
}

// After reports whether the time instant t is after u.
func (x *Timestamp) After(u *Timestamp) bool {
	if x != nil && u != nil {
		if x != u {
			return x.ToTime().After(u.ToTime())
		}
	}
	return false
}

// Before reports whether the time instant t is before u.
func (x *Timestamp) Before(u *Timestamp) bool {
	if x != nil && u != nil {
		if x != u {
			return x.ToTime().Before(u.ToTime())
		}
	}
	return false
}

// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 and 4:00 UTC are Equal.
func (x *Timestamp) Equal(u *Timestamp) bool {
	if x != nil && u != nil {
		if x == u {
			return true
		}
		return x.ToTime().Equal(u.ToTime())
	}
	return false
}

func (x *Timestamp) Compare(u *Timestamp) int {
	if x != nil {
		if u != nil {
			return x.ToTime().Compare(u.ToTime())
		} else {
			return 1
		}
	} else {
		if u != nil {
			return -1
		} else {
			return 0
		}
	}
}

func (x *Timestamp) Date() *Date {
	if x != nil {
		y, m, d := x.ToTime().Date()
		return &Date{
			Year:  int32(y),
			Month: int32(m),
			Day:   int32(d),
		}
	}
	return nil
}

func (x *Timestamp) Add(d *Duration) *Timestamp {
	if x != nil && d != nil {
		return FromTime(x.ToTime().Add(d.ToDuration()))
	}
	return nil
}

// AddDate returns the time corresponding to adding the
// given number of years, months, and days to t.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// AddDate normalizes its result in the same way that Date does,
// so, for example, adding one month to October 31 yields
// December 1, the normalized form for November 31.
func (x *Timestamp) AddDate(years int, months int, days int) *Timestamp {
	if x != nil {
		return FromTime(x.ToTime().AddDate(years, months, days))
	}
	return nil
}

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
func (x *Timestamp) Sub(u *Timestamp) *Duration {
	if x != nil && u != nil {
		return FromDuration(x.ToTime().Sub(u.ToTime()))
	}
	return nil
}
