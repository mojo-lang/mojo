package core

const TimeZoneTypeName = "TimeZone"
const TimeZoneTypeFullName = "mojo.core.TimeZone"

func NewTimeZone(name string, offset int32) *TimeZone {
	return &TimeZone{
		Offset: offset,
		Name:   name,
	}
}
