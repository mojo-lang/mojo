package core

import (
	"math"
	"time"
)

const DurationTypeName = "Duration"
const DurationTypeFullName = "mojo.core.Duration"

func FromDuration(d time.Duration) *Duration {
	duration := &Duration{}
	return duration.FromDuration(d)
}

func NewDurationFromSeconds(seconds float64) *Duration {
	duration := &Duration{}
	return duration.FromSeconds(seconds)
}

func Hours(hours int) *Duration {
	return FromDuration(time.Duration(hours) * time.Hour)
}

func Minute(minutes int) *Duration {
	return FromDuration(time.Duration(minutes) * time.Minute)
}

func Second(seconds int) *Duration {
	return FromDuration(time.Duration(seconds) * time.Second)
}

func (x *Duration) ToDuration() time.Duration {
	return time.Duration(x.Seconds)*time.Second + time.Duration(x.Nanoseconds)
}

func (x *Duration) FromDuration(d time.Duration) *Duration {
	if x != nil {
		x.Seconds = int64(d) / int64(time.Second)
		x.Nanoseconds = int32(int64(d) - x.Seconds*int64(time.Second))
	}
	return x
}

func (x *Duration) FromSeconds(seconds float64) *Duration {
	if x != nil {
		x.Seconds = int64(seconds)
		delta := seconds - float64(x.Seconds)
		x.Nanoseconds = int32(math.Round(delta * float64(time.Second)))
	}
	return x
}

func (x *Duration) ToHours() float64 {
	if x != nil {
		return x.ToDuration().Hours()
	}
	return 0
}

func (x *Duration) ToMinutes() float64 {
	if x != nil {
		return x.ToDuration().Minutes()
	}
	return 0
}

func (x *Duration) ToSeconds() float64 {
	if x != nil {
		return x.ToDuration().Seconds()
	}
	return 0
}

func (x *Duration) ToNanoseconds() int64 {
	if x != nil {
		return x.ToDuration().Nanoseconds()
	}
	return 0
}

func (x *Duration) Compare(d *Duration) int {
	if x != nil {
		if d != nil {
			if x.Seconds == d.Seconds {
				if x.Nanoseconds == d.Nanoseconds {
					return 0
				} else if x.Nanoseconds > d.Nanoseconds {
					return 1
				} else {
					return -1
				}
			} else if x.Seconds > d.Seconds {
				return 1
			} else {
				return -1
			}
		} else {
			return 1
		}
	} else if d != nil {
		return -1
	}
	return 0
}

func (x *Duration) FromFloat64(seconds float64) error {
	if x != nil {
		x.Seconds = int64(seconds)
		delta := seconds - float64(x.Seconds)
		x.Nanoseconds = int32(math.Round(delta * float64(time.Second)))
	}
	return nil
}
