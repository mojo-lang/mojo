package core

import (
	"testing"
	"time"

	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	TimestampString1       = "2019-10-13T19:52:25.000+08:00"
	TimestampString2       = "2019-10-13T19:53:25.000+08:00"
	TimestampString1Json   = `"2019-10-13T19:52:25.000+08:00"`
	TimestampString1Number = "1570967545"
	Timestamp1             = int64(1570967545)
)

var exampleFormats = []string{
	`"2019-10-13T19:52:25+08:00"`,
	`"2019-10-13 19:52:25+0800"`,
	`"2019-10-13 19:52:25 08:00"`,
	`"2019-10-13 11:52:25"`,
	// `"2019/10/13 19:52:25 CST"`,
	`"2019/10/13 11:52:25"`,
	`"2019/10/13 11:52:25+0000"`, // UTC 时间，等同于北京时间 2019/10/13 19:52:25+0800
	`"2019/10/13 11:52:25+00:00"`,
	`"2019-10-13 11:52:25+0000"`,
	`"2019-10-13 11:52:25+00:00"`,
	`"2019-10-13 18:52:25 07:00"`, // test escaped '+'
	`"1570967545"`,
	`"1570967545000"`,
	`1570967545`,
}

func TestTimestampUnmarshal(t *testing.T) {
	var ts Timestamp
	jsoniter.Unmarshal([]byte(TimestampString1Json), &ts)

	assert.Equal(t, Timestamp1, ts.Seconds)
	assert.Equal(t, int32(0), ts.Nanoseconds)

	jsoniter.Unmarshal([]byte(TimestampString1Number), &ts)

	assert.Equal(t, Timestamp1, ts.Seconds)
	assert.Equal(t, int32(0), ts.Nanoseconds)
}

func TestTimestampMarshal(t *testing.T) {
	ts := &Timestamp{Seconds: 1570967545, Nanoseconds: 0}
	str, _ := jsoniter.ConfigDefault.MarshalToString(ts)
	assert.Equal(t, TimestampString1Json, str)
}

func TestTimestampParse(t *testing.T) {
	var ts Timestamp
	for i := range exampleFormats {
		ts.Reset()
		jsoniter.Unmarshal([]byte(exampleFormats[i]), &ts)
		assert.Equal(t, Timestamp1, ts.Seconds, "%s seconds not equal", exampleFormats[i])
		assert.Equal(t, int32(0), ts.Nanoseconds, "%s nanos not equal", exampleFormats[i])
	}
}

func TestTimestamp_Equal(t *testing.T) {
	t1, _ := ParseTimestamp(TimestampString1)
	t12, _ := ParseTimestamp(TimestampString1)
	t2, _ := ParseTimestamp(TimestampString2)

	assert.True(t, t1.Equal(t12))
	assert.False(t, t1.Equal(t2))
}

func TestTimestamp_Sub(t *testing.T) {
	t1, _ := ParseTimestamp(TimestampString1)
	t2, _ := ParseTimestamp(TimestampString2)

	assert.Equal(t, 1.0, t2.Sub(t1).ToMinutes())
	assert.Equal(t, -1.0, t1.Sub(t2).ToMinutes())
}

func TestTimestamp_After(t *testing.T) {
	t1, _ := ParseTimestamp(TimestampString1)
	t2, _ := ParseTimestamp(TimestampString2)

	assert.True(t, t2.After(t1))
}

func TestTimestamp_Before(t *testing.T) {
	t1, _ := ParseTimestamp(TimestampString1)
	t2, _ := ParseTimestamp(TimestampString2)

	assert.True(t, t1.Before(t2))
}

func TestTimestampArithmetic(t *testing.T) {
	base := time.Date(2024, time.February, 28, 12, 30, 45, 123456789, time.Local)
	ts := FromTime(base)
	assert.Equal(t, base.Unix(), ts.Seconds)
	assert.Equal(t, int32(123456789), ts.Nanoseconds)
	assert.True(t, base.Equal(ts.ToTime()))
	assert.Equal(t, &Date{Year: 2024, Month: 2, Day: 28}, ts.Date())
	next := ts.Add(FromDuration(24 * time.Hour))
	assert.Equal(t, int32(29), next.Date().Day)
	assert.Equal(t, 24*time.Hour, next.Sub(ts).ToDuration())
	assert.Equal(t, base.AddDate(1, 2, 3).Unix(), ts.AddDate(1, 2, 3).Seconds)
	var copy Timestamp
	assert.Same(t, &copy, copy.FromTime(base))
	assert.True(t, ts.Equal(&copy))
	assert.True(t, ts.Equal(ts))
	assert.False(t, ts.Before(ts))
	assert.False(t, ts.After(ts))
	assert.Equal(t, -1, ts.Compare(next))
	assert.Equal(t, 1, next.Compare(ts))
	assert.Zero(t, ts.Compare(&copy))
	assert.Equal(t, base.Unix(), ts.Seconds, "arithmetic leaves the receiver unchanged")
}

func TestTimestampNilAndEpoch(t *testing.T) {
	var absent *Timestamp
	present := FromTime(time.Unix(-1, 500000000))
	assert.Equal(t, int64(-1), present.Seconds)
	assert.Equal(t, int32(500000000), present.Nanoseconds)
	assert.True(t, time.Unix(-1, 500000000).Equal(present.ToTime()))
	assert.True(t, absent.ToTime().IsZero())
	assert.Nil(t, absent.FromTime(time.Now()))
	assert.Nil(t, absent.Date())
	assert.Nil(t, absent.Add(Second(1)))
	assert.Nil(t, present.Add(nil))
	assert.Nil(t, absent.AddDate(1, 1, 1))
	assert.Nil(t, absent.Sub(present))
	assert.Nil(t, present.Sub(nil))
	assert.Nil(t, Since(nil))
	assert.Nil(t, Until(nil))
	assert.False(t, absent.Before(present))
	assert.False(t, present.Before(nil))
	assert.False(t, absent.After(present))
	assert.False(t, present.After(nil))
	assert.False(t, absent.Equal(present))
	assert.False(t, present.Equal(nil))
	assert.Zero(t, absent.Compare(nil))
	assert.Equal(t, -1, absent.Compare(present))
	assert.Equal(t, 1, present.Compare(nil))
}

func TestTimestampClockHelpers(t *testing.T) {
	before := time.Now()
	now := Now()
	after := time.Now()
	assert.False(t, now.ToTime().Before(before))
	assert.False(t, now.ToTime().After(after))
	target := FromTime(before.Add(-time.Hour))
	elapsed := Since(target).ToDuration()
	assert.GreaterOrEqual(t, elapsed, before.Sub(target.ToTime()))
	assert.LessOrEqual(t, elapsed, time.Now().Sub(target.ToTime()))
	future := FromTime(time.Now().Add(time.Hour))
	before = time.Now()
	remaining := Until(future).ToDuration()
	assert.LessOrEqual(t, remaining, future.ToTime().Sub(before))
	assert.GreaterOrEqual(t, remaining, future.ToTime().Sub(time.Now()))
}

func TestTimestampJSONBoundaries(t *testing.T) {
	for _, tt := range []struct {
		input   string
		seconds int64
		nanos   int32
	}{
		{`0`, 0, 0}, {`-1`, -1, 0}, {`2147483647`, 2147483647, 0}, {`2147483648`, 2147483648, 0},
		{`"2019-10-13T11:52:25.125Z"`, 1570967545, 125000000},
		{`"1570967545125"`, 1570967545, 125000000},
	} {
		t.Run(tt.input, func(t *testing.T) {
			ts := Timestamp{Seconds: 10, Nanoseconds: 999}
			require.NoError(t, jsoniter.UnmarshalFromString(tt.input, &ts))
			assert.Equal(t, tt.seconds, ts.Seconds)
			assert.Equal(t, tt.nanos, ts.Nanoseconds, "decoding must replace any previous fractional part")
		})
	}
	var ts Timestamp
	assert.Error(t, jsoniter.UnmarshalFromString(`"not-a-timestamp"`, &ts))
}
